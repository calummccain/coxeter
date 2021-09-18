package data

import (
	"coxeter/shared"
	"math"
)

func RectifiedIcosahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(P2)
	pVal := math.Inf(1)

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1.0 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	vv := 0.5*P2*cot + 0.5*P

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
	var a, b float64

	if metric == 'e' {

		a = 1
		b = 1

	} else {

		a = P2 * math.Sqrt(math.Abs(cot))
		b = math.Sqrt(math.Abs(P2*cot - P_2))

	}

	f = func(v [4]float64) [4]float64 {

		return [4]float64{a * v[0], b * v[1], b * v[2], b * v[3]}

	}

	return CellData{
		Metric:      metric,
		NumVertices: 30,
		NumEdges:    60,
		NumFaces:    32,
		FaceReflections: []string{
			"", "c", "bcbc", "bc", "cbc", "bacbcabacbc",
			"cbacbcabacbc", "bacabacbc", "bcabacbc",
			"cbcabacbc", "abcbc", "abc", "acbc",
			"abacabacbc", "abcabacbc", "bacaacbcabacbcbcbaca",
			"bacbc", "abacbc", "acabacbc", "cabacbc"},
		OuterReflection: "d",
		V:               [4]float64{0, 0, 0, 0},
		E:               [4]float64{0, 0, 0, 0},
		F:               [4]float64{0, 0, 0, 0},
		C:               [4]float64{0, 0, 0, 0},
		CellType:        "spherical",
		Vv:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices: [][4]float64{
			{1, P * 0.5, 0.5, P2 * 0.5}, {1, P * 0.5, -0.5, P2 * 0.5}, {1, P2 * 0.5, -P * 0.5, 0.5}, {1, P, 0, 0}, {1, P2 * 0.5, P * 0.5, 0.5},
			{1, 0, 0, P}, {1, 0.5, -P2 * 0.5, P * 0.5}, {1, P2 * 0.5, -P * 0.5, -0.5}, {1, P2 * 0.5, P * 0.5, -0.5}, {1, 0.5, P2 * 0.5, P * 0.5},
			{1, -P * 0.5, -0.5, P2 * 0.5}, {1, -0.5, -P2 * 0.5, P * 0.5}, {1, 0, -P, 0}, {1, 0.5, -P2 * 0.5, -P * 0.5}, {1, P * 0.5, -0.5, -P2 * 0.5},
			{1, P * 0.5, 0.5, -P2 * 0.5}, {1, 0.5, P2 * 0.5, -P * 0.5}, {1, 0, P, 0}, {1, -0.5, P2 * 0.5, P * 0.5}, {1, -P * 0.5, 0.5, P2 * 0.5},
			{1, -P2 * 0.5, -P * 0.5, 0.5}, {1, -0.5, -P2 * 0.5, -P * 0.5}, {1, 0, 0, -P}, {1, -0.5, P2 * 0.5, -P * 0.5}, {1, -P2 * 0.5, P * 0.5, 0.5},
			{1, -P2 * 0.5, -P * 0.5, -0.5}, {1, -P * 0.5, -0.5, -P2 * 0.5}, {1, -P * 0.5, 0.5, -P2 * 0.5}, {1, -P2 * 0.5, P * 0.5, -0.5}, {1, -P, 0, 0},
		},
		Edges: [][2]int{
			{0, 1}, {0, 4}, {0, 5}, {0, 9},
			{1, 5}, {1, 2}, {1, 6}, {2, 6},
			{2, 7}, {2, 3}, {3, 7}, {3, 8},
			{3, 4}, {4, 8}, {4, 9}, {5, 10},
			{5, 19}, {6, 11}, {6, 12}, {7, 13},
			{7, 14}, {8, 15}, {8, 16}, {9, 17},
			{9, 18}, {10, 11}, {10, 20}, {10, 19},
			{11, 20}, {11, 12}, {12, 13}, {12, 21},
			{13, 21}, {13, 14}, {14, 15}, {14, 22},
			{15, 16}, {15, 22}, {16, 17}, {16, 23},
			{17, 18}, {17, 23}, {18, 19}, {18, 24},
			{19, 24}, {20, 29}, {20, 25}, {21, 25},
			{21, 26}, {22, 26}, {22, 27}, {23, 27},
			{23, 28}, {24, 28}, {24, 29}, {25, 26},
			{26, 27}, {27, 28}, {28, 29}, {29, 25},
		},
		Faces: [][]int{
			{0, 1, 2, 3, 4}, {0, 5, 19, 18, 9}, {1, 6, 11, 10, 5}, {2, 7, 13, 12, 6},
			{3, 8, 15, 14, 7}, {4, 9, 17, 16, 8}, {11, 12, 21, 25, 20}, {13, 14, 22, 26, 21},
			{15, 16, 23, 27, 22}, {17, 18, 24, 28, 23}, {19, 10, 20, 29, 24}, {25, 26, 27, 28, 29},
			{0, 1, 5}, {1, 2, 6}, {2, 3, 7}, {3, 4, 8}, {4, 0, 9},
			{5, 10, 19}, {6, 11, 12}, {7, 13, 14}, {8, 15, 16}, {9, 17, 18},
			{10, 11, 20}, {12, 13, 21}, {14, 15, 22}, {16, 17, 23}, {18, 19, 24},
			{20, 25, 29}, {21, 25, 26}, {22, 26, 27}, {23, 27, 28}, {24, 28, 29},
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
