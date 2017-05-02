package math3

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X float64
	Y float64
}

func NewVector2() *Vector2 {

	return &Vector2{0, 0}

}

func Vector2_Max() *Vector2 {
	return &Vector2{
		math.MaxFloat64,
		math.MaxFloat64,
	}
}
func Vector2_Min() *Vector2 {
	return &Vector2{
		-math.MaxFloat64,
		-math.MaxFloat64,
	}
}

func (v *Vector2) String() string {
	return fmt.Sprintf("&Vector2{X: %.4f, Y: %.4f}", v.X, v.Y)
}

func (v *Vector2) Set(x, y float64) *Vector2 {

	v.X = x
	v.Y = y

	return v

}

func (v *Vector2) SetScalar(scalar float64) *Vector2 {

	v.X = scalar
	v.Y = scalar

	return v

}

func (v *Vector2) SetX(x float64) *Vector2 {

	v.X = x

	return v

}

func (v *Vector2) SetY(y float64) *Vector2 {

	v.Y = y

	return v

}

func (v *Vector2) SetComponent(index int, value float64) *Vector2 {

	switch index {

	case 0:
		v.X = value
	case 1:
		v.Y = value
	default:
		panic(fmt.Sprintf("index is out of range: %d", index))
	}

	return v

}

func (v *Vector2) GetComponent(index int) float64 {

	switch index {

	case 0:
		return v.X
	case 1:
		return v.Y
	default:
		panic(fmt.Sprintf("index is out of range: %d", index))

	}

}

func (v *Vector2) Clone() *Vector2 {

	return NewVector2().Set(v.X, v.Y)

}

func (v *Vector2) Copy(other *Vector2) *Vector2 {

	v.X = other.X
	v.Y = other.Y

	return v

}

func (v *Vector2) Add(other *Vector2) *Vector2 {

	v.X += other.X
	v.Y += other.Y

	return v

}

func (v *Vector2) AddScalar(s float64) *Vector2 {

	v.X += s
	v.Y += s

	return v

}

func (v *Vector2) AddVectors(a, b *Vector2) *Vector2 {

	v.X = a.X + b.X
	v.Y = a.Y + b.Y

	return v

}

func (v *Vector2) AddScaledVector(other *Vector2, s float64) *Vector2 {

	v.X += other.X * s
	v.Y += other.Y * s

	return v

}

func (v *Vector2) Sub(other *Vector2) *Vector2 {

	v.X -= other.X
	v.Y -= other.Y

	return v

}

func (v *Vector2) SubScalar(s float64) *Vector2 {

	v.X -= s
	v.Y -= s

	return v

}

func (v *Vector2) SubVectors(a, b *Vector2) *Vector2 {

	v.X = a.X - b.X
	v.Y = a.Y - b.Y

	return v

}

func (v *Vector2) Multiply(other *Vector2) *Vector2 {

	v.X *= other.X
	v.Y *= other.Y

	return v

}

func (v *Vector2) MultiplyScalar(scalar float64) *Vector2 {

	if false == math.IsInf(scalar, 0) {

		v.X *= scalar
		v.Y *= scalar

	} else {

		v.X = 0
		v.Y = 0

	}

	return v

}

func (v *Vector2) MultiplyVectors(a, b *Vector2) *Vector2 {

	v.X = a.X * b.X
	v.Y = a.Y * b.Y

	return v

}

/*func (v *Vector2) ApplyEuler(euler *Euler) *Vector2 {

	quaternion := NewQuaternion()

	return v.ApplyQuaternion(quaternion.SetFromEuler(euler, false))

}*/

/*func (v *Vector2) ApplyAxisAngle(axis *Vector2, angle float64) *Vector2 {

    quaternion := NewQuaternion()

    return v.ApplyQuaternion(quaternion.SetFromAxisAngle(axis, angle))

}*/

