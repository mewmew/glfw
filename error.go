package win

import (
	"log"

	"github.com/go-gl/glfw3"
)

func init() {
	// register the error callback function.
	glfw3.SetErrorCallback(onError)
}

// onError is the error callback function.
func onError(code glfw3.ErrorCode, desc string) {
	log.Println(desc)
}
