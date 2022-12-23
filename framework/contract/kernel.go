package contract

import "net/http"

const KernelKey = "hade:kernel"

type Kernel interface {
	// 多态
	HttpEngine() http.Handler
}
