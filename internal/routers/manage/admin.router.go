package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (ar *AdminRouter) InitAdminRouter(router *gin.RouterGroup) {
	// public routes
	adminRouterPublic := router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
	// private routes
	adminRouterPrivate := router.Group("/admin")
	{
		adminRouterPrivate.POST("/active_user")
	}
}
