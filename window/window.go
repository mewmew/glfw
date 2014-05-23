// Package window handles window creation, drawing and events. It uses a small
// subset of the features provided by the GLFW library version 3.0 [1].
//
// [1]: http://www.glfw.org/
package window

// #cgo LDFLAGS: -lglfw
// #include <GLFW/glfw3.h>
import "C"

import (
	"fmt"
	"image"
	"image/color"

	gl "github.com/chsc/gogl/gl43"
	"github.com/mewmew/glfw"
	"github.com/mewmew/wandi"
)

// TODO(u): Implement window styles.

// A Window represents a graphical window capable of handling draw operations
// and window events. It implements the wandi.Window interface.
type Window struct {
	// A window with an associated GPU context.
	win *C.GLFWwindow
}

// Open opens a new window of the specified dimensions.
//
// Note: This function may only be called from the main thread.
//
// Note: The Close method of the window must be called when finished using it.
func Open(width, height int) (win Window, err error) {
	// Initialize the GLFW library. It will only initialize if there are no other
	// active clients.
	err = glfw.Init()
	if err != nil {
		return Window{}, err
	}

	// Initialize OpenGL.
	gl.Init()

	// Create a new window of the specified dimensions.
	title := C.CString("untitled")
	win.win = C.glfwCreateWindow(C.int(width), C.int(height), title, nil, nil)
	if win.win == nil {
		return Window{}, fmt.Errorf("window.Open: %v", glfw.LastError())
	}

	return win, nil
}

// Close closes the window.
func (win Window) Close() {
	// Terminate the GLFW library. It will only terminate if there are no other
	// active clients.
	glfw.Quit()

	C.glfwDestroyWindow(win.win)
}

// SetTitle sets the title of the window.
//
// Note: The title will be updated on a future call to PollEvent.
func (win Window) SetTitle(title string) {
	C.glfwSetWindowTitle(win.win, C.CString(title))
}

// ShowCursor displays or hides the mouse cursor depending on the value of
// visible. It is visible by default.
func (win Window) ShowCursor(visible bool) {
	val := C.int(C.GLFW_CURSOR_HIDDEN)
	if visible {
		val = C.GLFW_CURSOR_NORMAL
	}
	C.glfwSetInputMode(win.win, C.GLFW_CURSOR, val)
}

// Width returns the width of the window.
func (win Window) Width() int {
	var width C.int
	C.glfwGetWindowSize(win.win, &width, nil)
	return int(width)
}

// Height returns the height of the window.
func (win Window) Height() int {
	var height C.int
	C.glfwGetWindowSize(win.win, nil, &height)
	return int(height)
}

// Draw draws the entire src image onto the window starting at the destination
// point dp.
func (win Window) Draw(dp image.Point, src wandi.Image) (err error) {
	sr := image.Rect(0, 0, src.Width(), src.Height())
	return win.DrawRect(dp, src, sr)
}

// DrawRect draws a subset of the src image, as defined by the source rectangle
// sr, onto the window starting at the destination point dp.
func (win Window) DrawRect(dp image.Point, src wandi.Image, sr image.Rectangle) (err error) {
	// Make the context of the window current. This operation is essentially a
	// nop if it's already current.
	C.glfwMakeContextCurrent(win.win)

	panic("not yet implemented")
}

// Fill fills the entire window with the provided color.
func (win Window) Fill(c color.Color) {
	// Make the context of the window current. This operation is essentially a
	// nop if it's already current.
	C.glfwMakeContextCurrent(win.win)

	// Fill the entire window with the color c.
	gl.ClearColor(glColor(c))
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

// Display displays what has been rendered so far to the window.
func (win Window) Display() {
	C.glfwSwapBuffers(win.win)
}