/*func (v *Vector2) ApplyMatrix3(m *Matrix3) *Vector2 {

	x := v.X
	y := v.Y
	z := v.Z
	e := m.Elements

	v.X = e[0]*x + e[3]*y + e[6]*z
	v.Y = e[1]*x + e[4]*y + e[7]*z
	v.Z = e[2]*x + e[5]*y + e[8]*z

	return v

}

func (v *Vector2) ApplyMatrix4(m *Matrix4) *Vector2 {

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

func (v *Vector2) ApplyProjection(m *Matrix4) *Vector2 {

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

func (v *Vector2) ApplyQuaternion(q *Quaternion) *Vector2 {

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

func (v *Vector2) Project(camera Projector) *Vector2 {

	matrix := NewMatrix4()

	matrix.MultiplyMatrices(camera.GetProjectionMatrix(), matrix.GetInverse(camera.GetMatrixWorld()))
	return v.ApplyProjection(matrix)

}

func (v *Vector2) Unproject(camera Projector) *Vector2 {

	matrix := NewMatrix4()

	matrix.MultiplyMatrices(camera.GetMatrixWorld(), matrix.GetInverse(camera.GetProjectionMatrix()))
	return v.ApplyProjection(matrix)

}

func (v *Vector2) TransformDirection(m *Matrix4) *Vector2 {

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

}*/

func (v *Vector2) Divide(other *Vector2) *Vector2 {

	v.X /= other.X
	v.Y /= other.Y

	return v

}

func (v *Vector2) DivideScalar(scalar float64) *Vector2 {

	return v.MultiplyScalar(1 / scalar)

}

func (v *Vector2) Min(other *Vector2) *Vector2 {

	v.X = math.Min(v.X, other.X)
	v.Y = math.Min(v.Y, other.Y)

	return v

}

func (v *Vector2) Max(other *Vector2) *Vector2 {

	v.X = math.Max(v.X, other.X)
	v.Y = math.Max(v.Y, other.Y)

	return v

}

func (v *Vector2) Clamp(min, max *Vector2) *Vector2 {

	// This function assumes min < max, if v assumption isn"t true it will not operate correctly

	v.X = math.Max(min.X, math.Min(max.X, v.X))
	v.Y = math.Max(min.Y, math.Min(max.Y, v.Y))

	return v

}

func (v *Vector2) ClampScalar(minVal, maxVal float64) *Vector2 {

	min := NewVector2()
	max := NewVector2()

	min.Set(minVal, minVal)
	max.Set(maxVal, maxVal)

	return v.Clamp(min, max)

}

func (v *Vector2) ClampLength(min, max float64) *Vector2 {

	length := v.Length()

	return v.MultiplyScalar(math.Max(min, math.Min(max, length)) / length)

}

func (v *Vector2) Floor() *Vector2 {

	v.X = math.Floor(v.X)
	v.Y = math.Floor(v.Y)

	return v

}

func (v *Vector2) Ceil() *Vector2 {

	v.X = math.Ceil(v.X)
	v.Y = math.Ceil(v.Y)

	return v

}

func (v *Vector2) Round() *Vector2 {

	v.X = Round(v.X)
	v.Y = Round(v.Y)

	return v

}

func (v *Vector2) RoundToZero() *Vector2 {

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
	return v

}

func (v *Vector2) Negate() *Vector2 {

	v.X = -v.X
	v.Y = -v.Y

	return v

}

/*func (v *Vector2) Dot(other *Vector2) float64 {

	return v.X*other.X + v.Y*other.Y

}*/

func (v *Vector2) LengthSq() float64 {

	return v.X*v.X + v.Y*v.Y

}

func (v *Vector2) Length() float64 {

	return math.Sqrt(v.LengthSq())

}

func (v *Vector2) LengthManhattan() float64 {

	return math.Abs(v.X) + math.Abs(v.Y)

}

func (v *Vector2) Normalize() *Vector2 {

	return v.DivideScalar(v.Length())

}

func (v *Vector2) SetLength(length float64) *Vector2 {

	return v.MultiplyScalar(length / v.Length())

}

func (v *Vector2) Lerp(other *Vector2, alpha float64) *Vector2 {

	v.X += (other.X - v.X) * alpha
	v.Y += (other.Y - v.Y) * alpha

	return v

}

func (v *Vector2) LerpVectors(v1, v2 *Vector2, alpha float64) *Vector2 {

	return v.SubVectors(v2, v1).MultiplyScalar(alpha).Add(v1)

}

