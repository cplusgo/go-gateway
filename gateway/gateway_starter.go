package gateway

import (
	"net/http"
	"github.com/cplusgo/go-gateway/filter"
	"github.com/cplusgo/go-gateway/service"
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/cplusgo/go-gateway/data"
)

type GatewayStarter struct {
	filters []filter.IFilter
	clusterMapping *service.ClusterMapping
}

func NewGatewayStarter() *GatewayStarter {
	starter := &GatewayStarter{
		filters: filter.HttpFilters,
		clusterMapping:service.NewClusterMapping(),
	}
	return starter
}

func (this *GatewayStarter) dispatch(w http.ResponseWriter, req *http.Request) {
	this.clusterMapping.ServeRequest(w, req)
}

func (this *GatewayStarter) filter(w http.ResponseWriter, req *http.Request) bool {
	var isPass bool = true
	for _, filter := range this.filters {
		passed := filter.Filter(req)
		if isPass {
			isPass = passed
		}
	}
	return isPass
}

func (this *GatewayStarter) gateway(w http.ResponseWriter, req *http.Request) {
	isPassed := this.filter(w, req)
	if isPassed {
		this.dispatch(w, req)
	}
}

func (this *GatewayStarter) parseSetting() {
	file, err := os.Open("/home/aron/workspace/go/src/github.com/cplusgo/go-gateway/data/config.json")
	if err != nil {
		log.Fatal("parse config.json failed !")
		os.Exit(0)
	}
	bytes, _ := ioutil.ReadAll(file)
	vo := &data.ClusterVoList{}
	json.Unmarshal(bytes, vo)
	fmt.Println(vo.Clusters[0].Name)

	for _,cluster := range vo.Clusters {
		for _, service := range cluster.Services {
			this.clusterMapping.AddService(cluster.Name, service.Name, service.Url)
		}
	}
}

func (this *GatewayStarter) Start() {
	this.parseSetting()
	http.HandleFunc("/", this.gateway)
	http.ListenAndServe(":8080", nil)
}
