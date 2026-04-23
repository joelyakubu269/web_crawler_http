package main

import (
	"encoding/json"
	"net/http"
)

var users = map[string]User{
	"1":{
		ID: "1",
		Name: "john",
		Mentions:[]string{
			"You were mentioned in a post",
			"Someone tagged you in a comment",
		},
	},
	"2":{
		ID: "1",
		Name: "bukowski",
		Mentions: []string{
			"you were mentioned in a post",
			"someone tagged you in a comment",
		},
	},
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
	w.Header().Set("Content-Type","application/json")
	switch section {
	case "mentions":
		json.NewEncoder(w).Encode(user.Mentions)
	default:
		json.NewEncoder(w).Encode(user)
	}
}
func main() {
	http.HandleFunc("/profile",handler)
	http.ListenAndServe("8080",nil)
}
