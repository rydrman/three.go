package math3_test

//ported from three.js by @rydrman

import (
	"math"
	"testing"

	math3 "github.com/rydrman/three.go/math"
)

func TestNewBox3(t *testing.T) {
	a := math3.NewBox3()
	if a == nil {
		t.Fail()
	}
}

func TestBox3_Set(t *testing.T) {
	a := math3.NewBox3()

	a.Set(zero3, one3)
	if !a.Min.Equals(zero3) || !a.Max.Equals(one3) {
		t.Fail()
	}
}

func TestBox3_Copy(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), one3.Clone())
	b := math3.NewBox3().Copy(a)
	if !b.Min.Equals(zero3) || !b.Max.Equals(one3) {
		t.FailNow()
	}

	// ensure that it is a true copy
	a.Min = zero3
	a.Max = one3
	if !b.Min.Equals(zero3) || !b.Max.Equals(one3) {
		t.Fail()
	}
}

func TestBox3_SetFromPoints(t *testing.T) {
	a := math3.NewBox3()

	a.SetFromPoints([]*math3.Vector3{zero3, one3, two3})
	if !a.Min.Equals(zero3) || !a.Max.Equals(two3) {
		t.FailNow()
	}

	a.SetFromPoints([]*math3.Vector3{one3})
	if !a.Min.Equals(one3) || !a.Max.Equals(one3) {
		t.FailNow()
	}

	a.SetFromPoints([]*math3.Vector3{})
	if !a.IsEmpty() {
		t.Fail()
	}
}

func TestBox3_MakeEmpty(t *testing.T) {
	a := math3.NewBox3()

	if !a.IsEmpty() {
		t.FailNow()
	}

	a = math3.NewBox3().Set(zero3.Clone(), one3.Clone())
	if a.IsEmpty() {
		t.Fail()
	}

	a.MakeEmpty()
	if !a.IsEmpty() {
		t.Fail()
	}
}

func TestBox3_Center(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), zero3.Clone())

	if !a.Center(nil).Equals(zero3) {
		t.Fail()
	}

	a = math3.NewBox3().Set(zero3.Clone(), one3.Clone())
	midpoint := one3.Clone().MultiplyScalar(0.5)
	if !a.Center(nil).Equals(midpoint) {
		t.Fail()
	}
}

func TestBox3_Size(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), zero3.Clone())

	if !a.Size(nil).Equals(zero3) {
		t.Fail()
	}

	a = math3.NewBox3().Set(zero3.Clone(), one3.Clone())
	if !a.Size(nil).Equals(one3) {
		t.Fail()
	}
}

func TestBox3_ExpandByPoint(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), zero3.Clone())

	a.ExpandByPoint(zero3)
	if !a.Size(nil).Equals(zero3) {
		t.FailNow()
	}

	a.ExpandByPoint(one3)
	if !a.Size(nil).Equals(one3) {
		t.FailNow()
	}

	a.ExpandByPoint(one3.Clone().Negate())
	if !a.Size(nil).Equals(one3.Clone().MultiplyScalar(2)) ||
		!a.Center(nil).Equals(zero3) {
		t.Fail()
	}
}

func TestBox3_ExpandByVector(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), zero3.Clone())

	a.ExpandByVector(zero3)
	if !a.Size(nil).Equals(zero3) {
		t.Fail()
	}

	a.ExpandByVector(one3)
	if !a.Size(nil).Equals(one3.Clone().MultiplyScalar(2)) ||
		!a.Center(nil).Equals(zero3) {
		t.Fail()
	}
}

func TestBox3_ExpandByScalar(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), zero3.Clone())

	a.ExpandByScalar(0)
	if !a.Size(nil).Equals(zero3) {
		t.Fail()
	}

	a.ExpandByScalar(1)
	if !a.Size(nil).Equals(one3.Clone().MultiplyScalar(2)) {
		t.Fail()
	}
	if !a.Center(nil).Equals(zero3) {
		t.Fail()
	}
}

