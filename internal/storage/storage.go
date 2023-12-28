package storage

import (
	"database/sql"
	"errors"
	"library/internal/models"
)

type storage struct {
	db *sql.DB
}

func New(db *sql.DB) Storage {
	return &storage{
		db,
	}
}

func (s *storage) GetBookById(bookId int64) (*models.Book, error) {
	book := models.Book{}
	err := s.db.QueryRow("select * from books where id = ?", bookId).Scan(&book.Id, &book.Title, &book.Author, &book.ReleasedYear)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}
	return &book, nil
}

func (s *storage) GetAllBooks() ([]models.Book, error) {
	book := models.Book{}
	books := []models.Book{}
	rows, err := s.db.Query("select * from books")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.ReleasedYear)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *storage) AddBook(in models.Book) (*models.Book, error) {
	result, err := s.db.Exec("insert into books (title, author, released_year) values (?,?,?)", in.Title, in.Author, in.ReleasedYear)
	if err != nil {
		return nil, err
	}
	in.Id, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &in, err
}

func (s *storage) UpdateBook(in models.Book) (*models.Book, error) {
	result, err := s.db.Exec("update books set title = ?, author = ?, released_year =? where id = ?", in.Title, in.Author, in.ReleasedYear, in.Id)
	if err != nil {
		return nil, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, ErrorNothingToUpdate
	}
	return &in, nil
}

func (s *storage) DeleteBookById(bookId int64) error {
	result, err := s.db.Exec("delete from books where id =?", bookId)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrorNotFound
	}
	return nil
}
