package domain

type PostDTO struct {
	Title    string `json:"title"`
	Status   string `json:"status"`
	Content  string `json:"content"`
}
