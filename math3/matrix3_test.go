package math3_test

import (
	"math"
	"testing"

	"github.com/rydrman/marshmallow"
)

func matrixEquals3(a, b *mm.Matrix3) bool {
	tolerance := 0.0001
	for i := range a.Elements {
		delta := a.Elements[i] - b.Elements[i]
		if delta > tolerance {
			return false
		}
	}
	return true
}

func toMatrix4(m3 *mm.Matrix3) *mm.Matrix4 {
	result := mm.NewMatrix4()
	re := result.Elements
	me := m3.Elements
	re[0] = me[0]
	re[1] = me[1]
	re[2] = me[2]
	re[4] = me[3]
	re[5] = me[4]
	re[6] = me[5]
	re[8] = me[6]
	re[9] = me[7]
	re[10] = me[8]

	return result
}

func TestNewMatrix3(t *testing.T) {
	a := mm.NewMatrix3()
	if a == nil {
		t.Fail()
	}
}

func TestMatrix3_Copy(t *testing.T) {
	a := mm.NewMatrix3().Set(0, 1, 2, 3, 4, 5, 6, 7, 8)
	b := mm.NewMatrix3().Copy(a)

	if !matrixEquals3(a, b) {
		t.Error("copied matrix does not equal original")
	}

	// ensure that it is a true copy
	a.Elements[0] = 2
	if matrixEquals3(a, b) {
		t.Error("copy was not a deep copy")
	}
}

func TestMatrix3_Set(t *testing.T) {
	b := mm.NewMatrix3()
	if b.Determinant() != 1 {
		t.Error("determinant should be 1")
	}

	b.Set(0, 1, 2, 3, 4, 5, 6, 7, 8)
	if b.Elements[0] != 0 ||
		b.Elements[1] != 3 ||
		b.Elements[2] != 6 ||
		b.Elements[3] != 1 ||
		b.Elements[4] != 4 ||
		b.Elements[5] != 7 ||
		b.Elements[6] != 2 ||
		b.Elements[7] != 5 ||
		b.Elements[8] != 8 {
		t.Fail()
	}
}

func TestMatrix3_Identity(t *testing.T) {
	b := mm.NewMatrix3().Set(0, 1, 2, 3, 4, 5, 6, 7, 8)

	a := mm.NewMatrix3()
	if matrixEquals3(a, b) {
		t.Fail()
	}

	b.Identity()
	if !matrixEquals3(a, b) {
		t.Error("expected a and b to both be identity matrices")
	}
}

func TestMatrix3_MultiplyScalar(t *testing.T) {
	b := mm.NewMatrix3().Set(0, 1, 2, 3, 4, 5, 6, 7, 8)
	b.MultiplyScalar(2)
	if b.Elements[0] != 0*2 ||
		b.Elements[1] != 3*2 ||
		b.Elements[2] != 6*2 ||
		b.Elements[3] != 1*2 ||
		b.Elements[4] != 4*2 ||
		b.Elements[5] != 7*2 ||
		b.Elements[6] != 2*2 ||
		b.Elements[7] != 5*2 ||
		b.Elements[8] != 8*2 {
		t.Fail()
	}
}

func TestMatrix3_Determinant(t *testing.T) {
	a := mm.NewMatrix3()
	if a.Determinant() != 1 {
		t.Fail()
	}

	a.Elements[0] = 2
	if a.Determinant() != 2 {
		t.Fail()
	}

	a.Elements[0] = 0
	if a.Determinant() != 0 {
		t.Fail()
	}

	a.Set(2, 3, 4, 5, 13, 7, 8, 9, 11)
	if a.Determinant() != -73 {
		t.Fail()
	}
}

func TestMatrix3_GetInverse(t *testing.T) {
	identity := mm.NewMatrix3()
	identity4 := mm.NewMatrix4()
	a := mm.NewMatrix3()
	b := mm.NewMatrix3()

	b.GetInverse(a)
	if !matrixEquals3(b, identity) {
		t.Error("inverse of identity should be identity")
	}

	testMatrices := []*mm.Matrix4{
		mm.NewMatrix4().MakeRotationX(0.3),
		mm.NewMatrix4().MakeRotationX(-0.3),
		mm.NewMatrix4().MakeRotationY(0.3),
		mm.NewMatrix4().MakeRotationY(-0.3),
		mm.NewMatrix4().MakeRotationZ(0.3),
		mm.NewMatrix4().MakeRotationZ(-0.3),
		mm.NewMatrix4().MakeScale(1, 2, 3),
		mm.NewMatrix4().MakeScale(1.0/8.0, 1.0/2.0, 1.0/3.0),
	}

	for _, m := range testMatrices {
		a.SetFromMatrix4(m)
		mInverse3 := b.GetInverse(a)

		mInverse := toMatrix4(mInverse3)

		// the determinant of the inverse should be the reciprocal
		if math.Abs(a.Determinant()*mInverse3.Determinant()-1) > 0.0001 {
			t.Errorf("%.4f", a.Determinant()*mInverse3.Determinant()-1)
			t.Errorf("expected reciprocal determinant: %.4f != %.4f", a.Determinant(), mInverse3.Determinant())
		}
		if math.Abs(m.Determinant()*mInverse.Determinant()-1) > 0.0001 {
			t.Errorf("%.2f", m.Determinant()*mInverse.Determinant()-1)
			t.Errorf("expected reciprocal determinant: %.2f, %.2f", m.Determinant(), mInverse.Determinant())
		}

		mProduct := mm.NewMatrix4().MultiplyMatrices(m, mInverse)
		if math.Abs(mProduct.Determinant()-1) > 0.0001 {
			t.Fail()
		}
		if !matrixEquals3(
			mm.NewMatrix3().SetFromMatrix4(mProduct),
			mm.NewMatrix3().SetFromMatrix4(identity4)) {
			t.Fail()
		}
	}
}

func TestMatrix3_MustGetInverse(t *testing.T) {

	m := mm.NewMatrix3().Set(0, 0, 0, 0, 0, 0, 0, 0, 0)

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected inverse to panic")
		}
	}()

	m.MustGetInverse(m)

}

func TestMatrix3_Transpose(t *testing.T) {
	a := mm.NewMatrix3()
	b := a.Clone().Transpose()
	if !matrixEquals3(a, b) {
		t.Error("transpose of identity should be identity")
	}

	b = mm.NewMatrix3().Set(0, 1, 2, 3, 4, 5, 6, 7, 8)
	c := b.Clone().Transpose()
	if matrixEquals3(b, c) {
		t.Error("expected matrices to not be equal")
		t.Error(c)
		t.Error(b)
	}
	c.Transpose()
	if !matrixEquals3(b, c) {
		t.Error("expected matrices to be equal")
		t.Error(c)
		t.Error(b)
	}
}

func TestMatrix3_Clone(t *testing.T) {
	a := mm.NewMatrix3().Set(0, 1, 2, 3, 4, 5, 6, 7, 8)
	b := a.Clone()

	if !matrixEquals3(a, b) {
		t.Fail()
	}

	// ensure that it is a true copy
	a.Elements[0] = 2
	if matrixEquals3(a, b) {
		t.Fail()
	}
}
