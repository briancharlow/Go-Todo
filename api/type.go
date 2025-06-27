package api

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Done      bool   `json:"done"`
	Priority  int    `json:"priority"`
	CreatedAt string `json:"createdAt"`
}
