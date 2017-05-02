package main

import (
	"time"

	"github.com/rydrman/three.go/geometries"
	"github.com/rydrman/three.go/math3"
	"github.com/rydrman/three.go/renderers"
	"github.com/rydrman/three.go/scenes"
)

func main() {

	win, err := renderers.NewWindowRenderer(500, 500, "Cube Example")
	if nil != err {
		panic(err)
	}
	defer win.Destroy()

	scene := scenes.NewScene()
	scene.BackgroundColor = math3.Colors("powderblue")

	geo := geometries.NewBoxGeometry(1, 1, 1)
	/*mat := three.NewBasicMaterial(three.MaterialProps{
		Color: three.Color().FromHex(0x001111),
	})

	mesh := three.NewMesh(geo, mat)
	scene.Add(mesh)*/

	//cam := three.NewPerspectiveCamera(1)

	for !win.ShouldClose() {
		win.Render(scene, nil)
		time.Sleep(time.Millisecond * 10)
	}

}
