package math3

import (
	"fmt"
	"math"
)

// Box3 represents a box or cube in 3D space. The main purpose of this is
// to represent the Minimum Bounding Boxes for objects.
type Box3 struct {
	Min *Vector3
	Max *Vector3
}

// NewBox3 constructs a default Box3 instance.
func NewBox3() *Box3 {

	return &Box3{
		Vector3_Max(),
		Vector3_Min(),
	}
}

func (b *Box3) String() string {
	return fmt.Sprintf("&Box3{Min: %s, Max: %s}", b.Min, b.Max)
}

// Set sets the lower and upper (x, y, z) boundaries of this box.
// min - Vector3 representing the lower (x, y, z) boundary of the box.
// max - Vector3 representing the lower upper (x, y, z) boundary of the box.
func (b *Box3) Set(min, max *Vector3) *Box3 {

	b.Min.Copy(min)
	b.Max.Copy(max)

	return b

}

// SetFromArray sets the upper and lower bounds of this box to include all
// of the data in array.
// array - An array of position data that the resulting box will envelop.
func (b *Box3) SetFromArray(array []float64) *Box3 {

	minX := math.MaxFloat64
	minY := math.MaxFloat64
	minZ := math.MaxFloat64

	maxX := math.SmallestNonzeroFloat64
	maxY := math.SmallestNonzeroFloat64
	maxZ := math.SmallestNonzeroFloat64

	for i, l := 0, len(array); i < l; i += 3 {

		x := array[i]
		y := array[i+1]
		z := array[i+2]

		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
		if z < minZ {
			minZ = z
		}

		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if z > maxZ {
			maxZ = z
		}

	}

	b.Min.Set(minX, minY, minZ)
	b.Max.Set(maxX, maxY, maxZ)

	return b

}

// SetFromPoints sets the upper and lower bounds of this box to include
// all of the points in points.
// points - Array of Vector3s that the resulting box will contain.
func (b *Box3) SetFromPoints(points []*Vector3) *Box3 {

	b.MakeEmpty()

	for _, point := range points {

		b.ExpandByPoint(point)

	}

	return b

}

// SetFromCenterAndSize centers this box on center and sets this box's width,
// height and depth to the values specified in size
// center - Desired center position of the box.
// size - Desired x, y and z dimensions of the box.
func (b *Box3) SetFromCenterAndSize(center, size *Vector3) *Box3 {

	v1 := NewVector3()

	halfSize := v1.Copy(size).MultiplyScalar(0.5)

	b.Min.Copy(center).Sub(halfSize)
	b.Max.Copy(center).Add(halfSize)

	return b

}

//TODO
/*func (b *Box3) SetFromObject(object Object) *Box3 {

    // Computes the world-axis-aligned bounding box of an object (including its children),
    // accounting for both the object's, and children's, world transforms

    v1 := NewVector3()

    object.UpdateMatrixWorld(true)

    b.MakeEmpty()

    object.Traverse(func(node Object) {

        var geometry *Geometry
        switch n := node.(type) {
        case *Mesh:
            geometry = n.Geometry
        }

        if geometry == nil {
            return
        }

        switch geo := ToGeneric(geometry).(type) {

        case BufferGeometry:
          attribute := geo.Attributes.Position

          if attribute == nil {
              return
          }

          var array []float64
          var offset, stride uint32

          switch attribute.(type) {

          case InterleavedBufferAttribute:
              iba := attribute.(InterleavedBufferAttribute)
              array = iba.data.array
              offset = iba.offset
              stride = iba.data.stride

          default:
              array = attribute.(Attribute).Array
              offset = 0
              stride = 3

          }

          for i := offset; i < len(array); i += stride {

              v1.FromArray(array, i)
              v1.ApplyMatrix4(node.MatrixWorld)

              b.expandByPoint(v1)

          }

        case Geometry:

            vertices := geo.Vertices

            for _, vertex := range vertices {

                v1.Copy(vertex)
                v1.ApplyMatrix4(node.(*Object3D).MatrixWorld)

                b.ExpandByPoint(v1)

            }

        }

    })

    return b

}*/

