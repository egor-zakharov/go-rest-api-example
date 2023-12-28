package storage

import (
	"library/internal/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetBookByIdSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := New(db)

	row := sqlmock.NewRows([]string{"id", "title", "author", "released_year"}).AddRow(1, "title", "author", 2020)
	mock.ExpectQuery("select (.+) from books").
		WithArgs(int64(1)).
		WillReturnRows(row)
	want := models.Book{Id: 1, Title: "title", Author: "author", ReleasedYear: 2020}
	got, _ := s.GetBookById(int64(1))
	if want != *got {
		t.Errorf("got %v want %v", got, want)
	}

}
