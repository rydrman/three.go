package math3_test

import (
	"math"
	"testing"

	"github.com/rydrman/marshmallow"
)

func matrixEquals4(a, b *mm.Matrix4) bool {
	tolerance := 0.0001
	for i, ea := range a.Elements {
		delta := ea - b.Elements[i]
		if delta > tolerance {
			return false
		}
	}
	return true
}

func TestNewMatrix4(t *testing.T) {
	m := mm.NewMatrix4()
	if m == nil {
		t.Error("expected non-nil value")
		return
	}
	if 1 != m.Determinant() {
		t.Error("expected identity matrix")
	}
}

func TestMatrix4_Set(t *testing.T) {
	m := mm.NewMatrix4().Set(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	if 0 != m.Elements[0] ||
		4 != m.Elements[1] ||
		8 != m.Elements[2] ||
		12 != m.Elements[3] ||
		1 != m.Elements[4] ||
		5 != m.Elements[5] ||
		9 != m.Elements[6] ||
		13 != m.Elements[7] ||
		2 != m.Elements[8] ||
		6 != m.Elements[9] ||
		10 != m.Elements[10] ||
		14 != m.Elements[11] ||
		3 != m.Elements[12] ||
		7 != m.Elements[13] ||
		11 != m.Elements[14] ||
		15 != m.Elements[15] {
		t.Errorf("set did not return valid element values")
	}
}

func TestMatrix4_Copy(t *testing.T) {

	a := mm.NewMatrix4().Set(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	b := mm.NewMatrix4().Copy(a)

	if !matrixEquals4(a, b) {
		t.Errorf("copied matrix does not equal source")
	}

	// ensure that it is a true copy
	a.Elements[0] = 2
	if matrixEquals4(a, b) {
		t.Errorf("matrix copy did not create a tru copy")
	}

}

func TestMatrix4_Identity(t *testing.T) {

	m := mm.NewMatrix4().Set(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	m.Identity()
	if 1 != m.Elements[0] ||
		0 != m.Elements[1] ||
		0 != m.Elements[2] ||
		0 != m.Elements[3] ||

		0 != m.Elements[4] ||
		1 != m.Elements[5] ||
		0 != m.Elements[6] ||
		0 != m.Elements[7] ||

		0 != m.Elements[8] ||
		0 != m.Elements[9] ||
		1 != m.Elements[10] ||
		0 != m.Elements[11] ||

		0 != m.Elements[12] ||
		0 != m.Elements[13] ||
		0 != m.Elements[14] ||
		1 != m.Elements[15] {
		t.Errorf("identity did not return valid identity matrix")
	}

}

func TestMatrix4_MuliplyScalar(t *testing.T) {

	m := mm.NewMatrix4().Set(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)

	m.MultiplyScalar(2)
	if m.Elements[0] != 0*2 ||
		m.Elements[1] != 4*2 ||
		m.Elements[2] != 8*2 ||
		m.Elements[3] != 12*2 ||
		m.Elements[4] != 1*2 ||
		m.Elements[5] != 5*2 ||
		m.Elements[6] != 9*2 ||
		m.Elements[7] != 13*2 ||
		m.Elements[8] != 2*2 ||
		m.Elements[9] != 6*2 ||
		m.Elements[10] != 10*2 ||
		m.Elements[11] != 14*2 ||
		m.Elements[12] != 3*2 ||
		m.Elements[13] != 7*2 ||
		m.Elements[14] != 11*2 ||
		m.Elements[15] != 15*2 {
		t.Error("elements do not have correct values")
	}

}

func TestMatrix4_Determinant(t *testing.T) {
	a := mm.NewMatrix4()
	if a.Determinant() != 1 {
		t.Error("determinant of identity should be 1")
	}

	a.Elements[0] = 2
	if a.Determinant() != 2 {
		t.Error("invalid determinant value")
	}

	a.Elements[0] = 0
	if a.Determinant() != 0 {
		t.Error("invalid determinant value")
	}

	a.Set(2, 3, 4, 5, -1, -21, -3, -4, 6, 7, 8, 10, -8, -9, -10, -12)
	if a.Determinant() != 76 {
		t.Error("invalid determinant value")
	}
}

func TestMatrix4_GetInverse(t *testing.T) {
	identity := mm.NewMatrix4()

	a := mm.NewMatrix4()
	b := mm.NewMatrix4().Set(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	if matrixEquals4(a, b) {
		t.Error("b should not be an identity matrix yet")
	}
	b.GetInverse(a)
	if !matrixEquals4(b, mm.NewMatrix4()) {
		t.Error("b should have become an identity matrix")
	}

	testMatrices := []*mm.Matrix4{
		mm.NewMatrix4().MakeRotationX(0.3),
		mm.NewMatrix4().MakeRotationX(-0.3),
		mm.NewMatrix4().MakeRotationY(0.3),
		mm.NewMatrix4().MakeRotationY(-0.3),
		mm.NewMatrix4().MakeRotationZ(0.3),
		mm.NewMatrix4().MakeRotationZ(-0.3),
		mm.NewMatrix4().MakeScale(1, 2, 3),
		mm.NewMatrix4().MakeScale(1/8, 1/2, 1/3),
		mm.NewMatrix4().MakeFrustum(-1, 1, -1, 1, 1, 1000),
		mm.NewMatrix4().MakeFrustum(-16, 16, -9, 9, 0.1, 10000),
		mm.NewMatrix4().MakeTranslation(1, 2, 3),
	}

	for _, m := range testMatrices {

		mInverse := mm.NewMatrix4().GetInverse(m)
		mSelfInverse := m.Clone()
		mSelfInverse.GetInverse(mSelfInverse)

		// self-inverse should the same as inverse
		if !matrixEquals4(mSelfInverse, mInverse) {
			t.Error()
		}

		// the determinant of the inverse should be the reciprocal
		if math.Abs(m.Determinant()*mInverse.Determinant()-1) > 0.0001 {
			t.Error()
		}

		mProduct := mm.NewMatrix4().MultiplyMatrices(m, mInverse)

		// the determinant of the identity matrix is 1
		if math.Abs(mProduct.Determinant()-1) > 0.0001 {
			t.Error()
		}
		if !matrixEquals4(mProduct, identity) {
			t.Error()
		}
	}
}

func TestMatrix4_MustGetInverse(t *testing.T) {

	m := mm.NewMatrix4().Set(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	defer func() {
		if r := recover(); r == nil {
			t.Error("inverse should have panicked")
		}
	}()

	m.MustGetInverse(nil)
}

func TestMatrix4_MakeExtractBasis(t *testing.T) {

	identityBasis := []*mm.Vector3{
		mm.NewVector3().Set(1, 0, 0),
		mm.NewVector3().Set(0, 1, 0),
		mm.NewVector3().Set(0, 0, 1),
	}

	a := mm.NewMatrix4().MakeBasis(identityBasis[0], identityBasis[1], identityBasis[2])
	identity := mm.NewMatrix4()
	if !matrixEquals4(a, identity) {
		t.Error()
	}

	testBases := [][]*mm.Vector3{
		{
			mm.NewVector3().Set(0, 1, 0),
			mm.NewVector3().Set(-1, 0, 0),
			mm.NewVector3().Set(0, 0, 1),
		},
	}
	for _, testBasis := range testBases {

		b := mm.NewMatrix4().MakeBasis(testBasis[0], testBasis[1], testBasis[2])
		outBasis := []*mm.Vector3{
			mm.NewVector3(),
			mm.NewVector3(),
			mm.NewVector3(),
		}
		b.ExtractBasis(outBasis[0], outBasis[1], outBasis[2])

		// check what goes in, is what comes out.
		for j, out := range outBasis {
			if !out.Equals(testBasis[j]) {
				t.Error("in basis should equal out basis")
			}
		}

		// get the basis out the hard way
		for j, basis := range identityBasis {
			outBasis[j].Copy(basis)
			outBasis[j].ApplyMatrix4(b)
		}

		// did the multiply method of basis extraction work?
		for j, out := range outBasis {
			if !out.Equals(testBasis[j]) {
				t.Error("multiplied basis should equal in basis")
			}
		}

	}

}

func TestMatrix4_Transpose(t *testing.T) {
	a := mm.NewMatrix4()
	b := a.Clone().Transpose()
	if !matrixEquals4(a, b) {
		t.Error("transpose of identity should equal identity")
	}

	b = mm.NewMatrix4().Set(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	c := b.Clone().Transpose()
	if matrixEquals4(b, c) {
		t.Error("non-transposed should no equal")
	}
	c.Transpose()
	if !matrixEquals4(b, c) {
		t.Error("transposed should equal")
		t.Error(b)
		t.Error(c)
	}
}

func TestMatrix4_Clone(t *testing.T) {
	a := mm.NewMatrix4().Set(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	b := a.Clone()

	if !matrixEquals4(a, b) {
		t.Error("clone of a should equal clone of b")
	}

	// ensure that it is a true copy
	a.Elements[0] = 2
	if matrixEquals4(a, b) {
		t.Error("clone should return deep copy")
	}
}

//TODO
/*func TestMatrix4_ComposeDecompose(t *testing.T) {
    tValues := [
        mm.NewVector3().Set(),
        mm.NewVector3().Set( 3, 0, 0 ),
        mm.NewVector3().Set( 0, 4, 0 ),
        mm.NewVector3().Set( 0, 0, 5 ),
        mm.NewVector3().Set( -6, 0, 0 ),
        mm.NewVector3().Set( 0, -7, 0 ),
        mm.NewVector3().Set( 0, 0, -8 ),
        mm.NewVector3().Set( -2, 5, -9 ),
        mm.NewVector3().Set( -2, -5, -9 )
    ]

    sValues := [
        mm.NewVector3().Set( 1, 1, 1 ),
        mm.NewVector3().Set( 2, 2, 2 ),
        mm.NewVector3().Set( 1, -1, 1 ),
        mm.NewVector3().Set( -1, 1, 1 ),
        mm.NewVector3().Set( 1, 1, -1 ),
        mm.NewVector3().Set( 2, -2, 1 ),
        mm.NewVector3().Set( -1, 2, -2 ),
        mm.NewVector3().Set( -1, -1, -1 ),
        mm.NewVector3().Set( -2, -2, -2 )
    ]

    rValues := [
        new THREE.Quaternion(),
        new THREE.Quaternion().SetFromEuler( new THREE.Euler( 1, 1, 0 ) ),
        new THREE.Quaternion().SetFromEuler( new THREE.Euler( 1, -1, 1 ) ),
        new THREE.Quaternion( 0, 0.9238795292366128, 0, 0.38268342717215614 )
    ]


    for( ti := 0 ti < tValues.length ti ++ ) {
        for( si := 0 si < sValues.length si ++ ) {
            for( ri := 0 ri < rValues.length ri ++ ) {
                t := tValues[ti]
                s := sValues[si]
                r := rValues[ri]

                m := mm.NewMatrix4().compose( t, r, s )
                t2 := mm.NewVector3().Set()
                r2 := new THREE.Quaternion()
                s2 := mm.NewVector3().Set()

                m.decompose( t2, r2, s2 )

                m2 := mm.NewMatrix4().compose( t2, r2, s2 )

                matrixIsSame := matrixEquals4( m, m2 )
                /* debug code
                if( ! matrixIsSame ) {
                    console.log( t, s, r )
                    console.log( t2, s2, r2 )
                    console.log( m, m2 )
                }*
                if !matrixEquals4( m, m2 ) {
 t.Error()
}

            }
        }
    }
}*/
