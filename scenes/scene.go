package scenes

import (
	"github.com/rydrman/three.go/math3"
	"github.com/rydrman/three.go/objects"
)

type Scene struct {
	*objects.Object

	BackgroundColor *math3.Color

	AutoUpdate bool
}

func NewScene() *Scene {

	return &Scene{
		Object: objects.NewObject(),

		BackgroundColor: math3.Colors("Black"),

		AutoUpdate: true,
	}

}
