package math3

import "math"

type Matrix4 struct {
	Elements []float64
}

func NewMatrix4() *Matrix4 {

	return &Matrix4{[]float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}}

}

func (m *Matrix4) Set(n11, n12, n13, n14, n21, n22, n23, n24, n31, n32, n33, n34, n41, n42, n43, n44 float64) *Matrix4 {

	te := m.Elements

	te[0], te[4], te[8], te[12] = n11, n12, n13, n14
	te[1], te[5], te[9], te[13] = n21, n22, n23, n24
	te[2], te[6], te[10], te[14] = n31, n32, n33, n34
	te[3], te[7], te[11], te[15] = n41, n42, n43, n44

	return m

}

func (m *Matrix4) Identity() *Matrix4 {

	return m.Set(

		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)

}

func (m *Matrix4) Clone() *Matrix4 {

	return NewMatrix4().Copy(m)

}

func (m *Matrix4) Copy(other *Matrix4) *Matrix4 {

	copy(m.Elements, other.Elements)

	return m

}

func (m *Matrix4) CopyPosition(other *Matrix4) *Matrix4 {

	te := m.Elements
	me := other.Elements

	te[12] = me[12]
	te[13] = me[13]
	te[14] = me[14]

	return m

}

func (m *Matrix4) ExtractBasis(xAxis, yAxis, zAxis *Vector3) *Matrix4 {

	xAxis.SetFromMatrixColumn(m, 0)
	yAxis.SetFromMatrixColumn(m, 1)
	zAxis.SetFromMatrixColumn(m, 2)

	return m

}

func (m *Matrix4) MakeBasis(xAxis, yAxis, zAxis *Vector3) *Matrix4 {

	return m.Set(
		xAxis.X, yAxis.X, zAxis.X, 0,
		xAxis.Y, yAxis.Y, zAxis.Y, 0,
		xAxis.Z, yAxis.Z, zAxis.Z, 0,
		0, 0, 0, 1,
	)

}

func (m *Matrix4) ExtractRotation(target *Matrix4) *Matrix4 {

	v1 := NewVector3()

	te := m.Elements
	me := target.Elements

	scaleX := 1 / v1.SetFromMatrixColumn(target, 0).Length()
	scaleY := 1 / v1.SetFromMatrixColumn(target, 1).Length()
	scaleZ := 1 / v1.SetFromMatrixColumn(target, 2).Length()

	te[0] = me[0] * scaleX
	te[1] = me[1] * scaleX
	te[2] = me[2] * scaleX

	te[4] = me[4] * scaleY
	te[5] = me[5] * scaleY
	te[6] = me[6] * scaleY

	te[8] = me[8] * scaleZ
	te[9] = me[9] * scaleZ
	te[10] = me[10] * scaleZ

	return target

}

func (m *Matrix4) MakeRotationFromEuler(euler *Euler) *Matrix4 {

	te := m.Elements

	x, y, z := euler.GetX(), euler.GetY(), euler.GetZ()
	a, b := math.Cos(x), math.Sin(x)
	c, d := math.Cos(y), math.Sin(y)
	e, f := math.Cos(z), math.Sin(z)

	order := euler.Order

	if order == XYZ {

		ae := a * e
		af := a * f
		be := b * e
		bf := b * f

		te[0] = c * e
		te[4] = -c * f
		te[8] = d

		te[1] = af + be*d
		te[5] = ae - bf*d
		te[9] = -b * c

		te[2] = bf - ae*d
		te[6] = be + af*d
		te[10] = a * c

	} else if order == YXZ {

		ce := c * e
		cf := c * f
		de := d * e
		df := d * f

		te[0] = ce + df*b
		te[4] = de*b - cf
		te[8] = a * d

		te[1] = a * f
		te[5] = a * e
		te[9] = -b

		te[2] = cf*b - de
		te[6] = df + ce*b
		te[10] = a * c

	} else if order == ZXY {

		ce := c * e
		cf := c * f
		de := d * e
		df := d * f

		te[0] = ce - df*b
		te[4] = -a * f
		te[8] = de + cf*b

		te[1] = cf + de*b
		te[5] = a * e
		te[9] = df - ce*b

		te[2] = -a * d
		te[6] = b
		te[10] = a * c

	} else if order == ZYX {

		ae := a * e
		af := a * f
		be := b * e
		bf := b * f

		te[0] = c * e
		te[4] = be*d - af
		te[8] = ae*d + bf

		te[1] = c * f
		te[5] = bf*d + ae
		te[9] = af*d - be

		te[2] = -d
		te[6] = b * c
		te[10] = a * c

	} else if order == YZX {

		ac := a * c
		ad := a * d
		bc := b * c
		bd := b * d

		te[0] = c * e
		te[4] = bd - ac*f
		te[8] = bc*f + ad

		te[1] = f
		te[5] = a * e
		te[9] = -b * e

		te[2] = -d * e
		te[6] = ad*f + bc
		te[10] = ac - bd*f

	} else if order == XZY {

		ac := a * c
		ad := a * d
		bc := b * c
		bd := b * d

		te[0] = c * e
		te[4] = -f
		te[8] = d * e

		te[1] = ac*f + bd
		te[5] = a * e
		te[9] = ad*f - bc

		te[2] = bc*f - ad
		te[6] = b * e
		te[10] = bd*f + ac

	}

	// last column
	te[3] = 0
	te[7] = 0
	te[11] = 0

	// bottom row
	te[12] = 0
	te[13] = 0
	te[14] = 0
	te[15] = 1

	return m

}

