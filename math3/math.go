package math3

import (
	"math"
	"math/rand"
)

const (
	Deg2Rad = math.Pi / 180.0
	Rad2Deg = 180.0 / math.Pi
)

var (
	Epsilon = math.Pow(2, -52)
)

func Max(a, b float64, vals ...float64) float64 {

	m := math.Max(a, b)

	for _, v := range vals {

		m = math.Max(m, v)

	}

	return m

}

func Round(v float64) float64 {

	return math.Floor(v + 0.5)

}

func Clamp(value, min, max float64) float64 {

	return math.Max(min, math.Min(max, value))

}

func EuclideanModulo(n, m float64) float64 {

	return math.Mod(math.Mod(n, m)+m, m)

}

func GenerateUUID() string {

	// http://www.broofa.com/Tools/math.Uuid.htm

	chars := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	uuid := make([]rune, 36)
	rnd := 0
	var r int

	for i := 0; i < 36; i++ {

		if i == 8 || i == 13 || i == 18 || i == 23 {

			uuid[i] = '-'

		} else if i == 14 {

			uuid[i] = '4'

		} else {

			if rnd <= 0x02 {
				rnd = int(0x2000000 + (rand.Float64() * 0x1000000))
			}
			r = rnd & 0xf
			rnd = rnd >> 4
			if i == 19 {
				uuid[i] = chars[r&0x3|0x8]
			} else {
				uuid[i] = chars[r]
			}

		}

	}

	return string(uuid)

}
