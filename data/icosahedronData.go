package data

import (
	"math"

	"github.com/calummccain/coxeter/shared"
)

func IcosahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(P2)
	pVal := 10.0 / 3.0

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	var vv float64

	if metric == 'p' {

		vv = 2 / (P + 2)

	} else {

		vv = (P2 - 1.0 + P4*cot) / math.Abs(P2+1.0-P4*cot)

	}

	var d func([4]float64) [4]float64

	if n == 3 {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				0.5 * ((P4-1.0)*v[0] - (P3-3.0*P_1)*v[2] - (P-3.0*P_3)*v[3]),
				v[1],
				0.5 * (P5*v[0] + (2.0-P4)*v[2] - P2*v[3]),
				0.5 * (P3*v[0] - P2*v[2] + v[3]),
			}

		}

	} else {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				(6.0*P2*cos-1.0)*v[0] + (2.0*P_1-6.0*P*cos)*v[2] + (2.0*P_3-6.0*cos*P_1)*v[3],
				v[1],
				2.0*P5*cos*v[0] + (1.0-2.0*P4*cos)*v[2] - 2.0*P2*cos*v[3],
				2.0*P3*cos*v[0] - 2.0*P2*cos*v[2] + (1.0-2.0*cos)*v[3],
			}

		}

	}

	var f func([4]float64) [4]float64

	f3_1 := P3 / 2.0
	f3_2 := math.Sqrt(3.0*P-1.0) / 2.0

	fp := 1.0 / math.Sqrt(P+2.0)

	fn_1 := P3 * math.Sqrt(math.Abs(cot/(P2+1.0-P4*cot)))
	fn_2 := math.Sqrt(math.Abs((P4*cot - 1.0) / (P2 + 1.0 - P4*cot)))

	if n == 3 {

		f = func(v [4]float64) [4]float64 {

			return [4]float64{f3_1 * v[0], f3_2 * v[1], f3_2 * v[2], f3_2 * v[3]}

		}

	} else if metric == 'e' {

		f = func(v [4]float64) [4]float64 {

			return [4]float64{v[0], v[1], v[2], v[3]}

		}

	} else if metric == 'p' {

		f = func(v [4]float64) [4]float64 {

			return [4]float64{v[0], fp * v[1], fp * v[2], fp * v[3]}

		}

	} else {

		f = func(v [4]float64) [4]float64 {

			return [4]float64{fn_1 * v[0], fn_2 * v[1], fn_2 * v[2], fn_2 * v[3]}

		}

	}

	return CellData{
		Metric:      metric,
		NumVertices: 12,
		NumEdges:    30,
		NumFaces:    20,
		FaceReflections: []string{
			"", "c", "bcbc", "bc", "cbc", "bacbcabacbc",
			"cbacbcabacbc", "bacabacbc", "bcabacbc",
			"cbcabacbc", "abcbc", "abc", "acbc",
			"abacabacbc", "abcabacbc", "cabcabacbc",
			"bacbc", "abacbc", "acabacbc", "cabacbc"},
		OuterReflection: "d",
		V:               [4]float64{1, 1, P, 0},
		E:               [4]float64{1, 0, P, 0},
		F:               [4]float64{3, 0, P3, P},
		C:               [4]float64{1, 0, 0, 0},
		CellType:        "spherical",
		Vv:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices: [][4]float64{
			{1, 1, P, 0}, {1, 1, -P, 0}, {1, -1, P, 0},
			{1, -1, -P, 0}, {1, P, 0, 1}, {1, -P, 0, 1},
			{1, P, 0, -1}, {1, -P, 0, -1}, {1, 0, 1, P},
			{1, 0, 1, -P}, {1, 0, -1, P}, {1, 0, -1, -P},
		},
		Edges: [][2]int{
			{0, 2}, {0, 4}, {0, 6}, {0, 8},
			{0, 9}, {1, 3}, {1, 4}, {1, 6},
			{1, 10}, {1, 11}, {2, 5}, {2, 7},
			{2, 8}, {2, 9}, {3, 5}, {3, 7},
			{3, 10}, {3, 11}, {4, 6}, {4, 8},
			{4, 10}, {5, 7}, {5, 8}, {5, 10},
			{6, 9}, {6, 11}, {7, 9}, {7, 11},
			{8, 10}, {9, 11},
		},
		Faces: [][]int{
			{0, 8, 2}, {0, 2, 9}, {0, 6, 4}, {0, 4, 8},
			{0, 9, 6}, {1, 3, 10}, {1, 11, 3}, {1, 4, 6},
			{1, 10, 4}, {1, 6, 11}, {2, 5, 7}, {2, 8, 5},
			{2, 7, 9}, {3, 7, 5}, {3, 5, 10}, {3, 11, 7},
			{4, 10, 8}, {5, 8, 10}, {6, 9, 11}, {7, 11, 9},
		},
		Matrices: shared.Matrices{
			A: func(v [4]float64) [4]float64 { return [4]float64{v[0], -v[1], v[2], v[3]} },
			B: func(v [4]float64) [4]float64 {
				return [4]float64{v[0], 0.5 * (v[1] + P_1*v[2] - P*v[3]), 0.5 * (P_1*v[1] + P*v[2] + v[3]), 0.5 * (-P*v[1] + v[2] - P_1*v[3])}
			},
			C: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[1], v[2], -v[3]} },
			D: d,
			E: func(v [4]float64) [4]float64 { return v },
			F: f,
		},
		Flip: func(v [4]float64) [4]float64 { return [4]float64{-v[0], v[1], v[2], v[3]} },
	}
}
