package routers

import (
	controllers "go-restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItem", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItem/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItem-order/:orderItem_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItem", controllers.CreateOrderItem())
	incomingRoutes.PATCH("orderItem/:orderItem_id", controllers.UpdateOrderItem())
}
