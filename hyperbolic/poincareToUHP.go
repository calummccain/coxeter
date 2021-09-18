package hyperbolic

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func PoincareToUHP(vec [3]float64) [3]float64 {

	eps := PoincareToUHPEps
	r := vector.NormSquared(vec[:])
	s := 1 / (r + 1.0 - 2.0*vec[2])

	if s < eps {

		return [3]float64{vec[0], vec[1], math.Inf(1)}

	} else if r > 1-eps {

		return vector.Scale3([3]float64{vec[0], vec[1], 0}, 2.0*s)

	} else {

		return vector.Scale3([3]float64{vec[0], vec[1], (1.0 - r) * 0.5}, 2.0*s)

	}

}
