package controller

import (
	"Clean_Architecture/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PodcastController struct {
	PodcastUseCase domain.PodcastUseCase
}

func (tc *PodcastController) Create(c *gin.Context) {
	var newPodcast domain.Podcast

	err := c.ShouldBind(&newPodcast)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	newPodcast.ID = uuid.New().String()

	err = tc.PodcastUseCase.Create(c, &newPodcast)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Podcast created successfully",
	})
}

func (tc *PodcastController) FetchByID(c *gin.Context) {
	id := c.Param("id")

	podcast, err := tc.PodcastUseCase.FetchByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, podcast)
}
