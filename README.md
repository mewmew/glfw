WIP
---

This project is a *work in progress*. The implementation is *incomplete* and
subject to change. The documentation can be inaccurate.

win
===

Package win handles window creation, drawing and events. The window events are
defined in a dedicated package located at:
	github.com/mewmew/we

The library uses a small subset of the features provided by [GLFW 3][glfw]. For
the sake of simplicity support for multiple windows has intentionally been left
out.

Channels are used instead of callbacks for event handling. Each event type has
its own dedicated channel and clients must register which events they are
interested in by calling the corresponding Enable*  functions.

All calls to this package must originate from the same dedicated OS thread.
Use runtime.LockOSThread to achieve this.

[glfw]: https://github.com/glfw/glfw/

Documentation
-------------

Documentation provided by GoDoc.

- glfw
   - [win][glfw/win]: handles window creation, drawing and events.

[glfw/win]: http://godoc.org/github.com/mewmew/glfw/win

Installation
------------

Install the [GLFW 3][glfw] library and run:

	go get github.com/mewmew/glfw/win

Examples
--------

simple is an example which opens a new window, hooks the close events and
displays two images.

	go get github.com/mewmew/win/examples/simple

![Screenshot - simple](https://raw.githubusercontent.com/mewmew/glfw/master/examples/simple/simple.png)

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
