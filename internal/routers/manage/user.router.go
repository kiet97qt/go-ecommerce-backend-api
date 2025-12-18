package manage

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {

	// private routes
	userRouterPrivate := router.Group("/admin/user")
	{
		userRouterPrivate.POST("/active_user")
	}
}
