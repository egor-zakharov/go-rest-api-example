package storage

import (
	"errors"
	"library/internal/models"
)

var ErrorNotFound = errors.New("Book not found")
var ErrorNothingToUpdate = errors.New("Nothing book to update")

type Storage interface {
	Find(bookId int64) (*models.Book, error)
	FindAll() ([]models.Book, error)
	Add(in models.Book) (*models.Book, error)
	Update(in models.Book) (*models.Book, error)
	Delete(bookId int64) error
}
