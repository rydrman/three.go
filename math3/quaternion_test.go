package math3_test

import (
	"math"
	"testing"

	"github.com/rydrman/marshmallow"
)

func quatEquals(a, b *mm.Quaternion) bool {
	tolerance := 0.0001
	diff := math.Abs(a.GetX()-b.GetX()) + math.Abs(a.GetY()-b.GetY()) + math.Abs(a.GetZ()-b.GetZ()) + math.Abs(a.GetW()-b.GetW())
	return (diff < tolerance)
}

var orders = []mm.EulerOrder{mm.XYZ, mm.YXZ, mm.ZXY, mm.ZYX, mm.YZX, mm.XZY}
var eulerAngles = mm.NewEuler().Set(0.1, -0.3, 0.25, mm.CurrentOrder)

func qSub(a, b *mm.Quaternion) *mm.Quaternion {
	result := mm.NewQuaternion()
	result.Copy(a)

	result.SetX(result.GetX() - b.GetX())
	result.SetY(result.GetY() - b.GetY())
	result.SetZ(result.GetZ() - b.GetZ())
	result.SetW(result.GetW() - b.GetW())

	return result
}

func checkQuatMembers(q *mm.Quaternion, x, y, z, w float64, t *testing.T) {

	if q.GetX() != x ||
		q.GetY() != y ||
		q.GetZ() != z ||
		q.GetW() != w {

		t.Errorf("expected %v, got: %.2f, %.2f, %.2f %.2f", q, x, y, z, w)

	}

}

func TestNewQuaternion(t *testing.T) {
	a := mm.NewQuaternion()
	if nil == a {
		t.Fail()
	}
}

func TestQuaternion_Set(t *testing.T) {
	a := mm.NewQuaternion()

	a.Set(x, y, z, w)
	checkQuatMembers(a, x, y, z, w, t)
}

func TestQuaternion_Copy(t *testing.T) {
	a := mm.NewQuaternion().Set(x, y, z, w)
	b := mm.NewQuaternion().Copy(a)
	if !quatEquals(a, b) {
		t.Fail()
	}

	// ensure that it is a true copy
	a.SetX(0)
	a.SetY(-1)
	a.SetZ(0)
	a.SetW(-1)
	if b.GetX() != x {
		t.Fail()
	}
}

func TestQuaternion_SetFromAxisAngle(t *testing.T) {

	zero := mm.NewQuaternion()

	a := mm.NewQuaternion().SetFromAxisAngle(mm.NewVector3().Set(1, 0, 0), 0)
	if !a.Equals(zero) {
		t.Fail()
	}
	a = mm.NewQuaternion().SetFromAxisAngle(mm.NewVector3().Set(0, 1, 0), 0)
	if !a.Equals(zero) {
		t.Fail()
	}
	a = mm.NewQuaternion().SetFromAxisAngle(mm.NewVector3().Set(0, 0, 1), 0)
	if !a.Equals(zero) {
		t.Fail()
	}

	b1 := mm.NewQuaternion().SetFromAxisAngle(mm.NewVector3().Set(1, 0, 0), math.Pi)
	if a.Equals(b1) {
		t.Fail()
	}
	b2 := mm.NewQuaternion().SetFromAxisAngle(mm.NewVector3().Set(1, 0, 0), -math.Pi)
	if a.Equals(b2) {
		t.Fail()
	}

	b1.Multiply(b2)
	if !a.Equals(b1) {
		t.Fail()
	}
}

func TestQuaternion_SetFromEulerSetFromQuaternion(t *testing.T) {

	angles := []*mm.Vector3{mm.NewVector3().Set(1, 0, 0), mm.NewVector3().Set(0, 1, 0), mm.NewVector3().Set(0, 0, 1)}

	// ensure euler conversion to/from Quaternion matches.
	for _, order := range orders {
		for _, angle := range angles {
			eulers2 := mm.NewEuler().SetFromQuaternion(
				mm.NewQuaternion().SetFromEuler(
					mm.NewEuler().Set(angle.X, angle.Y, angle.Z, order),
					false),
				order,
				false,
			)
			newAngle := mm.NewVector3().Set(eulers2.GetX(), eulers2.GetY(), eulers2.GetZ())
			if newAngle.DistanceTo(angle) > 0.001 {
				t.Fail()
			}
		}
	}

}

