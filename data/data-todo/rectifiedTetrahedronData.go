package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func RectifiedTetrahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(Rt_2)
	pVal := math.Inf(1)

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	vv := cot / 2.0

	var d func(vector.Vec4) vector.Vec4

	if n == 3 {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				-0.25*v.W + 1.25*v.X + 1.25*v.Y - 1.25*v.Z,
				0.25*v.W + 0.75*v.X - 0.25*v.Y + 0.25*v.Z,
				0.25*v.W - 0.25*v.X + 0.75*v.Y + 0.25*v.Z,
				-0.25*v.W + 0.25*v.X + 0.25*v.Y + 0.75*v.Z,
			}

		}

	} else if n == 4 {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				0.5 * (v.W + v.X + v.Y - v.Z),
				0.5 * (v.W + v.X - v.Y + v.Z),
				0.5 * (v.W - v.X + v.Y + v.Z),
				0.5 * (-v.W + v.X + v.Y + v.Z),
			}

		}

	} else if n == 6 {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				1.25*v.W - 0.25*v.X - 0.25*v.Y + 0.25*v.Z,
				0.75*v.W + 0.25*v.X - 0.75*v.Y + 0.75*v.Z,
				0.75*v.W - 0.75*v.X + 0.25*v.Y + 0.75*v.Z,
				-0.75*v.W + 0.75*v.X + 0.75*v.Y + 0.25*v.Z,
			}

		}

	} else {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				(3*sin-1)*(-v.W+v.X+v.Y-v.Z) + v.W,
				cos*(v.W-v.X-v.Y+v.Z) + v.X,
				cos*(v.W-v.X-v.Y+v.Z) + v.Y,
				cos*(-v.W+v.X+v.Y-v.Z) + v.Z,
			}

		}

	}

	var f func(vector.Vec4) vector.Vec4
	var a, b float64

	if metric == 'e' {

		a = 1
		b = 1

	} else {

		a = math.Sqrt(math.Abs(cot / 2.0))
		b = math.Sqrt(math.Abs((cot - 2) / 2.0))

	}

	f = func(v vector.Vec4) vector.Vec4 {

		return vector.Vec4{a * v.W, b * v.X, b * v.Y, b * v.Z}

	}

	return CellData{
		Metric:          metric,
		NumVertices:     6,
		NumEdges:        12,
		NumFaces:        8,
		FaceReflections: []string{"", "abc", "bc", "c"},
		OuterReflection: "d",
		V:               vector.Vec4{0, 0, 0, 0},
		E:               vector.Vec4{0, 0, 0, 0},
		F:               vector.Vec4{0, 0, 0, 0},
		C:               vector.Vec4{0, 0, 0, 0},
		CellType:        "spherical",
		VV:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices: []vector.Vec4{
			{1, 1, 0, 0},
			{1, -1, 0, 0},
			{1, 0, 1, 0},
			{1, 0, -1, 0},
			{1, 0, 0, 1},
			{1, 0, 0, -1},
		},
		Edges: [][2]int{
			{0, 2}, {0, 3}, {0, 4}, {0, 5},
			{1, 2}, {1, 3}, {1, 4}, {1, 5},
			{2, 4}, {4, 3}, {3, 5}, {5, 2},
		},
		Faces: [][]int{
			{0, 2, 4}, {0, 5, 2},
			{0, 4, 3}, {0, 3, 5},
			{1, 4, 2}, {1, 2, 5},
			{1, 3, 4}, {1, 5, 3},
		},
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, -v.Z, -v.Y} },
		Bmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.Y, v.X, v.Z} },
		Cmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Z, v.Y} },
		Dmat: d,
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}
}
