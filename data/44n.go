package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func GoursatTetrahedron44n(n float64) GoursatTetrahedron {

	cos := math.Pow(math.Cos(math.Pi/n), 2.0)

	gt := GoursatTetrahedron{
		P:    4.0,
		Q:    4.0,
		R:    n,
		V:    vector.Vec4{W: 1, X: 0, Y: 2, Z: 0},
		E:    vector.Vec4{W: 2, X: 0, Y: 1, Z: 1},
		F:    vector.Vec4{W: 1, X: 0, Y: 0, Z: 0},
		C:    vector.Vec4{W: cos, X: 1, Y: 0, Z: 0},
		EVal: 2.0, // not really
		PVal: 4.0,
	}

	gt.Metric = Boundaries(n, gt.EVal, gt.PVal)

	if gt.Metric == "p" {
		gt.Scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: 2.0 * v.W, X: v.X, Y: 2.0 * v.Y, Z: 2.0 * v.Z}
		}
	} else {
		den := 1.0 / math.Sqrt(math.Abs(1.0-2.0*cos))
		gt.Scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W * den, X: cos * v.X * den, Y: math.Sqrt(cos) * v.Y * den, Z: math.Sqrt(cos) * v.Z * den}
		}
	}

	gt.Populate()

	return gt

}
