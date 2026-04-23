package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	section := query.Get("sk")
	user, exists := users[id]
	if !exists {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
}
