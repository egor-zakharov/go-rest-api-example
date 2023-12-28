package main

import (
	"database/sql"
	"fmt"
	"library/config"
	"library/internal/models"
	"library/internal/storage"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.New()
	//DSN full form username:password@protocol(address)/dbname?param=value
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(:%s)/%s", cfg.DBUserName, cfg.DBPassword, cfg.DBPort, cfg.DBDatabaseName))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookStorage := storage.New(db)

	book, err := bookStorage.GetBookById(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("")
		fmt.Printf("%+v", book)
	}

	books, err := bookStorage.GetAllBooks()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", books)
	}

	newBook := models.Book{Title: "Lol", Author: "KEL", ReleasedYear: 2020}

	addedBook, err := bookStorage.AddBook(newBook)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", addedBook)
	}

	bookUpdate := models.Book{Id: 1, Title: "Lol", Author: "KEL1", ReleasedYear: 2020}
	updatedBook, err := bookStorage.UpdateBook(bookUpdate)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", updatedBook)
	}

	deletedBookId := int64(18)
	err = bookStorage.DeleteBookById(deletedBookId)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}

}
