package routers

import (
	"fmt"
	"gopkgSpider/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/gopkg/{name}", GetGopkg).Methods("GET")

	return router
}

// GET - /api/gopkg/{name}, reponse go package content form grpc service.
func GetGopkg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		name := vars["name"]
		jPkg := models.RequestToPB(name)
		fmt.Println(jPkg)
		fmt.Fprintf(w, jPkg)

	case "POST":
		fmt.Fprintf(w, "please sending Go package name as name key by get.")
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
