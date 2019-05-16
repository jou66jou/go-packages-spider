package main

import (
	"fmt"
	"gopkgSpider/routers"
	"log"
	"net/http"
)

func main() {
	url := "https://golang.google.cn/pkg/"
	fmt.Println("Main: url set-" + url)

	// run spider gopkgs and insert to mysql
	// md.SpiderGopkgs(url)

	// starts an api service with a given routers.
	log.Fatal(http.ListenAndServe(":8081", routers.Routers()))
}
