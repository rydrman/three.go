package math3_test

import (
	"math"
	"testing"

	"github.com/rydrman/marshmallow"
)

func TestNewVector2(t *testing.T) {

	a := mm.NewVector2()
	if a.X != 0 ||
		a.Y != 0 ||
		a.Z != 0 {
		t.Error("new vector elements should all be 0")
	}

}

func TestVector2_Set(t *testing.T) {

	a := mm.NewVector2()
	a.Set(x, y, z)
	if a.X != x ||
		a.Y != y ||
		a.Z != z {
		t.Error("set should change the values properly")
	}

}

func TestVector2_Copy(t *testing.T) {

	a := mm.NewVector2().Set(x, y, z)
	b := mm.NewVector2().Copy(a)
	if b.X != a.X ||
		b.Y != a.Y ||
		b.Z != a.Z {
		t.Error("copy should equal original")
		t.Error(a)
		t.Error(b)
	}

	// ensure that it is a true copy
	a.Y = -1
	if a.Y == b.Y {
		t.Error("copy should create a deep copy")
	}

}

func TestVector2_SetXYZ(t *testing.T) {
	a := mm.NewVector2()

	a.SetX(x)
	a.SetY(y)
	a.SetZ(z)

	if a.X != x ||
		a.Y != y ||
		a.Z != z {
		t.Error("setX, Y, and Z should change the values properly")
	}

}

func TestVector2_GetSetComponent(t *testing.T) {

	a := mm.NewVector2()

	a.SetComponent(0, x)
	a.SetComponent(1, y)
	a.SetComponent(2, z)
	if a.X != x ||
		a.Y != y ||
		a.Z != z {
		t.Error("SetComponent should change the values properly")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("out of range should panic")
		}
	}()
	a.SetComponent(3, 0)

}

func TestVector2_Add(t *testing.T) {
	a := mm.NewVector2().Set(x, y, z)
	b := mm.NewVector2().Set(-x, -y, -z)

	a.Add(b)

	if a.X != 0 ||
		a.Y != 0 ||
		a.Z != 0 {
		t.Error("add returned invalid value")
	}

}

func TestVector2_AddVectors(t *testing.T) {

	a := mm.NewVector2().Set(x, y, z)
	b := mm.NewVector2().Set(-x, -y, -z)

	c := mm.NewVector2().AddVectors(a, b)

	if c.X != 0 ||
		c.Y != 0 ||
		c.Z != 0 {
		t.Error("add vectors returned invalid value")
	}
}

func TestVector2_Sub(t *testing.T) {
	a := mm.NewVector2().Set(x, y, z)
	b := mm.NewVector2().Set(-x, -y, -z)

	a.Sub(b)

	if a.X != 2*x ||
		a.Y != 2*y ||
		a.Z != 2*z {
		t.Error("sub returned invalid value")
	}

}

func TestVector2_SubVectors(t *testing.T) {

	a := mm.NewVector2().Set(x, y, z)
	b := mm.NewVector2().Set(-x, -y, -z)

	c := mm.NewVector2().SubVectors(a, b)

	if c.X != 2*x ||
		c.Y != 2*y ||
		c.Z != 2*z {
		t.Error("sub vectors returned invalid value")
	}
}

func TestVector2_MultiplyDivideScalar(t *testing.T) {

	a := mm.NewVector2().Set(x, y, z)
	b := mm.NewVector2().Set(-x, -y, -z)

	a.MultiplyScalar(-2)
	if a.X != x*-2 ||
		a.Y != y*-2 ||
		a.Z != z*-2 {
		t.Error("mulitply scalar returned incorrect value")
	}

	b.MultiplyScalar(-2)
	if b.X != 2*x ||
		b.Y != 2*y ||
		b.Z != 2*z {
		t.Error("mulitply scalar returned incorrect value")
	}

	a.DivideScalar(-2)
	if a.X != x ||
		a.Y != y ||
		a.Z != z {
		t.Error("divide scalar returned incorrect value")
	}

	b.DivideScalar(-2)
	if b.X != -x ||
		b.Y != -y ||
		b.Z != -z {
		t.Error("divide scalar returned incorrect value")
	}

}

func TestVector2_MinMaxClamp(t *testing.T) {

	a := mm.NewVector2().Set(x, y, z)
	b := mm.NewVector2().Set(-x, -y, -z)
	c := mm.NewVector2()

	c.Copy(a).Min(b)
	if c.X != -x ||
		c.Y != -y ||
		c.Z != -z {
		t.Error("min returned incorrect value")
	}

	c.Copy(a).Max(b)
	if c.X != x ||
		c.Y != y ||
		c.Z != z {
		t.Error("max returned incorrect value")
	}

	c.Set(-2*x, 2*y, -2*z)
	c.Clamp(b, a)
	if c.X != -x ||
		c.Y != y ||
		c.Z != -z {
		t.Error("clamp returned incorrect value")
	}

}

