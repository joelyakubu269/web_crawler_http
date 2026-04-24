package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var users = map[string]User{
	"1": {
		ID:   "1",
		Name: "john",
		Mentions: []string{
			"You were mentioned in a post",
			"Someone tagged you in a comment",
		},
	},
	"2": {
		ID:   "1",
		Name: "bukowski",
		Mentions: []string{
			"you were mentioned in a post",
			"someone tagged you in a comment",
		},
	},
}
var posts = []Post{
	{
		ID:        1,
		UserID:    2,
		Title:     "My first day at school",
		Content:   "today was full of anxiety but it ended up being fun",
		Tags:      []string{"school", "first-day", "life"},
		CreatedAt: time.Now(),
	},
	{
		ID:        2,
		UserID:    3,
		Title:     "Its my convocation",
		Content:   "Today a journey i started 5 years ago comes to an end",
		Tags:      []string{"convocation", "last day in school", "proud engineer"},
		CreatedAt: time.Now(),
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	section := query.Get("sk")
	user, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	switch section {
	case "mentions":
		json.NewEncoder(w).Encode(user.Mentions)
	default:
		json.NewEncoder(w).Encode(user)
	}

}
func postHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	section := query.Get("sk")
	newId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "error in converting to integer", http.StatusBadRequest)
		return
	}
	var result Post
	found := false
	for _, t := range posts {
		if t.ID == newId {
			found = true
			result = t
			break

		}

	}
	if !found {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	switch section {
	case "post":
		json.NewEncoder(w).Encode(result)
	default:
		json.NewEncoder(w).Encode(posts)

	}

}
func main() {
	http.HandleFunc("/profile", handler)
	http.HandleFunc("/post", postHandler)
	fmt.Println("server is up and running")
	http.ListenAndServe(":8080", nil)
}
