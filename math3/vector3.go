package math3

import (
	"fmt"
	"math"
)

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func NewVector3() *Vector3 {

	return &Vector3{0, 0, 0}

}

func Vector3_Max() *Vector3 {
	return &Vector3{
		math.MaxFloat64,
		math.MaxFloat64,
		math.MaxFloat64,
	}
}
func Vector3_Min() *Vector3 {
	return &Vector3{
		-math.MaxFloat64,
		-math.MaxFloat64,
		-math.MaxFloat64,
	}
}

func (v *Vector3) String() string {
	return fmt.Sprintf("&Vector3{X: %.4f, Y: %.4f, Z: %.4f}", v.X, v.Y, v.Z)
}

func (v *Vector3) Set(x, y, z float64) *Vector3 {

	v.X = x
	v.Y = y
	v.Z = z

	return v

}

func (v *Vector3) SetScalar(scalar float64) *Vector3 {

	v.X = scalar
	v.Y = scalar
	v.Z = scalar

	return v

}

func (v *Vector3) SetX(x float64) *Vector3 {

	v.X = x

	return v

}

func (v *Vector3) SetY(y float64) *Vector3 {

	v.Y = y

	return v

}

func (v *Vector3) SetZ(z float64) *Vector3 {

	v.Z = z

	return v

}

func (v *Vector3) SetComponent(index int, value float64) *Vector3 {

	switch index {

	case 0:
		v.X = value
	case 1:
		v.Y = value
	case 2:
		v.Z = value
	default:
		panic(fmt.Sprintf("index is out of range: %d", index))
	}

	return v

}

func (v *Vector3) GetComponent(index int) float64 {

	switch index {

	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	default:
		panic(fmt.Sprintf("index is out of range: %d", index))

	}

}

func (v *Vector3) Clone() *Vector3 {

	return NewVector3().Set(v.X, v.Y, v.Z)

}

func (v *Vector3) Copy(other *Vector3) *Vector3 {

	v.X = other.X
	v.Y = other.Y
	v.Z = other.Z

	return v

}

func (v *Vector3) Add(other *Vector3) *Vector3 {

	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z

	return v

}

func (v *Vector3) AddScalar(s float64) *Vector3 {

	v.X += s
	v.Y += s
	v.Z += s

	return v

}

func (v *Vector3) AddVectors(a, b *Vector3) *Vector3 {

	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	v.Z = a.Z + b.Z

	return v

}

func (v *Vector3) AddScaledVector(other *Vector3, s float64) *Vector3 {

	v.X += other.X * s
	v.Y += other.Y * s
	v.Z += other.Z * s

	return v

}

func (v *Vector3) Sub(other *Vector3) *Vector3 {

	v.X -= other.X
	v.Y -= other.Y
	v.Z -= other.Z

	return v

}

func (v *Vector3) SubScalar(s float64) *Vector3 {

	v.X -= s
	v.Y -= s
	v.Z -= s

	return v

}

func (v *Vector3) SubVectors(a, b *Vector3) *Vector3 {

	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	v.Z = a.Z - b.Z

	return v

}

func (v *Vector3) Multiply(other *Vector3) *Vector3 {

	v.X *= other.X
	v.Y *= other.Y
	v.Z *= other.Z

	return v

}

func (v *Vector3) MultiplyScalar(scalar float64) *Vector3 {

	if false == math.IsInf(scalar, 0) {

		v.X *= scalar
		v.Y *= scalar
		v.Z *= scalar

	} else {

		v.X = 0
		v.Y = 0
		v.Z = 0

	}

	return v

}

func (v *Vector3) MultiplyVectors(a, b *Vector3) *Vector3 {

	v.X = a.X * b.X
	v.Y = a.Y * b.Y
	v.Z = a.Z * b.Z

	return v

}

func (v *Vector3) ApplyEuler(euler *Euler) *Vector3 {

	quaternion := NewQuaternion()

	return v.ApplyQuaternion(quaternion.SetFromEuler(euler, false))

}

