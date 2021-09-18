package hyperbolic

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func KleinToPoincare(vec [3]float64) [3]float64 {

	eps := KleinToPoincareEps

	if vector.Norm3(vec) < 1-eps {

		return vec

	} else {

		return vector.Scale3(vec, 1.0/(1.0+math.Sqrt(1.0-vector.NormSquared(vec[:]))))

	}

}
