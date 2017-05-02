package math3_test

import (
	"math"

	math3 "github.com/rydrman/three.go/math"
)

//these are all constants to be used in testing

const (
	x = 2
	y = 3
	z = 4
	w = 5
)

var (

	//var negInf2 = new THREE.Vector2( math.Inf(-1), math.Inf(-1) )
	//var posInf2 = new THREE.Vector2( math.Inf(1), math.Inf(1) )

	//var zero2 = new THREE.Vector2()
	//var one2 = new THREE.Vector2( 1, 1 )
	//var two2 = new THREE.Vector2( 2, 2 )

	negInf3 = math3.NewVector3().Set(math.Inf(-1), math.Inf(-1), math.Inf(-1))
	posInf3 = math3.NewVector3().Set(math.Inf(1), math.Inf(1), math.Inf(1))

	zero3 = math3.NewVector3()
	one3  = math3.NewVector3().Set(1, 1, 1)
	two3  = math3.NewVector3().Set(2, 2, 2)
)
