package framework

// provider.go 所有提供服务的方法都写在这个文件里

type NewInstance func(...interface{}) (interface{}, error)

type ServiceProvider interface {
	Register(Container) NewInstance
	Name() string
	Params(container Container) []interface{}
	IsDefer() bool
	Boot(Container) error
}
