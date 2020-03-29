//package main
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//
//	"github.com/julienschmidt/httprouter"
//)
//
//func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	fmt.Fprint(w, "Welcome!\n")
//}
//
//func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
//}
//
//func main() {
//	router := httprouter.New()
//	router.GET("/", Index)
//	router.GET("/hello/:name", Hello)
//
//	log.Fatal(http.ListenAndServe(":8080", router))
//}

package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)

	return router
}

func main(){
	r := RegisterHandlers()
	log.Fatal(http.ListenAndServe(":8000", r))
}