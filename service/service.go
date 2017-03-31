package service

import (
	"net/http"
	"github.com/cplusgo/go-gateway/proxy"
)

/**
 *服务实体
 */
type Service struct {
	name          string `json:"name"`
	url           string `json:"url"`
	connectionNum int32 `json:"connection_num"`
	httpProxy     *proxy.HttpProxy
}

func (this *Service) serveRequest(w http.ResponseWriter, req *http.Request) {
	this.connectionNum++
	this.httpProxy.ServeRequest(w, req, this.url)
	this.connectionNum--
}
