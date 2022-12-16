package framework

//
//import (
//	"strings"
//)
//
//type IGroup interface {
//	Get(pattern string, handler ControllerHandler)
//	Post(pattern string, handler ControllerHandler)
//	Put(pattern string, handler ControllerHandler)
//	Delete(pattern string, handler ControllerHandler)
//
//	GetDynamic(pattern string, handler ...ControllerHandler)
//	PostDynamic(pattern string, handler ...ControllerHandler)
//	PutDynamic(pattern string, handler ...ControllerHandler)
//	DeleteDynamic(pattern string, handler ...ControllerHandler)
//
//	// Group 嵌套
//
//	Group(prefix string) IGroup
//}
//
//type Group struct {
//	core       *Core
//	prefix     string
//	parent     *Group
//	middleware []ControllerHandler
//}
//
//func (group *Group) Group(prefix string) IGroup {
//	return &Group{group.core, prefix, nil, nil}
//}
//
//func NewGroup(core *Core, prefix string) *Group {
//	return &Group{
//		core:   core,
//		prefix: prefix,
//	}
//}
//
//func (group *Group) Get(component string, handler ControllerHandler) {
//	url := strings.ToUpper(group.prefix + component)
//	group.core.Get(url, handler)
//}
//
//func (group *Group) Post(component string, handler ControllerHandler) {
//	url := strings.ToUpper(group.prefix + component)
//	group.core.Post(url, handler)
//}
//
//func (group *Group) Put(component string, handler ControllerHandler) {
//	url := strings.ToUpper(group.prefix + component)
//	group.core.Put(url, handler)
//}
//
//func (group *Group) Delete(component string, handler ControllerHandler) {
//	url := strings.ToUpper(group.prefix + component)
//	group.core.Delete(url, handler)
//}
//
//// dynamic router
//
//func (group *Group) GetDynamic(component string, handlers ...ControllerHandler) {
//	url := strings.ToUpper(group.prefix + component)
//	allHandlers := append(group.getMiddleWare(), handlers...)
//	group.core.GetDynamic(url, allHandlers...)
//}
//
//func (group *Group) PostDynamic(component string, handlers ...ControllerHandler) {
//	url := strings.ToUpper(group.prefix + component)
//	allHandlers := append(group.getMiddleWare(), handlers...)
//	group.core.PostDynamic(url, allHandlers...)
//}
//
//func (group *Group) PutDynamic(component string, handlers ...ControllerHandler) {
//	url := strings.ToUpper(group.prefix + component)
//	allHandlers := append(group.getMiddleWare(), handlers...)
//	group.core.PutDynamic(url, allHandlers...)
//}
//
//func (group *Group) DeleteDynamic(component string, handlers ...ControllerHandler) {
//	url := strings.ToUpper(group.prefix + component)
//	allHandlers := append(group.getMiddleWare(), handlers...)
//	group.core.DeleteDynamic(url, allHandlers...)
//}
//
//// middleware
//
//func (group *Group) Use(handlers ...ControllerHandler) {
//	group.middleware = append(group.middleware, handlers...)
//}
//
//func (group *Group) getMiddleWare() []ControllerHandler {
//	if group.parent != nil {
//		return group.parent.getMiddleWare()
//	}
//	return group.middleware
//}
