package controller

import (
	"Clean_Architecture/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ChapterController struct {
	ChapterUseCase domain.ChapterUseCase
}

func (tc *ChapterController) Create(c *gin.Context) {
	var newChapter domain.Chapter

	err := c.ShouldBind(&newChapter)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	newChapter.ID = uuid.New().String()

	err = tc.ChapterUseCase.Create(c, &newChapter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Chapter created successfully",
	})
}

func (tc *ChapterController) FetchByID(c *gin.Context) {
	id := c.Param("id")

	chapter, err := tc.ChapterUseCase.FetchByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, chapter)
}