func TestBox3_ContainsPoint(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), zero3.Clone())

	if !a.ContainsPoint(zero3) ||
		a.ContainsPoint(one3) {
		t.Fail()
	}

	a.ExpandByScalar(1)
	if !a.ContainsPoint(zero3) ||
		!a.ContainsPoint(one3) ||
		!a.ContainsPoint(one3.Clone().Negate()) {
		t.Fail()
	}
}

func TestBox3_ContainsBox(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), zero3.Clone())
	b := math3.NewBox3().Set(zero3.Clone(), one3.Clone())
	c := math3.NewBox3().Set(one3.Clone().Negate(), one3.Clone())

	if !a.ContainsBox(a) {
		t.Error("expected a to contain itself")
	}
	if a.ContainsBox(b) {
		t.Error("expected a not to contain b")
	}
	if a.ContainsBox(c) {
		t.Error("expected a not to contain c")
	}

	if !b.ContainsBox(a) {
		t.Error("expected b to contain a")
	}
	if !c.ContainsBox(a) {
		t.Error("expected c to contain a")
	}
	if b.ContainsBox(c) {
		t.Error("expected c not to conatin c")
	}
}

func TestBox3_GetParameter(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), one3.Clone())
	b := math3.NewBox3().Set(one3.Clone().Negate(), one3.Clone())

	if !a.GetParameter(
		math3.NewVector3().Set(0, 0, 0),
		nil).Equals(math3.NewVector3().Set(0, 0, 0)) {
		t.Fail()
	}
	if !a.GetParameter(
		math3.NewVector3().Set(1, 1, 1),
		nil).Equals(math3.NewVector3().Set(1, 1, 1)) {
		t.Fail()
	}

	if !b.GetParameter(
		math3.NewVector3().Set(-1, -1, -1),
		nil).Equals(math3.NewVector3().Set(0, 0, 0)) {
		t.Fail()
	}
	if !b.GetParameter(
		math3.NewVector3().Set(0, 0, 0),
		nil).Equals(math3.NewVector3().Set(0.5, 0.5, 0.5)) {
		t.Fail()
	}
	if !b.GetParameter(
		math3.NewVector3().Set(1, 1, 1),
		nil).Equals(math3.NewVector3().Set(1, 1, 1)) {
		t.Fail()
	}
}

func TestBox3_ClampPoint(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), zero3.Clone())
	b := math3.NewBox3().Set(one3.Clone().Negate(), one3.Clone())

	if !a.ClampPoint(
		math3.NewVector3().Set(0, 0, 0), nil).Equals(
		math3.NewVector3().Set(0, 0, 0)) {
		t.Fail()
	}
	if !a.ClampPoint(
		math3.NewVector3().Set(1, 1, 1), nil).Equals(
		math3.NewVector3().Set(0, 0, 0)) {
		t.Fail()
	}
	if !a.ClampPoint(
		math3.NewVector3().Set(-1, -1, -1), nil).Equals(
		math3.NewVector3().Set(0, 0, 0)) {
		t.Fail()
	}

	if !b.ClampPoint(
		math3.NewVector3().Set(2, 2, 2), nil).Equals(
		math3.NewVector3().Set(1, 1, 1)) {
		t.Fail()
	}
	if !b.ClampPoint(
		math3.NewVector3().Set(1, 1, 1), nil).Equals(
		math3.NewVector3().Set(1, 1, 1)) {
		t.Fail()
	}
	if !b.ClampPoint(
		math3.NewVector3().Set(0, 0, 0), nil).Equals(
		math3.NewVector3().Set(0, 0, 0)) {
		t.Fail()
	}
	if !b.ClampPoint(
		math3.NewVector3().Set(-1, -1, -1), nil).Equals(
		math3.NewVector3().Set(-1, -1, -1)) {
		t.Fail()
	}
	if !b.ClampPoint(
		math3.NewVector3().Set(-2, -2, -2), nil,
	).Equals(math3.NewVector3().Set(-1, -1, -1)) {
		t.Fail()
	}
}