func (m *Matrix4) MakeRotationFromQuaternion(q *Quaternion) *Matrix4 {

	te := m.Elements

	x := q.GetX()
	y := q.GetY()
	z := q.GetZ()
	w := q.GetW()
	x2 := x + x
	y2 := y + y
	z2 := z + z
	xx := x * x2
	xy := x * y2
	xz := x * z2
	yy := y * y2
	yz := y * z2
	zz := z * z2
	wx := w * x2
	wy := w * y2
	wz := w * z2

	te[0] = 1 - (yy + zz)
	te[4] = xy - wz
	te[8] = xz + wy

	te[1] = xy + wz
	te[5] = 1 - (xx + zz)
	te[9] = yz - wx

	te[2] = xz - wy
	te[6] = yz + wx
	te[10] = 1 - (xx + yy)

	// last column
	te[3] = 0
	te[7] = 0
	te[11] = 0

	// bottom row
	te[12] = 0
	te[13] = 0
	te[14] = 0
	te[15] = 1

	return m

}

func (m *Matrix4) LookAt(eye, target, up *Vector3) *Matrix4 {

	x := NewVector3()
	y := NewVector3()
	z := NewVector3()

	te := m.Elements

	z.SubVectors(eye, target).Normalize()

	if z.LengthSq() == 0 {

		z.Z = 1

	}

	x.CrossVectors(up, z).Normalize()

	if x.LengthSq() == 0 {

		z.Z += 0.0001
		x.CrossVectors(up, z).Normalize()

	}

	y.CrossVectors(z, x)

	te[0] = x.X
	te[4] = y.X
	te[8] = z.X
	te[1] = x.Y
	te[5] = y.Y
	te[9] = z.Y
	te[2] = x.Z
	te[6] = y.Z
	te[10] = z.Z

	return m

}

func (m *Matrix4) Multiply(other *Matrix4) *Matrix4 {

	return m.MultiplyMatrices(m, other)

}

func (m *Matrix4) Premultiply(other *Matrix4) *Matrix4 {

	return m.MultiplyMatrices(other, m)

}

