package routers

import (
	"src/internal/routers/manage"
	"src/internal/routers/users"
)

type RouterGroup struct {
	User   users.UserRouterGroup
	Manage manage.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
