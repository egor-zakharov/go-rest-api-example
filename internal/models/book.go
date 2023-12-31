package models

import "errors"

var ErrorNegativeId = errors.New("id must be more than 1")
var ErrorIncorrectId = errors.New("id must be absent")
var ErrorIncorrectTitle = errors.New("title must be absent")
var ErrorIncorrectAuthor = errors.New("auhtor must be absent")
var ErrorIncorrectYear = errors.New("releasedYear must be in the range '1901' to '2155'")

type Book struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	ReleasedYear int64  `json:"releasedYear"`
}

func (b *Book) IsValidReleasedYear() bool {
	return b.ReleasedYear >= 1901 && b.ReleasedYear <= 2155
}

func (b *Book) IsDraft() bool {
	return b.Id == 0
}

func (b *Book) IsValidId() bool {
	return b.Id > 0
}

func (b *Book) IsValidTitle() bool {
	return b.Title != ""
}

func (b *Book) IsValidAuthor() bool {
	return b.Author != ""
}