func (b *Box3) Clone() *Box3 {

	return NewBox3().Copy(b)

}

func (b *Box3) Copy(src *Box3) *Box3 {

	b.Min.Copy(src.Min)
	b.Max.Copy(src.Max)

	return b

}

func (b *Box3) MakeEmpty() *Box3 {

	b.Min = Vector3_Max()
	b.Max = Vector3_Min()

	return b

}

func (b *Box3) IsEmpty() bool {

	// b is a more robust check for empty than ( volume <= 0 ) because volume can get positive with two negative axes

	return (b.Max.X < b.Min.X) || (b.Max.Y < b.Min.Y) || (b.Max.Z < b.Min.Z)

}

func (b *Box3) Center(target *Vector3) *Vector3 {

	if nil == target {
		target = NewVector3()
	}

	return target.AddVectors(b.Min, b.Max).MultiplyScalar(0.5)

}

func (b *Box3) Size(target *Vector3) *Vector3 {

	if nil == target {
		target = NewVector3()
	}
	return target.SubVectors(b.Max, b.Min)

}

func (b *Box3) ExpandByPoint(point *Vector3) *Box3 {

	b.Min.Min(point)
	b.Max.Max(point)

	return b

}

func (b *Box3) ExpandByVector(vector *Vector3) *Box3 {

	b.Min.Sub(vector)
	b.Max.Add(vector)

	return b

}

func (b *Box3) ExpandByScalar(scalar float64) *Box3 {

	b.Min.AddScalar(-scalar)
	b.Max.AddScalar(scalar)

	return b

}

func (b *Box3) ContainsPoint(point *Vector3) bool {

	if point.X < b.Min.X || point.X > b.Max.X ||
		point.Y < b.Min.Y || point.Y > b.Max.Y ||
		point.Z < b.Min.Z || point.Z > b.Max.Z {

		return false

	}

	return true

}

func (b *Box3) ContainsBox(box *Box3) bool {

	if (b.Min.X <= box.Min.X) && (box.Max.X <= b.Max.X) &&
		(b.Min.Y <= box.Min.Y) && (box.Max.Y <= b.Max.Y) &&
		(b.Min.Z <= box.Min.Z) && (box.Max.Z <= b.Max.Z) {

		return true

	}

	return false

}

func (b *Box3) GetParameter(point, target *Vector3) *Vector3 {

	// This can potentially have a divide by zero if the box
	// has a size dimension of 0.

	if nil == target {
		target = NewVector3()
	}

	return target.Set(
		(point.X-b.Min.X)/(b.Max.X-b.Min.X),
		(point.Y-b.Min.Y)/(b.Max.Y-b.Min.Y),
		(point.Z-b.Min.Z)/(b.Max.Z-b.Min.Z),
	)

}

func (b *Box3) IntersectsBox(other *Box3) bool {

	// using 6 splitting planes to rule out intersections.

	if other.Max.X < b.Min.X || other.Min.X > b.Max.X ||
		other.Max.Y < b.Min.Y || other.Min.Y > b.Max.Y ||
		other.Max.Z < b.Min.Z || other.Min.Z > b.Max.Z {

		return false

	}

	return true

}

//TODO
/*func (b *Box3) IntersectsSphere(sphere *Sphere) bool {

    closestPoint := NewVector3()

    // Find the point on the AABB closest to the sphere center.
    b.ClampPoint(sphere.Center, closestPoint)

    // If that point is inside the sphere, the AABB and sphere intersect.
    return closestPoint.DistanceToSquared(sphere.Center) <= (sphere.Radius * sphere.Radius)

}*/

