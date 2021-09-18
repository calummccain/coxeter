package hyperbolic

import (
	"coxeter/vector"
)

func UHPToKlein(vec [3]float64) [3]float64 {

	r := vector.NormSquared(vec[:])

	return vector.Scale3([3]float64{vec[0], vec[1], (r - 1.0) * 0.5}, 2.0/(r+1.0))

}