func (m *Matrix4) MultiplyMatrices(a, b *Matrix4) *Matrix4 {

	ae := a.Elements
	be := b.Elements
	te := m.Elements

	a11, a12, a13, a14 := ae[0], ae[4], ae[8], ae[12]
	a21, a22, a23, a24 := ae[1], ae[5], ae[9], ae[13]
	a31, a32, a33, a34 := ae[2], ae[6], ae[10], ae[14]
	a41, a42, a43, a44 := ae[3], ae[7], ae[11], ae[15]

	b11, b12, b13, b14 := be[0], be[4], be[8], be[12]
	b21, b22, b23, b24 := be[1], be[5], be[9], be[13]
	b31, b32, b33, b34 := be[2], be[6], be[10], be[14]
	b41, b42, b43, b44 := be[3], be[7], be[11], be[15]

	te[0] = a11*b11 + a12*b21 + a13*b31 + a14*b41
	te[4] = a11*b12 + a12*b22 + a13*b32 + a14*b42
	te[8] = a11*b13 + a12*b23 + a13*b33 + a14*b43
	te[12] = a11*b14 + a12*b24 + a13*b34 + a14*b44

	te[1] = a21*b11 + a22*b21 + a23*b31 + a24*b41
	te[5] = a21*b12 + a22*b22 + a23*b32 + a24*b42
	te[9] = a21*b13 + a22*b23 + a23*b33 + a24*b43
	te[13] = a21*b14 + a22*b24 + a23*b34 + a24*b44

	te[2] = a31*b11 + a32*b21 + a33*b31 + a34*b41
	te[6] = a31*b12 + a32*b22 + a33*b32 + a34*b42
	te[10] = a31*b13 + a32*b23 + a33*b33 + a34*b43
	te[14] = a31*b14 + a32*b24 + a33*b34 + a34*b44

	te[3] = a41*b11 + a42*b21 + a43*b31 + a44*b41
	te[7] = a41*b12 + a42*b22 + a43*b32 + a44*b42
	te[11] = a41*b13 + a42*b23 + a43*b33 + a44*b43
	te[15] = a41*b14 + a42*b24 + a43*b34 + a44*b44

	return m

}

func (m *Matrix4) MultiplyToArray(a, b *Matrix4, target []float64) []float64 {

	te := m.Elements

	if nil == target {
		target = make([]float64, 16)
	}

	m.MultiplyMatrices(a, b)

	copy(target, te)

	return target

}

func (m *Matrix4) MultiplyScalar(s float64) *Matrix4 {

	te := m.Elements

	te[0] *= s
	te[4] *= s
	te[8] *= s
	te[12] *= s
	te[1] *= s
	te[5] *= s
	te[9] *= s
	te[13] *= s
	te[2] *= s
	te[6] *= s
	te[10] *= s
	te[14] *= s
	te[3] *= s
	te[7] *= s
	te[11] *= s
	te[15] *= s

	return m

}

//TODO
/*func (m *Matrix4) ApplyToVector3Array(array []*Vector3, offset, length int) []*Vector3 {

    v1 := NewVector3()

    if length == -1 {
        length = len(array)
    }

    for i, j := 0, offset; i < length; i, j = i+3, j+3 {

        v1.FromArray(array, j)
        v1.ApplyMatrix4(m)
        v1.ToArray(array, j)

    }

    return array

}*/

//TODO
/*func (m *Matrix4) ApplyToBuffer() *Matrix4 {

    var v1

    return function applyToBuffer( buffer, offset, length ) *Matrix4 {

        if v1 == nil { 1 = NewVector3().Set(); }
        if offset == nil { ffset = 0; }
        if length == nil { ength = buffer.Length / buffer.ItemSize; }

        for  i := 0; j := offset; i < length; i ++, j ++  {

            v1.X = buffer.GetX( j )
            v1.Y = buffer.GetY( j )
            v1.Z = buffer.GetZ( j )

            v1.ApplyMatrix4( m )

            buffer.SetXYZ( v1.X, v1.Y, v1.Z )

        }

        return buffer

    }

}(),*/

func (m *Matrix4) Determinant() float64 {

	te := m.Elements

	n11, n12, n13, n14 := te[0], te[4], te[8], te[12]
	n21, n22, n23, n24 := te[1], te[5], te[9], te[13]
	n31, n32, n33, n34 := te[2], te[6], te[10], te[14]
	n41, n42, n43, n44 := te[3], te[7], te[11], te[15]

	return (n41*(+n14*n23*n32-n13*n24*n32-n14*n22*n33+n12*n24*n33+n13*n22*n34-n12*n23*n34) +
		n42*(+n11*n23*n34-n11*n24*n33+n14*n21*n33-n13*n21*n34+n13*n24*n31-n14*n23*n31) +
		n43*(+n11*n24*n32-n11*n22*n34-n14*n21*n32+n12*n21*n34+n14*n22*n31-n12*n24*n31) +
		n44*(-n13*n22*n31-n11*n23*n32+n11*n22*n33+n13*n21*n32-n12*n21*n33+n12*n23*n31))

}

