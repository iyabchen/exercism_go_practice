package triangle

import "math"

const testVersion = 3

// Notice it returns this type.  Pick something suitable.
type Kind int

// Pick values for the following identifiers used by the test program.
const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// a + b == c is ok to be a triangle, but a+b < c is Nat
func KindFromSides(a, b, c float64) Kind {
	sides := [...]float64{a, b, c}
	for _, s := range sides {
		if math.IsNaN(s) || math.IsInf(s, 0) || s <= 0 {
			return NaT
		}
	}
	if a+b < c || a+c < b || b+c < a {
		return NaT
	} else if a == b && b == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	}
	return Sca

}