func TestQuaternion_SetFromEulerSetFromRotationMatrix(t *testing.T) {

	// ensure euler conversion for Quaternion matches that of Matrix4
	for _, order := range orders {

		eulerAngles.Order = order
		q := mm.NewQuaternion().SetFromEuler(eulerAngles, false)
		m := mm.NewMatrix4().MakeRotationFromEuler(eulerAngles)
		q2 := mm.NewQuaternion().SetFromRotationMatrix(m)

		if qSub(q, q2).Length() > 0.001 {
			t.Errorf("expected euler and matrix4 to convert eulers the same way (%s)", order)
		}

	}

}

func TestQuaternion_NormalizeLengthLengthSq(t *testing.T) {
	a := mm.NewQuaternion().Set(x, y, z, w)

	if a.Length() == 1 ||
		a.LengthSq() == 1 {
		t.Fail()
	}

	a.Normalize()
	if a.Length() != 1 ||
		a.LengthSq() != 1 {
		t.Fail()
	}

	a.Set(0, 0, 0, 0)
	if a.LengthSq() != 0 ||
		a.Length() != 0 {
		t.Fail()
	}

	a.Normalize()
	if a.LengthSq() != 1 ||
		a.Length() != 1 {
		t.Fail()
	}
}

func TestQuaternion_InverseConjugate(t *testing.T) {
	a := mm.NewQuaternion().Set(x, y, z, w)

	b := a.Clone().Conjugate()

	checkQuatMembers(a, -b.GetX(), -b.GetY(), -b.GetZ(), b.GetW(), t)
}

func TestQuaternion_MultiplyQuaternionsMultiply(t *testing.T) {

	angles := []*mm.Euler{
		mm.NewEuler().Set(1, 0, 0, mm.CurrentOrder),
		mm.NewEuler().Set(0, 1, 0, mm.CurrentOrder),
		mm.NewEuler().Set(0, 0, 1, mm.CurrentOrder),
	}

	q1 := mm.NewQuaternion().SetFromEuler(angles[0], false)
	q2 := mm.NewQuaternion().SetFromEuler(angles[1], false)
	q3 := mm.NewQuaternion().SetFromEuler(angles[2], false)

	q := mm.NewQuaternion().MultiplyQuaternions(q1, q2).Multiply(q3)

	m1 := mm.NewMatrix4().MakeRotationFromEuler(angles[0])
	m2 := mm.NewMatrix4().MakeRotationFromEuler(angles[1])
	m3 := mm.NewMatrix4().MakeRotationFromEuler(angles[2])

	m := mm.NewMatrix4().MultiplyMatrices(m1, m2).Multiply(m3)

	qFromM := mm.NewQuaternion().SetFromRotationMatrix(m)

	if qSub(q, qFromM).Length() > 0.001 {
		t.Error("expected quaternions to have same length")
	}
}

func TestQuaternion_MultiplyVector3(t *testing.T) {

	angles := []*mm.Euler{
		mm.NewEuler().Set(1, 0, 0, mm.CurrentOrder),
		mm.NewEuler().Set(0, 1, 0, mm.CurrentOrder),
		mm.NewEuler().Set(0, 0, 1, mm.CurrentOrder),
	}

	// ensure euler conversion for Quaternion matches that of Matrix4
	for _, order := range orders {
		for _, angle := range angles {
			angle.Order = order
			q := mm.NewQuaternion().SetFromEuler(angle, false)
			m := mm.NewMatrix4().MakeRotationFromEuler(angle)

			v0 := mm.NewVector3().Set(1, 0, 0)
			qv := v0.Clone().ApplyQuaternion(q)
			mv := v0.Clone().ApplyMatrix4(m)

			if qv.DistanceTo(mv) > 0.001 {
				t.Fail()
			}
		}
	}

}

func TestQuaternion_Equals(t *testing.T) {
	a := mm.NewQuaternion().Set(x, y, z, w)
	b := mm.NewQuaternion().Set(-x, -y, -z, -w)

	if a.GetX() == b.GetX() ||
		a.GetY() == b.GetY() {
		t.Fail()
	}

	if a.Equals(b) ||
		b.Equals(a) {
		t.Fail()
	}

	a.Copy(b)
	if a.GetX() != b.GetX() ||
		a.GetY() != b.GetY() ||
		!a.Equals(b) ||
		!b.Equals(a) {
		t.Fail()
	}
}

type slerpResult struct {
	Equals func(x, y, z, w, e float64) bool
	Length float64
	DotA   float64
	DotB   float64
}

