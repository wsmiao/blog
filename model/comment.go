package model

// Comment model.
type Comment struct {
	Model

	ArticleID uint64
	Content   string
	IP        string
}
