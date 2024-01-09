package controllers

import (
	"fmt"
	"media-devoted/handlers"
	"media-devoted/types"
	"media-devoted/types/custom_error"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RocketController interface {
	GetRockets(ctx *gin.Context)
	GetRocket(ctx *gin.Context)
	AddRocket(ctx *gin.Context)
	UpdateRocket(ctx *gin.Context)
	DeleteRocket(ctx *gin.Context)
}

type RocketControllerImpl struct {
	RockerHandler handlers.RocketHandler
}

func RocketControllerInstance() RocketController {
	return &RocketControllerImpl{
		RockerHandler: handlers.RocketHandlerInstance(),
	}
}

func (c *RocketControllerImpl) GetRockets(ctx *gin.Context) {
	rockets, err := c.RockerHandler.GetRockets(ctx)
	if err != nil {
		custom_error.RocketError(ctx, custom_error.RocketResponseError{
			Err:    err.Error(),
			Msg:    "error occurred retrieving rockets",
			Status: http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, rockets)
}

func (c *RocketControllerImpl) GetRocket(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		custom_error.RocketError(ctx, custom_error.RocketResponseError{
			Msg:    "error occurred by wrong uuid",
			Status: http.StatusBadRequest,
		})
		return
	}

	rocket, err := c.RockerHandler.GetRocket(ctx, &id)
	if err != nil {
		custom_error.RocketError(ctx, custom_error.RocketResponseError{
			Err:    err.Error(),
			Msg:    fmt.Sprintf("error occurred retrieving rocket with id: %v", id),
			Status: http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusOK, &rocket)
}

func (c *RocketControllerImpl) AddRocket(ctx *gin.Context) {
	var rocket *types.Rocket
	if bindErr := ctx.BindJSON(&rocket); bindErr != nil {
		custom_error.RocketError(ctx, custom_error.RocketResponseError{
			Err:    bindErr.Error(),
			Msg:    "error occurred binding json",
			Status: http.StatusBadRequest,
		})
		return
	}

	err := c.RockerHandler.AddRocket(ctx, rocket)
	if err != nil {
		custom_error.RocketError(ctx, custom_error.RocketResponseError{
			Err:    err.Error(),
			Msg:    "error occurred creating rocket",
			Status: http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"res": "OK"})
}

func (c *RocketControllerImpl) UpdateRocket(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		custom_error.RocketError(ctx, custom_error.RocketResponseError{
			Msg:    "error occurred parsing uuid",
			Status: http.StatusBadRequest,
		})
		return
	}

	var rocket *types.Rocket
	if bindErr := ctx.BindJSON(&rocket); bindErr != nil {
		custom_error.RocketError(ctx, custom_error.RocketResponseError{
			Err:    bindErr.Error(),
			Msg:    "error occurred binding rocket json",
			Status: http.StatusBadRequest,
		})
		return
	}

	rocket.Id = id

	updateErr := c.RockerHandler.UpdateRocket(ctx, rocket)
	if updateErr != nil {
		custom_error.RocketError(ctx, custom_error.RocketResponseError{
			Err:    updateErr.Error(),
			Msg:    "error occurred updating rocket",
			Status: http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"res": "OK"})
}

func (c *RocketControllerImpl) DeleteRocket(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		custom_error.RocketError(ctx, custom_error.RocketResponseError{
			Msg:    "error occurred parsing uuid",
			Status: http.StatusBadRequest,
		})
		return
	}

	deleteErr := c.RockerHandler.DeleteRocket(ctx, id)
	if deleteErr != nil {
		custom_error.RocketError(ctx, custom_error.RocketResponseError{
			Msg:    "error occurred deleting rocket",
			Status: http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"res": "Rocket has been deleted"})
}
