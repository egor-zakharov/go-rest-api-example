package service

import (
	"library/internal/models"
	"library/internal/storage"
)

type service struct {
	storage storage.Storage
}

func New(storage storage.Storage) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) GetBookById(bookId int64) (*models.Book, error) {
	if bookId < 1 {
		return nil, ErrorNegativeId
	}
	return s.storage.GetBookById(bookId)
}

func (s *service) GetAllBooks() ([]models.Book, error) {
	return s.storage.GetAllBooks()
}

func (s *service) AddBook(in models.Book) (*models.Book, error) {
	//TODO required fields
	if in.Id != 0 {
		return nil, ErrorIncorrectId
	}
	if in.ReleasedYear < 1901 || in.ReleasedYear > 2155 {
		return nil, ErrorYear
	}
	return s.storage.AddBook(in)
}

func (s *service) UpdateBook(in models.Book) (*models.Book, error) {
	if in.Id < 1 {
		return nil, ErrorNegativeId
	}
	return s.storage.UpdateBook(in)
}

func (s *service) DeleteBookById(bookId int64) error {
	return s.storage.DeleteBookById(bookId)
}

//
