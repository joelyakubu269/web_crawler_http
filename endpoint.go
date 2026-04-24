package main

import (
	"encoding/json"
	"net/http"
	"time"
)

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

func pstHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post method allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Method == http.MethodPost {
		var post = Post{}
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, "Invalid json request", http.StatusBadRequest)
			return
		}
		post.ID = len(posts) + 1
		posts = append(posts, post)

	}
}
