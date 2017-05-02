package math3_test

import (
	"math"
	"testing"

	"github.com/rydrman/marshmallow"
)

func validateColor(c *mm.Color, r, g, b float64, t *testing.T) {

	if math.Abs(c.R-r) > 0.0001 {

		t.Errorf("expected r to be %0.2f, got: %0.2f", r, c.R)

	}
	if math.Abs(c.G-g) > 0.0001 {

		t.Errorf("expected g to be %0.2f, got: %0.2f", g, c.G)

	}
	if math.Abs(c.B-b) > 0.0001 {

		t.Errorf("expected b to be %0.2f, got: %0.2f", b, c.B)

	}

}

func TestNewColor(t *testing.T) {
	c := mm.NewColor()
	if nil == c {
		t.Fail()
	}
}

/*func TestcopyHex(t *testing.T) {
    c := mm.NewColor()
    c2 := mm.NewColor(0xF5FFFA)
    c.Copy(c2)
    ok(c.GetHex() == c2.GetHex(), "Hex c: " + c.GetHex() + " Hex c2: " + c2.GetHex())
}

func TestcopyColorString(t *testing.T) {
    c := mm.NewColor()
    c2 := mm.NewColor("ivory")
    c.Copy(c2)
    ok(c.GetHex() == c2.GetHex(), "Hex c: " + c.GetHex() + " Hex c2: " + c2.GetHex())
}*/

func TestColor_SetRGB(t *testing.T) {
	c := mm.NewColor()
	c.SetRGB(1, 0.2, 0.1)
	validateColor(c, 1, 0.2, 0.1, t)
}

func TestColor_CopyGammaToLinear(t *testing.T) {
	c := mm.NewColor()
	c2 := mm.NewColor()
	c2.SetRGB(0.3, 0.5, 0.9)
	c.CopyGammaToLinear(c2)
	validateColor(c, 0.09, 0.25, 0.81, t)
}

func TestColor_CopyLinearToGamma(t *testing.T) {
	c := mm.NewColor()
	c2 := mm.NewColor()
	c2.SetRGB(0.09, 0.25, 0.81)
	c.CopyLinearToGamma(c2)
	validateColor(c, 0.3, 0.5, 0.9, t)
}

func TestColor_ConvertGammaToLinear(t *testing.T) {
	c := mm.NewColor()
	c.SetRGB(0.3, 0.5, 0.9)
	c.ConvertGammaToLinear()
	validateColor(c, 0.09, 0.25, 0.81, t)
}

func Test_ConvertLinearToGamma(t *testing.T) {
	c := mm.NewColor()
	c.SetRGB(4, 9, 16)
	c.ConvertLinearToGamma()
	validateColor(c, 2, 3, 4, t)
}

func TestColor_SetStyle(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("silver")
	g := 0.75294117647
	validateColor(c, g, g, g, t)
}

func TestColor_Clone(t *testing.T) {
	c := mm.NewColor().SetStyle("teal")
	c2 := c.Clone()
	validateColor(c2, c.R, c.G, c.B, t)
}

func TestColor_Lerp(t *testing.T) {
	c := mm.NewColor()
	c2 := mm.NewColor()
	c.SetRGB(0, 0, 0)
	c.Lerp(c2, 0.2)
	validateColor(c, 0.2, 0.2, 0.2, t)
}

func TestColor_GetHex(t *testing.T) {
	c := mm.NewColor().SetStyle("red")
	res := c.GetHex()
	if res != 0xFF0000 {
		t.Error("expected %x, got %x", 0xFF0000, res)
	}
}

func TestColor_SetHex(t *testing.T) {
	c := mm.NewColor()
	c.SetHex(0xFA8072)
	res := c.GetHex()
	if res != 0xFA8072 {
		t.Error("expected %x, got %x", 0xFA8072, res)
	}
}

func TestColor_SetStyle_RGBRed(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("rgb(255,0,0)")
	validateColor(c, 1, 0, 0, t)
}

func TestColor_SetStyle_RGBARed(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("rgba(255,0,0, 0.5)")
	validateColor(c, 1, 0, 0, t)
}

