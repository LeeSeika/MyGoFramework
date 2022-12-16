package gin

import (
	"encoding/json"
	"net/http"
)

// IResponse 代表返回方法
type IResponse interface {
	// Json 输出
	IJson(obj interface{}) IResponse

	// Jsonp 输出
	IJsonp(obj interface{}) IResponse

	//xml 输出
	IXml(obj interface{}) IResponse

	// html 输出
	IHtml(template string, obj interface{}) IResponse

	// string
	IText(format string, values ...interface{}) IResponse

	// 重定向
	IRedirect(path string) IResponse

	// header
	ISetHeader(key string, val string) IResponse

	// Cookie
	ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	ISetStatus(code int) IResponse

	// 设置 200 状态
	ISetOkStatus() IResponse
}

func (c *Context) IJsonp(obj interface{}) IResponse {
	//TODO implement me
	panic("implement me")
}

func (c *Context) IXml(obj interface{}) IResponse {
	//TODO implement me
	panic("implement me")
}

func (c *Context) IHtml(template string, obj interface{}) IResponse {
	//TODO implement me
	panic("implement me")
}

func (c *Context) IText(format string, values ...interface{}) IResponse {
	//TODO implement me
	panic("implement me")
}

func (c *Context) IRedirect(path string) IResponse {
	//TODO implement me
	panic("implement me")
}

func (c *Context) ISetHeader(key string, val string) IResponse {
	//TODO implement me
	panic("implement me")
}

func (c *Context) ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	//TODO implement me
	panic("implement me")
}

func (c *Context) ISetOkStatus() IResponse {
	c.Writer.WriteHeader(200)
	return c
}

func (ctx *Context) ISetStatus(code int) IResponse {
	ctx.Writer.WriteHeader(code)
	return ctx
}

func (ctx *Context) IJson(obj interface{}) IResponse {
	jsonRespBody, err := json.Marshal(obj)
	if err != nil {
		ctx.ISetStatus(http.StatusInternalServerError)
	}
	ctx.Writer.Write(jsonRespBody)
	return ctx
}
