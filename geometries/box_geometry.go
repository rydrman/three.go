/*
 * Proted from three.js by @rydrman
 */

package geometries

import (
	"github.com/rydrman/three.go/core"
	"github.com/rydrman/three.go/math3"
)

type BoxGeometry struct {
	*core.Geometry
	Type string
}

func NewBoxGeometry(width, height, depth float32, widthSegments, heightSegments, depthSegments int) *BoxGeometry {

	geo := &BoxGeometry{
		Type: "BoxGeometry",
	}

	/*geo.Parameters = map[string]float{
		"width":          width,
		"height":         height,
		"depth":          depth,
		"widthSegments":  widthSegments,
		"heightSegments": heightSegments,
		"depthSegments":  depthSegments,
	}*/

	geo.FromBufferGeometry(
		NewBoxBufferGeometry(
			width,
			height,
			depth,
			widthSegments,
			heightSegments,
			depthSegments,
		),
	)
	geo.MergeVertices()

}

type BoxBufferGeometry struct {
	*core.BufferGeometry
}

func NewBoxBufferGeometry(width, height, depth float32, widthSegments, heightSegments, depthSegments int) *BoxBufferGeometry {

	geo := &BoxBufferGeometry{
		Type: "BoXbufferGeometry",
	}

	/*this.parameters = {
		width: width,
		height: height,
		depth: depth,
		widthSegments: widthSegments,
		heightSegments: heightSegments,
		depthSegments: depthSegments
	}
	*/

	// segments

	if widthSegments == 0 {
		widthSegments = 1
	}
	if heightSegments == 0 {
		heightSegments = 1
	}
	if depthSegments == 0 {
		depthSegments = 1
	}

	// buffers

	var (
		indices  []int
		vertices []*math3.Vector3
		normals  []*math3.Vector3
		uvs      []*math3.Vector2
	)

	// helper variables

	var (
		numberOfVertices = 0
		groupStart       = 0
	)

	// build each side of the box geometry

	geo.BuildPlane("z", "y", "x", -1, -1, depth, height, width, depthSegments, heightSegments, 0)  // px
	geo.BuildPlane("z", "y", "x", 1, -1, depth, height, -width, depthSegments, heightSegments, 1)  // nx
	geo.BuildPlane("x", "z", "y", 1, 1, width, depth, height, widthSegments, depthSegments, 2)     // py
	geo.BuildPlane("x", "z", "y", 1, -1, width, depth, -height, widthSegments, depthSegments, 3)   // ny
	geo.BuildPlane("x", "y", "z", 1, -1, width, height, depth, widthSegments, heightSegments, 4)   // pz
	geo.BuildPlane("x", "y", "z", -1, -1, width, height, -depth, widthSegments, heightSegments, 5) // nz

	// build geometry

	this.setIndex(indices)
	this.addAttribute("position", NewFloat32BufferAttribute(vertices, 3))
	this.addAttribute("normal", NewFloat32BufferAttribute(normals, 3))
	this.addAttribute("uv", NewFloat32BufferAttribute(uvs, 2))

}

func (geo *BoxBufferGeometry) buildPlane(u, v, w, udir, vdir, width, height, depth float32, gridX, gridY, materialIndex int) {

	var (
		segmentWidth  = width / gridX
		segmentHeight = height / gridY

		widthHalf  = width / 2.0
		heightHalf = height / 2.0
		depthHalf  = depth / 2.0

		gridX1 = gridX + 1
		gridY1 = gridY + 1

		vertexCounter = 0
		groupCount    = 0

		ix, iy int

		vector = math3.NewVector3()
	)

	// generate vertices, normals and uvs

	for iy = 0; iy < gridY1; iy++ {

		var y = iy*segmentHeight - heightHalf

		for ix = 0; ix < gridX1; ix++ {

			var x = ix*segmentWidth - widthHalf

			// set values to correct vector component

			vector[u] = x * udir
			vector[v] = y * vdir
			vector[w] = depthHalf

			// now apply vector to vertex buffer

			vertices.push(vector.x, vector.y, vector.z)

			// set values to correct vector component

			vector[u] = 0
			vector[v] = 0
			if depth > 0 {
				vector[w] = 1
			} else {
				vector[w] = -1
			}

			// now apply vector to normal buffer

			normals.push(vector.x, vector.y, vector.z)

			// uvs

			uvs.push(ix / gridX)
			uvs.push(1 - (iy / gridY))

			// counters

			vertexCounter++

		}

	}

	// indices

	// 1. you need three indices to draw a single face
	// 2. a single segment consists of two faces
	// 3. so we need to generate six (2*3) indices per segment

	for iy = 0; iy < gridY; iy++ {

		for ix = 0; ix < gridX; ix++ {

			var a = numberOfVertices + ix + gridX1*iy
			var b = numberOfVertices + ix + gridX1*(iy+1)
			var c = numberOfVertices + (ix + 1) + gridX1*(iy+1)
			var d = numberOfVertices + (ix + 1) + gridX1*iy

			// faces

			indices.push(a, b, d)
			indices.push(b, c, d)

			// increase counter

			groupCount += 6

		}

	}

	// add a group to the geometry. this will ensure multi material support

	geo.AddGroup(groupStart, groupCount, materialIndex)

	// calculate Newstart value for groups

	groupStart += groupCount

	// update total number of vertices

	numberOfVertices += vertexCounter

}
