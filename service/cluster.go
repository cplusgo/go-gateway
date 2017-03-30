package service

import (
	"net/http"
	"github.com/cplusgo/go-gateway/proxy"
)

/**
 * 业务集群，集群的名称必须是全局唯一的
 * 这个比较好区分，如账号服务account,订单服务order,商品服务goods
 */
type Cluster struct {
	name     string `json:"name"`
	services []*Service
}

func (this *Cluster) addService(name, url string) {
	service := &Service{
		name: name,
		url: url,
		connectionNum:0,
		httpProxy:proxy.NewHttpProxy(url),
	}
	this.services = append(this.services, service)
}

func (this *Cluster) serveRequest(w http.ResponseWriter, req *http.Request)  {
	this.services[0].serveRequest(w, req)
}