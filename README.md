WIP
---

This project is a *work in progress*. The implementation is *incomplete* and
subject to change. The documentation can be inaccurate.

win
===

This package provides a simplified Go binding for [GLFW][] 3. Channels are used
instead of callbacks for event handling.

For the sake of simplicity this package only allows the use of one window. Each
event type has it's own dedicated channel and clients must register which events
they are interested in by calling the corresponding Enable* functions.


[GLFW]: https://github.com/glfw/glfw/

Examples
--------

simple is an example which opens a new window, hooks the close events and
displays two images.

	go get github.com/mewmew/win/examples/simple

![Screenshot - simple](https://github.com/mewmew/win/blob/master/examples/simple/simple.png?raw=true)

Documentation
-------------

Documentation provided by GoDoc.

   - [win][]: provides a simplified Go binding for GLFW 3.

[win]: http://godoc.org/github.com/mewmew/win

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
