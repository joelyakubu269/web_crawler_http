package main
import(
	"time"
)
type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Mentions []string `json:"mentions"`
}
type Post struct {
	ID  int  `json:"id"`
	UserID int `json:"user_id"`
	Title string  `json:"title"`
	Content string `json:"content"`
	Tags  []string `json:"tags"`
	CreatedAt  time.Time  `json:"created_at"`
}