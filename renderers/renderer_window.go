package renderers

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/rydrman/three.go/math3"
	"github.com/rydrman/three.go/scenes"
)

type WindowRenderer struct {
	title  string
	window *glfw.Window
}

func NewWindowRenderer(w, h int, title string) (*WindowRenderer, error) {

	r := &WindowRenderer{
		title: title,
	}

	err := r.setupWindow(w, h)
	if nil != err {
		return nil, err
	}

	return r, nil

}

func (r *WindowRenderer) setupWindow(w, h int) error {

	if r.window != nil {

		r.window.Destroy()

	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	//glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	//glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	win, err := glfw.CreateWindow(w, h, r.title, nil, nil)
	if err != nil {
		return err
	}
	r.window = win

	return nil

}

func (r *WindowRenderer) Render(scene *scenes.Scene, camera math3.Projector) {

	if nil == r.window {
		return
	}

	r.window.MakeContextCurrent()

	if scene.BackgroundColor != nil {

		gl.ClearColor(scene.BackgroundColor.R32(), scene.BackgroundColor.G32(), scene.BackgroundColor.B32(), 1)

	}

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	r.window.SwapBuffers()

}

func (r *WindowRenderer) ShouldClose() bool {

	glfw.PollEvents()

	if r.window != nil {

		return r.window.ShouldClose()

	}

	return true
}

func (r *WindowRenderer) Destroy() {

	if r.window != nil {

		r.window.Destroy()

	}

}
