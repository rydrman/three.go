package math3

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/golang/glog"
	"github.com/rydrman/three.go"
)

type Color struct {
	R float64
	G float64
	B float64
}

func NewColor() *Color {

	return &Color{
		1, 1, 1,
	}

}

func (color *Color) R32() float32 {

	return float32(color.R)

}

func (color *Color) G32() float32 {

	return float32(color.G)

}

func (color *Color) B32() float32 {

	return float32(color.B)

}

func (color *Color) Set(r, g, b float64) *Color {

	color.R = r
	color.G = g
	color.B = b

	return color

}

func (color *Color) SetScalar(scalar float64) *Color {

	color.R = scalar
	color.G = scalar
	color.B = scalar

	return color

}

func (color *Color) SetHex(hex int) *Color {

	hex = int(math.Floor(float64(hex)))

	color.R = float64((hex >> 16 & 255)) / 255
	color.G = float64((hex >> 8 & 255)) / 255
	color.B = float64((hex & 255)) / 255

	return color

}

func (color *Color) SetRGB(r, g, b float64) *Color {

	color.R = r
	color.G = g
	color.B = b

	return color

}

func (color *Color) SetHSL(h, s, l float64) *Color {

	// h,s,l ranges are in 0.0 - 1.0
	h = EuclideanModulo(h, 1)
	s = Clamp(s, 0, 1)
	l = Clamp(l, 0, 1)

	if s == 0 {

		color.R = l
		color.G = l
		color.B = l

	} else {

		var p float64
		if l <= 0.5 {
			p = l * (1 + s)
		} else {
			p = l + s - (l * s)
		}
		q := (2 * l) - p

		color.R = hue2RGB(q, p, h+1.0/3.0)
		color.G = hue2RGB(q, p, h)
		color.B = hue2RGB(q, p, h-1.0/3.0)

	}

	return color

}

func (color *Color) SetStyle(style string) *Color {

	match := regexp.MustCompile(`^((?:rgb|hsl)a?)\(\s*([^\)]*)\)`).
		FindStringSubmatch(style)

	if len(match) > 0 {

		// rgb / hsl

		name := match[1]
		components := match[2]

		switch name {

		case "rgb":
			fallthrough
		case "rgba":

			colorMatch := regexp.MustCompile(
				`^(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*(,\s*([0-9]*\.?[0-9]+)\s*)?$`).
				FindStringSubmatch(components)
			if len(colorMatch) > 0 {

				// rgb(255,0,0) rgba(255,0,0,0.5)
				r, _ := strconv.ParseInt(colorMatch[1], 10, 32)
				g, _ := strconv.ParseInt(colorMatch[2], 10, 32)
				b, _ := strconv.ParseInt(colorMatch[3], 10, 32)
				color.R = math.Min(255, float64(r)) / 255
				color.G = math.Min(255, float64(g)) / 255
				color.B = math.Min(255, float64(b)) / 255

				handleAlpha(colorMatch[5])

				return color

			}

			colorMatch = regexp.MustCompile(`^(\d+)\%\s*,\s*(\d+)\%\s*,\s*(\d+)\%\s*(,\s*([0-9]*\.?[0-9]+)\s*)?$`).
				FindStringSubmatch(components)
			if len(colorMatch) > 0 {

				// rgb(100%,0%,0%) rgba(100%,0%,0%,0.5)
				r, _ := strconv.ParseInt(colorMatch[1], 10, 32)
				g, _ := strconv.ParseInt(colorMatch[2], 10, 32)
				b, _ := strconv.ParseInt(colorMatch[3], 10, 32)
				color.R = math.Min(100, float64(r)) / 100
				color.G = math.Min(100, float64(g)) / 100
				color.B = math.Min(100, float64(b)) / 100

				handleAlpha(colorMatch[5])

				return color

			}

			break

		case "hsl":
			fallthrough
		case "hsla":

			colorMatch := regexp.MustCompile(`^([0-9]*\.?[0-9]+)\s*,\s*(\d+)\%\s*,\s*(\d+)\%\s*(,\s*([0-9]*\.?[0-9]+)\s*)?$`).
				FindStringSubmatch(components)
			if len(colorMatch) > 0 {

				// hsl(120,50%,50%) hsla(120,50%,50%,0.5)
				h, _ := strconv.ParseFloat(colorMatch[1], 10)
				h = h / 360
				s, _ := strconv.ParseInt(colorMatch[2], 10, 32)
				sf := float64(s) / 100
				l, _ := strconv.ParseInt(colorMatch[3], 10, 32)
				lf := float64(l) / 100

				handleAlpha(colorMatch[5])

				return color.SetHSL(h, sf, lf)

			}

			break

		}

	} else {
		match = regexp.MustCompile(`^\#([A-Fa-f0-9]+)$`).
			FindStringSubmatch(style)
		if len(match) > 0 {

			// hex color

			hex := match[1]
			size := len(hex)

			if size == 3 {

				// #ff0
				r, _ := strconv.ParseInt(three.CharAt(hex, 0)+three.CharAt(hex, 0), 16, 32)
				g, _ := strconv.ParseInt(three.CharAt(hex, 1)+three.CharAt(hex, 1), 16, 32)
				b, _ := strconv.ParseInt(three.CharAt(hex, 2)+three.CharAt(hex, 2), 16, 32)

				color.R = float64(r) / 255
				color.G = float64(g) / 255
				color.B = float64(b) / 255

				return color

			} else if size == 6 {

				// #ff0000
				r, _ := strconv.ParseInt(three.CharAt(hex, 0)+three.CharAt(hex, 1), 16, 34)
				g, _ := strconv.ParseInt(three.CharAt(hex, 2)+three.CharAt(hex, 3), 16, 32)
				b, _ := strconv.ParseInt(three.CharAt(hex, 4)+three.CharAt(hex, 5), 16, 32)

				color.R = float64(r) / 255
				color.G = float64(g) / 255
				color.B = float64(b) / 255

				return color

			}

		}

	}

	if len(style) > 0 {

		// color keywords
		hex := ColorKeywords[style]

		if hex > 0 {

			// red
			color.SetHex(hex)

		} else {

			// unknown color
			glog.Warningf("three.Color: Unknown color %s", style)

		}

	}

	return color

}

