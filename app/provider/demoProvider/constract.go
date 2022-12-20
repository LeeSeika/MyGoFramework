package demoProvider

const Key = "hade:demo"

// demo服务的接口
type Service interface {
	GetFoo() Foo
}

// 虚构的一个业务数据结构
type Foo struct {
	Name string
}
