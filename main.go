package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joao-fontenele/go-rest-api/handlers"
)

var port = 8080

func main() {
	log.Printf("Server listening on port \"%v\"\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), &handlers.Search{})
	if err != nil {
		log.Fatal(err)
	}
}
