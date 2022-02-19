package hyperbolic

import (
	"math"
)

func GeodesicEndpoints(a, b HPoint, l float64) (HPoint, HPoint) {

	cosh := -l
	sinh := math.Sqrt(l*l - 1.0)

	p1 := HPoint{H: Scale4(Sum4(Scale4(a.H, sinh-cosh), b.H), 1.0/sinh)}
	p2 := HPoint{H: Scale4(Diff4(Scale4(a.H, sinh+cosh), b.H), 1.0/sinh)}

	return p1, p2

}
