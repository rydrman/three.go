# three.go
Port of the three.js library to golang

## Versions

This port is currently targeting three.js r85

## Status

This library is in early stages of development, with only some minor core libraries working.

Currently the main math and core types are defined on some level, allowing a simple open gl window to render a color. The current goal is to flesh out the [simple cube example](examples/simple_cube.go) so that the basic geometry rendering works.
