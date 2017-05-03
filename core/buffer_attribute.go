/*
 * Ported from three.js by @rydrman
 */

package core

import "github.com/rydrman/three.go/math3"

// BufferAttribute stores data for an attribute (such as vertex positions, face indices, normals, colors, UVs,
// and any custom attributes) associated with a BufferGeometry, which allows for more efficient passing of data
// to the GPU. See that page for details and a usage example.
//
// Data is stored as vectors of any length (defined by itemSize), and in general in the methods outlined below
// if passing in an index, this is automatically multiplied by the vector length.
type BufferAttribute struct {
	UUID string

	ItemSize   int
	Count      int
	Normalized bool

	Dynamic           bool
	UpdateRangeOffset int
	UpdateRangeCount  int

	array            math3.TypeArray
	needsUpdate      bool
	onUploadCallback func()

	Version int
}

func newBufferAttribute(array math3.TypeArray, itemSize int, normalized bool) *BufferAttribute {

	attr := &BufferAttribute{

		UUID: math3.GenerateUUID(),

		ItemSize:   itemSize,
		Normalized: normalized,

		Dynamic:           false,
		UpdateRangeOffset: 0,
		UpdateRangeCount:  -1,

		needsUpdate:      true,
		onUploadCallback: nil,

		Version: 0,
	}

	attr.SetArray(array)
	return attr

}

func (attr *BufferAttribute) NeedsUpdate() bool {

	return attr.needsUpdate

}

func (attr *BufferAttribute) SetNeedsUpdate(value bool) {

	if value {
		attr.Version++
	}
	attr.needsUpdate = value

}

func (attr *BufferAttribute) Array() math3.TypeArray {

	return attr.array

}

func (attr *BufferAttribute) SetArray(array math3.TypeArray) {

	attr.array = array

	switch array.(type) := arr {

		case Int8Array:
			attr.Count = len(arr) / itemSize
			break

		case Uint8Array:
			attr.Count = len(arr) / itemSize
			break

		case Int16Array:
			attr.Count = len(arr) / itemSize
			break

		case Uint16Array:
			attr.Count = len(arr) / itemSize
			break

		case Int32Array:
			attr.Count = len(arr) / itemSize
			break

		case Uint32Array:
			attr.Count = len(arr) / itemSize
			break

		case Float32Array:
			attr.Count = len(arr) / itemSize
			break

		case Float64Array:
			attr.Count = len(arr) / itemSize
			break

		default:
			glog.Fatalln("three/core/BufferAttribute.SetArray must be passed one of math3.Int8Array, math3.Uint8Array, " + 
				"math3.Int16Array, math3.Uint16Array, math3.Int32Array, math3.Uint32Array, " + 
				"math3.Float32Array, math3.Float64Array")

	}

}

func (attr *BufferAttribute) Copy(source *BufferAttribute) {

	attr.array = make([]interface{}, len(source.array))
	copy(source.array, attr.array)
	attr.ItemSize = source.ItemSize
	attr.Count = source.Count
	attr.Normalized = source.Normalized

	attr.Dynamic = source.Dynamic

}

func (attr *BufferAttribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) {

	index1 *= attr.ItemSize
	index2 *= attribute.itemSize

	for i := 0; i < attr.ItemSize; i++ {

		attr.array[index1+i] = attribute.array[index2+i]

	}

}

func (attr *BufferAttribute) CopyArray(array []interface{}) {

	attr.array = make([]interface{}, len(array))
	copy(array, attr.array)

}

func (attr *BufferAttribute) CopyColorsArray(colors []*math3.Color) {

	offset := 0

	for i := 0; i < colors.length; i++ {

		color := colors[i]

		if color == nil {

			glog.Warningf("three.core.BufferAttribute.CopyColorsArray(): color is nil %d\n", i)
			color = math3.NewColor()

		}

		array[offset] = color.R
		offset++
		array[offset] = color.G
		offset++
		array[offset] = color.B
		offset++

	}

}

/*func (attr *BufferAttribute) CopyIndicesArray(indices []*math3.Vector3) {

	offset := 0

	for i := 0; i < indices.length; i ++ {

		index := indices[i]

		array[offset] = index.A; offset++
		array[offset] = index.B; offset++
		array[offset] = index.C; offset++

	}

}*/

func (attr *BufferAttribute) CopyVector2sArray(vectors []*math3.Vector3) {

	offset := 0

	for i := 0; i < len(vectors); i++ {

		vector := vectors[i]

		if vector == nil {

			glog.Warningf("three.core.BufferAttribute.CopyVector2sArray(): vector is nil, %d\n", i)
			vector = math3.NewVector2()

		}

		array[offset] = vector.X
		offset++
		array[offset] = vector.Y
		offset++

	}

}

func (attr *BufferAttribute) CopyVector3sArray(vectors []*math3.Vector3) {

	offset := 0

	for i := 0; i < len(vectors); i++ {

		vector := vectors[i]

		if vector == nil {

			glog.Warningf("three.core.BufferAttribute.CopyVector3sArray(): vector is nil, %d\n", i)
			vector = math3.NewVector3()

		}

		array[offset] = vector.X
		offset++
		array[offset] = vector.Y
		offset++
		array[offset] = vector.Z
		offset++

	}

}

