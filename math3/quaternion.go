package math3

import (
	"fmt"
	"math"
)

type Quaternion struct {
	x                float64
	y                float64
	z                float64
	w                float64
	OnChangeCallback func()
}

func NewQuaternion() *Quaternion {

	return &Quaternion{0, 0, 0, 1, func() {}}

}

func (q *Quaternion) GetX() float64 {

	return q.x

}

func (q *Quaternion) SetX(value float64) {

	q.x = value
	q.OnChangeCallback()

}

func (q *Quaternion) GetY() float64 {

	return q.y

}

func (q *Quaternion) SetY(value float64) {

	q.y = value
	q.OnChangeCallback()

}

func (q *Quaternion) GetZ() float64 {

	return q.z

}

func (q *Quaternion) SetZ(value float64) {

	q.z = value
	q.OnChangeCallback()

}

func (q *Quaternion) GetW() float64 {

	return q.w

}

func (q *Quaternion) SetW(value float64) {

	q.w = value
	q.OnChangeCallback()

}

func (q *Quaternion) String() string {
	return fmt.Sprintf("&Quaternion{ %.2f, %.2f, %.2f, %.2f }", q.x, q.y, q.z, q.w)
}

func (q *Quaternion) Set(x, y, z, w float64) *Quaternion {

	q.x = x
	q.y = y
	q.z = z
	q.w = w

	q.OnChangeCallback()

	return q

}

func (q *Quaternion) Clone() *Quaternion {

	return NewQuaternion().Set(q.x, q.y, q.z, q.w)

}

func (q *Quaternion) Copy(quaternion *Quaternion) *Quaternion {

	q.x = quaternion.x
	q.y = quaternion.y
	q.z = quaternion.z
	q.w = quaternion.w

	q.OnChangeCallback()

	return q

}

func (q *Quaternion) SetFromEuler(euler *Euler, update bool) *Quaternion {

	// http://www.mathworks.Com/matlabcentral/fileexchange/
	//  20696-function-to-convert-between-dcm-euler-angles-quaternions-and-euler-vectors/
	//  content/SpinCalc.M

	c1 := math.Cos(euler.x / 2)
	c2 := math.Cos(euler.y / 2)
	c3 := math.Cos(euler.z / 2)
	s1 := math.Sin(euler.x / 2)
	s2 := math.Sin(euler.y / 2)
	s3 := math.Sin(euler.z / 2)

	order := euler.Order

	if order == XYZ {

		q.x = s1*c2*c3 + c1*s2*s3
		q.y = c1*s2*c3 - s1*c2*s3
		q.z = c1*c2*s3 + s1*s2*c3
		q.w = c1*c2*c3 - s1*s2*s3

	} else if order == YXZ {

		q.x = s1*c2*c3 + c1*s2*s3
		q.y = c1*s2*c3 - s1*c2*s3
		q.z = c1*c2*s3 - s1*s2*c3
		q.w = c1*c2*c3 + s1*s2*s3

	} else if order == ZXY {

		q.x = s1*c2*c3 - c1*s2*s3
		q.y = c1*s2*c3 + s1*c2*s3
		q.z = c1*c2*s3 + s1*s2*c3
		q.w = c1*c2*c3 - s1*s2*s3

	} else if order == ZYX {

		q.x = s1*c2*c3 - c1*s2*s3
		q.y = c1*s2*c3 + s1*c2*s3
		q.z = c1*c2*s3 - s1*s2*c3
		q.w = c1*c2*c3 + s1*s2*s3

	} else if order == YZX {

		q.x = s1*c2*c3 + c1*s2*s3
		q.y = c1*s2*c3 + s1*c2*s3
		q.z = c1*c2*s3 - s1*s2*c3
		q.w = c1*c2*c3 - s1*s2*s3

	} else if order == XZY {

		q.x = s1*c2*c3 - c1*s2*s3
		q.y = c1*s2*c3 - s1*c2*s3
		q.z = c1*c2*s3 + s1*s2*c3
		q.w = c1*c2*c3 + s1*s2*s3

	}

	if update != false {
		q.OnChangeCallback()
	}

	return q

}

func (q *Quaternion) SetFromAxisAngle(axis *Vector3, angle float64) *Quaternion {

	// http://www.Euclideanspace.Com/maths/geometry/rotations/conversions/angleToQuaternion/index.Htm

	// assumes axis is normalized

	halfAngle := angle / 2
	s := math.Sin(halfAngle)

	q.x = axis.X * s
	q.y = axis.Y * s
	q.z = axis.Z * s
	q.w = math.Cos(halfAngle)

	q.OnChangeCallback()

	return q

}

