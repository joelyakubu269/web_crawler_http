package main

import (
	"net/http"
)

var users = map[string]User{
	"1":{
		ID: "1",
		Name: "john",
		Mentions:{
			"You were mentioned in a post",
			"Someone tagged you in a comment",
		},
	}
}

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
