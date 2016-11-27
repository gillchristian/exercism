package triangle

import (
	"math"
	"sort"
)

const testVersion = 3

type Kind struct {
	string
}

var NaT = Kind{"NaT"} // Not a triangle.
var Equ = Kind{"Equ"} // Equilateral.
var Iso = Kind{"Iso"} // Isosceles.
var Sca = Kind{"Sca"} // Scalene.

func KindFromSides(a, b, c float64) Kind {
	if isNat([]float64{a, b, c}) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}

func isNat(sides []float64) bool {
	hasInf := math.IsInf(sides[0], 0) || math.IsInf(sides[1], 0) || math.IsInf(sides[2], 0)
	hasNan := math.IsNaN(sides[0]) || math.IsNaN(sides[1]) || math.IsNaN(sides[2])
	hasNegative := sides[0] <= 0 || sides[1] <= 0 || sides[2] <= 0
	if hasInf || hasNan || hasNegative {
		return true
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(sides)))

	if sides[0] > (sides[1] + sides[2]) {
		return true
	}
	return false
}
