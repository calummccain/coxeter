package data

import (
	"math"

	"github.com/calummccain/coxeter/shared"
)

func RectifiedTetrahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(Rt_2)
	pVal := math.Inf(1)

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	vv := cot / 2.0

	var d func([4]float64) [4]float64

	if n == 3 {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				-0.25*v[0] + 1.25*v[1] + 1.25*v[2] - 1.25*v[3],
				0.25*v[0] + 0.75*v[1] - 0.25*v[2] + 0.25*v[3],
				0.25*v[0] - 0.25*v[1] + 0.75*v[2] + 0.25*v[3],
				-0.25*v[0] + 0.25*v[1] + 0.25*v[2] + 0.75*v[3],
			}

		}

	} else if n == 4 {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				0.5 * (v[0] + v[1] + v[2] - v[3]),
				0.5 * (v[0] + v[1] - v[2] + v[3]),
				0.5 * (v[0] - v[1] + v[2] + v[3]),
				0.5 * (-v[0] + v[1] + v[2] + v[3]),
			}

		}

	} else if n == 6 {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				1.25*v[0] - 0.25*v[1] - 0.25*v[2] + 0.25*v[3],
				0.75*v[0] + 0.25*v[1] - 0.75*v[2] + 0.75*v[3],
				0.75*v[0] - 0.75*v[1] + 0.25*v[2] + 0.75*v[3],
				-0.75*v[0] + 0.75*v[1] + 0.75*v[2] + 0.25*v[3],
			}

		}

	} else {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				(3*sin-1)*(-v[0]+v[1]+v[2]-v[3]) + v[0],
				cos*(v[0]-v[1]-v[2]+v[3]) + v[1],
				cos*(v[0]-v[1]-v[2]+v[3]) + v[2],
				cos*(-v[0]+v[1]+v[2]-v[3]) + v[3],
			}

		}

	}

	var f func([4]float64) [4]float64
	var a, b float64

	if metric == 'e' {

		a = 1
		b = 1

	} else {

		a = math.Sqrt(math.Abs(cot / 2.0))
		b = math.Sqrt(math.Abs((cot - 2) / 2.0))

	}

	f = func(v [4]float64) [4]float64 {

		return [4]float64{a * v[0], b * v[1], b * v[2], b * v[3]}

	}

	return CellData{
		Metric:          metric,
		NumVertices:     6,
		NumEdges:        12,
		NumFaces:        8,
		FaceReflections: []string{"", "abc", "bc", "c"},
		OuterReflection: "d",
		V:               [4]float64{0, 0, 0, 0},
		E:               [4]float64{0, 0, 0, 0},
		F:               [4]float64{0, 0, 0, 0},
		C:               [4]float64{0, 0, 0, 0},
		CellType:        "spherical",
		Vv:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices: [][4]float64{
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
		Matrices: shared.Matrices{
			A: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[1], -v[3], -v[2]} },
			B: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[2], v[1], v[3]} },
			C: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[1], v[3], v[2]} },
			D: d,
			E: func(v [4]float64) [4]float64 { return v },
			F: f,
		},
		Flip: func(v [4]float64) [4]float64 { return [4]float64{-v[0], v[1], v[2], v[3]} },
	}
}
