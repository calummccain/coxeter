package hyperbolic

import (
	"coxeter/vector"
	"math"
)

func GeodesicEndpoints(a, b [4]float64, l float64) ([4]float64, [4]float64) {

	cosh := -l
	sinh := math.Sqrt(l*l - 1.0)

	p1 := vector.Scale4(vector.Sum4(vector.Scale4(a, sinh-cosh), b), 1.0/sinh)
	p2 := vector.Scale4(vector.Diff4(b, vector.Scale4(a, sinh+cosh)), 1.0/sinh)

	return p1, p2

}
