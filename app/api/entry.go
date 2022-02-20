package api

import (
	"main/app/api/backend"
	"main/app/api/frontend"
)

var insFrontend = frontend.Group{}
var insBackend = backend.Group{}

func Frontend() *frontend.Group {
	return &insFrontend
}

func Backend() *backend.Group {
	return &insBackend
}
