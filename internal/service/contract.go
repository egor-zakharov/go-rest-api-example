package service

import (
	"library/internal/models"
)

type Service interface {
	Find(bookId int64) (*models.Book, error)
	FindAll() ([]models.Book, error)
	Add(in models.Book) (*models.Book, error)
	Update(in models.Book) (*models.Book, error)
	Delete(bookdId int64) error
}
