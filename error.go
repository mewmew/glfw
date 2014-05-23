package glfw

// #cgo LDFLAGS: -lglfw
// #include <GLFW/glfw3.h>
//
// void init(void);
import "C"

import (
	"errors"
)

// LastError returns the last error reported by the GLFW library.
func LastError() (err error) {
	select {
	case err = <-lastError:
	default:
		err = errors.New("glfw: unknown error")
	}
	return err
}

// lastError is a channel which holds the last GLFW error value.
var lastError = make(chan error, 1)

//export onError
// onError is the callback function which handles all GLFW errors. It does so by
// sending incomming error messages to the lastError channel.
func onError(code C.int, desc *C.char) {
	lastError <- errors.New(C.GoString(desc))
}

// init initializes the error handling callback.
func init() {
	C.init()
}