/*func (v *Vector3) ApplyAxisAngle(axis *Vector3, angle float64) *Vector3 {

    quaternion := NewQuaternion()

    return v.ApplyQuaternion(quaternion.SetFromAxisAngle(axis, angle))

}*/

func (v *Vector3) ApplyMatrix3(m *Matrix3) *Vector3 {

	x := v.X
	y := v.Y
	z := v.Z
	e := m.Elements

	v.X = e[0]*x + e[3]*y + e[6]*z
	v.Y = e[1]*x + e[4]*y + e[7]*z
	v.Z = e[2]*x + e[5]*y + e[8]*z

	return v

}

func (v *Vector3) ApplyMatrix4(m *Matrix4) *Vector3 {

	// input: THREE.Matrix4 affine matrix

	x := v.X
	y := v.Y
	z := v.Z
	e := m.Elements

	v.X = e[0]*x + e[4]*y + e[8]*z + e[12]
	v.Y = e[1]*x + e[5]*y + e[9]*z + e[13]
	v.Z = e[2]*x + e[6]*y + e[10]*z + e[14]

	return v

}

func (v *Vector3) ApplyProjection(m *Matrix4) *Vector3 {

	// input: THREE.Matrix4 projection matrix

	x := v.X
	y := v.Y
	z := v.Z
	e := m.Elements
	d := 1 / (e[3]*x + e[7]*y + e[11]*z + e[15]) // perspective divide

	v.X = (e[0]*x + e[4]*y + e[8]*z + e[12]) * d
	v.Y = (e[1]*x + e[5]*y + e[9]*z + e[13]) * d
	v.Z = (e[2]*x + e[6]*y + e[10]*z + e[14]) * d

	return v

}

func (v *Vector3) ApplyQuaternion(q *Quaternion) *Vector3 {

	x := v.X
	y := v.Y
	z := v.Z
	qx := q.GetX()
	qy := q.GetY()
	qz := q.GetZ()
	qw := q.GetW()

	// calculate quat * vector

	ix := qw*x + qy*z - qz*y
	iy := qw*y + qz*x - qx*z
	iz := qw*z + qx*y - qy*x
	iw := -qx*x - qy*y - qz*z

	// calculate result * inverse quat

	v.X = ix*qw + iw*-qx + iy*-qz - iz*-qy
	v.Y = iy*qw + iw*-qy + iz*-qx - ix*-qz
	v.Z = iz*qw + iw*-qz + ix*-qy - iy*-qx

	return v

}

func (v *Vector3) Project(camera Projector) *Vector3 {

	matrix := NewMatrix4()

	matrix.MultiplyMatrices(camera.GetProjectionMatrix(), matrix.GetInverse(camera.GetMatrixWorld()))
	return v.ApplyProjection(matrix)

}

func (v *Vector3) Unproject(camera Projector) *Vector3 {

	matrix := NewMatrix4()

	matrix.MultiplyMatrices(camera.GetMatrixWorld(), matrix.GetInverse(camera.GetProjectionMatrix()))
	return v.ApplyProjection(matrix)

}

func (v *Vector3) TransformDirection(m *Matrix4) *Vector3 {

	// input: THREE.Matrix4 affine matrix
	// vector interpreted as a direction

	x := v.X
	y := v.Y
	z := v.Z
	e := m.Elements

	v.X = e[0]*x + e[4]*y + e[8]*z
	v.Y = e[1]*x + e[5]*y + e[9]*z
	v.Z = e[2]*x + e[6]*y + e[10]*z

	return v.Normalize()

}

func (v *Vector3) Divide(other *Vector3) *Vector3 {

	v.X /= other.X
	v.Y /= other.Y
	v.Z /= other.Z

	return v

}

func (v *Vector3) DivideScalar(scalar float64) *Vector3 {

	return v.MultiplyScalar(1 / scalar)

}

func (v *Vector3) Min(other *Vector3) *Vector3 {

	v.X = math.Min(v.X, other.X)
	v.Y = math.Min(v.Y, other.Y)
	v.Z = math.Min(v.Z, other.Z)

	return v

}

