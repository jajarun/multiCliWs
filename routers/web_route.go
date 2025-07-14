package routers

import (
	"multicliws/controllers"

	"github.com/gin-gonic/gin"
)

func InitWebRouter(r *gin.Engine) {
	forwardRoute := r.Group("/mocksocket")
	{
		forwardRoute.GET("/forward", controllers.Forward)
	}
}
