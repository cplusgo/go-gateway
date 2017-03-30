#go-gateway

ApiGateway 网站服务网关
看了很多实现，大多数的实现需要配置比较多，
所以这里创建一个网关尽量减少配置的数据，减少部署的压力

关键点在于ApiGateway上不需要配置任何url映射，
只要给每个服务定一个名字，
客户端在请求时会附带这个名字(比如: http://example.com?service=ServiceName)，
网关根据这个名字然后把请求转发到具体的服务器上。

这样便减少了大量的url配置流程。

这样做网关的日志统计依然可以正常进行。

![网关架构图](imgs/api_gateway.png =800x500)