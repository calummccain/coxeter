package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func GoursatTetrahedron34n(n float64) GoursatTetrahedron {

	tan := math.Tan(math.Pi / n)
	ipVal := math.Sqrt(math.Abs(tan*tan - 2.0))

	gt := GoursatTetrahedron{
		P:      3.0,
		Q:      4.0,
		R:      n,
		V:      vector.Vec4{W: 1, X: 1, Y: 0, Z: 0},
		E:      vector.Vec4{W: 2, X: 1, Y: 1, Z: 0},
		F:      vector.Vec4{W: 3, X: 1, Y: 1, Z: 1},
		C:      vector.Vec4{W: 1, X: 0, Y: 0, Z: 0},
		Metric: "s",
		Scale: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W * Rt2, X: v.X * ipVal, Y: v.Y * ipVal, Z: v.Z * ipVal}
		},
	}

	gt.Populate()

	//TODO Mertic and IP
	return gt

}
