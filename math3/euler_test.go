package math3_test

import (
	"math"
	"testing"

	"github.com/rydrman/marshmallow"
)

var (
	eulerZero = mm.NewEuler().Set(0, 0, 0, mm.XYZ)
	eulerAxyz = mm.NewEuler().Set(1, 0, 0, mm.XYZ)
	eulerAzyx = mm.NewEuler().Set(0, 1, 0, mm.ZYX)
)

func eulerEquals(a, b *mm.Euler) bool {
	tolerance := 0.0001
	diff := math.Abs(a.GetX()-b.GetX()) + math.Abs(a.GetY()-b.GetY()) + math.Abs(a.GetZ()-b.GetZ())
	return (diff < tolerance)
}

func TestNewEuler(t *testing.T) {
	e := mm.NewEuler()
	if e == nil {
		t.Fail()
	}
}

func TestEuler_Equals(t *testing.T) {
	a := mm.NewEuler()
	if !a.Equals(eulerZero) {
		t.Fail()
	}
	if a.Equals(eulerAxyz) ||
		a.Equals(eulerAzyx) {
		t.Fail()
	}
}

func TestEuler_CloneCopy(t *testing.T) {
	a := eulerAxyz.Clone()
	if !a.Equals(eulerAxyz) ||
		a.Equals(eulerZero) ||
		a.Equals(eulerAzyx) {
		t.Fail()
	}

	a.Copy(eulerAzyx)
	if !a.Equals(eulerAzyx) ||
		a.Equals(eulerAxyz) ||
		a.Equals(eulerZero) {
		t.Fail()
	}

}

func TestEuler_SetSetFromVector3ToVector3(t *testing.T) {
	a := mm.NewEuler()

	a.Set(0, 1, 0, "ZYX")
	if !a.Equals(eulerAzyx) ||
		a.Equals(eulerAxyz) ||
		a.Equals(eulerZero) {
		t.Fail()
	}

	vec := mm.NewVector3().Set(0, 1, 0)

	b := mm.NewEuler().SetFromVector3(vec, "ZYX")
	if !a.Equals(b) {
		t.Fail()
	}

	c := b.ToVector3(nil)
	if !c.Equals(vec) {
		t.Fail()
	}
}

func TestEuler_QuaternionSetFromEulerEulerFromQuaternion(t *testing.T) {
	testValues := []*mm.Euler{eulerZero, eulerAxyz, eulerAzyx}
	for _, v := range testValues {
		q := mm.NewQuaternion().SetFromEuler(v, false)

		v2 := mm.NewEuler().SetFromQuaternion(q, v.Order, false)
		q2 := mm.NewQuaternion().SetFromEuler(v2, false)
		if !quatEquals(q, q2) {
			t.Error("expected quaternions to equal")
			t.Error(q)
			t.Error(q2)
		}
	}
}

func TestEuler_Matrix4SetFromEulerEulerFromRotationMatrix(t *testing.T) {
	testValues := []*mm.Euler{eulerZero, eulerAxyz, eulerAzyx}
	for _, v := range testValues {
		m := mm.NewMatrix4().MakeRotationFromEuler(v)

		v2 := mm.NewEuler().SetFromRotationMatrix(m, v.Order, false)
		m2 := mm.NewMatrix4().MakeRotationFromEuler(v2)
		if !matrixEquals4(m, m2) {
			t.Error("expected matrices to be equal")
			t.Error(m)
			t.Error(m2)
		}
	}
}

func TestEuler_Reorder(t *testing.T) {
	testValues := []*mm.Euler{eulerZero, eulerAxyz, eulerAzyx}
	for _, v := range testValues {
		q := mm.NewQuaternion().SetFromEuler(v, false)

		v.Reorder(mm.YZX)
		q2 := mm.NewQuaternion().SetFromEuler(v, false)
		if !quatEquals(q, q2) {
			t.Fail()
		}

		v.Reorder(mm.ZXY)
		q3 := mm.NewQuaternion().SetFromEuler(v, false)
		if !quatEquals(q, q3) {
			t.Fail()
		}
	}
}

func TestEuler_GimbalLocalQuat(t *testing.T) {
	// known problematic quaternions
	q1 := mm.NewQuaternion().Set(0.5207769385244341, -0.4783214164122354, 0.520776938524434, 0.47832141641223547)
	//q2 := mm.NewQuaternion().Set(0.11284905712620674, 0.6980437630368944, -0.11284905712620674, 0.6980437630368944)

	var eulerOrder mm.EulerOrder
	eulerOrder = mm.ZYX

	// create Euler directly from a Quaternion
	eViaQ1 := mm.NewEuler().SetFromQuaternion(q1, eulerOrder, false) // there is likely a bug here

	// create Euler from Quaternion via an intermediate Matrix4
	mViaQ1 := mm.NewMatrix4().MakeRotationFromQuaternion(q1)
	eViaMViaQ1 := mm.NewEuler().SetFromRotationMatrix(mViaQ1, eulerOrder, false)

	// the results here are different
	if !eulerEquals(eViaQ1, eViaMViaQ1) { // this result is correct
		t.Fail()
	}

}
