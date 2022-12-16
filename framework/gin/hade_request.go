package gin

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/spf13/cast"
	"io/ioutil"
	"mime/multipart"
)

// 代表请求包含的方法

type IRequest interface {

	// 请求地址 url 中带的参数
	// 形如: foo.com?a=1&b=bar&c[]=bar

	DefaultQueryInt(key string, def int) (int, bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)

	// 路由匹配中带的参数
	// 形如 /book/:id

	DefaultParamInt(key string, def int) (int, bool)
	DefaultParamInt64(key string, def int64) (int64, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParamBool(key string, def bool) (bool, bool)
	DefaultParamString(key string, def string) (string, bool)
	DefaultParam(key string) interface{}

	// form 表单中带的参数

	DefaultFormInt(key string, def int) (int, bool)
	DefaultFormInt64(key string, def int64) (int64, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormFloat32(key string, def float32) (float32, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultFormString(key string, def string) (string, bool)
	DefaultFormStringSlice(key string, def []string) ([]string, bool)
	DefaultFormFile(key string) (*multipart.FileHeader, error)
	DefaultForm(key string) interface{}

	// json body
	BindJson(obj interface{}) error

	// xml body
	BindXml(obj interface{}) error

	// 其他格式
	GetRawData() ([]byte, error)

	// 基础信息
	Uri() string
	Method() string
	Host() string
	ClientIp() string

	// header
	Headers() map[string][]string
	Header(key string) (string, bool)

	// cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

// request

func (ctx *Context) QueryAll() map[string][]string {
	ctx.initQueryCache()
	return ctx.queryCache
}

func (ctx *Context) DefaultQueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len := len(values); len > 0 {
			return cast.ToInt(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len := len(values); len > 0 {
			return cast.ToInt64(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len := len(values); len > 0 {
			return cast.ToFloat32(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len := len(values); len > 0 {
			return cast.ToFloat64(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len := len(values); len > 0 {
			return cast.ToBool(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len := len(values); len > 0 {
			return values[0], true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryStringLice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return values, true
	}
	return def, false
}

//func (ctx *Context) Query(key string) interface{} {
//	params := ctx.QueryAll()
//	if values, ok := params[key]; ok {
//		if len := len(values); len > 0 {
//			return values[0]
//		}
//	}
//	return nil
//}

//func (ctx *Context) QueryArray(key string, def []string) []string {
//	params := ctx.QueryAll()
//	if values, ok := params[key]; ok {
//		if len := len(values); len > 0 {
//			return values
//		}
//	}
//	return def
//}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.Request != nil {
		return ctx.Request.PostForm
	}
	return map[string][]string{}
}

func (ctx *Context) DefaultFormInt(key string, def int) (int, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len := len(values); len > 0 {
			return cast.ToInt(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormInt64(key string, def int64) (int64, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len := len(values); len > 0 {
			return cast.ToInt64(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormFloat32(key string, def float32) (float32, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len := len(values); len > 0 {
			return cast.ToFloat32(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormFloat64(key string, def float64) (float64, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len := len(values); len > 0 {
			return cast.ToFloat64(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormBool(key string, def bool) (bool, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len := len(values); len > 0 {
			return cast.ToBool(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormString(key string, def string) (string, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len := len(values); len > 0 {
			return values[len-1], true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormStringLice(key string, def []string) ([]string, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		return values, true
	}
	return def, false
}

func (ctx *Context) DefaultFormFile(key string) (*multipart.FileHeader, error) {
	return nil, nil
}

func (ctx *Context) DefaultForm(key string) interface{} {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len := len(values); len > 0 {
			return values[0]
		}
	}
	return nil
}

func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len := len(values); len > 0 {
			return values
		}
	}
	return def
}

func (ctx *Context) DefaultParamInt(key string, def int) (int, bool) {
	if value, ok := ctx.Params.Get(key); ok {
		return cast.ToInt(value), true
	}
	return def, false
}

func (ctx *Context) DefaultParamInt64(key string, def int64) (int64, bool) {
	if value, ok := ctx.Params.Get(key); ok {
		return cast.ToInt64(value), true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat64(key string, def float64) (float64, bool) {
	if value, ok := ctx.Params.Get(key); ok {
		return cast.ToFloat64(value), true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat32(key string, def float32) (float32, bool) {
	if value, ok := ctx.Params.Get(key); ok {
		return cast.ToFloat32(value), true
	}
	return def, false
}

func (ctx *Context) DefaultParamBool(key string, def bool) (bool, bool) {
	if value, ok := ctx.Params.Get(key); ok {
		return cast.ToBool(value), true
	}
	return def, false
}

func (ctx *Context) DefaultParamString(key string, def string) (string, bool) {
	if value, ok := ctx.Params.Get(key); ok {
		return value, true
	}
	return def, false
}

func (ctx *Context) DefaultParam(key string) interface{} {
	if value, ok := ctx.Params.Get(key); ok {
		return value
	}
	return nil
}

func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.Request == nil {
		return errors.New("context.Request cannot be nil")
	}
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	err = json.Unmarshal(body, obj)
	if err != nil {
		return err
	}
	return nil
}

func (ctx *Context) BindXml(obj interface{}) error {
	return nil
}

// 其他格式
//func (ctx *Context) GetRawData() ([]byte, error) {
//	return nil, nil
//}

// 基础信息
func (ctx *Context) Uri() string {
	return ctx.Request.URL.Path
}

func (ctx *Context) Method() string {
	return ctx.Request.Method
}
func (ctx *Context) Host() string {
	return ctx.Request.Host
}

func (ctx *Context) ClientIp() string {
	return ctx.Request.RemoteAddr
}

// header
func (ctx *Context) Headers() map[string][]string {
	return ctx.Request.Header
}

//func (ctx *Context) Header(key string) (string, bool) {
//	if value, ok := ctx.Request.Header[key]; ok {
//		len := len(value)
//		if len > 0 {
//			return value[0], true
//		}
//	}
//	return "", false
//}

// cookie
func (ctx *Context) Cookies() map[string]string {
	cookies := ctx.Request.Cookies()
	cookieMap := map[string]string{}
	for _, cookie := range cookies {
		cookieMap[cookie.Name] = cookie.Value
	}
	return cookieMap
}

//func (ctx *Context) Cookie(key string) (string, bool) {
//	cookieMap := ctx.Cookies()
//	if value, ok := cookieMap[key]; ok {
//		return value, true
//	}
//	return "", false
//}