func (v *Vector3) Max(other *Vector3) *Vector3 {

	v.X = math.Max(v.X, other.X)
	v.Y = math.Max(v.Y, other.Y)
	v.Z = math.Max(v.Z, other.Z)

	return v

}

func (v *Vector3) Clamp(min, max *Vector3) *Vector3 {

	// This function assumes min < max, if v assumption isn"t true it will not operate correctly

	v.X = math.Max(min.X, math.Min(max.X, v.X))
	v.Y = math.Max(min.Y, math.Min(max.Y, v.Y))
	v.Z = math.Max(min.Z, math.Min(max.Z, v.Z))

	return v

}

func (v *Vector3) ClampScalar(minVal, maxVal float64) *Vector3 {

	min := NewVector3()
	max := NewVector3()

	min.Set(minVal, minVal, minVal)
	max.Set(maxVal, maxVal, maxVal)

	return v.Clamp(min, max)

}

func (v *Vector3) ClampLength(min, max float64) *Vector3 {

	length := v.Length()

	return v.MultiplyScalar(math.Max(min, math.Min(max, length)) / length)

}

func (v *Vector3) Floor() *Vector3 {

	v.X = math.Floor(v.X)
	v.Y = math.Floor(v.Y)
	v.Z = math.Floor(v.Z)

	return v

}

func (v *Vector3) Ceil() *Vector3 {

	v.X = math.Ceil(v.X)
	v.Y = math.Ceil(v.Y)
	v.Z = math.Ceil(v.Z)

	return v

}

func (v *Vector3) Round() *Vector3 {

	v.X = Round(v.X)
	v.Y = Round(v.Y)
	v.Z = Round(v.Z)

	return v

}

func (v *Vector3) RoundToZero() *Vector3 {

	if v.X < 0 {
		v.X = math.Ceil(v.X)
	} else {
		v.X = math.Floor(v.X)
	}
	if v.Y < 0 {
		v.Y = math.Ceil(v.Y)
	} else {
		v.Y = math.Floor(v.Y)
	}
	if v.Z < 0 {
		v.Z = math.Ceil(v.Z)
	} else {
		v.Z = math.Floor(v.Z)
	}
	return v

}

func (v *Vector3) Negate() *Vector3 {

	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z

	return v

}

func (v *Vector3) Dot(other *Vector3) float64 {

	return v.X*other.X + v.Y*other.Y + v.Z*other.Z

}

func (v *Vector3) LengthSq() float64 {

	return v.X*v.X + v.Y*v.Y + v.Z*v.Z

}

func (v *Vector3) Length() float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)

}

func (v *Vector3) LengthManhattan() float64 {

	return math.Abs(v.X) + math.Abs(v.Y) + math.Abs(v.Z)

}

func (v *Vector3) Normalize() *Vector3 {

	return v.DivideScalar(v.Length())

}

func (v *Vector3) SetLength(length float64) *Vector3 {

	return v.MultiplyScalar(length / v.Length())

}

func (v *Vector3) Lerp(other *Vector3, alpha float64) *Vector3 {

	v.X += (other.X - v.X) * alpha
	v.Y += (other.Y - v.Y) * alpha
	v.Z += (other.Z - v.Z) * alpha

	return v

}

func (v *Vector3) LerpVectors(v1, v2 *Vector3, alpha float64) *Vector3 {

	return v.SubVectors(v2, v1).MultiplyScalar(alpha).Add(v1)

}

func (v *Vector3) Cross(other *Vector3) *Vector3 {

	x := v.X
	y := v.Y
	z := v.Z

	v.X = y*other.Z - z*other.Y
	v.Y = z*other.X - x*other.Z
	v.Z = x*other.Y - y*other.X

	return v

}

func (v *Vector3) CrossVectors(a, b *Vector3) *Vector3 {

	ax := a.X
	ay := a.Y
	az := a.Z
	bx := b.X
	by := b.Y
	bz := b.Z

	v.X = ay*bz - az*by
	v.Y = az*bx - ax*bz
	v.Z = ax*by - ay*bx

	return v

}