func TestVector2_Negate(t *testing.T) {

	a := mm.NewVector2().Set(x, y, z)

	a.Negate()
	if a.X != -x ||
		a.Y != -y ||
		a.Z != -z {
		t.Error("negate returned incorrect value")
	}

}

func TestVector2_Dot(t *testing.T) {
	a := mm.NewVector2().Set(x, y, z)
	b := mm.NewVector2().Set(-x, -y, -z)
	c := mm.NewVector2()

	result := a.Dot(b)
	if result != -x*x-y*y-z*z {
		t.Error("dot returned incorrect value")
	}

	result = a.Dot(c)
	if result != 0 {
		t.Error("dot with 0 vector should be 0")
	}
}

func TestVector2_LengthLengthSq(t *testing.T) {
	a := mm.NewVector2().Set(x, 0, 0)
	b := mm.NewVector2().Set(0, -y, 0)
	c := mm.NewVector2().Set(0, 0, z)
	d := mm.NewVector2()

	if a.Length() != x ||
		a.LengthSq() != x*x {
		t.Error("Length of Length squared returned incorrect value")
	}
	if b.Length() != y ||
		b.LengthSq() != y*y {
		t.Error("Length of Length squared returned incorrect value")
	}
	if c.Length() != z ||
		c.LengthSq() != z*z {
		t.Error("Length of Length squared returned incorrect value")
	}
	if d.Length() != 0 ||
		d.LengthSq() != 0 {
		t.Error("Length of Length squared returned incorrect value")
	}

	a.Set(x, y, z)
	if a.Length() != math.Sqrt(x*x+y*y+z*z) ||
		a.LengthSq() != (x*x+y*y+z*z) {
		t.Error("Length of Length squared returned incorrect value")
	}
}

func TestVector2_Normalize(t *testing.T) {
	a := mm.NewVector2().Set(x, 0, 0)
	b := mm.NewVector2().Set(0, -y, 0)
	c := mm.NewVector2().Set(0, 0, z)

	a.Normalize()
	if a.Length() != 1 ||
		a.X != 1 {
		t.Error("normalize failed to normalize a properly")
	}

	b.Normalize()
	if b.Length() != 1 ||
		b.Y != -1 {
		t.Error("normalize failed to normalize b properly")
	}

	c.Normalize()
	if c.Length() != 1 ||
		c.Z != 1 {
		t.Error("normalize failed to normalize c properly")
	}
}

func TestVector2_DistanceToDistanceToSquared(t *testing.T) {
	a := mm.NewVector2().Set(x, 0, 0)
	b := mm.NewVector2().Set(0, -y, 0)
	c := mm.NewVector2().Set(0, 0, z)
	d := mm.NewVector2()

	if a.DistanceTo(d) != x ||
		a.DistanceToSquared(d) != x*x {
		t.Error("distance to or distance to squared returned incorrect value")
	}

	if b.DistanceTo(d) != y ||
		b.DistanceToSquared(d) != y*y {
		t.Error("distance to or distance to squared returned incorrect value")
	}

	if c.DistanceTo(d) != z ||
		c.DistanceToSquared(d) != z*z {
		t.Error("distance to or distance to squared returned incorrect value")
	}
}

func TestVector2_SetLength(t *testing.T) {
	a := mm.NewVector2().Set(x, 0, 0)

	a.SetLength(y)
	if a.Length() != y {
		t.Fail()
	}

	a = mm.NewVector2().Set(0, 0, 0)
	a.SetLength(y)
	if a.Length() != 0 {
		t.Fail()
	}

}

func TestVector2_ProjectOnVector(t *testing.T) {
	a := mm.NewVector2().Set(1, 0, 0)
	b := mm.NewVector2()
	normal := mm.NewVector2().Set(10, 0, 0)

	if !b.Copy(a).ProjectOnVector(normal).Equals(mm.NewVector2().Set(1, 0, 0)) {
		t.Fail()
	}

	a.Set(0, 1, 0)
	if !b.Copy(a).ProjectOnVector(normal).Equals(mm.NewVector2().Set(0, 0, 0)) {
		t.Fail()
	}

	a.Set(0, 0, -1)
	if !b.Copy(a).ProjectOnVector(normal).Equals(mm.NewVector2().Set(0, 0, 0)) {
		t.Fail()
	}

	a.Set(-1, 0, 0)
	if !b.Copy(a).ProjectOnVector(normal).Equals(mm.NewVector2().Set(-1, 0, 0)) {
		t.Fail()
	}

}

