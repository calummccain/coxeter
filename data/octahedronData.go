package data

import (
	"math"

	"github.com/calummccain/coxeter/shared"
)

func OctahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(Rt2)
	pVal := 4.0

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	var vv float64

	if metric == 'p' {

		vv = 1

	} else {

		vv = cot / math.Abs(1-cot)

	}

	var d func([4]float64) [4]float64

	if n == 3 {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				0.5 * (v[0] + v[1] + v[2] + v[3]),
				0.5 * (v[0] + v[1] - v[2] - v[3]),
				0.5 * (v[0] - v[1] + v[2] - v[3]),
				0.5 * (v[0] - v[1] - v[2] + v[3]),
			}

		}

	} else if n == 4 {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				2*v[0] - v[1] - v[2] - v[3],
				v[0] - v[2] - v[3],
				v[0] - v[1] - v[3],
				v[0] - v[1] - v[2],
			}

		}

	} else {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				(6*cos-2)*(v[0]-v[1]-v[2]-v[3]) + v[0],
				2*cos*(v[0]-v[1]-v[2]-v[3]) + v[1],
				2*cos*(v[0]-v[1]-v[2]-v[3]) + v[2],
				2*cos*(v[0]-v[1]-v[2]-v[3]) + v[3],
			}

		}

	}

	var f func([4]float64) [4]float64
	var a, b float64

	if n == 3 {

		a = Rt_2
		b = Rt_2

	} else if n == 4 {

		a = 1
		b = 1

	} else if metric == 'e' {

		a = 1
		b = 1

	} else {

		a = math.Sqrt(math.Abs(cot / (1 - cot)))
		b = math.Sqrt(math.Abs((1 - 2*cot) / (1 - cot)))

	}

	f = func(v [4]float64) [4]float64 {

		return [4]float64{a * v[0], b * v[1], b * v[2], b * v[3]}

	}

	return CellData{
		Metric:          metric,
		NumVertices:     6,
		NumEdges:        12,
		NumFaces:        8,
		FaceReflections: []string{"", "c", "bc", "cbc", "abc", "cabc", "bcabc", "cbcabc"},
		OuterReflection: "d",
		V:               [4]float64{1, 1, 0, 0},
		E:               [4]float64{2, 1, 1, 0},
		F:               [4]float64{3, 1, 1, 1},
		C:               [4]float64{1, 0, 0, 0},
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
			A: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[2], v[1], v[3]} },
			B: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[1], v[3], v[2]} },
			C: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[1], v[2], -v[3]} },
			D: d,
			E: func(v [4]float64) [4]float64 { return v },
			F: f,
		},
		Flip: func(v [4]float64) [4]float64 { return [4]float64{-v[0], v[1], v[2], v[3]} },
	}
}
