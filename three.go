/*
Package three is the main package for the three.go graphics library
for golang. It attempts to port the three.js architecture as accurately
as possible while following proper golang practices.
*/
package three

import (
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {

	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	if err := gl.Init(); err != nil {
		panic(err)
	}

}

func CharAt(s string, i int) string {

	return string([]rune(s)[i])

}
