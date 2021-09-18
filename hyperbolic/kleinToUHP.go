package hyperbolic

import (
	"coxeter/vector"
	"math"
)

func KleinToUHP(vec [3]float64) [3]float64 {

	return vector.Scale3([3]float64{vec[0], vec[1], math.Sqrt(1.0 - vector.NormSquared(vec[:]))}, 1.0/(1.0-vec[2]))

}
