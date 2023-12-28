package storage

import (
	"errors"
	"library/internal/models"
)

var ErrorNotFound = errors.New("Book not found")
var ErrorNothingToUpdate = errors.New("Nothing book to update")

type Storage interface {
	GetBookById(bookId int64) (*models.Book, error)
	GetAllBooks() ([]models.Book, error)
	AddBook(in models.Book) (*models.Book, error)
	UpdateBook(in models.Book) (*models.Book, error)
	DeleteBookById(bookId int64) error
}

//
