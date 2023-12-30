package service

import (
	"library/internal/models"
	"library/internal/storage"
	"library/internal/utils"
)

type service struct {
	storage storage.Storage
}

func New(storage storage.Storage) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) Find(bookId int64) (*models.Book, error) {
	return s.storage.Find(bookId)
}

func (s *service) FindAll() ([]models.Book, error) {
	return s.storage.FindAll()
}

func (s *service) Add(in models.Book) (*models.Book, error) {
	err := utils.ValidateAddableBook(in)
	if err != nil {
		return nil, err
	}
	return s.storage.Add(in)
}

func (s *service) Update(in models.Book) (*models.Book, error) {
	err := utils.ValidateUpdatableBook(in)
	if err != nil {
		return nil, err
	}
	return s.storage.Update(in)
}

func (s *service) Delete(bookId int64) error {
	return s.storage.Delete(bookId)
}
