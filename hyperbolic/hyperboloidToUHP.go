package hyperbolic

import (
	"coxeter/vector"
	"math"
)

func HyperboloidToUHP(vec [4]float64) [3]float64 {

	eps := HyperboloidToUHPEps

	if math.Abs(vec[0]-vec[3]) < eps {

		return [3]float64{vec[1] / vec[0], vec[2] / vec[0], math.Inf(1)}

	} else if HyperbolicNorm(vec) < eps {

		return vector.Scale3([3]float64{vec[1], vec[2], 0.0}, 1.0/(vec[0]-vec[3]))

	} else {

		return vector.Scale3([3]float64{vec[1], vec[2], 1.0}, 1.0/(vec[0]-vec[3]))

	}

}
