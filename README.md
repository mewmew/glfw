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

Documentation
-------------

Documentation provided by GoDoc.

   - [win][]

[win]: http://godoc.org/github.com/mewmew/win

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
