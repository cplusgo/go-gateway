package main

import (
	"github.com/cplusgo/go-gateway/gateway"
	"os"
	"log"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(0)
	}
	path = path + "/web"
	os.Chdir(path)
	log.Println(os.Getwd())
	starter := gateway.NewGatewayStarter()
	starter.Start()
}
