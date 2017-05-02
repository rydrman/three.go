package math3

import (
	"math"

	"github.com/golang/glog"
)

type EulerChangeCallback func()

type Euler struct {
	x                float64
	y                float64
	z                float64
	Order            EulerOrder
	OnChangeCallback EulerChangeCallback
}

func NewEuler() *Euler {

	return &Euler{
		0, 0, 0,
		EulerDefaultOrder,
		EulerEmptyCallback,
	}

}

type EulerOrder string

const (
	XYZ          EulerOrder = "XYZ"
	YZX                     = "YZX"
	ZXY                     = "ZXY"
	XZY                     = "XZY"
	YXZ                     = "YXZ"
	ZYX                     = "ZYX"
	CurrentOrder            = "current"
)

var EulerRotationOrders = []EulerOrder{XYZ, YZX, ZXY, XZY, YXZ, ZYX}

const EulerDefaultOrder = XYZ

func (e *Euler) GetX() float64 {

	return e.x

}

func (e *Euler) SetX(value float64) {

	e.x = value
	e.OnChangeCallback()

}

func (e *Euler) GetY() float64 {

	return e.y

}

func (e *Euler) SetY(value float64) {

	e.y = value
	e.OnChangeCallback()

}

func (e *Euler) GetZ() float64 {

	return e.z

}

func (e *Euler) SetZ(value float64) {

	e.z = value
	e.OnChangeCallback()

}

func (e *Euler) GetOder() EulerOrder {

	return e.Order

}

func (e *Euler) SetOrder(value EulerOrder) {

	e.Order = value
	e.OnChangeCallback()

}

func (e *Euler) Set(x, y, z float64, order EulerOrder) *Euler {

	e.x = x
	e.y = y
	e.z = z
	if order != CurrentOrder {
		e.Order = order
	}
	e.OnChangeCallback()

	return e

}

func (e *Euler) Clone() *Euler {

	return NewEuler().Copy(e)

}

func (e *Euler) Copy(src *Euler) *Euler {

	e.x = src.x
	e.y = src.y
	e.z = src.z
	e.Order = src.Order

	e.OnChangeCallback()

	return e

}

func (e *Euler) SetFromRotationMatrix(
	m *Matrix4, order EulerOrder, update bool) *Euler {

	// assumes the upper 3x3 of m is a pure rotation matrix (i.E, unscaled)

	te := m.Elements
	m11 := te[0]
	m12 := te[4]
	m13 := te[8]
	m21 := te[1]
	m22 := te[5]
	m23 := te[9]
	m31 := te[2]
	m32 := te[6]
	m33 := te[10]

	if order == CurrentOrder {
		order = e.Order
	}

	if order == XYZ {

		e.y = math.Asin(Clamp(m13, -1, 1))

		if math.Abs(m13) < 0.99999 {

			e.x = math.Atan2(-m23, m33)
			e.z = math.Atan2(-m12, m11)

		} else {

			e.x = math.Atan2(m32, m22)
			e.z = 0

		}

	} else if order == YXZ {

		e.x = math.Asin(-Clamp(m23, -1, 1))

		if math.Abs(m23) < 0.99999 {

			e.y = math.Atan2(m13, m33)
			e.z = math.Atan2(m21, m22)

		} else {

			e.y = math.Atan2(-m31, m11)
			e.z = 0

		}

	} else if order == ZXY {

		e.x = math.Asin(Clamp(m32, -1, 1))

		if math.Abs(m32) < 0.99999 {

			e.y = math.Atan2(-m31, m33)
			e.z = math.Atan2(-m12, m22)

		} else {

			e.y = 0
			e.z = math.Atan2(m21, m11)

		}

	} else if order == ZYX {

		e.y = math.Asin(-Clamp(m31, -1, 1))

		if math.Abs(m31) < 0.99999 {

			e.x = math.Atan2(m32, m33)
			e.z = math.Atan2(m21, m11)

		} else {

			e.x = 0
			e.z = math.Atan2(-m12, m22)

		}

	} else if order == YZX {

		e.z = math.Asin(Clamp(m21, -1, 1))

		if math.Abs(m21) < 0.99999 {

			e.x = math.Atan2(-m23, m22)
			e.y = math.Atan2(-m31, m11)

		} else {

			e.x = 0
			e.y = math.Atan2(m13, m33)

		}

	} else if order == XZY {

		e.z = math.Asin(-Clamp(m12, -1, 1))

		if math.Abs(m12) < 0.99999 {

			e.x = math.Atan2(m32, m22)
			e.y = math.Atan2(m13, m11)

		} else {

			e.x = math.Atan2(-m23, m33)
			e.y = 0

		}

	} else {

		glog.Warningf("three.Euler: SetFromRotationMatrix() given unsupported order: %s", string(order))

	}

	e.Order = order

	if update != false {
		e.OnChangeCallback()
	}

	return e

}

func (e *Euler) SetFromQuaternion(
	q *Quaternion, order EulerOrder, update bool) *Euler {

	matrix := NewMatrix4()

	matrix.MakeRotationFromQuaternion(q)

	return e.SetFromRotationMatrix(matrix, order, update)

}

func (e *Euler) SetFromVector3(v *Vector3, order EulerOrder) *Euler {

	return e.Set(v.X, v.Y, v.Z, order)

}

func (e *Euler) Reorder(newOrder EulerOrder) *Euler {

	// WARNING: e discards revolution information -bhouston

	q := NewQuaternion()

	q.SetFromEuler(e, false)

	return e.SetFromQuaternion(q, newOrder, false)

}

func (e *Euler) Equals(other *Euler) bool {

	return (e.x == other.x) && (e.y == other.y) && (e.z == other.z) && (e.Order == other.Order)

}

func (e *Euler) FromArray(array []float64, order EulerOrder) *Euler {

	e.x = array[0]
	e.y = array[1]
	e.z = array[2]
	if order == CurrentOrder {
		e.Order = order
	}

	e.OnChangeCallback()

	return e

}

func (e *Euler) ToArray(array []float64, offset int) ([]float64, EulerOrder) {

	if array == nil {
		array = make([]float64, offset+3)
	}

	array[offset] = e.x
	array[offset+1] = e.y
	array[offset+2] = e.z

	return array, e.Order

}

func (e *Euler) ToVector3(target *Vector3) *Vector3 {

	if target == nil {
		target = NewVector3()
	}
	return target.Set(e.x, e.y, e.z)

}

func (e *Euler) OnChange(callback EulerChangeCallback) *Euler {

	e.OnChangeCallback = callback

	return e

}

func EulerEmptyCallback() {}
