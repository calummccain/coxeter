package hyperbolic

import (
	"github.com/calummccain/coxeter/vector"
)

func PoincareToKlein(vec [3]float64) [3]float64 {

	return vector.Scale3(vec, 2.0/(1.0+vector.NormSquared(vec[:])))

}
