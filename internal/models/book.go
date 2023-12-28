package models

type Book struct {
	Id           int64  `db:"id" json:"id"`
	Title        string `db:"title" json:"title"`
	Author       string `db:"author" json:"author"`
	ReleasedYear int64  `db:"released_year" json:"releasedYear"`
}