func (color *Color) Clone() *Color {

	return NewColor().Copy(color)

}

func (color *Color) Copy(src *Color) *Color {

	color.R = src.R
	color.G = src.G
	color.B = src.B

	return color

}

func (color *Color) CopyGammaToLinear(gammaColor *Color) *Color {

	gammaFactor := 2.0

	color.R = math.Pow(gammaColor.R, gammaFactor)
	color.G = math.Pow(gammaColor.G, gammaFactor)
	color.B = math.Pow(gammaColor.B, gammaFactor)

	return color

}

func (color *Color) CopyLinearToGamma(gammaColor *Color) *Color {

	gammaFactor := 2.0

	if gammaFactor > 0 {
		gammaFactor = 1.0 / gammaFactor
	} else {
		gammaFactor = 1.0
	}

	color.R = math.Pow(gammaColor.R, gammaFactor)
	color.G = math.Pow(gammaColor.G, gammaFactor)
	color.B = math.Pow(gammaColor.B, gammaFactor)

	return color

}

func (color *Color) ConvertGammaToLinear() *Color {

	r := color.R
	g := color.G
	b := color.B

	color.R = r * r
	color.G = g * g
	color.B = b * b

	return color

}

func (color *Color) ConvertLinearToGamma() *Color {

	color.R = math.Sqrt(color.R)
	color.G = math.Sqrt(color.G)
	color.B = math.Sqrt(color.B)

	return color

}

func (color *Color) GetHex() int {

	return int(color.R*255)<<16 ^ int(color.G*255)<<8 ^ int(color.B*255)<<0

}

func (color *Color) GetHexString() string {

	return fmt.Sprintf("%x", color.GetHex())

}

func (color *Color) GetHSL() (hue, saturation, lightness float64) {

	// h,s,l ranges are in 0.0 - 1.0

	r := color.R
	g := color.G
	b := color.B

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))

	hue = (min + max) / 2.0
	saturation = hue
	lightness = hue

	if min == max {

		hue = 0
		saturation = 0

	} else {

		delta := max - min

		if lightness <= 0.5 {
			saturation = delta / (max + min)
		} else {
			saturation = delta / (2 - max - min)
		}

		switch max {

		case r:
			if g < b {
				hue = (g-b)/delta + 6
			} else {
				hue = (g-b)/delta + 0
			}
		case g:
			hue = (b-r)/delta + 2
		case b:
			hue = (r-g)/delta + 4

		}

		hue /= 6

	}

	return

}

func (color *Color) GetStyle() string {

	return fmt.Sprintf(
		"rgb(%003d,%003d,%003d)",
		int(color.R*255), int(color.G*255), int(color.B*255))

}

