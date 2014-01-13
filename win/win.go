// Package win handles window creation, drawing and events. The window events
// are defined in a dedicated package located at:
//    github.com/mewmew/we
//
// The library uses a small subset of the features provided by GLFW 3. For the
// sake of simplicity support for multiple windows has intentionally been left
// out.
//
// Channels are used instead of callbacks for event handling. Each event type
// has its own dedicated channel and clients must register which events they are
// interested in by calling the corresponding Enable*  functions.
//
// All calls to this package must originate from the same dedicated OS thread.
// Use runtime.LockOSThread to achieve this.
package win

import (
	"errors"
	"fmt"
	"image"

	gl "github.com/chsc/gogl/gl21"
	"github.com/go-gl/glfw3"
)

// w represents the window containing the OpenGL context. It is opened through a
// call to Open and it's this single window that is utilized throughout the
// entire library.
var w *glfw3.Window

// Open opens a window with the specified dimensions. Only one window can be
// open at the same time. It is this single window that is utilized throughout
// the entire library.
//
// Note: The Close function must be called when finished using the window.
func Open(width, height int) (err error) {
	if w != nil {
		panic("win.Open: the window has already been opened.")
	}

	if !glfw3.Init() {
		return errors.New("win.Open: glfw3.Init failed")
	}

	w, err = glfw3.CreateWindow(width, height, "untitled", nil, nil)
	if err != nil {
		return fmt.Errorf("win.Open: glfw3.CreateWindow failed; %v", err)
	}
	w.MakeContextCurrent()

	// Enable vsync.
	glfw3.SwapInterval(1)

	gl.Init()
	gl.Enable(gl.TEXTURE_2D)

	// Enable alpha channel for transparency.
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.BLEND)

	// left:   0
	// right:  width
	// bottom: hight
	// top:    0
	gl.Ortho(0, gl.Double(width), gl.Double(height), 0, -1, 1)

	// Disable depth testing and lighting.
	/// ### [ todo ] ###
	///    - benchmark if this enhances the performance.
	/// ### [/ todo ] ###
	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.LIGHTING)

	gl.Disable(gl.DITHER)

	hookEvents()

	go pollEvents()

	return nil
}

// Close closes the window.
func Close() {
	if w == nil {
		panic("win.Close: the window has already been closed.")
	}

	w.Destroy()
	w = nil
	glfw3.Terminate()
}

// SetTitle sets the title of the window.
func SetTitle(title string) {
	w.SetTitle(title)
}

// Size returns the size, in screen coordinates, of the client area of the
// window.
func Size() (width, height int) {
	return w.GetSize()
}

// SetSize sets the size, in screen coordinates, of the client area of the
// window.
func SetSize(width, height int) {
	w.SetSize(width, height)
}

// MousePos returns the last reported position of the cursor.
func MousePos() (pt image.Point) {
	x, y := w.GetCursorPosition()
	pt.X = int(x)
	pt.Y = int(y)
	return pt
}

// SetMousePos sets the position of the cursor. The window must be focused. If
// the window does not have focus when this function is called, it fails
// silently.
func SetMousePos(pt image.Point) {
	x := float64(pt.X)
	y := float64(pt.Y)
	w.SetCursorPosition(x, y)
}

// SwapBuffers swaps the front and back buffers of the window. The GPU driver
// will wait for 1 screen update before swapping the buffers, since vsync is
// enabled.
func SwapBuffers() {
	w.SwapBuffers()
}