func (attr *BufferAttribute) CopyVector4sArray(vectors []*math3.Vector4) {

	offset := 0

	for i := 0; i < len(vectors); i++ {

		var vector = vectors[i]

		if vector == nil {

			glog.Warningf("three.core.BufferAttribute.CopyVector4sArray(): vector is nil, %d\n", i)
			vector = math3.NewVector4()

		}

		array[offset] = vector.X
		offset++
		array[offset] = vector.Y
		offset++
		array[offset] = vector.Z
		offset++
		array[offset] = vector.W
		offset++

	}

}

func (attr *BufferAttribute) Set(value []interface{}, offset int) {

	attr.array = make([]interface{}, len(value)-offset)

	for i := 0; i < len(value)-offset; i++ {

		attr.array[i] = value[i+offset]

	}

}

/*func (attr *BufferAttribute) GetX(index int) interface{} {

	return attr.array[index * attr.ItemSize]

}

func (attr *BufferAttribute) SetX(index int, x interface{}) {

	attr.array[index * attr.ItemSize] = x

}

func (attr *BufferAttribute) GetY(index int) interface{} {

	return attr.array[index * attr.ItemSize + 1]

}

func (attr *BufferAttribute) SetY(index int, y interface{}) {

	attr.array[index * attr.ItemSize + 1] = y

}

func (attr *BufferAttribute) GetZ(index) {

	return attr.array[index * attr.ItemSize + 2]

}

func (attr *BufferAttribute) SetZ(index, z) {

	attr.array[index * attr.ItemSize + 2] = z

	return attr

}

func (attr *BufferAttribute) GetW(index) {

	return attr.array[index * attr.ItemSize + 3]

}

func (attr *BufferAttribute) SetW(index, w) {

	attr.array[index * attr.ItemSize + 3] = w

	return attr

}

func (attr *BufferAttribute) SetXY(index, x, y) {

	index *= attr.ItemSize

	attr.array[index + 0] = x
	attr.array[index + 1] = y

	return attr

}

func (attr *BufferAttribute) SetXYZ(index, x, y, z) {

	index *= attr.ItemSize

	attr.array[index + 0] = x
	attr.array[index + 1] = y
	attr.array[index + 2] = z

	return attr

}

func (attr *BufferAttribute) SetXYZW(index, x, y, z, w) {

	index *= attr.ItemSize

	attr.array[index + 0] = x
	attr.array[index + 1] = y
	attr.array[index + 2] = z
	attr.array[index + 3] = w

	return attr

}*/

func (attr *BufferAttribute) OnUpload(callback func()) {

	attr.onUploadCallback = callback

}

func (attr *BufferAttribute) Clone() *BufferAttribute {

	newAttr := NewBufferAttribute(attr.array, attr.ItemSize)
	newAttr.Copy(attr)
	return newAttr

}

type Int8BufferAttribute struct {
	*BufferAttribute
}

func (attr *Int8BufferAttribute) Array() []int8 {
	return attr.array.([]int8)
}

func (attr *Int8BufferAttribute) SetArray(array []int8) {
	attr.array = array
}

type Uint8BufferAttribute struct {
	*BufferAttribute
}

func (attr *Uint8BufferAttribute) Array() []uint8 {
	return attr.array.([]uint8)
}

func (attr *Uint8BufferAttribute) SetArray(array []uint8) {
	attr.array = array
}

type Uint8ClampedBufferAttribute struct {
	*BufferAttribute
}

func (attr *Uint8ClampedBufferAttribute) Array() []uint8 {
	return attr.array.([]uint8)
}

func (attr *Uint8ClampedBufferAttribute) SetArray(array []uint8) {
	attr.array = array
}

type Int16BufferAttribute struct {
	*BufferAttribute
}

func (attr *Int16BufferAttribute) Array() []int16 {
	return attr.array.([]int16)
}

func (attr *Int16BufferAttribute) SetArray(array []int16) {
	attr.array = array
}

type Uint16BufferAttribute struct {
	*BufferAttribute
}

func (attr *Uint16BufferAttribute) Array() []uint16 {
	return attr.array.([]uint16)
}

func (attr *Uint16BufferAttribute) SetArray(array []uint16) {
	attr.array = uarray
}

type Int32BufferAttribute struct {
	*BufferAttribute
}

func (attr *Int32BufferAttribute) Array() []int32 {
	return attr.array.([]int32)
}

func (attr *Int32BufferAttribute) SetArray(array []int32) {
	attr.array = array
}

type Uint32BufferAttribute struct {
	*BufferAttribute
}

func (attr *Uint32BufferAttribute) Array() []uint32 {
	return attr.array.([]uint32)
}

func (attr *Uint32BufferAttribute) SetArray(array []uint32) {
	attr.array = uarray
}

type Float32BufferAttribute struct {
	*BufferAttribute
}

func (attr *Float32BufferAttribute) Array() []float32 {
	return attr.array.([]float32)
}

func (attr *Float32BufferAttribute) SetArray(array []float32) {
	attr.array = array
}

type Float64BufferAttribute struct {
	*BufferAttribute
}

func (attr *Float64BufferAttribute) Array() []float64 {
	return attr.array.([]float64)
}

func (attr *Float64BufferAttribute) SetArray(array []float64) {
	attr.array = array
}
