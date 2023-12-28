package handler

import (
	"errors"
	"fmt"
	"library/internal/models"
	"library/internal/service"
	"library/internal/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type restHanlder struct {
	service service.Service
}

func NewHandler(service service.Service) RestHandler {
	return &restHanlder{
		service: service,
	}
}

// GetAllBooks godoc
// @Summary Get all books
// @Produce json
// @Success 200 {array} models.Book
// @Router /book/ [get]
func (h *restHanlder) GetAllBooks(c *gin.Context) {
	result, err := h.service.GetAllBooks()
	//Проверять err!=nil или err == nil?
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, SuccessResponse{Result: result})
	}
}

// GetBookById godoc
// @Summary Get book
// @Produce json
// @Param id path int true "Book Id"
// @Success 200 {object} models.Book
// @Router /book/{id} [get]
func (h *restHanlder) GetBookById(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)
	result, err := h.service.GetBookById(int64(id))
	//ОК ли тянуть ошибки из service и storage в handler?
	if errors.Is(err, service.ErrorNegativeId) {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	} else if errors.Is(err, storage.ErrorNotFound) {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, SuccessResponse{Result: result})
	}
}

// AddBook godoc
// @Summary Add book
// @Produce json
// @Param book body models.Book true "Add book"
// @Success 200 {object} models.Book
// @Router /book/ [post]
func (h *restHanlder) AddBook(c *gin.Context) {
	param := &models.Book{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	result, err := h.service.AddBook(*param)
	if err != nil {
		if errors.Is(err, service.ErrorIncorrectId) {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		} else if errors.Is(err, service.ErrorYear) {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, SuccessResponse{Result: result})
	}
}

// UpdateBook godoc
// @Summary Update book
// @Produce json
// @Param book body models.Book true "Update book"
// @Success 200 {object} models.Book
// @Router /book/ [put]
func (h *restHanlder) UpdateBook(c *gin.Context) {
	param := &models.Book{}
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	result, err := h.service.UpdateBook(*param)
	if err != nil {
		if errors.Is(err, storage.ErrorNothingToUpdate) {
			c.JSON(http.StatusNoContent, ErrorResponse{Message: err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, SuccessResponse{Result: result})
	}
}

// DeleteBookById godoc
// @Summary Delete book
// @Produce json
// @Param id path int true "Book Id"
// @Success 200
// @Router /book/{id} [delete]
func (h *restHanlder) DeleteBookById(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)
	err := h.service.DeleteBookById(int64(id))
	if err != nil {
		if errors.Is(err, storage.ErrorNotFound) {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, SuccessResponse{Result: fmt.Sprintf("Book:%d. successfully deleted", id)})
	}
}
