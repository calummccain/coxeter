package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func GoursatTetrahedron43n(n float64) GoursatTetrahedron {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)
	cn := math.Cos(2.0 * math.Pi / n)

	return GoursatTetrahedron{
		P: 4.0,
		Q: 3.0,
		R: n,
		V: vector.Vec4{W: 1, X: 1, Y: 1, Z: 1},
		E: vector.Vec4{W: 1, X: 1, Y: 1, Z: 0},
		F: vector.Vec4{W: 1, X: 1, Y: 0, Z: 0},
		C: vector.Vec4{W: 1, X: 0, Y: 0, Z: 0},
		Scale: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: 2.0 * cot * v.W, X: v.X, Y: v.Y, Z: v.Z}
		},
		IP: func(v, w vector.Vec4) float64 { return v.Dot(w) / 4.0 },
		Matrices: ReflectionMatrices{
			A: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: -v.Z} },
			B: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Z, Z: v.Y} },
			C: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.Y, Y: v.X, Z: v.Z} },
			D: func(v vector.Vec4) vector.Vec4 {
				return vector.Vec4{
					W: (1.0+2.0*cn)*v.W - 2.0*cn*v.X,
					X: (2.0+2.0*cn)*v.W - (1.0+2.0*cn)*v.X,
					Y: v.Y,
					Z: v.Z,
				}
			},
		},
		FaceReflections: []string{"bc", "c", "cbabc", "abc", "", "babc"},
	}

}