func (v *Vector3) ProjectOnVector(vector *Vector3) *Vector3 {

	scalar := vector.Dot(v) / vector.LengthSq()

	return v.Copy(vector).MultiplyScalar(scalar)

}

func (v *Vector3) ProjectOnPlane(planeNormal *Vector3) *Vector3 {

	v1 := NewVector3()

	v1.Copy(v).ProjectOnVector(planeNormal)

	return v.Sub(v1)

}

func (v *Vector3) Reflect(normal *Vector3) *Vector3 {

	// reflect incident vector off plane orthogonal to normal
	// normal is assumed to have unit length

	v1 := NewVector3()

	return v.Sub(v1.Copy(normal).MultiplyScalar(2 * v.Dot(normal)))

}

func (v *Vector3) AngleTo(other *Vector3) float64 {

	theta := v.Dot(other) / (math.Sqrt(v.LengthSq() * other.LengthSq()))

	// clamp, to handle numerical problems

	return math.Acos(Clamp(theta, -1, 1))

}

func (v *Vector3) DistanceTo(other *Vector3) float64 {

	return math.Sqrt(v.DistanceToSquared(other))

}

func (v *Vector3) DistanceToSquared(other *Vector3) float64 {

	dx := v.X - other.X
	dy := v.Y - other.Y
	dz := v.Z - other.Z

	return dx*dx + dy*dy + dz*dz

}

func (v *Vector3) DistanceToManhattan(other *Vector3) float64 {

	return math.Abs(v.X-other.X) + math.Abs(v.Y-other.Y) + math.Abs(v.Z-other.Z)

}

/*func (v *Vector3) SetFromSpherical(s *Spherical) *Vector3 {

    sinPhiRadius := math.Sin(s.Phi) * s.Radius

    v.X = sinPhiRadius * math.Sin(s.Theta)
    v.Y = math.Cos(s.Phi) * s.Radius
    v.Z = sinPhiRadius * math.Cos(s.Theta)

    return v

}*/

func (v *Vector3) SetFromMatrixPosition(m *Matrix4) *Vector3 {

	return v.SetFromMatrixColumn(m, 3)

}

func (v *Vector3) SetFromMatrixScale(m *Matrix4) *Vector3 {

	sx := v.SetFromMatrixColumn(m, 0).Length()
	sy := v.SetFromMatrixColumn(m, 1).Length()
	sz := v.SetFromMatrixColumn(m, 2).Length()

	v.X = sx
	v.Y = sy
	v.Z = sz

	return v

}

func (v *Vector3) SetFromMatrixColumn(m *Matrix4, index int) *Vector3 {

	return v.FromArray(m.Elements, index*4)

}

func (v *Vector3) Equals(other *Vector3) bool {

	return ((other.X == v.X) && (other.Y == v.Y) && (other.Z == v.Z))

}

func (v *Vector3) FromArray(array []float64, offset int) *Vector3 {

	v.X = array[offset]
	v.Y = array[offset+1]
	v.Z = array[offset+2]

	return v

}

func (v *Vector3) ToArray(target []float64, offset int) []float64 {

	if target == nil {
		target = make([]float64, 3)
	}

	target[offset] = v.X
	target[offset+1] = v.Y
	target[offset+2] = v.Z

	return target

}

func (v *Vector3) ToArray32(target []float32, offset int) []float32 {

	if target == nil {
		target = make([]float32, 3)
	}

	target[offset] = float32(v.X)
	target[offset+1] = float32(v.Y)
	target[offset+2] = float32(v.Z)

	return target

}

/*func (v *Vector3) FromAttribute(attribute *Attribute, index int, offset int) *Vector3 {

    index = index*attribute.ItemSize + offset

    v.X = attribute.Array[index]
    v.Y = attribute.Array[index+1]
    v.Z = attribute.Array[index+2]

    return v

}*/
