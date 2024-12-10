package model

type Note struct {
	ID      int    `json:"id"`
	UserId  int    `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
