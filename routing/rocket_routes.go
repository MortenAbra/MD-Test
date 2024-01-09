package routing

import (
	"media-devoted/controllers"

	"github.com/gin-gonic/gin"
)

func AddRocketGroup(group *gin.RouterGroup) {

	rocketController := controllers.RocketControllerInstance()

	rocketGroup := group.Group("/rockets")
	{
		rocketGroup.GET("", rocketController.GetRockets)
		rocketGroup.GET("/:id", rocketController.GetRocket)
		rocketGroup.POST("", rocketController.AddRocket)
		rocketGroup.POST("/:id", rocketController.UpdateRocket)
	}
}
