package model

type Checklist struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Items    []string `json:"items"`
	Complete bool     `json:"complete"`
}
