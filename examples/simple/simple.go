package main

import (
	"image"
	"log"
	"runtime"
	"time"

	"github.com/mewmew/win"
)

func main() {
	err := simple()
	if err != nil {
		log.Fatalln(err)
	}
}

// fps corresponds to the number of frames per second that should be drawn.
const fps = 60

func simple() (err error) {
	// OpenGL requires a dedicated OS thread.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Open window with the specified dimensions.
	err = win.Open(480, 270)
	if err != nil {
		return err
	}
	defer win.Close()

	// Register that we are interested in receiving close events.
	win.EnableCloseChan()

	// http://www.publicdomainpictures.net/pictures/40000/velka/pyrotechnics.jpg
	imgA, err := win.OpenImage("a.png")
	if err != nil {
		return err
	}
	defer imgA.Close()

	imgB, err := win.OpenImage("b.png")
	if err != nil {
		return err
	}
	defer imgB.Close()

	c := time.Tick(time.Second / fps)
	var drawA bool
	for {
		if drawA {
			imgA.Draw(image.ZP)
			drawA = false
		} else {
			imgB.Draw(image.ZP)
			drawA = true
		}

		// Swap buffers to display all drawings since last screen update.
		win.SwapBuffers()

		select {
		case <-win.CloseChan:
			// handle close events.
			return nil
		case <-c:
			// very simple implementation to update 60 times per second.
		}
	}

	return nil
}
