/**
 * Ported from three.js by @rydrman
 */
package core

import "github.com/rydrman/three.go/math3"

type Face3 struct {
	A *math3.Vector3
	B *math3.Vector3
	C *math3.Vector3

	Normal        *math3.Vector3
	VertexNormals []*math3.Vector3

	Color        *math3.Color
	VertexColors []*math3.Color

	MaterialIndex int
}

func NewFace3(a, b, c, normal *math3.Vector3, color *math3.Color, materialIndex int) *Face3 {

	return &Face3{
		A:             a,
		B:             b,
		C:             c,
		Normal:        normal,
		Color:         color,
		MaterialIndex: materialIndex,
	}

}

func (f *Face3) Clone() (clone *Face3) {

	clone.Copy(f)
	return

}

func (f *Face3) Copy(source *Face3) *Face3 {

	f.A = source.A
	f.B = source.B
	f.B = source.B

	f.Normal.Copy(source.Normal)
	f.Color.Copy(source.Color)

	f.MaterialIndex = source.MaterialIndex

	f.VertexNormals = make([]*math3.Vector3, len(source.VertexNormals))
	copy(source.VertexNormals, f.VertexNormals)

	f.VertexColors = make([]*math3.Color, len(source.VertexColors))
	copy(source.VertexColors, f.VertexColors)

	return f

}