func (color *Color) OffsetHSL(h, s, l float64) *Color {

	hc, sc, lc := color.GetHSL()

	hc += h
	sc += s
	lc += l

	color.SetHSL(hc, sc, lc)

	return color

}

func (color *Color) Add(c *Color) *Color {

	color.R += c.R
	color.G += c.G
	color.B += c.B

	return color

}

func (color *Color) AddColors(color1, color2 *Color) *Color {

	color.R = color1.R + color2.R
	color.G = color1.G + color2.G
	color.B = color1.B + color2.B

	return color

}

func (color *Color) AddScalar(s float64) *Color {

	color.R += s
	color.G += s
	color.B += s

	return color

}

func (color *Color) Sub(c *Color) *Color {

	color.R = math.Max(0, color.R-c.R)
	color.G = math.Max(0, color.G-c.G)
	color.B = math.Max(0, color.B-c.B)

	return color

}

func (color *Color) Multiply(c *Color) *Color {

	color.R *= c.R
	color.G *= c.G
	color.B *= c.B

	return color

}

func (color *Color) MultiplyScalar(s float64) *Color {

	color.R *= s
	color.G *= s
	color.B *= s

	return color

}

func (color *Color) Lerp(c *Color, alpha float64) *Color {

	color.R += (c.R - color.R) * alpha
	color.G += (c.G - color.G) * alpha
	color.B += (c.B - color.B) * alpha

	return color

}

func (color *Color) Equals(c *Color) bool {

	return (c.R == color.R) && (c.G == color.G) && (c.B == color.B)

}

func (color *Color) FromArray(array []float64, offset int) *Color {

	color.R = array[offset]
	color.G = array[offset+1]
	color.B = array[offset+2]

	return color

}

func (color *Color) ToArray(array []float64, offset int) []float64 {

	if array == nil {
		array = make([]float64, offset+3)
	}

	array[offset] = color.R
	array[offset+1] = color.G
	array[offset+2] = color.B

	return array

}

func hue2RGB(p, q, t float64) float64 {

	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*6*(2/3-t)
	}
	return p

}

func handleAlpha(str string) {

	if str == "" {
		return
	}

	if f, _ := strconv.ParseFloat(str, 10); f < 1 {

		glog.Warningf("three.Color: Alpha component of %s will be ignored.", str)

	}

}

func Colors(name string) *Color {
	val := ColorKeywords[name]
	return NewColor().SetHex(val)
}

