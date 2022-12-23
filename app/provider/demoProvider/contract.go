package demoProvider

const DemoBizKey = "hade:demo"

// demo服务的接口
type IService interface {
	GetAllStudent() []Student
}

// 虚构的一个业务数据结构
type Student struct {
	ID   int
	Name string
}
