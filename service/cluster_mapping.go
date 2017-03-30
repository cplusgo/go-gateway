package service

import (
	"net/http"
	"regexp"
)

//包含业务集群列表，不同的服务集群使用不同的名称来区分
type ClusterMapping struct {
	clusterMapping map[string]*Cluster
}

func NewClusterMapping() *ClusterMapping  {
	mapping := make(map[string]*Cluster)
	return &ClusterMapping{clusterMapping:mapping}
}

func (this *ClusterMapping) AddService(clusterName, serviceName, url string) {
	if _, isExist := this.clusterMapping[clusterName]; !isExist {
		services := make([]*Service, 0)
		this.clusterMapping[clusterName] = &Cluster{name: clusterName, services: services}
	}
	this.clusterMapping[clusterName].addService(serviceName, url)
}

func (this *ClusterMapping) ServeRequest(w http.ResponseWriter, req *http.Request) {
	//这里直接解析出来该请求的服务类型
	reg, _ := regexp.Compile("service=([a-z0-9A-Z]+)")
	seg := reg.FindStringSubmatch(req.URL.RawQuery)
	serviceName := seg[1]
	if cluster, ok := this.clusterMapping[serviceName]; ok {
		cluster.serveRequest(w, req)
	}
}
