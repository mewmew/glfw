WIP
---

This project is a *work in progress*. The implementation is *incomplete* and
subject to change. The documentation can be inaccurate.

win
===

This package provides a simplified Go binding for [GLFW 3][glfw]. Channels are
used instead of callbacks for event handling.

For the sake of simplicity this package only allows the use of one window. Each
event type has it's own dedicated channel and clients must register which events
they are interested in by calling the corresponding Enable* functions.


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

![Screenshot - simple](https://raw.github.com/mewmew/glfw/master/examples/simple/simple.png)

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
