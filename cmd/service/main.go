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

	//убрать после ревью
	//_ "github.com/egor-zakharov/go-rest-api-example/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title  Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @BasePath /api/v1
func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	cfg := config.New()
	//DSN full form username:password@protocol(address)/dbname?param=value
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(:%s)/%s", cfg.DBUserName, cfg.DBPassword, cfg.DBPort, cfg.DBDatabaseName))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	h := handler.NewHandler(service.New(storage.New(db)))
	v1 := router.Group("api/v1")
	{
		v1.GET("/book", h.GetAllBooks)
		v1.GET("/book/:id", h.GetBookById)
		v1.POST("/book/", h.AddBook)
		v1.PUT("/book/:id", h.UpdateBook)
		v1.DELETE("/book/:id", h.DeleteBookById)
	}

	router.Run(fmt.Sprintf(":%d", cfg.Port))

}
