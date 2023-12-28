package service

import (
	"errors"
	"library/internal/models"
)

var ErrorNegativeId = errors.New("Id must be more than 1")
var ErrorIncorrectId = errors.New("Id must be absent")
var ErrorYear = errors.New("Year must be in the range '1901' to '2155'")

type Service interface {
	GetBookById(bookId int64) (*models.Book, error)
	GetAllBooks() ([]models.Book, error)
	AddBook(in models.Book) (*models.Book, error)
	UpdateBook(in models.Book) (*models.Book, error)
	DeleteBookById(bookdId int64) error
}
