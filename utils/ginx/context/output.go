package context

import "github.com/gin-gonic/gin"

// Output does work for sending response header.
type Output struct {
	Context    *gin.Context
	Status     int
	EnableGzip bool
}

// SetStatus sets response status code.
// It writes response header directly.
func (output *Output) SetStatus(status int) {
	output.Status = status
}
