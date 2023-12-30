package handler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	GetBookById(c *gin.Context)
	GetAllBooks(c *gin.Context)
	AddBook(c *gin.Context)
	UpdateBook(c *gin.Context)
	DeleteBookById(c *gin.Context)
}