func TestColor_SetStyle_RGBRedWithSpaces(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("rgb( 255 , 0,   0 )")
	validateColor(c, 1, 0, 0, t)
}

func TestColor_SetStyle_RGBARedWithSpaces(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("rgba( 255,  0,  0  , 1 )")
	validateColor(c, 1, 0, 0, t)
}

func TestColor_SetStyle_RGBPercent(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("rgb(100%,50%,10%)")
	validateColor(c, 1, 0.5, 0.1, t)
}

func TestColor_SetStyle_RGBAPercent(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("rgba(100%,50%,10%, 0.5)")
	validateColor(c, 1, 0.5, 0.1, t)
}

func TestColor_SetStyle_RGBPercentWithSpaces(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("rgb( 100% ,50%  , 10% )")
	validateColor(c, 1, 0.5, 0.1, t)
}

func TestColor_SetStyle_RGBAPercentWithSpaces(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("rgba( 100% ,50%  ,  10%, 0.5 )")
	validateColor(c, 1, 0.5, 0.1, t)
}

func TestColor_SetStyle_HSLRed(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("hsl(360,100%,50%)")
	validateColor(c, 1, 0, 0, t)
}

func TestColor_SetStyle_HSLARed(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("hsla(360,100%,50%,0.5)")
	validateColor(c, 1, 0, 0, t)
}

func TestColor_SetStyle_HSLRedWithSpaces(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("hsl(360,  100% , 50% )")
	validateColor(c, 1, 0, 0, t)
}

func TestColor_SetStyle_HSLARedWithSpaces(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("hsla( 360,  100% , 50%,  0.5 )")
	validateColor(c, 1, 0, 0, t)
}

func TestColor_SetStyle_HexSkyBlue(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("#87CEEB")
	res := c.GetHex()
	if res != 0x87CEEB {
		t.Error("expected %x, got %x", 0x87CEEB, res)
	}
}

func TestColor_SetStyle_HexSkyBlueMixed(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("#87cEeB")
	res := c.GetHex()
	if res != 0x87CEEB {
		t.Error("expected %x, got %x", 0x87CEEB, res)
	}
}

func TestColor_SetStyle_Hex2Olive(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("#F00")
	res := c.GetHex()
	if res != 0xFF0000 {
		t.Error("expected %x, got %x", 0xFF0000, res)
	}
}

func TestColor_SetStyle_Hex2OliveMixed(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("#f00")
	res := c.GetHex()
	if res != 0xFF0000 {
		t.Error("expected %x, got %x", 0xFF0000, res)
	}
}

func TestColor_SetStyle_ColorName(t *testing.T) {
	c := mm.NewColor()
	c.SetStyle("powderblue")
	res := c.GetHex()
	if res != 0xB0E0E6 {
		t.Error("expected %x, got %x", 0xB0E0E6, res)
	}
}

func TestColor_GetHexString(t *testing.T) {
	c := mm.NewColor().SetStyle("tomato")
	res := c.GetHexString()
	if res != "ff6347" {
		t.Errorf("expected ff6347, got: %s", res)
	}
}

func TestColor_GetStyle(t *testing.T) {
	c := mm.Colors("plum")
	res := c.GetStyle()
	if res != "rgb(221,160,221)" {
		t.Errorf("expected rgb(221,160,221), got %s", res)
	}
}

func TestColor_GetHSL(t *testing.T) {
	c := mm.NewColor().SetHex(0x80ffff)
	h, s, l := c.GetHSL()

	if h != 0.5 || s != 1.0 || math.Abs(l-0.75) > 0.001 {
		t.Errorf("expected 0.5, 1.0, 0.75, got: %0.2f, %0.2f, %0.2f", h, s, l)
	}
}

func TestColor_SetHSL(t *testing.T) {
	c := mm.NewColor()
	c.SetHSL(0.75, 1.0, 0.25)
	h, s, l := c.GetHSL()

	if h != 0.75 || s != 1.0 || l != 0.25 {
		t.Errorf("expected 0.75, 1.0, 0.25, got: %0.2f, %0.2f, %0.2f", h, s, l)
	}
}