func TestBox3_DistanceToPoint(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), zero3.Clone())
	b := math3.NewBox3().Set(one3.Clone().Negate(), one3.Clone())

	if a.DistanceToPoint(math3.NewVector3().Set(0, 0, 0)) != 0 {
		t.Fail()
	}
	if a.DistanceToPoint(math3.NewVector3().Set(1, 1, 1)) != math.Sqrt(3) {
		t.Fail()
	}
	if a.DistanceToPoint(math3.NewVector3().Set(-1, -1, -1)) != math.Sqrt(3) {
		t.Fail()
	}

	if b.DistanceToPoint(math3.NewVector3().Set(2, 2, 2)) != math.Sqrt(3) {
		t.Fail()
	}
	if b.DistanceToPoint(math3.NewVector3().Set(1, 1, 1)) != 0 {
		t.Fail()
	}
	if b.DistanceToPoint(math3.NewVector3().Set(0, 0, 0)) != 0 {
		t.Fail()
	}
	if b.DistanceToPoint(math3.NewVector3().Set(-1, -1, -1)) != 0 {
		t.Fail()
	}
	if b.DistanceToPoint(math3.NewVector3().Set(-2, -2, -2)) != math.Sqrt(3) {
		t.Fail()
	}
}

func TestBox3_IntersectsBox(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), zero3.Clone())
	b := math3.NewBox3().Set(zero3.Clone(), one3.Clone())
	c := math3.NewBox3().Set(one3.Clone().Negate(), one3.Clone())

	if !a.IntersectsBox(a) {
		t.FailNow()
	}
	if !a.IntersectsBox(b) {
		t.FailNow()
	}
	if !a.IntersectsBox(c) {
		t.FailNow()
	}

	if !b.IntersectsBox(a) {
		t.FailNow()
	}
	if !c.IntersectsBox(a) {
		t.FailNow()
	}
	if !b.IntersectsBox(c) {
		t.FailNow()
	}

	b.Translate(math3.NewVector3().Set(2, 2, 2))
	if a.IntersectsBox(b) {
		t.FailNow()
	}
	if b.IntersectsBox(a) {
		t.FailNow()
	}
	if b.IntersectsBox(c) {
		t.FailNow()
	}
}

//TODO
/*func TestBox3_IntersectsSphere(t *testing.T) {
    a := math3.NewBox3().Set(zero3.Clone(), one3.Clone())
    b := math3.NewSphere().Set(zero3.Clone(), 1)

    if !a.IntersectsSphere(b) {
        t.Fail()
    }

    b.Translate(math3.NewVector3().Set(2, 2, 2))
    if a.IntersectsSphere(b) {
        t.Fail()
    }
}*/

//TODO
/*func TestBox3_IntersectsPlane(t *testing.T) {
    a := math3.NewBox3().Set( zero3.Clone(), one3.Clone() )
    b := new THREE.Plane( math3.NewVector3().Set( 0, 1, 0 ), 1 )
    c := new THREE.Plane( math3.NewVector3().Set( 0, 1, 0 ), 1.25 )
    d := new THREE.Plane( math3.NewVector3().Set( 0, -1, 0 ), 1.25 )

    ok( a.IntersectsPlane( b ) , "Passed!" )
    ok( ! a.IntersectsPlane( c ) , "Passed!" )
    ok( ! a.IntersectsPlane( d ) , "Passed!" )
}/*

//TODO
/*func TestBox3_GetBoundingSphere(t *testing.T) {
    a := math3.NewBox3().Set( zero3.Clone(), zero3.Clone() )
    b := math3.NewBox3().Set( zero3.Clone(), one3.Clone() )
    c := math3.NewBox3().Set( one3.Clone().Negate(), one3.Clone() )

    ok( a.GetBoundingSphere().Equals( new THREE.Sphere( zero3, 0 ) ), "Passed!" )
    ok( b.GetBoundingSphere().Equals( new THREE.Sphere( one3.Clone().MultiplyScalar( 0.5 ), math.Sqrt( 3 ) * 0.5 ) ), "Passed!" )
    ok( c.GetBoundingSphere().Equals( new THREE.Sphere( zero3, math.Sqrt( 12 ) * 0.5 ) ), "Passed!" )
}*/