func (q *Quaternion) SetFromRotationMatrix(m *Matrix4) *Quaternion {

	// http://www.Euclideanspace.Com/maths/geometry/rotations/conversions/matrixToQuaternion/index.Htm

	// assumes the upper 3x3 of m is a pure rotation matrix (i.E, unscaled)

	te := m.Elements

	m11, m12, m13 := te[0], te[4], te[8]
	m21, m22, m23 := te[1], te[5], te[9]
	m31, m32, m33 := te[2], te[6], te[10]

	trace := m11 + m22 + m33
	var s float64

	if trace > 0 {

		s = 0.5 / math.Sqrt(trace+1.0)

		q.w = 0.25 / s
		q.x = (m32 - m23) * s
		q.y = (m13 - m31) * s
		q.z = (m21 - m12) * s

	} else if m11 > m22 && m11 > m33 {

		s = 2.0 * math.Sqrt(1.0+m11-m22-m33)

		q.w = (m32 - m23) / s
		q.x = 0.25 * s
		q.y = (m12 + m21) / s
		q.z = (m13 + m31) / s

	} else if m22 > m33 {

		s = 2.0 * math.Sqrt(1.0+m22-m11-m33)

		q.w = (m13 - m31) / s
		q.x = (m12 + m21) / s
		q.y = 0.25 * s
		q.z = (m23 + m32) / s

	} else {

		s = 2.0 * math.Sqrt(1.0+m33-m11-m22)

		q.w = (m21 - m12) / s
		q.x = (m13 + m31) / s
		q.y = (m23 + m32) / s
		q.z = 0.25 * s

	}

	q.OnChangeCallback()

	return q

}

func (q *Quaternion) SetFromUnitVectors(vFrom, vTo *Vector3) *Quaternion {

	// http://lolengine.Net/blog/2014/02/24/quaternion-from-two-vectors-final

	// assumes direction vectors vFrom and vTo are normalized

	v1 := NewVector3()

	EPS := 0.000001

	r := vFrom.Dot(vTo) + 1

	if r < EPS {

		r = 0

		if math.Abs(vFrom.X) > math.Abs(vFrom.Z) {

			v1.Set(-vFrom.Y, vFrom.X, 0)

		} else {

			v1.Set(0, -vFrom.Z, vFrom.Y)

		}

	} else {

		v1.CrossVectors(vFrom, vTo)

	}

	q.x = v1.X
	q.y = v1.Y
	q.z = v1.Z
	q.w = r

	return q.Normalize()

}

func (q *Quaternion) Inverse() *Quaternion {

	return q.Conjugate().Normalize()

}

func (q *Quaternion) Conjugate() *Quaternion {

	q.x *= -1
	q.y *= -1
	q.z *= -1

	q.OnChangeCallback()

	return q

}

func (q *Quaternion) Dot(v *Quaternion) float64 {

	return q.x*v.x + q.y*v.y + q.z*v.z + q.w*v.w

}

func (q *Quaternion) LengthSq() float64 {

	return q.x*q.x + q.y*q.y + q.z*q.z + q.w*q.w

}

func (q *Quaternion) Length() float64 {

	return math.Sqrt(q.x*q.x + q.y*q.y + q.z*q.z + q.w*q.w)

}

func (q *Quaternion) Normalize() *Quaternion {

	l := q.Length()

	if l == 0 {

		q.x = 0
		q.y = 0
		q.z = 0
		q.w = 1

	} else {

		l = 1 / l

		q.x = q.x * l
		q.y = q.y * l
		q.z = q.z * l
		q.w = q.w * l

	}

	q.OnChangeCallback()

	return q

}

func (q *Quaternion) Multiply(other *Quaternion) *Quaternion {

	return q.MultiplyQuaternions(q, other)

}

func (q *Quaternion) Premultiply(other *Quaternion) *Quaternion {

	return q.MultiplyQuaternions(other, q)

}

func (q *Quaternion) MultiplyQuaternions(a, b *Quaternion) *Quaternion {

	// from http://www.Euclideanspace.Com/maths/algebra/realNormedAlgebra/quaternions/code/index.Htm

	qax := a.x
	qay := a.y
	qaz := a.z
	qaw := a.w
	qbx := b.x
	qby := b.y
	qbz := b.z
	qbw := b.w

	q.x = qax*qbw + qaw*qbx + qay*qbz - qaz*qby
	q.y = qay*qbw + qaw*qby + qaz*qbx - qax*qbz
	q.z = qaz*qbw + qaw*qbz + qax*qby - qay*qbx
	q.w = qaw*qbw - qax*qbx - qay*qby - qaz*qbz

	q.OnChangeCallback()

	return q

}

