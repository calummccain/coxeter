package data

import (
	"math"

	"github.com/calummccain/coxeter/shared"
)

func TruncatedOctahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(Rt2)
	pVal := math.Pi / math.Atan(Rt_5)

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1.0 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	third := 1.0 / 3.0
	twoThird := 2.0 / 3.0

	var vv float64

	if metric == 'p' {

		vv = 1.0

	} else {

		vv = cot / math.Abs(1.0-cot)

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
				2.0*v[0] - v[1] - v[2] - v[3],
				v[0] - v[2] - v[3],
				v[0] - v[1] - v[3],
				v[0] - v[1] - v[2],
			}

		}

	} else {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				(6.0*cos-2.0)*(v[0]-v[1]-v[2]-v[3]) + v[0],
				2.0*cos*(v[0]-v[1]-v[2]-v[3]) + v[1],
				2.0*cos*(v[0]-v[1]-v[2]-v[3]) + v[2],
				2.0*cos*(v[0]-v[1]-v[2]-v[3]) + v[3],
			}

		}

	}

	var f func([4]float64) [4]float64
	var a, b float64

	if metric == 'p' {

		a = 3.0 * Rt_5
		b = 3.0 * Rt_5

	} else if metric == 'e' {

		a = 1.0
		b = 1.0

	} else {

		a = 3.0 * math.Sqrt(math.Abs(cot/(5.0-cot)))
		b = 3.0 * math.Sqrt(math.Abs((1.0-2.0*cot)/(5.0-cot)))

	}

	f = func(v [4]float64) [4]float64 {

		return [4]float64{a * v[0], b * v[1], b * v[2], b * v[3]}

	}

	return CellData{
		Metric:          metric,
		NumVertices:     24,
		NumEdges:        36,
		NumFaces:        14,
		FaceReflections: []string{"", "c", "bc", "cbc", "abc", "cabc", "bcabc", "cbcabc"},
		OuterReflection: "d",
		V:               [4]float64{0, 0, 0, 0},
		E:               [4]float64{0, 0, 0, 0},
		F:               [4]float64{0, 0, 0, 0},
		C:               [4]float64{0, 0, 0, 0},
		CellType:        "spherical",
		Vv:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices: [][4]float64{
			{1, twoThird, third, 0}, {1, twoThird, 0, third}, {1, twoThird, -third, 0}, {1, twoThird, 0, -third},
			{1, 0, twoThird, third}, {1, third, twoThird, 0}, {1, 0, twoThird, -third}, {1, -third, twoThird, 0},
			{1, third, 0, twoThird}, {1, 0, third, twoThird}, {1, -third, 0, twoThird}, {1, 0, -third, twoThird},
			{1, -twoThird, third, 0}, {1, -twoThird, 0, third}, {1, -twoThird, -third, 0}, {1, -twoThird, 0, -third},
			{1, 0, -twoThird, third}, {1, third, -twoThird, 0}, {1, 0, -twoThird, -third}, {1, -third, -twoThird, 0},
			{1, third, 0, -twoThird}, {1, 0, third, -twoThird}, {1, -third, 0, -twoThird}, {1, 0, -third, -twoThird},
		},
		Edges: [][2]int{
			{0, 1}, {0, 3}, {0, 5}, {1, 2}, {1, 8}, {2, 3}, {2, 17}, {3, 20},
			{4, 5}, {4, 7}, {4, 9}, {5, 6}, {6, 7}, {6, 21}, {7, 12}, {8, 9},
			{8, 11}, {9, 10}, {10, 11}, {10, 13}, {11, 16}, {12, 13}, {12, 15},
			{13, 14}, {14, 15}, {14, 19}, {15, 22}, {16, 17}, {16, 19}, {17, 18},
			{18, 19}, {18, 23}, {20, 21}, {20, 23}, {21, 22}, {22, 23},
		},
		Faces: [][]int{
			{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11},
			{12, 13, 14, 15}, {16, 17, 18, 19}, {20, 21, 22, 23},
			{0, 1, 8, 9, 4, 5}, {0, 3, 20, 21, 6, 5},
			{4, 7, 12, 13, 10, 9}, {1, 2, 17, 16, 11, 8},
			{2, 3, 20, 23, 18, 17}, {6, 7, 12, 15, 22, 21},
			{10, 11, 16, 19, 14, 13}, {14, 15, 22, 23, 18, 19},
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
