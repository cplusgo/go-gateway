package filter

import (
	"net/http"
	"log"
)

type IFilter interface {
	Filter(req *http.Request) bool
}

var HttpFilters []IFilter = []IFilter{
	LogFilter{},
}

/**
 *注册一个简单的过滤器
 *如果想要实现自己的过滤器
 *可以参考这个实现方式
 */
type LogFilter struct {

}

func (this LogFilter) Filter(req *http.Request) bool {
	log.Println(req.RemoteAddr, req.RequestURI, req.Referer())
	return true
}