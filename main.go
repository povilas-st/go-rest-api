package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-rest-api/users"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/ping", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "pong")
	})
	router.GET("/users", users.Index)
	router.GET("/users/:id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users/:id", users.Update)
	router.DELETE("/users/:id", users.Delete)

	address := ":80"
	log.Println("Running http server on " + address)
	log.Fatal(http.ListenAndServe(address, router))
}
