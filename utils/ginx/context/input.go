package context

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"sync"
)

type Input struct {
	Context *gin.Context
	//CruSession    session.Store
	pnames        []string
	pvalues       []string
	data          map[interface{}]interface{} // store some values in this context when calling context in filter or controller.
	dataLock      sync.RWMutex
	RequestBody   []byte
	RunMethod     string
	RunController reflect.Type
}

// Param returns router param by a given key.
func (input *Input) Param(key string) string {
	for i, v := range input.pnames {
		if v == key && i <= len(input.pvalues) {
			// we cannot use url.PathEscape(input.pvalues[i])
			// for example, if the value is /a/b
			// after url.PathEscape(input.pvalues[i]), the value is %2Fa%2Fb
			// However, the value is used in ControllerRegister.ServeHTTP
			// and split by "/", so function crash...
			return input.pvalues[i]
		}
	}
	return ""
}

// Query returns input data item string by a given string.
func (input *Input) Query(c *gin.Context, key string) string {
	input.Context = c
	if val := input.Param(key); val != "" {
		return val
	}
	if input.Context.Request.Form == nil {
		input.Context.Request.ParseForm()
	}
	return input.Context.Request.Form.Get(key)
}
