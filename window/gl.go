package window

import "C"

import (
	"image/color"

	gl "github.com/chsc/gogl/gl43"
)

// sfmlColor returns a gl color based on the provided Go color.Color.
func glColor(c color.Color) (gl.Float, gl.Float, gl.Float, gl.Float) {
	r, g, b, a := c.RGBA()
	return gl.Float(r), gl.Float(g), gl.Float(b), gl.Float(a)
}