/*func (v *Vector2) Cross(other *Vector2) *Vector2 {

	x := v.X
	y := v.Y
	z := v.Z

	v.X = y*other.Z - z*other.Y
	v.Y = z*other.X - x*other.Z
	v.Z = x*other.Y - y*other.X

	return v

}*/

/*func (v *Vector2) CrossVectors(a, b *Vector2) *Vector2 {

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

}*/

/*func (v *Vector2) ProjectOnVector(vector *Vector2) *Vector2 {

	scalar := vector.Dot(v) / vector.LengthSq()

	return v.Copy(vector).MultiplyScalar(scalar)

}*/

/*func (v *Vector2) ProjectOnPlane(planeNormal *Vector2) *Vector2 {

	v1 := NewVector2()

	v1.Copy(v).ProjectOnVector(planeNormal)

	return v.Sub(v1)

}

func (v *Vector2) Reflect(normal *Vector2) *Vector2 {

	// reflect incident vector off plane orthogonal to normal
	// normal is assumed to have unit length

	v1 := NewVector2()

	return v.Sub(v1.Copy(normal).MultiplyScalar(2 * v.Dot(normal)))

}*/

/*func (v *Vector2) AngleTo(other *Vector2) float64 {

	theta := v.Dot(other) / (math.Sqrt(v.LengthSq() * other.LengthSq()))

	// clamp, to handle numerical problems

	return math.Acos(Clamp(theta, -1, 1))

}*/

func (v *Vector2) DistanceTo(other *Vector2) float64 {

	return math.Sqrt(v.DistanceToSquared(other))

}

func (v *Vector2) DistanceToSquared(other *Vector2) float64 {

	dx := v.X - other.X
	dy := v.Y - other.Y

	return dx*dx + dy*dy

}

func (v *Vector2) DistanceToManhattan(other *Vector2) float64 {

	return math.Abs(v.X-other.X) + math.Abs(v.Y-other.Y)

}

/*func (v *Vector2) SetFromSpherical(s *Spherical) *Vector2 {

    sinPhiRadius := math.Sin(s.Phi) * s.Radius

    v.X = sinPhiRadius * math.Sin(s.Theta)
    v.Y = math.Cos(s.Phi) * s.Radius
    v.Z = sinPhiRadius * math.Cos(s.Theta)

    return v

}*/

/*func (v *Vector2) SetFromMatrixPosition(m *Matrix4) *Vector2 {

	return v.SetFromMatrixColumn(m, 3)

}

func (v *Vector2) SetFromMatrixScale(m *Matrix4) *Vector2 {

	sx := v.SetFromMatrixColumn(m, 0).Length()
	sy := v.SetFromMatrixColumn(m, 1).Length()
	sz := v.SetFromMatrixColumn(m, 2).Length()

	v.X = sx
	v.Y = sy
	v.Z = sz

	return v

}

func (v *Vector2) SetFromMatrixColumn(m *Matrix4, index int) *Vector2 {

	return v.FromArray(m.Elements, index*4)

}*/

func (v *Vector2) Equals(other *Vector2) bool {

	return ((other.X == v.X) && (other.Y == v.Y))

}

func (v *Vector2) FromArray(array []float64, offset int) *Vector2 {

	v.X = array[offset]
	v.Y = array[offset+1]

	return v

}

func (v *Vector2) ToArray(target []float64, offset int) []float64 {

	if target == nil {
		target = make([]float64, 2)
	}

	target[offset] = v.X
	target[offset+1] = v.Y

	return target

}

func (v *Vector2) ToArray32(target []float32, offset int) []float32 {

	if target == nil {
		target = make([]float32, 2)
	}

	target[offset] = float32(v.X)
	target[offset+1] = float32(v.Y)

	return target

}

/*func (v *Vector2) FromAttribute(attribute *Attribute, index int, offset int) *Vector2 {

    index = index*attribute.ItemSize + offset

    v.X = attribute.Array[index]
    v.Y = attribute.Array[index+1]
    v.Z = attribute.Array[index+2]

    return v

}*/
