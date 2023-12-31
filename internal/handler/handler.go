package handler

import (
	"fmt"
	"library/internal/models"
	"library/internal/service"
	u "library/internal/utils"
	"net/http"

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
	result, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Result: result})
}

// GetBookById godoc
// @Summary Get book
// @Produce json
// @Param id path int true "Book Id"
// @Success 200 {object} models.Book
// @Router /book/{id} [get]
func (h *restHanlder) GetBookById(c *gin.Context) {
	id, err := u.GetInt64FromContext(c, "id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	result, err := h.service.Find(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Result: result})
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
		return
	}
	result, err := h.service.Add(*param)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Result: result})
}

// UpdateBook godoc
// @Summary Update book
// @Produce json
// @Param id path int true "Book Id"
// @Param book body models.Book true "Update book"
// @Success 200 {object} models.Book
// @Router /book/{id} [put]
func (h *restHanlder) UpdateBook(c *gin.Context) {
	id, err := u.GetInt64FromContext(c, "id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	param := &models.Book{}
	err = c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	param.Id = id
	result, err := h.service.Update(*param)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Result: result})
}

// DeleteBookById godoc
// @Summary Delete book
// @Produce json
// @Param id path int true "Book Id"
// @Success 200
// @Router /book/{id} [delete]
func (h *restHanlder) DeleteBookById(c *gin.Context) {
	id, err := u.GetInt64FromContext(c, "id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	err = h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return

	}
	c.JSON(http.StatusOK, SuccessResponse{Result: fmt.Sprintf("Book: %d. successfully deleted", id)})
}
