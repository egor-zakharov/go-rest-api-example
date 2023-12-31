package storage

import (
	"errors"
	"library/internal/models"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFindSuccess(t *testing.T) {
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
	got, _ := s.Find(int64(1))
	if want != *got {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFindNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := New(db)

	mock.ExpectQuery("select (.+) from books").
		WithArgs(int64(1)).
		WillReturnError(ErrorNotFound)
	want := ErrorNotFound
	_, got := s.Find(int64(1))
	if want != got {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFindAllSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := New(db)

	rows := sqlmock.NewRows([]string{"id", "title", "author", "released_year"}).
		AddRow(1, "title", "author", 2020).
		AddRow(2, "title_2", "author_2", 2021)

	mock.ExpectQuery("select (.+) from books order by id").
		WillReturnRows(rows)

	want := []models.Book{
		{Id: 1, Title: "title", Author: "author", ReleasedYear: 2020},
		{Id: 2, Title: "title_2", Author: "author_2", ReleasedYear: 2021},
	}
	got, _ := s.FindAll()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFindAllError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := New(db)

	err = errors.New("some error")
	want := err

	mock.ExpectQuery("select (.+) from books order by id").
		WillReturnError(err)

	_, got := s.FindAll()
	if want != got {
		t.Errorf("got %v want %v", got, want)
	}
}
