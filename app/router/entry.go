package router

import (
	"main/app/router/backend"
	"main/app/router/frontend"
)

type Group struct {
	Frontend frontend.RouterGroup
	Backend  backend.RouterGroup
}

var GroupApp = new(Group)
