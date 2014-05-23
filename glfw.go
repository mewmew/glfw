// Package glfw encapsulates low-level details of the GLFW library. Its main
// purpose is to keep the API of other packages clean.
package glfw

// #cgo LDFLAGS: -lglfw
// #include <GLFW/glfw3.h>
import "C"

import (
	"fmt"
)

// clients represent the number of clients using the GLFW library, which is
// tracked by the number of calls to Init. If a call to Quit is made when no
// other clients are using the library, it will terminate and deallocate
// resources dedicated to GLFW.
var clients int

// Init initializes the GLFW library. It tracks the number of active clients so
// that a call to Quit only terminates the GLFW library if no other clients are
// using it.
//
// Note: This function may only be called from the main thread.
//
// Note: The Quit function must be called when finished using the library.
func Init() error {
	if clients == 0 {
		// Initialize the GLFW library if there are no active clients.
		if C.glfwInit() == C.GL_FALSE {
			return fmt.Errorf("glfw.Init: %v", LastError())
		}
	}
	clients++
	return nil
}

// Quit destroys all remaining windows and frees any allocated resources of the
// library. It does so only if there are no other active clients of the GLFW
// library.
//
// Note: This function may only be called from the main thread.
func Quit() {
	clients--
	if clients == 0 {
		// Terminate the GLFW library if there are no active clients.
		C.glfwTerminate()
	}
}
