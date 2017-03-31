package data

type ServiceVo struct {
	Name string `json:"name"`
	Domain  string `json:"domain"`
}

type ClusterVo struct {
	Name     string `json:"name"`
	Services []ServiceVo `json:"services"`
}

type ClusterVoList struct {
	Clusters []ClusterVo `json:"clusters"`
}