func doSlerpObject(aArr, bArr []float64, t float64) slerpResult {

	a := mm.NewQuaternion().FromArray(aArr, 0)
	b := mm.NewQuaternion().FromArray(bArr, 0)
	c := mm.NewQuaternion().FromArray(aArr, 0)

	c.Slerp(b, t)

	return slerpResult{

		Equals: func(x, y, z, w, maxError float64) bool {

			return math.Abs(x-c.GetX()) <= maxError &&
				math.Abs(y-c.GetY()) <= maxError &&
				math.Abs(z-c.GetZ()) <= maxError &&
				math.Abs(w-c.GetW()) <= maxError

		},

		Length: c.Length(),

		DotA: c.Dot(a),
		DotB: c.Dot(b),
	}

}

func doSlerpArray(a, b []float64, t float64) slerpResult {

	result := []float64{0, 0, 0, 0}

	mm.QuaternionSlerpFlat(result, 0, a, 0, b, 0, t)

	arrDot := func(a, b []float64) float64 {

		return a[0]*b[0] + a[1]*b[1] +
			a[2]*b[2] + a[3]*b[3]

	}

	return slerpResult{

		Equals: func(x, y, z, w, maxError float64) bool {

			return math.Abs(x-result[0]) <= maxError &&
				math.Abs(y-result[1]) <= maxError &&
				math.Abs(z-result[2]) <= maxError &&
				math.Abs(w-result[3]) <= maxError

		},

		Length: math.Sqrt(arrDot(result, result)),

		DotA: arrDot(result, a),
		DotB: arrDot(result, b),
	}

}

func slerpTestSkeleton(
	doSlerp func(a, b []float64, t float64) slerpResult,
	maxError float64,
	t *testing.T,
) {

	a := []float64{
		0.6753410084407496,
		0.4087830051091744,
		0.32856700410659473,
		0.5185120064806223,
	}

	b := []float64{
		0.6602792107657797,
		0.43647413932562285,
		0.35119011210236006,
		0.5001871596632682,
	}

	maxNormError := 0.0

	isNormal := func(result slerpResult) bool {

		normError := math.Abs(1 - result.Length)
		maxNormError = math.Max(maxNormError, normError)
		return normError <= maxError

	}

	result := doSlerp(a, b, 0)
	if !result.Equals(a[0], a[1], a[2], a[3], 0) {
		t.Error("expected exactly A @ t = 0")
	}

	result = doSlerp(a, b, 1)
	if !result.Equals(b[0], b[1], b[2], b[3], 0) {
		t.Error("expected exactly B @ t = 1")
	}

	result = doSlerp(a, b, 0.5)
	if math.Abs(result.DotA-result.DotB) > mm.Epsilon {
		t.Error("expected Symmetry at 0.5")
	}
	if !isNormal(result) {
		t.Error("expected Approximately normal (at 0.5)")
	}

	result = doSlerp(a, b, 0.25)
	if result.DotA < result.DotB {
		t.Error("expected Interpolating at 0.25")
	}
	if !isNormal(result) {
		t.Error("expected Approximately normal (at 0.25)")
	}

	result = doSlerp(a, b, 0.75)
	if result.DotA > result.DotB {
		t.Error("expected Interpolating at 0.75")
	}
	if !isNormal(result) {
		t.Error("expected Approximately normal (at 0.75)")
	}

	D := math.Sqrt(0.5)

	result = doSlerp([]float64{1, 0, 0, 0}, []float64{0, 0, 1, 0}, 0.5)
	if !result.Equals(D, 0, D, 0, mm.Epsilon) {
		t.Error("expected X/Z diagonal from axes")
	}
	if !isNormal(result) {
		t.Error("expected Approximately normal (X/Z diagonal)")
	}

	result = doSlerp([]float64{0, D, 0, D}, []float64{0, -D, 0, D}, 0.5)
	if !result.Equals(0, 0, 0, 1, mm.Epsilon) {
		t.Error("expected W-Unit from diagonals")
	}
	if !isNormal(result) {
		t.Error("expected Approximately normal (W-Unit)")
	}
}

func TestQuaternion_Slerp(t *testing.T) {

	slerpTestSkeleton(doSlerpObject, mm.Epsilon, t)

}

func TestQuaternion_SlerpFlat(t *testing.T) {

	slerpTestSkeleton(doSlerpArray, mm.Epsilon, t)

}
