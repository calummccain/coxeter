package data

import (
	"math"

	"github.com/calummccain/coxeter/shared"
)

func TruncatedHexahedronData(n float64) CellData {

	eVal := 4.0
	pVal := math.Pi / math.Atan(math.Sqrt(1.0/(7.0+4.0*Rt2)))

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1.0 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	factor := Rt2 - 1.0

	var vv float64

	if metric == 'p' {

		vv = 2.0 / (7.0 + 4.0*Rt2)

	} else {

		vv = ((3.0-2.0*Rt2)*cot + 2.0*Rt2 - 1.0) / math.Abs(-(3.0-2.0*Rt2)*cot+5.0-2.0*Rt2)

	}

	var d func([4]float64) [4]float64

	if n == 3 {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				v[1],
				v[0],
				v[2],
				v[3],
			}

		}

	} else if n == 4 {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				v[0],
				2.0*v[0] - v[1],
				v[2],
				v[3],
			}

		}

	} else if n == 5 {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				P*v[0] - P_1*v[1],
				P2*v[0] - P*v[1],
				v[2],
				v[3],
			}

		}

	} else if n == 6 {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				2.0*v[0] - v[1],
				3.0*v[0] - 2.0*v[1],
				v[2],
				v[3],
			}

		}

	} else {

		d = func(v [4]float64) [4]float64 {

			c := math.Cos(2.0 * math.Pi / n)

			return [4]float64{
				(1.0+2.0*c)*v[0] - 2.0*c*v[1],
				2.0 + 2.0*c*v[0] - (1.0+2.0*c)*v[1],
				v[2],
				v[3],
			}

		}

	}

	var f func([4]float64) [4]float64
	var a, b float64

	if n == 4 {

		a = 1.0
		b = 1.0

	} else if metric == 'p' {

		a = 1.0
		b = 1.0 / math.Sqrt(5.0-2.0*Rt2)

	} else {

		a = math.Sqrt(math.Abs(2.0 * cot / (5.0 - 2.0*Rt2 - (3.0-2.0*Rt2)*cot)))
		b = math.Sqrt(math.Abs((cot - 1.0) / (5.0 - 2.0*Rt2 - (3.0-2.0*Rt2)*cot)))

	}

	f = func(v [4]float64) [4]float64 {

		return [4]float64{a * v[0], b * v[1], b * v[2], b * v[3]}

	}

	return CellData{
		Metric:          metric,
		NumVertices:     24,
		NumEdges:        36,
		NumFaces:        14,
		FaceReflections: []string{"bc", "c", "cbabc", "abc", "", "babc"},
		OuterReflection: "d",
		V:               [4]float64{1, 1, 1, 1},
		E:               [4]float64{1, 1, 1, 0},
		F:               [4]float64{1, 1, 0, 0},
		C:               [4]float64{1, 0, 0, 0},
		CellType:        "spherical",
		Vv:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices: [][4]float64{
			{1, 1, 1, factor}, {1, 1, factor, 1}, {1, factor, 1, 1},
			{1, 1, 1, -factor}, {1, 1, factor, -1}, {1, factor, 1, -1},
			{1, 1, -1, factor}, {1, 1, -factor, 1}, {1, factor, -1, 1},
			{1, 1, -1, -factor}, {1, 1, -factor, -1}, {1, factor, -1, -1},
			{1, -1, 1, factor}, {1, -1, factor, 1}, {1, -factor, 1, 1},
			{1, -1, 1, -factor}, {1, -1, factor, -1}, {1, -factor, 1, -1},
			{1, -1, -1, factor}, {1, -1, -factor, 1}, {1, -factor, -1, 1},
			{1, -1, -1, -factor}, {1, -1, -factor, -1}, {1, -factor, -1, -1},
		},
		Edges: [][2]int{
			{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 7}, {2, 14}, {3, 4}, {3, 5},
			{4, 5}, {4, 10}, {5, 17}, {6, 7}, {6, 8}, {6, 9}, {7, 8}, {8, 20},
			{9, 10}, {9, 11}, {10, 11}, {11, 23}, {12, 13}, {12, 14}, {12, 15}, {13, 14},
			{13, 19}, {15, 16}, {15, 17}, {16, 17}, {16, 22}, {18, 19}, {18, 20}, {18, 21},
			{19, 20}, {21, 22}, {21, 23}, {22, 23},
		},
		Faces: [][]int{
			{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, 10, 11},
			{12, 13, 14}, {15, 16, 17}, {18, 19, 20}, {21, 22, 23},
			{0, 1, 7, 6, 9, 10, 4, 3}, {0, 2, 14, 12, 15, 17, 5, 3},
			{1, 2, 14, 13, 19, 20, 8, 7}, {6, 8, 20, 18, 21, 23, 11, 9},
			{4, 5, 17, 16, 22, 23, 11, 10}, {12, 13, 19, 18, 21, 22, 16, 15},
		},
		Matrices: shared.Matrices{
			A: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[1], v[2], -v[3]} },
			B: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[1], v[3], v[2]} },
			C: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[2], v[1], v[3]} },
			D: d,
			E: func(v [4]float64) [4]float64 { return v },
			F: f,
		},
		Flip: func(v [4]float64) [4]float64 { return [4]float64{-v[0], v[1], v[2], v[3]} },
	}
}
