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

func (s *storage) Find(bookId int64) (*models.Book, error) {
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

func (s *storage) FindAll() ([]models.Book, error) {
	books := []models.Book{}
	rows, err := s.db.Query("select * from books order by id")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := models.Book{}
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.ReleasedYear)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *storage) Add(in models.Book) (*models.Book, error) {
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

func (s *storage) Update(in models.Book) (*models.Book, error) {
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

func (s *storage) Delete(bookId int64) error {
	_, err := s.db.Exec("delete from books where id =?", bookId)
	if err != nil {
		return err
	}
	return nil
}