func (m *Matrix4) Transpose() *Matrix4 {

	te := m.Elements
	var tmp float64

	tmp = te[1]
	te[1] = te[4]
	te[4] = tmp
	tmp = te[2]
	te[2] = te[8]
	te[8] = tmp
	tmp = te[6]
	te[6] = te[9]
	te[9] = tmp

	tmp = te[3]
	te[3] = te[12]
	te[12] = tmp
	tmp = te[7]
	te[7] = te[13]
	te[13] = tmp
	tmp = te[11]
	te[11] = te[14]
	te[14] = tmp

	return m

}

func (m *Matrix4) SetPosition(v *Vector3) *Matrix4 {

	te := m.Elements

	te[12] = v.X
	te[13] = v.Y
	te[14] = v.Z

	return m

}

func (m *Matrix4) GetInverse(src *Matrix4) *Matrix4 {

	defer func() {
		if r := recover(); r != nil {
			m.Identity()
		}
	}()

	m.MustGetInverse(src)

	return m

}

func (m Matrix4) MustGetInverse(src *Matrix4) *Matrix4 {

	if nil == src {
		src = NewMatrix4()
	}

	// based on http://www.Euclideanspace.Com/maths/algebra/matrix/functions/inverse/fourD/index.Htm
	me := m.Elements
	te := src.Elements

	n11, n21, n31, n41 := me[0], me[1], me[2], me[3]
	n12, n22, n32, n42 := me[4], me[5], me[6], me[7]
	n13, n23, n33, n43 := me[8], me[9], me[10], me[11]
	n14, n24, n34, n44 := me[12], me[13], me[14], me[15]

	t11 := n23*n34*n42 - n24*n33*n42 + n24*n32*n43 - n22*n34*n43 - n23*n32*n44 + n22*n33*n44
	t12 := n14*n33*n42 - n13*n34*n42 - n14*n32*n43 + n12*n34*n43 + n13*n32*n44 - n12*n33*n44
	t13 := n13*n24*n42 - n14*n23*n42 + n14*n22*n43 - n12*n24*n43 - n13*n22*n44 + n12*n23*n44
	t14 := n14*n23*n32 - n13*n24*n32 - n14*n22*n33 + n12*n24*n33 + n13*n22*n34 - n12*n23*n34

	det := n11*t11 + n21*t12 + n31*t13 + n41*t14

	if det == 0 {

		panic("Matrix4.MustGetInverse(): can't invert matrix, determinant is 0")

	}

	detInv := 1 / det

	te[0] = t11 * detInv
	te[1] = (n24*n33*n41 - n23*n34*n41 - n24*n31*n43 + n21*n34*n43 + n23*n31*n44 - n21*n33*n44) * detInv
	te[2] = (n22*n34*n41 - n24*n32*n41 + n24*n31*n42 - n21*n34*n42 - n22*n31*n44 + n21*n32*n44) * detInv
	te[3] = (n23*n32*n41 - n22*n33*n41 - n23*n31*n42 + n21*n33*n42 + n22*n31*n43 - n21*n32*n43) * detInv

	te[4] = t12 * detInv
	te[5] = (n13*n34*n41 - n14*n33*n41 + n14*n31*n43 - n11*n34*n43 - n13*n31*n44 + n11*n33*n44) * detInv
	te[6] = (n14*n32*n41 - n12*n34*n41 - n14*n31*n42 + n11*n34*n42 + n12*n31*n44 - n11*n32*n44) * detInv
	te[7] = (n12*n33*n41 - n13*n32*n41 + n13*n31*n42 - n11*n33*n42 - n12*n31*n43 + n11*n32*n43) * detInv

	te[8] = t13 * detInv
	te[9] = (n14*n23*n41 - n13*n24*n41 - n14*n21*n43 + n11*n24*n43 + n13*n21*n44 - n11*n23*n44) * detInv
	te[10] = (n12*n24*n41 - n14*n22*n41 + n14*n21*n42 - n11*n24*n42 - n12*n21*n44 + n11*n22*n44) * detInv
	te[11] = (n13*n22*n41 - n12*n23*n41 - n13*n21*n42 + n11*n23*n42 + n12*n21*n43 - n11*n22*n43) * detInv

	te[12] = t14 * detInv
	te[13] = (n13*n24*n31 - n14*n23*n31 + n14*n21*n33 - n11*n24*n33 - n13*n21*n34 + n11*n23*n34) * detInv
	te[14] = (n14*n22*n31 - n12*n24*n31 - n14*n21*n32 + n11*n24*n32 + n12*n21*n34 - n11*n22*n34) * detInv
	te[15] = (n12*n23*n31 - n13*n22*n31 + n13*n21*n32 - n11*n23*n32 - n12*n21*n33 + n11*n22*n33) * detInv

	return src

}

