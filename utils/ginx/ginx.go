package ginx

import (
	"main/utils/ginx/context"
)

var insInput = context.Input{}
var insOutput = context.Output{}

func Input() *context.Input {
	return &insInput
}

func Output() *context.Output {
	return &insOutput
}