//TODO
/*func (b *Box3) IntersectsPlane(plane *Plane) bool {

    // We compute the minimum and maximum dot product values. If those values
    // are on the same side (back or front) of the plane, then there is no intersection.

    var min, max float64

    if plane.Normal.X > 0 {

        min = plane.Normal.X * b.Min.X
        max = plane.Normal.X * b.Max.X

    } else {

        min = plane.Normal.X * b.Max.X
        max = plane.Normal.X * b.Min.X

    }

    if plane.Normal.Y > 0 {

        min += plane.Normal.Y * b.Min.Y
        max += plane.Normal.Y * b.Max.Y

    } else {

        min += plane.Normal.Y * b.Max.Y
        max += plane.Normal.Y * b.Min.Y

    }

    if plane.Normal.Z > 0 {

        min += plane.Normal.Z * b.Min.Z
        max += plane.Normal.Z * b.Max.Z

    } else {

        min += plane.Normal.Z * b.Max.Z
        max += plane.Normal.Z * b.Min.Z

    }

    return (min <= plane.Constant && max >= plane.Constant)

}*/

func (b *Box3) ClampPoint(point, target *Vector3) *Vector3 {

	if nil == target {
		target = NewVector3()
	}
	return target.Copy(point).Clamp(b.Min, b.Max)

}

func (b *Box3) DistanceToPoint(point *Vector3) float64 {

	v1 := NewVector3()

	clampedPoint := v1.Copy(point).Clamp(b.Min, b.Max)
	return clampedPoint.Sub(point).Length()

}

//TODO
/*func (b *Box3) GetBoundingSphere(target *Sphere) *Sphere {

    v1 := NewVector3()

    if nil == target {
        target = NewSphere()
    }

    target.Center = b.Center(nil)
    target.Radius = b.Size(v1).Length() * 0.5

    return target

}*/

func (b *Box3) Intersect(other *Box3) *Box3 {

	b.Min.Max(other.Min)
	b.Max.Min(other.Max)

	// ensure that if there is no overlap, the result is fully empty, not slightly empty with non-inf/+inf values that will cause subsequence intersects to erroneously return valid values.
	if b.IsEmpty() {
		b.MakeEmpty()
	}

	return b

}

func (b *Box3) Union(other *Box3) *Box3 {

	b.Min.Min(other.Min)
	b.Max.Max(other.Max)

	return b

}

func (b *Box3) ApplyMatrix4(matrix *Matrix4) *Box3 {

	points := []*Vector3{
		NewVector3(),
		NewVector3(),
		NewVector3(),
		NewVector3(),
		NewVector3(),
		NewVector3(),
		NewVector3(),
		NewVector3(),
	}

	// transform of empty box is an empty box.
	if b.IsEmpty() {
		return b
	}

	// NOTE: I am using a binary pattern to specify all 2^3 combinations below
	points[0].Set(b.Min.X, b.Min.Y, b.Min.Z).ApplyMatrix4(matrix) // 000
	points[1].Set(b.Min.X, b.Min.Y, b.Max.Z).ApplyMatrix4(matrix) // 001
	points[2].Set(b.Min.X, b.Max.Y, b.Min.Z).ApplyMatrix4(matrix) // 010
	points[3].Set(b.Min.X, b.Max.Y, b.Max.Z).ApplyMatrix4(matrix) // 011
	points[4].Set(b.Max.X, b.Min.Y, b.Min.Z).ApplyMatrix4(matrix) // 100
	points[5].Set(b.Max.X, b.Min.Y, b.Max.Z).ApplyMatrix4(matrix) // 101
	points[6].Set(b.Max.X, b.Max.Y, b.Min.Z).ApplyMatrix4(matrix) // 110
	points[7].Set(b.Max.X, b.Max.Y, b.Max.Z).ApplyMatrix4(matrix) // 111

	return b.SetFromPoints(points)

}

func (b *Box3) Translate(offset *Vector3) *Box3 {

	b.Min.Add(offset)
	b.Max.Add(offset)

	return b

}

func (b *Box3) Equals(box *Box3) bool {

	return b.Min.Equals(box.Min) && b.Max.Equals(box.Max)

}
