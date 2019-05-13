package main

import (
	"fmt"
	"gopkgSpider/models"
)

func main() {
	url := "https://golang.google.cn/pkg/"
	fmt.Println("Main: url set-" + url)

	data, getE := models.GetGopkgsList(url) // get golang pkgs standard lib description list
	if getE != nil {
		fmt.Println(getE)
		return
	}
	fmt.Println("Main: get gopkgslist success")

	insertDataE := models.InsertPkgData(data) // insert pkgs list to db
	if insertDataE != nil {
		fmt.Println(insertDataE)
		return
	}
	fmt.Println("Main: gopkgslist insert to db success")

	insertOverViewE := models.GetOVAndInsert() // routine get pkgs overview and insert to db

	if insertOverViewE != nil {
		fmt.Println(getE)
		return
	}
	fmt.Println("Main: pkgs overview insert to db success")
	models.CloseDB()
	fmt.Println("db closs")

}
