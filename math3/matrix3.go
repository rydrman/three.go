package math3

type Matrix3 struct {
	Elements []float64
}

func NewMatrix3() *Matrix3 {

	return &Matrix3{
		[]float64{
			1, 0, 0,
			0, 1, 0,
			0, 0, 1,
		},
	}

}

func (m *Matrix3) Set(n11, n12, n13, n21, n22, n23, n31, n32, n33 float64) *Matrix3 {

	te := m.Elements

	te[0], te[1], te[2] = n11, n21, n31
	te[3], te[4], te[5] = n12, n22, n32
	te[6], te[7], te[8] = n13, n23, n33

	return m

}

func (m *Matrix3) Identity() *Matrix3 {

	m.Set(

		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	)

	return m

}

func (m *Matrix3) Clone() *Matrix3 {

	return NewMatrix3().FromArray(m.Elements)

}

func (m *Matrix3) Copy(src *Matrix3) *Matrix3 {

	me := src.Elements

	m.Set(

		me[0], me[3], me[6],
		me[1], me[4], me[7],
		me[2], me[5], me[8],
	)

	return m

}

func (m *Matrix3) SetFromMatrix4(other *Matrix4) *Matrix3 {

	me := other.Elements

	m.Set(

		me[0], me[4], me[8],
		me[1], me[5], me[9],
		me[2], me[6], me[10],
	)

	return m

}

func (m *Matrix3) ApplyToVector3Array(array []float64, offset int, length int) []float64 {

	v1 := NewVector3()

	if offset < 0 {
		offset = 0
	}
	if length < 0 {
		length = len(array)
	}

	for i, j := 0, offset; i < length; i, j = i+3, j+3 {

		v1.FromArray(array, j)
		v1.ApplyMatrix3(m)
		v1.ToArray(array, j)

	}

	return array

}

/*func (m *Matrix3) ApplyToBuffer() {

    var v1

    return function applyToBuffer( buffer, offset, length ) {

        if v1 == nil { v1 = NewVector3() }
        if offset == nil { offset = 0 }
        if length == nil { length = len(buffer) / buffer.ItemSize }

        for ( i := 0, j = offset i < length i ++, j ++ ) {

            v1.X = buffer.GetX( j )
            v1.Y = buffer.GetY( j )
            v1.Z = buffer.GetZ( j )

            v1.ApplyMatrix3( mat )

            buffer.SetXYZ( v1.X, v1.Y, v1.Z )

        }

        return buffer

    }

}(),*/

func (m *Matrix3) MultiplyScalar(s float64) *Matrix3 {

	te := m.Elements

	te[0] *= s
	te[3] *= s
	te[6] *= s
	te[1] *= s
	te[4] *= s
	te[7] *= s
	te[2] *= s
	te[5] *= s
	te[8] *= s

	return m

}

func (m *Matrix3) Determinant() float64 {

	te := m.Elements

	a, b, c := te[0], te[1], te[2]
	d, e, f := te[3], te[4], te[5]
	g, h, i := te[6], te[7], te[8]

	return a*e*i - a*f*h - b*d*i + b*f*g + c*d*h - c*e*g

}

func (m *Matrix3) GetInverse(matrix *Matrix3) *Matrix3 {

	defer func() {
		if r := recover(); r != nil {
			m.Identity()
		}
	}()

	m.MustGetInverse(matrix)

	return m
}

func (m *Matrix3) MustGetInverse(matrix *Matrix3) *Matrix3 {

	me := matrix.Elements
	te := m.Elements

	n11, n21, n31 := me[0], me[1], me[2]
	n12, n22, n32 := me[3], me[4], me[5]
	n13, n23, n33 := me[6], me[7], me[8]

	t11 := n33*n22 - n32*n23
	t12 := n32*n13 - n33*n12
	t13 := n23*n12 - n22*n13

	det := n11*t11 + n21*t12 + n31*t13

	if det == 0 {
		panic("Matrix3.GetInverse(): can't invert matrix, determinant is 0")
	}

	detInv := 1 / det

	te[0] = t11 * detInv
	te[1] = (n31*n23 - n33*n21) * detInv
	te[2] = (n32*n21 - n31*n22) * detInv

	te[3] = t12 * detInv
	te[4] = (n33*n11 - n31*n13) * detInv
	te[5] = (n31*n12 - n32*n11) * detInv

	te[6] = t13 * detInv
	te[7] = (n21*n13 - n23*n11) * detInv
	te[8] = (n22*n11 - n21*n12) * detInv

	return m

}

func (m *Matrix3) Transpose() *Matrix3 {

	var tmp float64
	te := m.Elements

	tmp = te[1]
	te[1] = te[3]
	te[3] = tmp
	tmp = te[2]
	te[2] = te[6]
	te[6] = tmp
	tmp = te[5]
	te[5] = te[7]
	te[7] = tmp

	return m

}

func (m *Matrix3) GetNormalMatrix(matrix4 *Matrix4) *Matrix3 {

	return m.SetFromMatrix4(matrix4).GetInverse(m).Transpose()

}

func (m *Matrix3) TransposeIntoArray(r []float64) *Matrix3 {

	te := m.Elements

	r[0] = te[0]
	r[1] = te[3]
	r[2] = te[6]
	r[3] = te[1]
	r[4] = te[4]
	r[5] = te[7]
	r[6] = te[2]
	r[7] = te[5]
	r[8] = te[8]

	return m

}

func (m *Matrix3) FromArray(array []float64) *Matrix3 {

	copy(m.Elements, array)

	return m

}

func (m *Matrix3) ToArray(array []float64, offset int) []float64 {

	if array == nil {
		array = make([]float64, 0)
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

	return array

}
