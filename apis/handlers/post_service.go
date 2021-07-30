package handlers

import (
	"context"
	"net/http"

	"github.com/OnlyAloner/api_gateway/apis/models"
	"github.com/OnlyAloner/api_gateway/proto/post_service"
	"github.com/gin-gonic/gin"
)

// @Router /handlers/post_service/create [POST]
// @Summary create
// @Description Create
// @Tags post_service
// @Accept json
// @Produce  json
// @Param post_service body models.CreateRequest true "post_service"
// @Success 200 {object} models.CreateResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *handlerV1) Create(c *gin.Context) {
	var (
		payload  post_service.CreateRequest
		response models.CreateResponse
	)
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Code:    http.StatusBadRequest,
			Message: "Error while binding json to struct",
			Reason:  err.Error(),
		})
		return
	}

	res, err := h.grpcClient.PostService().Create(
		context.Background(),
		&payload,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error while creating post",
			Reason:  err.Error(),
		})
		return
	}

	err = ProtoToStruct(&response, res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error while converting proto to struct",
			Reason:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