func (m *Matrix4) Scale(v *Vector3) *Matrix4 {

	te := m.Elements
	x := v.X
	y := v.Y
	z := v.Z

	te[0] *= x
	te[4] *= y
	te[8] *= z
	te[1] *= x
	te[5] *= y
	te[9] *= z
	te[2] *= x
	te[6] *= y
	te[10] *= z
	te[3] *= x
	te[7] *= y
	te[11] *= z

	return m

}

func (m *Matrix4) GetMaxScaleOnAxis() float64 {

	te := m.Elements

	scaleXSq := te[0]*te[0] + te[1]*te[1] + te[2]*te[2]
	scaleYSq := te[4]*te[4] + te[5]*te[5] + te[6]*te[6]
	scaleZSq := te[8]*te[8] + te[9]*te[9] + te[10]*te[10]

	return math.Sqrt(Max(scaleXSq, scaleYSq, scaleZSq))

}

func (m *Matrix4) MakeTranslation(x, y, z float64) *Matrix4 {

	return m.Set(

		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	)

}

func (m *Matrix4) MakeRotationX(theta float64) *Matrix4 {

	c := math.Cos(theta)
	s := math.Sin(theta)

	return m.Set(

		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	)

}

func (m *Matrix4) MakeRotationY(theta float64) *Matrix4 {

	c := math.Cos(theta)
	s := math.Sin(theta)

	return m.Set(

		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	)

}

func (m *Matrix4) MakeRotationZ(theta float64) *Matrix4 {

	c := math.Cos(theta)
	s := math.Sin(theta)

	return m.Set(

		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)

}

func (m *Matrix4) MakeRotationAxis(axis *Vector3, angle float64) *Matrix4 {

	// Based on http://www.Gamedev.Net/reference/articles/article1199.Asp

	c := math.Cos(angle)
	s := math.Sin(angle)
	t := 1 - c
	x := axis.X
	y := axis.Y
	z := axis.Z
	tx := t * x
	ty := t * y

	return m.Set(

		tx*x+c, tx*y-s*z, tx*z+s*y, 0,
		tx*y+s*z, ty*y+c, ty*z-s*x, 0,
		tx*z-s*y, ty*z+s*x, t*z*z+c, 0,
		0, 0, 0, 1,
	)

}

func (m *Matrix4) MakeScale(x, y, z float64) *Matrix4 {

	return m.Set(

		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	)

}

//TODO
/*func (m *Matrix4) Compose(position *Vector3, quaternion *Quaternion, scale *Vector3) *Matrix4 {

    m.MakeRotationFromQuaternion(quaternion)
    m.Scale(scale)
    m.SetPosition(position)

}*/

//TODO
/*func (m *Matrix4) Decompose() (position *Vector3, quaternion *Quaternion, scale *Vector3) *Matrix4 {

    vector := NewVector3()
    matrix := NewMatrix4()

    te := m.Elements

    sx := vector.Set(te[0], te[1], te[2]).Length()
    sy := vector.Set(te[4], te[5], te[6]).Length()
    sz := vector.Set(te[8], te[9], te[10]).Length()

    // if determine is negative, we need to invert one scale
    det := m.Determinant()
    if det < 0 {

        sx = -sx

    }

    position.X = te[12]
    position.Y = te[13]
    position.Z = te[14]

    // scale the rotation part

    copy(m.Elements, matrix.Elements) // at this point matrix is incomplete so we can"t use .Copy()

    invSX := 1 / sx
    invSY := 1 / sy
    invSZ := 1 / sz

    matrix.Elements[0] *= invSX
    matrix.Elements[1] *= invSX
    matrix.Elements[2] *= invSX

    matrix.Elements[4] *= invSY
    matrix.Elements[5] *= invSY
    matrix.Elements[6] *= invSY

    matrix.Elements[8] *= invSZ
    matrix.Elements[9] *= invSZ
    matrix.Elements[10] *= invSZ

    quaternion.SetFromRotationMatrix(matrix)

    scale.X = sx
    scale.Y = sy
    scale.Z = sz

    return

}*/

