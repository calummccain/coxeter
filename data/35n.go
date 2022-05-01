package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func GoursatTetrahedron35n(n float64) GoursatTetrahedron {

	tan := math.Tan(math.Pi / n)
	ipVal := math.Sqrt(math.Abs(tan*tan - P4))

	gt := GoursatTetrahedron{
		P:    3.0,
		Q:    5.0,
		R:    n,
		V:    vector.Vec4{W: 1, X: 1, Y: P, Z: 0},
		E:    vector.Vec4{W: 1, X: 0, Y: P, Z: 0},
		F:    vector.Vec4{W: 3, X: 0, Y: P3, Z: P},
		C:    vector.Vec4{W: 1, X: 0, Y: 0, Z: 0},
		EVal: 2.6051148443, //math.Pi / math.Atan(phi^2)
		PVal: 3.3333333333, // 10 / 3
		Scale: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W * P3, X: v.X * ipVal, Y: v.Y * ipVal, Z: v.Z * ipVal}
		},
	}

	gt.Metric = Boundaries(n, gt.EVal, gt.PVal)

	if gt.Metric == "e" {
		gt.Scale = func(v vector.Vec4) vector.Vec4 {
			return v
		}
	}

	gt.Populate()

	return gt

}