func TestVector2_ProjectOnPlane(t *testing.T) {
	a := mm.NewVector2().Set(1, 0, 0)
	b := mm.NewVector2()
	normal := mm.NewVector2().Set(1, 0, 0)

	if !b.Copy(a).ProjectOnPlane(normal).Equals(mm.NewVector2().Set(0, 0, 0)) {
		t.Fail()
	}

	a.Set(0, 1, 0)
	if !b.Copy(a).ProjectOnPlane(normal).Equals(mm.NewVector2().Set(0, 1, 0)) {
		t.Fail()
	}

	a.Set(0, 0, -1)
	if !b.Copy(a).ProjectOnPlane(normal).Equals(mm.NewVector2().Set(0, 0, -1)) {
		t.Fail()
	}

	a.Set(-1, 0, 0)
	if !b.Copy(a).ProjectOnPlane(normal).Equals(mm.NewVector2().Set(0, 0, 0)) {
		t.Fail()
	}

}

func TestVector2_Reflect(t *testing.T) {
	a := mm.NewVector2()
	normal := mm.NewVector2().Set(0, 1, 0)
	b := mm.NewVector2()

	a.Set(0, -1, 0)
	if !b.Copy(a).Reflect(normal).Equals(mm.NewVector2().Set(0, 1, 0)) {
		t.Fail()
	}

	a.Set(1, -1, 0)
	if !b.Copy(a).Reflect(normal).Equals(mm.NewVector2().Set(1, 1, 0)) {
		t.Fail()
	}

	a.Set(1, -1, 0)
	normal.Set(0, -1, 0)
	if !b.Copy(a).Reflect(normal).Equals(mm.NewVector2().Set(1, 1, 0)) {
		t.Fail()
	}
}

func TestVector2_angleTo(t *testing.T) {
	a := mm.NewVector2().Set(0, -0.18851655680720186, 0.9820700116639124)
	b := mm.NewVector2().Set(0, 0.18851655680720186, -0.9820700116639124)

	if a.AngleTo(a) != 0 {
		t.Fail()
	}
	if a.AngleTo(b) != math.Pi {
		t.Fail()
	}

	x := mm.NewVector2().Set(1, 0, 0)
	y := mm.NewVector2().Set(0, 1, 0)
	z := mm.NewVector2().Set(0, 0, 1)

	if x.AngleTo(y) != math.Pi/2 {
		t.Fail()
	}
	if x.AngleTo(z) != math.Pi/2 {
		t.Fail()
	}
	if z.AngleTo(x) != math.Pi/2 {
		t.Fail()
	}

	if math.Abs(x.AngleTo(mm.NewVector2().Set(1, 1, 0))-(math.Pi/4)) > 0.0000001 {
		t.Fail()
	}
}

func TestVector2_LerpClone(t *testing.T) {

	a := mm.NewVector2().Set(x, 0, z)
	b := mm.NewVector2().Set(0, -y, 0)

	if !a.Lerp(a, 0).Equals(a.Lerp(a, 0.5)) {
		t.Fail()
	}
	if !a.Lerp(a, 0).Equals(a.Lerp(a, 1)) {
		t.Fail()
	}

	if !a.Clone().Lerp(b, 0).Equals(a) {
		t.Fail()
	}

	if a.Clone().Lerp(b, 0.5).X != x*0.5 {
		t.Fail()
	}
	if a.Clone().Lerp(b, 0.5).Y != -y*0.5 {
		t.Fail()
	}
	if a.Clone().Lerp(b, 0.5).Z != z*0.5 {
		t.Fail()
	}

	if !a.Clone().Lerp(b, 1).Equals(b) {
		t.Fail()
	}

}

func TestVector2_Equals(t *testing.T) {

	a := mm.NewVector2().Set(x, 0, z)
	b := mm.NewVector2().Set(0, -y, 0)

	if a.X == b.X {
		t.Fail()
	}
	if a.Y == b.Y {
		t.Fail()
	}
	if a.Z == b.Z {
		t.Fail()
	}

	if a.Equals(b) {
		t.Fail()
	}
	if b.Equals(a) {
		t.Fail()
	}

	a.Copy(b)
	if a.X != b.X {
		t.Fail()
	}
	if a.Y != b.Y {
		t.Fail()
	}
	if a.Z != b.Z {
		t.Fail()
	}

	if !a.Equals(b) {
		t.Fail()
	}
	if !b.Equals(a) {
		t.Fail()
	}

}