func (q *Quaternion) Slerp(qb *Quaternion, t float64) *Quaternion {

	if t == 0 {
		return q
	}
	if t == 1 {
		return q.Copy(qb)
	}

	x := q.x
	y := q.y
	z := q.z
	w := q.w

	// http://www.Euclideanspace.Com/maths/algebra/realNormedAlgebra/quaternions/slerp/

	cosHalfTheta := w*qb.w + x*qb.x + y*qb.y + z*qb.z

	if cosHalfTheta < 0 {

		q.w = -qb.w
		q.x = -qb.x
		q.y = -qb.y
		q.z = -qb.z

		cosHalfTheta = -cosHalfTheta

	} else {

		q.Copy(qb)

	}

	if cosHalfTheta >= 1.0 {

		q.w = w
		q.x = x
		q.y = y
		q.z = z

		return q

	}

	sinHalfTheta := math.Sqrt(1.0 - cosHalfTheta*cosHalfTheta)

	if math.Abs(sinHalfTheta) < 0.001 {

		q.w = 0.5 * (w + q.w)
		q.x = 0.5 * (x + q.x)
		q.y = 0.5 * (y + q.y)
		q.z = 0.5 * (z + q.z)

		return q

	}

	halfTheta := math.Atan2(sinHalfTheta, cosHalfTheta)
	ratioA := math.Sin((1-t)*halfTheta) / sinHalfTheta
	ratioB := math.Sin(t*halfTheta) / sinHalfTheta

	q.w = (w*ratioA + q.w*ratioB)
	q.x = (x*ratioA + q.x*ratioB)
	q.y = (y*ratioA + q.y*ratioB)
	q.z = (z*ratioA + q.z*ratioB)

	q.OnChangeCallback()

	return q

}

func (q *Quaternion) Equals(quaternion *Quaternion) bool {

	return (quaternion.x == q.x) && (quaternion.y == q.y) && (quaternion.z == q.z) && (quaternion.w == q.w)

}

func (q *Quaternion) FromArray(array []float64, offset int) *Quaternion {

	q.x = array[offset]
	q.y = array[offset+1]
	q.z = array[offset+2]
	q.w = array[offset+3]

	q.OnChangeCallback()

	return q

}

func (q *Quaternion) ToArray(array []float64, offset int) []float64 {

	if array == nil {
		array = make([]float64, 4)
	}

	array[offset] = q.x
	array[offset+1] = q.y
	array[offset+2] = q.z
	array[offset+3] = q.w

	return array

}

func (q *Quaternion) OnChange(callback func()) *Quaternion {

	q.OnChangeCallback = callback

	return q

}

func QuaternionSlerp(qa, qb, qm *Quaternion, t float64) *Quaternion {

	return qm.Copy(qa).Slerp(qb, t)

}

func QuaternionSlerpFlat(
	dst []float64, dstOffset int, src0 []float64, srcOffset0 int, src1 []float64, srcOffset1 int, t float64) {

	// fuzz-free, array-based Quaternion SLERP operation

	x0 := src0[srcOffset0+0]
	y0 := src0[srcOffset0+1]
	z0 := src0[srcOffset0+2]
	w0 := src0[srcOffset0+3]

	x1 := src1[srcOffset1+0]
	y1 := src1[srcOffset1+1]
	z1 := src1[srcOffset1+2]
	w1 := src1[srcOffset1+3]

	if w0 != w1 || x0 != x1 || y0 != y1 || z0 != z1 {

		s := 1 - t

		cos := x0*x1 + y0*y1 + z0*z1 + w0*w1

		dir := 1.0
		if cos < 0 {
			dir = -1
		}

		sqrSin := 1 - cos*cos

		// Skip the Slerp for tiny steps to avoid numeric problems:
		if sqrSin > Epsilon {

			sin := math.Sqrt(sqrSin)
			leng := math.Atan2(sin, cos*dir)

			s = math.Sin(s*leng) / sin
			t = math.Sin(t*leng) / sin

		}

		tDir := t * dir

		x0 = x0*s + x1*tDir
		y0 = y0*s + y1*tDir
		z0 = z0*s + z1*tDir
		w0 = w0*s + w1*tDir

		// Normalize in case we just did a lerp:
		if s == 1-t {

			f := 1 / math.Sqrt(x0*x0+y0*y0+z0*z0+w0*w0)

			x0 *= f
			y0 *= f
			z0 *= f
			w0 *= f

		}

	}

	dst[dstOffset] = x0
	dst[dstOffset+1] = y0
	dst[dstOffset+2] = z0
	dst[dstOffset+3] = w0

}
