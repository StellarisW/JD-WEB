package service

import (
	"main/app/internal/service/frontend"
)

var insFrontend = frontend.Group{}

func Frontend() *frontend.Group {
	return &insFrontend
}
