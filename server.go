package main

import (
	route "github.com/heaptracetechnology/microservice-urban-airship/route"
	"log"
	"net/http"
)

func main() {
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
