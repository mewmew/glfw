package win

import (
	"image"

	"github.com/go-gl/glfw3"
	"github.com/mewmew/we"
)

// Event channels.
var (
	// CloseChan receives close events when the window is closed.
	CloseChan chan we.Close

	// ResizeChan receives resize events when the window is resized.
	ResizeChan chan we.Resize

	// KeyPressChan receives key press events when a keyboard key is pressed.
	KeyPressChan chan we.KeyPress
	// KeyReleaseChan receives key release events when a keyboard key is
	// released.
	KeyReleaseChan chan we.KeyRelease
	// KeyRepeatChan receives key repeat events when a keyboard key was held down
	// until it repeated.
	KeyRepeatChan chan we.KeyRepeat
	// KeyRuneChan receives key rune events when a unicode character is typed on
	// the keyboard. For instance if `a` and `shift` are held down on the
	// keyboard KeyRune will correspond to 'A'.
	KeyRuneChan chan we.KeyRune

	// MousePressChan receives mouse press events when a mouse button is pressed.
	MousePressChan chan we.MousePress
	// MouseReleaseChan receives mouse release events when a mouse button is
	// released.
	MouseReleaseChan chan we.MouseRelease
	// MouseMoveChan receives mouse move events when the mouse is moved from one
	// location to another.
	MouseMoveChan chan we.MouseMove
	// MouseDragChan receives mouse drag events when the mouse is moved from one
	// location to another while a mouse button is held down.
	MouseDragChan chan we.MouseDrag
	// MouseEnterChan receives mouse enter events when the mouse enters or leaves
	// the window. On mouse enter the value is true, otherwise it is false.
	MouseEnterChan chan we.MouseEnter

	// ScrollXChan receives scroll x events when the mouse wheel is scrolled on
	// the horizontal axis.
	ScrollXChan chan we.ScrollX
	// ScrollYChan receives scroll y events when the mouse wheel is scrolled on
	// the vertical axis.
	ScrollYChan chan we.ScrollY
)

// EnableCloseChan enables CloseChan. The client is responsible for receiving
// events on CloseChan, in order to prevent deadlocks.
func EnableCloseChan() {
	CloseChan = make(chan we.Close)
}

// EnableResizeChan enables ResizeChan. The client is responsible for receiving
// events on ResizeChan, in order to prevent deadlocks.
func EnableResizeChan() {
	ResizeChan = make(chan we.Resize)
}

// EnableKeyPressChan enables KeyPressChan. The client is responsible for
// receiving events on KeyPressChan, in order to prevent deadlocks.
func EnableKeyPressChan() {
	KeyPressChan = make(chan we.KeyPress)
}

// EnableKeyReleaseChan enables KeyReleaseChan. The client is responsible for
// receiving events on KeyReleaseChan, in order to prevent deadlocks.
func EnableKeyReleaseChan() {
	KeyReleaseChan = make(chan we.KeyRelease)
}

// EnableKeyRepeatChan enables KeyRepeatChan. The client is responsible for
// receiving events on KeyRepeatChan, in order to prevent deadlocks.
func EnableKeyRepeatChan() {
	KeyRepeatChan = make(chan we.KeyRepeat)
}

// EnableKeyRuneChan enables KeyRuneChan. The client is responsible for
// receiving events on KeyRuneChan, in order to prevent deadlocks.
func EnableKeyRuneChan() {
	KeyRuneChan = make(chan we.KeyRune)
}

// EnableMousePressChan enables MousePressChan. The client is responsible for
// receiving events on MousePressChan, in order to prevent deadlocks.
func EnableMousePressChan() {
	MousePressChan = make(chan we.MousePress)
}

// EnableMouseReleaseChan enables MouseReleaseChan. The client is responsible
// for receiving events on MouseReleaseChan, in order to prevent deadlocks.
func EnableMouseReleaseChan() {
	MouseReleaseChan = make(chan we.MouseRelease)
}

// EnableMouseMoveChan enables MouseMoveChan. The client is responsible for
// receiving events on MouseMoveChan, in order to prevent deadlocks.
func EnableMouseMoveChan() {
	MouseMoveChan = make(chan we.MouseMove)
}

// EnableMouseDragChan enables MouseDragChan. The client is responsible for
// receiving events on MouseDragChan, in order to prevent deadlocks.
func EnableMouseDragChan() {
	MouseDragChan = make(chan we.MouseDrag)
}

// EnableMouseEnterChan enables MouseEnterChan. The client is responsible for
// receiving events on MouseEnterChan, in order to prevent deadlocks.
func EnableMouseEnterChan() {
	MouseEnterChan = make(chan we.MouseEnter)
}

// EnableScrollXChan enables ScrollXChan. The client is responsible for
// receiving events on ScrollXChan, in order to prevent deadlocks.
func EnableScrollXChan() {
	ScrollXChan = make(chan we.ScrollX)
}

// EnableScrollYChan enables ScrollYChan. The client is responsible for
// receiving events on ScrollYChan, in order to prevent deadlocks.
func EnableScrollYChan() {
	ScrollYChan = make(chan we.ScrollY)
}

