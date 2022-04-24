package data

import (
	"github.com/calummccain/coxeter/vector"
)

func GoursatTetrahedron43n(n float64) GoursatTetrahedron {

	//cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	gt := GoursatTetrahedron{
		P:      4.0,
		Q:      3.0,
		R:      n,
		V:      vector.Vec4{W: 1, X: 1, Y: 1, Z: 1},
		E:      vector.Vec4{W: 1, X: 1, Y: 1, Z: 0},
		F:      vector.Vec4{W: 1, X: 1, Y: 0, Z: 0},
		C:      vector.Vec4{W: 1, X: 0, Y: 0, Z: 0},
		Metric: "s",
		Scale: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W / 2.0, X: v.X / 2.0, Y: v.Y / 2.0, Z: v.Z / 2.0}
		},
	}

	gt.Populate()

	//TODO SCALE
	return gt

}
