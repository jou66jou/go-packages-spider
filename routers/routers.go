package routers

import (
	"fmt"
	"gopkgSpider/auth"
	"gopkgSpider/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	fmt.Println("HTTP Method list:")
	fmt.Println("GET method : /api/gopkg/{name} - call the api whit JWT authorization to get go package description")
	register("GET", "/api/gopkg/{name}", controllers.GetGopkg, auth.TokenMiddleware)
	fmt.Println("POST method : /user/login - post json{username:\"test\"} to get JWT authorization, the auth time out after 2 hours")
	register("POST", "/user/login", controllers.Login, nil)
	fmt.Println("")

}

func Routers() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		r := router.Methods(route.Method).
			Path(route.Pattern)
		if route.Middleware != nil { // JWT valid
			r.Handler(route.Middleware(route.Handler))
		} else {
			r.Handler(route.Handler)
		}
	}
	return router

}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
