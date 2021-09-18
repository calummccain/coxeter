package hyperbolic

import (
	"coxeter/vector"
	"math"
)

func PoincareToHyperboloid(vec [3]float64) [4]float64 {

	eps := PoincareToHyperboloidEps
	r := vector.NormSquared(vec[:])

	if math.Abs(r-1) < eps {

		return [4]float64{1, vec[0], vec[1], vec[2]}

	} else {

		return vector.Scale4([4]float64{(1.0 + r) * 0.5, vec[0], vec[1], vec[2]}, 2.0/(1.0-r))

	}

}
