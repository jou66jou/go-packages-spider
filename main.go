package main

import (
	"fmt"
	"gopkgSpider/models"
	"gopkgSpider/routers"
	"log"
	"net/http"
)

func main() {
	url := "https://golang.google.cn/pkg/"
	fmt.Println("Main: url set-" + url)

	// run spider gopkgs and insert to mysql
	models.SpiderGopkgs(url)

	// starts an api service with a given routers.
	log.Fatal(http.ListenAndServe(":8081", routers.Routers()))
}
