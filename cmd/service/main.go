package main

import (
	"database/sql"
	"fmt"
	"library/config"
	"library/internal/handler"

	"library/internal/service"
	"library/internal/storage"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	cfg := config.New()
	//DSN full form username:password@protocol(address)/dbname?param=value
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(:%s)/%s", cfg.DBUserName, cfg.DBPassword, cfg.DBPort, cfg.DBDatabaseName))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	h := handler.NewHandler(service.New(storage.New(db)))

	router.GET("/book", h.GetAllBooks)
	router.GET("/book/:id", h.GetBookById)
	router.POST("/book/", h.AddBook)
	router.PUT("/book/:id", h.UpdateBook)
	router.DELETE("/book/:id", h.DeleteBookById)
	router.Run(fmt.Sprintf(":%d", cfg.Port))

}