// hookEvents hooks all the event callback functions.
func hookEvents() {
	// close events.
	w.SetCloseCallback(onClose)

	// resize events.
	w.SetSizeCallback(onSize)

	// key events.
	w.SetKeyCallback(onKey)
	w.SetCharacterCallback(onChar)

	// mouse events.
	w.SetMouseButtonCallback(onMouseButton)
	w.SetCursorPositionCallback(onCursorPos)
	w.SetCursorEnterCallback(onCursorEnter)

	// scroll events.
	w.SetScrollCallback(onScroll)
}

// pollEvents will loop forever polling events.
func pollEvents() {
	for {
		glfw3.WaitEvents()
	}
}

// onClose is the close event callback function.
func onClose(w *glfw3.Window) {
	if CloseChan == nil {
		return
	}
	CloseChan <- we.Close{}
}

// onSize is the resize event callback function.
func onSize(w *glfw3.Window, width, height int) {
	if ResizeChan == nil {
		return
	}
	e := we.Resize{
		Width:  width,
		Height: height,
	}
	ResizeChan <- e
}

// onKey is the key event callback function.
func onKey(w *glfw3.Window, key glfw3.Key, scancode int, action glfw3.Action, mod glfw3.ModifierKey) {
	switch action {
	case glfw3.Press:
		if KeyPressChan == nil {
			return
		}
		e := we.KeyPress{
			Key: we.Key(key),
			Mod: we.Mod(mod),
		}
		KeyPressChan <- e
	case glfw3.Release:
		if KeyReleaseChan == nil {
			return
		}
		e := we.KeyRelease{
			Key: we.Key(key),
			Mod: we.Mod(mod),
		}
		KeyReleaseChan <- e
	case glfw3.Repeat:
		if KeyRepeatChan == nil {
			return
		}
		e := we.KeyRepeat{
			Key: we.Key(key),
			Mod: we.Mod(mod),
		}
		KeyRepeatChan <- e
	}
}

// onChar is the char event callback function.
func onChar(w *glfw3.Window, r uint) {
	if KeyRuneChan == nil {
		return
	}
	KeyRuneChan <- we.KeyRune(r)
}

// btn tracks mouse button states. If a mouse button is held down btn has the
// value of `button + 1`, otherwise it is 0.
var btn int

// onMouseButton is the mouse button event callback function.
func onMouseButton(w *glfw3.Window, button glfw3.MouseButton, action glfw3.Action, mod glfw3.ModifierKey) {
	switch action {
	case glfw3.Press:
		if MousePressChan == nil {
			return
		}
		e := we.MousePress{
			Point:  MousePos(),
			Button: we.Button(button),
			Mod:    we.Mod(mod),
		}
		MousePressChan <- e

		// store mouse button state which is used for possible mouse drag event.
		btn = int(button + 1)
	case glfw3.Release:
		if MouseReleaseChan == nil {
			return
		}
		e := we.MouseRelease{
			Point:  MousePos(),
			Button: we.Button(button),
			Mod:    we.Mod(mod),
		}
		MouseReleaseChan <- e

		// clear mouse button state.
		btn = 0
	}
}

// prevPt is the previous coordiates of the cursor.
var prevPt image.Point

// onCursorPos is the cursor position event callback function.
func onCursorPos(w *glfw3.Window, x float64, y float64) {
	// mouse move event.
	pt := image.Pt(int(x), int(y))
	if MouseMoveChan != nil {
		e := we.MouseMove{
			Point: pt,
			From:  prevPt,
		}
		MouseMoveChan <- e
	}

	// mouse drag event.
	if btn > 0 {
		if MouseDragChan != nil {
			e := we.MouseDrag{
				Point:  pt,
				From:   prevPt,
				Button: we.Button(btn - 1),
				Mod:    getMod(),
			}
			MouseDragChan <- e
		}
	}
	prevPt = pt
}

// onCursorEnter is the cursor event event callback function.
func onCursorEnter(w *glfw3.Window, entered bool) {
	if MouseEnterChan == nil {
		return
	}
	MouseEnterChan <- we.MouseEnter(entered)
}

// onScroll is the scroll event callback function.
func onScroll(w *glfw3.Window, xoff, yoff float64) {
	// TODO(u): Add cursor coordinates to the ScrollX and ScrollY events.
	if xoff != 0 {
		if ScrollXChan != nil {
			e := we.ScrollX{
				Off: int(xoff),
				Mod: getMod(),
			}
			ScrollXChan <- e
		}
	}
	if yoff != 0 {
		if ScrollYChan != nil {
			e := we.ScrollY{
				Off: int(yoff),
				Mod: getMod(),
			}
			ScrollYChan <- e
		}
	}
}

// getMod returns the keyboard modifiers that are held down.
func getMod() (mod we.Mod) {
	if isKeyDown(glfw3.KeyLeftShift) || isKeyDown(glfw3.KeyRightShift) {
		mod |= we.ModShift
	}
	if isKeyDown(glfw3.KeyLeftControl) || isKeyDown(glfw3.KeyRightControl) {
		mod |= we.ModControl
	}
	if isKeyDown(glfw3.KeyLeftAlt) || isKeyDown(glfw3.KeyRightAlt) {
		mod |= we.ModAlt
	}
	if isKeyDown(glfw3.KeyLeftSuper) || isKeyDown(glfw3.KeyRightSuper) {
		mod |= we.ModSuper
	}
	return mod
}

// isKeyDown returns true if the key is held down.
func isKeyDown(key glfw3.Key) bool {
	return w.GetKey(key) == glfw3.Press
}
