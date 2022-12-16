package framework

//
//import (
//	"github.com/godemo/coredemo/framework/gin"
//	"net/http"
//	"strings"
//)
//
//type Core struct {
//	router        map[string]map[string]ControllerHandler
//	dynamicRouter map[string]*Tree
//	middleware    []ControllerHandler
//}
//
//func NewCore() *Core {
//	getRouter := map[string]ControllerHandler{}
//	postRouter := map[string]ControllerHandler{}
//	putRouter := map[string]ControllerHandler{}
//	deleteRouter := map[string]ControllerHandler{}
//
//	core := Core{router: map[string]map[string]ControllerHandler{}}
//	core.router["GET"] = getRouter
//	core.router["POST"] = postRouter
//	core.router["PUT"] = putRouter
//	core.router["DELETE"] = deleteRouter
//
//	core.dynamicRouter = map[string]*Tree{}
//	core.dynamicRouter["GET"] = gin.newTree()
//	core.dynamicRouter["POST"] = gin.newTree()
//	core.dynamicRouter["PUT"] = gin.newTree()
//	core.dynamicRouter["DELETE"] = gin.newTree()
//
//	return &core
//}
//
//// register router
//
//func (c *Core) Get(pattern string, handler ControllerHandler) {
//	upper := strings.ToUpper(pattern)
//	c.router["GET"][upper] = handler
//}
//
//func (c *Core) Post(pattern string, handler ControllerHandler) {
//	upper := strings.ToUpper(pattern)
//	c.router["POST"][upper] = handler
//}
//
//func (c *Core) Put(pattern string, handler ControllerHandler) {
//	upper := strings.ToUpper(pattern)
//	c.router["PUT"][upper] = handler
//}
//
//func (c *Core) Delete(pattern string, handler ControllerHandler) {
//	upper := strings.ToUpper(pattern)
//	c.router["DELETE"][upper] = handler
//}
//
//func (c *Core) GetDynamic(pattern string, handlers ...ControllerHandler) {
//	upper := strings.ToUpper(pattern)
//	allHandlers := append(c.middleware, handlers...)
//	c.dynamicRouter["GET"].AddRouter(upper, allHandlers...)
//}
//
//func (c *Core) PostDynamic(pattern string, handlers ...ControllerHandler) {
//	upper := strings.ToUpper(pattern)
//	allHandlers := append(c.middleware, handlers...)
//	c.dynamicRouter["POST"].AddRouter(upper, allHandlers...)
//}
//
//func (c *Core) PutDynamic(pattern string, handlers ...ControllerHandler) {
//	upper := strings.ToUpper(pattern)
//	allHandlers := append(c.middleware, handlers...)
//	c.dynamicRouter["PUT"].AddRouter(upper, allHandlers...)
//}
//
//func (c *Core) DeleteDynamic(pattern string, handlers ...ControllerHandler) {
//	upper := strings.ToUpper(pattern)
//	allHandlers := append(c.middleware, handlers...)
//	c.dynamicRouter["DELETE"].AddRouter(upper, allHandlers...)
//}
//
//func (c *Core) Group(prefix string) IGroup {
//	group := NewGroup(c, prefix)
//	return group
//}
//
//// find router
//
//func (c *Core) FindRouterByRequest(request *http.Request) ControllerHandler {
//	method := strings.ToUpper(request.Method)
//	url := strings.ToUpper(request.URL.Path)
//	if methodMap, ok := c.router[method]; ok {
//		if handler, ok := methodMap[url]; ok {
//			return handler
//		}
//	}
//	return nil
//}
//
//func (c *Core) FindDynamicRouterByRequest(request *http.Request) *gin.node {
//	method := strings.ToUpper(request.Method)
//	url := strings.ToUpper(request.URL.Path)
//	if methodMap, ok := c.dynamicRouter[method]; ok {
//		if node, error := methodMap.FindHandler(url); error == nil {
//			return node
//		}
//	}
//	return nil
//}
//
//func (c *Core) Use(handlers ...ControllerHandler) {
//	c.middleware = append(c.middleware, handlers...)
//}
//
//func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
//	context := NewContext(request, response)
//
//	node := c.FindDynamicRouterByRequest(request)
//	params := node.parseParamFromEndNode(request.URL.Path)
//	if node.handlers == nil {
//		context.Json(500, "not found")
//		return
//	}
//	context.SetHandlers(node.handlers...)
//	context.SetParams(params)
//
//	err := context.Next()
//	if err != nil {
//		context.Json(500, "inner error")
//		return
//	}
//
//}
