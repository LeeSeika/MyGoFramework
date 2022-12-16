package gin

import "golang.org/x/net/context"

//type Context struct {
//	request        *http.Request
//	responseWriter http.ResponseWriter
//	hasTimeout     bool
//
//	index    int
//	handlers []ControllerHandler
//	params   map[string]string
//
//	writerMux *sync.Mutex
//}
//
//func NewContext(r *http.Request, w http.ResponseWriter) *Context {
//	return &Context{
//		request:        r,
//		responseWriter: w,
//		hasTimeout:     false,
//		writerMux:      &sync.Mutex{},
//		index:          -1,
//	}
//}

// base

//func (ctx *Context) WriterMux() *sync.Mutex {
//	return ctx.writerMux
//}
//
//func (ctx *Context) GetRequest() *http.Request {
//	return ctx.request
//}
//
//func (ctx *Context) GetResponse() http.ResponseWriter {
//	return ctx.responseWriter
//}
//
//func (ctx *Context) SetHasTimeout() {
//	ctx.hasTimeout = true
//}
//
//func (ctx *Context) HasTimeOut() bool {
//	return ctx.hasTimeout
//}
//
//// middleware
//
//func (ctx *Context) SetHandlers(handlers ...ControllerHandler) {
//	ctx.handlers = append(ctx.handlers, handlers...)
//}
//
//func (ctx *Context) Next() error {
//	if ctx.index < len(ctx.handlers)-1 {
//		ctx.index++
//		err := ctx.handlers[ctx.index](ctx)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//// context
//
//func (ctx *Context) BaseContext() context.Context {
//	return ctx.request.Context()
//}
//
//func (ctx *Context) Done() <-chan struct{} {
//	return ctx.BaseContext().Done()
//}
//
//func (ctx *Context) DeadLine() (time.Time, bool) {
//	return ctx.BaseContext().Deadline()
//}
//
//func (ctx *Context) Err() error {
//	return ctx.BaseContext().Err()
//}
//
//func (ctx *Context) Value(key interface{}) interface{} {
//	return ctx.BaseContext().Value(key)
//}
//
//func (ctx *Context) SetParams(params map[string]string) {
//	ctx.params = params
//}

func (c *Context) BaseContext() context.Context {
	return c.Request.Context()
}
