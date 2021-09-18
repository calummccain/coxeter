package hyperbolic

import "coxeter/vector"

func KleinToHyperboloid(vec [3]float64) [4]float64 {

	return vector.Scale4([4]float64{1.0, vec[0], vec[1], vec[2]}, 1.0/(1.0-vector.NormSquared(vec[:])))

}