var ColorKeywords = map[string]int{
	"aliceblue":            0xF0F8FF,
	"antiquewhite":         0xFAEBD7,
	"aqua":                 0x00FFFF,
	"aquamarine":           0x7FFFD4,
	"azure":                0xF0FFFF,
	"beige":                0xF5F5DC,
	"bisque":               0xFFE4C4,
	"black":                0x000000,
	"blanchedalmond":       0xFFEBCD,
	"blue":                 0x0000FF,
	"blueviolet":           0x8A2BE2,
	"brown":                0xA52A2A,
	"burlywood":            0xDEB887,
	"cadetblue":            0x5F9EA0,
	"chartreuse":           0x7FFF00,
	"chocolate":            0xD2691E,
	"coral":                0xFF7F50,
	"cornflowerblue":       0x6495ED,
	"cornsilk":             0xFFF8DC,
	"crimson":              0xDC143C,
	"cyan":                 0x00FFFF,
	"darkblue":             0x00008B,
	"darkcyan":             0x008B8B,
	"darkgoldenrod":        0xB8860B,
	"darkgray":             0xA9A9A9,
	"darkgreen":            0x006400,
	"darkgrey":             0xA9A9A9,
	"darkkhaki":            0xBDB76B,
	"darkmagenta":          0x8B008B,
	"darkolivegreen":       0x556B2F,
	"darkorange":           0xFF8C00,
	"darkorchid":           0x9932CC,
	"darkred":              0x8B0000,
	"darksalmon":           0xE9967A,
	"darkseagreen":         0x8FBC8F,
	"darkslateblue":        0x483D8B,
	"darkslategray":        0x2F4F4F,
	"darkslategrey":        0x2F4F4F,
	"darkturquoise":        0x00CED1,
	"darkviolet":           0x9400D3,
	"deeppink":             0xFF1493,
	"deepskyblue":          0x00BFFF,
	"dimgray":              0x696969,
	"dimgrey":              0x696969,
	"dodgerblue":           0x1E90FF,
	"firebrick":            0xB22222,
	"floralwhite":          0xFFFAF0,
	"forestgreen":          0x228B22,
	"fuchsia":              0xFF00FF,
	"gainsboro":            0xDCDCDC,
	"ghostwhite":           0xF8F8FF,
	"gold":                 0xFFD700,
	"goldenrod":            0xDAA520,
	"gray":                 0x808080,
	"green":                0x008000,
	"greenyellow":          0xADFF2F,
	"grey":                 0x808080,
	"honeydew":             0xF0FFF0,
	"hotpink":              0xFF69B4,
	"indianred":            0xCD5C5C,
	"indigo":               0x4B0082,
	"ivory":                0xFFFFF0,
	"khaki":                0xF0E68C,
	"lavender":             0xE6E6FA,
	"lavenderblush":        0xFFF0F5,
	"lawngreen":            0x7CFC00,
	"lemonchiffon":         0xFFFACD,
	"lightblue":            0xADD8E6,
	"lightcoral":           0xF08080,
	"lightcyan":            0xE0FFFF,
	"lightgoldenrodyellow": 0xFAFAD2,
	"lightgray":            0xD3D3D3,
	"lightgreen":           0x90EE90,
	"lightgrey":            0xD3D3D3,
	"lightpink":            0xFFB6C1,
	"lightsalmon":          0xFFA07A,
	"lightseagreen":        0x20B2AA,
	"lightskyblue":         0x87CEFA,
	"lightslategray":       0x778899,
	"lightslategrey":       0x778899,
	"lightsteelblue":       0xB0C4DE,
	"lightyellow":          0xFFFFE0,
	"lime":                 0x00FF00,
	"limegreen":            0x32CD32,
	"linen":                0xFAF0E6,
	"magenta":              0xFF00FF,
	"maroon":               0x800000,
	"mediumaquamarine":     0x66CDAA,
	"mediumblue":           0x0000CD,
	"mediumorchid":         0xBA55D3,
	"mediumpurple":         0x9370DB,
	"mediumseagreen":       0x3CB371,
	"mediumslateblue":      0x7B68EE,
	"mediumspringgreen":    0x00FA9A,
	"mediumturquoise":      0x48D1CC,
	"mediumvioletred":      0xC71585,
	"midnightblue":         0x191970,
	"mintcream":            0xF5FFFA,
	"mistyrose":            0xFFE4E1,
	"moccasin":             0xFFE4B5,
	"navajowhite":          0xFFDEAD,
	"navy":                 0x000080,
	"oldlace":              0xFDF5E6,
	"olive":                0x808000,
	"olivedrab":            0x6B8E23,
	"orange":               0xFFA500,
	"orangered":            0xFF4500,
	"orchid":               0xDA70D6,
	"palegoldenrod":        0xEEE8AA,
	"palegreen":            0x98FB98,
	"paleturquoise":        0xAFEEEE,
	"palevioletred":        0xDB7093,
	"papayawhip":           0xFFEFD5,
	"peachpuff":            0xFFDAB9,
	"peru":                 0xCD853F,
	"pink":                 0xFFC0CB,
	"plum":                 0xDDA0DD,
	"powderblue":           0xB0E0E6,
	"purple":               0x800080,
	"red":                  0xFF0000,
	"rosybrown":            0xBC8F8F,
	"royalblue":            0x4169E1,
	"saddlebrown":          0x8B4513,
	"salmon":               0xFA8072,
	"sandybrown":           0xF4A460,
	"seagreen":             0x2E8B57,
	"seashell":             0xFFF5EE,
	"sienna":               0xA0522D,
	"silver":               0xC0C0C0,
	"skyblue":              0x87CEEB,
	"slateblue":            0x6A5ACD,
	"slategray":            0x708090,
	"slategrey":            0x708090,
	"snow":                 0xFFFAFA,
	"springgreen":          0x00FF7F,
	"steelblue":            0x4682B4,
	"tan":                  0xD2B48C,
	"teal":                 0x008080,
	"colortle":             0xD8BFD8,
	"tomato":               0xFF6347,
	"turquoise":            0x40E0D0,
	"violet":               0xEE82EE,
	"wheat":                0xF5DEB3,
	"white":                0xFFFFFF,
	"whitesmoke":           0xF5F5F5,
	"yellow":               0xFFFF00,
	"yellowgreen":          0x9ACD32,
}
