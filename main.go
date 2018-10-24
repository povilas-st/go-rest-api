package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	address := ":80"
	log.Println("Running http server on " + address)

	log.Fatal(http.ListenAndServe(address, nil))
}
