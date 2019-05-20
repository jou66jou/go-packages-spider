package main

import (
	"fmt"
	"gopkg-spider/models"
	"gopkg-spider/protoserver"
	"gopkg-spider/routers"
	"log"
	"net/http"
)

func main() {
	var port string = "8080"
	fmt.Println("Http listen port (default 8080):")
	fmt.Scanln(&port)
	// run spider gopkgs and insert to mysql
	fmt.Println("run spider gopkgs...")
	models.SpiderGopkgs()

	// starts an gRPC service and return server done.
	go protoserver.Start_gRPC()

	// starts an api service with a given routers.
	fmt.Println("HTTP Listening on " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, routers.Routers()))

}
