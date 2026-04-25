package main

import (
	"encoding/json"
	"fmt"
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
		var post = Post{} // this is declared here so that it does not overwrite each time we run the program
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, "Invalid json request", http.StatusBadRequest)
			return
		}
		post.ID = len(posts) + 1
		posts = append(posts, post)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(posts)
		if err != nil {
			http.Error(w, "invalid struct", http.StatusBadRequest)
			return
		}
	}
}
func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "only put method is allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Method == http.MethodPut {
		var change Post
		err := json.NewDecoder(r.Body).Decode(&change)
		if err != nil {
			http.Error(w, "Invalid json request", http.StatusBadRequest)
			return
		}
		for i, r := range posts {

			if r.ID == change.ID {
				posts[i].Title = change.Title
				break

			}
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(posts)
		if err != nil {
			http.Error(w, "invalid struct", http.StatusBadRequest)
			return
		}
	}
}
func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "This handler handles only delete", http.StatusMethodNotAllowed)
		return
	}
	if r.Method == http.MethodDelete {
		var del Post
		err := json.NewDecoder(r.Body).Decode(&del)
		if err != nil {
			http.Error(w, "Invalid json format", http.StatusBadRequest)
			return
		}
		n := 0
		for i, r := range posts {
			if r.ID == del.ID {
				n = i

			}
			posts = append(posts[:n], posts[:n+1]...)
			break
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(posts)
		if err != nil {
			http.Error(w, "Invalid json format", http.StatusBadRequest)
			return
		}
	}
}
func main() {
	http.HandleFunc("/profile", handler)
	http.HandleFunc("/post", pstHandler)
	http.HandleFunc("/put", updateHandler)
	fmt.Println("server is up and running")
	http.ListenAndServe(":8080", nil)
}
