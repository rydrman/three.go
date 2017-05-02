/**
 * Ported from three.js by @rydrman
 */

package core

type Uniform struct {
	Value int
}

func NewUniform(value int) *Uniform {

	return &Uniform{
		Value: value,
	}

}

func (u *Uniform) Clone() *Uniform {

	return NewUniform(u.Value)

}
