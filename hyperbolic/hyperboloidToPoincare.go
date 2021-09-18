package hyperbolic

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func HyperboloidToPoincare(vec [4]float64) [3]float64 {

	eps := HyperboloidToPoincareEps

	var inv float64

	norm := HyperbolicNorm(vec)

	if math.Abs(norm) < eps {

		inv = 1.0 / vec[0]

	} else if norm > eps {

		inv = 1.0 / (1.0 + vec[0])

	} else {

		inv = 1.0 / vector.Norm(vec[1:4])

	}

	return [3]float64{vec[1] * inv, vec[2] * inv, vec[3] * inv}

}
