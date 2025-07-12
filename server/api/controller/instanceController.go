package controller

import (
	"Clean_Architecture/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InstanceController struct {
	InstanceUseCase domain.InstanceUseCase
}

func (tc *InstanceController) Create(c *gin.Context) {
	var newInstance domain.Instance

	err := c.ShouldBind(&newInstance)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")

	newInstance.ID = uuid.New().String()
	newInstance.UserID = userID

	err = tc.InstanceUseCase.Create(c, &newInstance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Instance created successfully",
	})
}

func (tc *InstanceController) FetchByID(c *gin.Context) {
	id := c.Param("id")

	instance, err := tc.InstanceUseCase.FetchByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, instance)
}

func (tc *InstanceController) FetchByUserID(c *gin.Context) {
	userID := c.GetString("x-user-id")

	instances, err := tc.InstanceUseCase.FetchByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, instances)
}
