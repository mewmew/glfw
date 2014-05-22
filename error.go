package glfw

// #cgo LDFLAGS: -lglfw
// #include <GLFW/glfw3.h>
//
// void init(void);
import "C"

import (
	"errors"
)

// LastError is a channel which holds the last GLFW error value.
var LastError = make(chan error, 1)

//export onError
// onError is the callback function which handles all GLFW errors. It does so by
// sending incomming error messages to the LastError channel.
func onError(code C.int, desc *C.char) {
	LastError <- errors.New(C.GoString(desc))
}

// init initializes the error handling callback.
func init() {
	C.init()
}
