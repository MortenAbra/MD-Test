package custom_error

import "github.com/gin-gonic/gin"

type RocketResponseError struct {
	Status int         `json:"-"`
	Msg    interface{} `json:"msg"`
	Err    string      `json:"err"`
}

func RocketError(ctx *gin.Context, err RocketResponseError) {
	ctx.JSON(err.Status, err)
}
