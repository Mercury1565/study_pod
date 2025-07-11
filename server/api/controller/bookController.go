package controller

import (
	"Clean_Architecture/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookController struct {
	BookUseCase domain.BookUseCase
}

func (tc *BookController) Create(c *gin.Context) {
	var newBook domain.Book

	err := c.ShouldBind(&newBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")

	newBook.ID = uuid.New().String()
	newBook.UserID = userID

	err = tc.BookUseCase.Create(c, &newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Book created successfully",
	})
}

func (tc *BookController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	books, err := tc.BookUseCase.FetchByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