func (m *Matrix4) MakeFrustum(left, right, bottom, top, near, far float64) *Matrix4 {

	te := m.Elements
	x := 2 * near / (right - left)
	y := 2 * near / (top - bottom)

	a := (right + left) / (right - left)
	b := (top + bottom) / (top - bottom)
	c := -(far + near) / (far - near)
	d := -2 * far * near / (far - near)

	te[0] = x
	te[4] = 0
	te[8] = a
	te[12] = 0
	te[1] = 0
	te[5] = y
	te[9] = b
	te[13] = 0
	te[2] = 0
	te[6] = 0
	te[10] = c
	te[14] = d
	te[3] = 0
	te[7] = 0
	te[11] = -1
	te[15] = 0

	return m

}

func (m *Matrix4) MakePerspective(fov, aspect, near, far float64) *Matrix4 {

	ymax := near * math.Tan(Deg2Rad*fov*0.5)
	ymin := -ymax
	xmin := ymin * aspect
	xmax := ymax * aspect

	return m.MakeFrustum(xmin, xmax, ymin, ymax, near, far)

}

func (m *Matrix4) MakeOrthographic(left, right, top, bottom, near, far float64) *Matrix4 {

	te := m.Elements
	w := 1.0 / (right - left)
	h := 1.0 / (top - bottom)
	p := 1.0 / (far - near)

	x := (right + left) * w
	y := (top + bottom) * h
	z := (far + near) * p

	te[0] = 2 * w
	te[4] = 0
	te[8] = 0
	te[12] = -x
	te[1] = 0
	te[5] = 2 * h
	te[9] = 0
	te[13] = -y
	te[2] = 0
	te[6] = 0
	te[10] = -2 * p
	te[14] = -z
	te[3] = 0
	te[7] = 0
	te[11] = 0
	te[15] = 1

	return m

}

func (m *Matrix4) Equals(matrix *Matrix4) bool {

	te := m.Elements
	me := matrix.Elements

	for i := 0; i < 16; i++ {

		if te[i] != me[i] {
			return false
		}

	}

	return true

}

func (m *Matrix4) FromArray(array []float64) *Matrix4 {

	copy(array, m.Elements)

	return m

}

func (m Matrix4) ToArray(array []float64, offset int) []float64 {

	if array == nil {
		array = make([]float64, 16)
	}

	te := m.Elements

	array[offset] = te[0]
	array[offset+1] = te[1]
	array[offset+2] = te[2]
	array[offset+3] = te[3]

	array[offset+4] = te[4]
	array[offset+5] = te[5]
	array[offset+6] = te[6]
	array[offset+7] = te[7]

	array[offset+8] = te[8]
	array[offset+9] = te[9]
	array[offset+10] = te[10]
	array[offset+11] = te[11]

	array[offset+12] = te[12]
	array[offset+13] = te[13]
	array[offset+14] = te[14]
	array[offset+15] = te[15]

	return array

}

func (m Matrix4) ToArray32(array []float32, offset int) []float32 {

	if array == nil {
		array = make([]float32, 16)
	}

	te := m.Elements

	array[offset] = float32(te[0])
	array[offset+1] = float32(te[1])
	array[offset+2] = float32(te[2])
	array[offset+3] = float32(te[3])

	array[offset+4] = float32(te[4])
	array[offset+5] = float32(te[5])
	array[offset+6] = float32(te[6])
	array[offset+7] = float32(te[7])

	array[offset+8] = float32(te[8])
	array[offset+9] = float32(te[9])
	array[offset+10] = float32(te[10])
	array[offset+11] = float32(te[11])

	array[offset+12] = float32(te[12])
	array[offset+13] = float32(te[13])
	array[offset+14] = float32(te[14])
	array[offset+15] = float32(te[15])

	return array

}
