package framework

import (
	"errors"
	"fmt"
	"sync"
)

type Container interface {
	IsBind(key string) bool
	Bind(provider ServiceProvider) error
	Make(key string) (interface{}, error)
	MustMake(key string) interface{}
	MakeNew(key string, params []interface{}) (interface{}, error)
}

type HadeContainer struct {
	Container
	providers map[string]ServiceProvider
	instances map[string]interface{}
	lock      sync.RWMutex
}

func NewHadeContainer() *HadeContainer {
	return &HadeContainer{
		instances: map[string]interface{}{},
		providers: map[string]ServiceProvider{},
		lock:      sync.RWMutex{},
	}
}

func (ctn *HadeContainer) Bind(provider ServiceProvider) error {
	ctn.lock.Lock()
	key := provider.Name()
	ctn.providers[key] = provider
	ctn.lock.Unlock()

	if provider.IsDefer() == false {
		err := provider.Boot(ctn)
		if err != nil {
			return err
		}
		makeInstance := provider.Register(ctn)
		params := provider.Params(ctn)
		instance, err := makeInstance(params...)
		if err != nil {
			return err
		}
		ctn.instances[key] = instance
	}
	return nil
}

func (ctn *HadeContainer) IsBind(key string) bool {
	ctn.lock.RLock()
	defer ctn.lock.RUnlock()
	_, ok := ctn.providers[key]
	return ok
}

func (ctn *HadeContainer) Make(key string) (interface{}, error) {
	return ctn.make(key, nil, false)
}

func (ctn *HadeContainer) MustMake(key string) interface{} {
	instance, _ := ctn.make(key, nil, false)
	return instance
}

// 不使用map里的单例，使用方法传入的params作为参数创建一个实例

func (ctn *HadeContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return ctn.make(key, params, true)
}

// private

func (ctn *HadeContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	ctn.lock.RLock()
	defer func() {
		ctn.lock.RUnlock()
		//fmt.Println("error!!!!!!!!!!!!!!!")
	}()

	provider := ctn.findServiceProvider(key)
	if provider == nil {
		fmt.Println("error!!!!!!!!!!!!!!!")
		return nil, errors.New("provider " + key + " not found")
	}

	if forceNew {
		return ctn.newInstance(provider, params)
	}

	if instance, ok := ctn.instances[key]; ok {
		return instance, nil
	}

	instance, err := ctn.newInstance(provider, nil)
	if err != nil {
		return nil, err
	}

	ctn.instances[key] = instance
	return instance, nil
}

func (ctn *HadeContainer) newInstance(provider ServiceProvider, params []interface{}) (interface{}, error) {

	err := provider.Boot(ctn)
	if err != nil {
		return nil, err
	}

	if params == nil {
		params = provider.Params(ctn)
	}

	makeInstance := provider.Register(ctn)
	instance, err := makeInstance(params...)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (ctn *HadeContainer) findServiceProvider(key string) ServiceProvider {
	if provider, ok := ctn.providers[key]; ok {
		return provider
	}
	return nil
}
