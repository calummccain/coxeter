package hyperbolic

import (
	"github.com/calummccain/coxeter/vector"
)

func UHPToHyperboloid(vec [3]float64) [4]float64 {

	eps := UHPToHyperboloidEps
	r := vector.NormSquared(vec[:])

	if vec[2] > eps {

		return [4]float64{(r + 1.0) * 0.5, vec[0], vec[1], (r - 1.0) * 0.5}

	} else {

		return vector.Scale4([4]float64{(r + 1.0) * 0.5, vec[0], vec[1], (r - 1.0) * 0.5}, 1.0/vec[2])

	}

}