func TestBox3_Intersect(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), zero3.Clone())
	b := math3.NewBox3().Set(zero3.Clone(), one3.Clone())
	c := math3.NewBox3().Set(one3.Clone().Negate(), one3.Clone())

	if !a.Clone().Intersect(a).Equals(a) {
		t.Fail()
	}
	if !a.Clone().Intersect(b).Equals(a) {
		t.Fail()
	}
	if !b.Clone().Intersect(b).Equals(b) {
		t.Fail()
	}
	if !a.Clone().Intersect(c).Equals(a) {
		t.Fail()
	}
	if !b.Clone().Intersect(c).Equals(b) {
		t.Fail()
	}
	if !c.Clone().Intersect(c).Equals(c) {
		t.Fail()
	}
}

func TestBox3_Union(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), zero3.Clone())
	b := math3.NewBox3().Set(zero3.Clone(), one3.Clone())
	c := math3.NewBox3().Set(one3.Clone().Negate(), one3.Clone())

	if !a.Clone().Union(a).Equals(a) {
		t.Fail()
	}
	if !a.Clone().Union(b).Equals(b) {
		t.Fail()
	}
	if !a.Clone().Union(c).Equals(c) {
		t.Fail()
	}
	if !b.Clone().Union(c).Equals(c) {
		t.Fail()
	}
}

func compareBox(a, b *math3.Box3) bool {
	threshold := 0.0001
	return (a.Min.DistanceTo(b.Min) < threshold &&
		a.Max.DistanceTo(b.Max) < threshold)
}

func TestBox3_ApplyMatrix4(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), zero3.Clone())
	b := math3.NewBox3().Set(zero3.Clone(), one3.Clone())
	c := math3.NewBox3().Set(one3.Clone().Negate(), one3.Clone())
	d := math3.NewBox3().Set(one3.Clone().Negate(), zero3.Clone())

	m := math3.NewMatrix4().MakeTranslation(1, -2, 1)
	t1 := math3.NewVector3().Set(1, -2, 1)

	if !compareBox(a.Clone().ApplyMatrix4(m), a.Clone().Translate(t1)) {
		t.Errorf("error 1")
	}
	if !compareBox(b.Clone().ApplyMatrix4(m), b.Clone().Translate(t1)) {
		t.Errorf("error 2")
	}
	if !compareBox(c.Clone().ApplyMatrix4(m), c.Clone().Translate(t1)) {
		t.Errorf("error 3")
	}
	if !compareBox(d.Clone().ApplyMatrix4(m), d.Clone().Translate(t1)) {
		t.Errorf("error 4")
	}
}

func TestBox3_Translate(t *testing.T) {
	a := math3.NewBox3().Set(zero3.Clone(), zero3.Clone())
	b := math3.NewBox3().Set(zero3.Clone(), one3.Clone())
	d := math3.NewBox3().Set(one3.Clone().Negate(), zero3.Clone())

	if !a.Clone().Translate(one3).Equals(math3.NewBox3().Set(one3, one3)) {
		t.Error("error 1")
	}
	if !a.Clone().Translate(one3).Translate(one3.Clone().Negate()).Equals(a) {
		t.Error("error 2")
	}
	if !d.Clone().Translate(one3).Equals(b) {
		t.Error("error 3")
	}
	if !b.Clone().Translate(one3.Clone().Negate()).Equals(d) {
		t.Error("error 4")
	}
}
