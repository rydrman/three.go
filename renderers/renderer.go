package renderers

import (
	"github.com/rydrman/three.go/math3"
	"github.com/rydrman/three.go/scenes"
)

type Renderer interface {
	Render(*scenes.Scene, math3.Projector)
}
