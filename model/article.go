package model

// Article model.
type Article struct {
	Model

	Title   string
	Content string
	Comment Comment
	IP      string
}
