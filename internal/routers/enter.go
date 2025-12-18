package routers

import (
	"go-ecommerce-backend-api/internal/routers/manage"
	"go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.UserRouterGroup
}

// var RouterGroupApp *RouterGroup
// RouterGroupApp = &RouterGroup{}
var RouterGroupApp = new(RouterGroup)
