package win

import (
	"fmt"
	"image"
	_ "image/jpeg" // decode jpeg images.
	_ "image/png"  // decode png images.
	"io"
	"log"
	"os"

	gl "github.com/chsc/gogl/gl21"
)

// Image is an image dedicated for OpenGL operations.
type Image struct {
	// Width of the image in pixels.
	Width int
	// Height of the image in pixels.
	Height int
	// id is the OpenGL texture id of the image.
	id gl.Uint
}

// Close closes the image and deletes it's OpenGL texture.
func (img *Image) Close() {
	gl.DeleteTextures(1, &img.id)
}

// curTexID represent the current texture id.
var curTexID = gl.Uint(1<<32 - 1)

// bind binds the provided texture id. BindTexture is an expensive operation and
// therefor we only bind if the texture id doesn't match the current one.
func bind(id gl.Uint) {
	if curTexID == id {
		return
	}
	curTexID = id
	gl.BindTexture(gl.TEXTURE_2D, id)
}

// Draw draws the entire src image onto the screen starting at the destination
// point dp.
func (src *Image) Draw(dp image.Point) {
	r := image.Rect(dp.X, dp.Y, dp.X+src.Width, dp.Y+src.Height)

	bind(src.id)

	gl.Begin(gl.QUADS)

	// top left
	gl.TexCoord2i(0, 0)
	gl.Vertex2i(gl.Int(r.Min.X), gl.Int(r.Min.Y))

	// top right
	gl.TexCoord2i(1, 0)
	gl.Vertex2i(gl.Int(r.Max.X), gl.Int(r.Min.Y))

	// bottom right
	gl.TexCoord2i(1, 1)
	gl.Vertex2i(gl.Int(r.Max.X), gl.Int(r.Max.Y))

	// bottom left
	gl.TexCoord2i(0, 1)
	gl.Vertex2i(gl.Int(r.Min.X), gl.Int(r.Max.Y))

	gl.End()
}

// DrawRect fills the destination rectangle dr of the screen with corresponding
// pixels from the src image starting at the source point sp.
func (src *Image) DrawRect(dr image.Rectangle, sp image.Point) {
	bind(src.id)

	xmin := gl.Float(sp.X) / gl.Float(src.Width)
	ymin := gl.Float(sp.Y) / gl.Float(src.Height)
	xmax := gl.Float(sp.X+dr.Dx()) / gl.Float(src.Width)
	ymax := gl.Float(sp.Y+dr.Dx()) / gl.Float(src.Height)

	gl.Begin(gl.QUADS)

	// top left
	gl.TexCoord2f(xmin, ymin)
	gl.Vertex2i(gl.Int(dr.Min.X), gl.Int(dr.Min.Y))

	// top right
	gl.TexCoord2f(xmax, ymin)
	gl.Vertex2i(gl.Int(dr.Max.X), gl.Int(dr.Min.Y))

	// bottom right
	gl.TexCoord2f(xmax, ymax)
	gl.Vertex2i(gl.Int(dr.Max.X), gl.Int(dr.Max.Y))

	// bottom left
	gl.TexCoord2f(xmin, ymax)
	gl.Vertex2i(gl.Int(dr.Min.X), gl.Int(dr.Max.Y))

	gl.End()
}

// OpenImage opens the provided image and returns a parsed OpenGL texture. The
// client must close the image when finished with it.
func OpenImage(filePath string) (img *Image, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadImage(f)
}

// ReadImage reads from r and returns a parsed OpenGL texture. The client must
// close the image when finished with it.
func ReadImage(r io.Reader) (img *Image, err error) {
	// Decode the image to an RGBA image.
	i, _, err := image.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("win.ReadImage: unable to decode image; %s", err)
	}
	rgba, ok := i.(*image.RGBA)
	if !ok {
		log.Printf("win.ReadImage: no fast path implemented for image type %T.\n", i)
		rgba = convertImage(i)
	}

	// load texture.
	bounds := i.Bounds()
	img = &Image{
		Width:  bounds.Dx(),
		Height: bounds.Dy(),
	}
	gl.GenTextures(1, &img.id)
	bind(img.id)

	// This library will mostly be used to handle images that doesn't require
	// scaling and thus uses gl.NEAREST instead of gl.LINEAR.

	// Texture minifying function:
	//
	// Returns the value of the texture element that is nearest (in Manhattan
	// distance) to the specified texture coordinates.
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)

	// Texture magnification function:
	//
	// Returns the value of the texture element that is nearest (in Manhattan
	// distance) to the specified texture coordinates.
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	gl.TexImage2D(gl.TEXTURE_2D, 0, 4 /* rgba */, gl.Sizei(img.Width), gl.Sizei(img.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Pointer(&rgba.Pix[0]))

	return img, nil
}

// convertImage converts the provided image to an RGBA image.
func convertImage(img image.Image) (rgba *image.RGBA) {
	i, ok := img.(*image.RGBA)
	if ok {
		return i
	}

	rect := img.Bounds()
	rgba = image.NewRGBA(rect)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	return rgba
}
