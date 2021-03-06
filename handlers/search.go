package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type searchRequest struct {
	Query string `json:"query"`
}

// Search is a handler for the search funcionality
type Search struct{}

func (s *Search) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	request := new(searchRequest)
	err := decoder.Decode(request)
	if err != nil || len(request.Query) < 1 {
		log.Println("hadlers.Search: err unmarshaling,", err)
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
}
