package data

import (
	"math"

	"github.com/calummccain/coxeter/shared"
)

func TruncatedIcosahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(P2)
	pVal := math.Pi / math.Atan(P2/math.Sqrt(9.0*P2+1.0))

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1.0 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	A := P / 3.0
	B := (P + 2.0) / 3.0
	C := 1.0 / 3.0
	D := P3 / 3.0
	E := 2.0 * P / 3.0
	F := 2.0 / 3.0

	var vv float64

	if metric == 'p' {

		vv = 2.0 / (9.0*P2 + 1.0)

	} else {

		vv = (P4*cot + 9.0*P2 - 1.0) / math.Abs(9.0*P2+1.0-P4*cot)

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
	var a, b float64

	if metric == 'e' {

		a = 1
		b = 1

	} else if metric == 'p' {

		a = 1
		b = 3.0 / math.Sqrt(9.0*P2+1.0)

	} else {

		a = P3 * 3.0 * math.Sqrt(math.Abs(cot/(9*P2+1.0-P4*cot)))
		b = 3.0 * math.Sqrt(math.Abs((P4*cot-1.0)/(9.0*P2+1.0-P4*cot)))

	}

	f = func(v [4]float64) [4]float64 {

		return [4]float64{a * v[0], b * v[1], b * v[2], b * v[3]}

	}

	return CellData{
		Metric:      metric,
		NumVertices: 60,
		NumEdges:    90,
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
			{1, F, D, A}, {1, B, E, C}, {1, B, E, -C}, {1, F, D, -A}, {1, C, P, 0},
			{1, C, B, E}, {1, A, F, D}, {1, 0, C, P}, {1, -A, F, D}, {1, -C, B, E},
			{1, -F, D, A}, {1, -C, P, 0}, {1, -F, D, -A}, {1, -B, E, -C}, {1, -B, E, C},
			{1, -D, A, F}, {1, -E, C, B}, {1, -E, -C, B}, {1, -D, -A, F}, {1, -P, 0, C},
			{1, -C, B, -E}, {1, C, B, -E}, {1, A, F, -D}, {1, 0, C, -P}, {1, -A, F, -D},
			{1, D, A, -F}, {1, E, C, -B}, {1, E, -C, -B}, {1, D, -A, -F}, {1, P, 0, -C},
			{1, D, A, F}, {1, P, 0, C}, {1, D, -A, F}, {1, E, -C, B}, {1, E, C, B},
			{1, A, -F, D}, {1, 0, -C, P}, {1, -A, -F, D}, {1, -C, -B, E}, {1, C, -B, E},
			{1, -D, -A, -F}, {1, -P, 0, -C}, {1, -D, A, -F}, {1, -E, C, -B}, {1, -E, -C, -B},
			{1, -F, -D, -A}, {1, -B, -E, -C}, {1, -B, -E, C}, {1, -F, -D, A}, {1, -C, -P, 0},
			{1, C, -P, 0}, {1, F, -D, A}, {1, B, -E, C}, {1, B, -E, -C}, {1, F, -D, -A},
			{1, -C, -B, -E}, {1, -A, -F, -D}, {1, 0, -C, -P}, {1, A, -F, -D}, {1, C, -B, -E},
		},
		Edges: [][2]int{
			{0, 1}, {0, 4}, {0, 5}, {1, 2}, {1, 30},
			{2, 3}, {2, 25}, {3, 4}, {3, 21}, {4, 11},
			{5, 6}, {5, 9}, {6, 7}, {6, 34}, {7, 8},
			{7, 36}, {8, 9}, {8, 16}, {9, 10}, {10, 11},
			{10, 14}, {11, 12}, {12, 13}, {12, 20}, {13, 14},
			{13, 42}, {14, 15}, {15, 16}, {15, 19}, {16, 17},
			{17, 18}, {17, 37}, {18, 19}, {18, 47}, {19, 41},
			{20, 21}, {20, 24}, {21, 22}, {22, 23}, {22, 26},
			{23, 24}, {23, 57}, {24, 43}, {25, 26}, {25, 29},
			{26, 27}, {27, 28}, {27, 58}, {28, 29}, {28, 53},
			{29, 31}, {30, 31}, {30, 34}, {31, 32}, {32, 33},
			{32, 52}, {33, 34}, {33, 35}, {35, 36}, {35, 39},
			{36, 37}, {37, 38}, {38, 39}, {38, 48}, {39, 51},
			{40, 41}, {40, 44}, {40, 46}, {41, 42}, {42, 43},
			{43, 44}, {44, 56}, {45, 46}, {45, 49}, {45, 55},
			{46, 47}, {47, 48}, {48, 49}, {49, 50}, {50, 51},
			{50, 54}, {51, 52}, {52, 53}, {53, 54}, {54, 59},
			{55, 56}, {55, 59}, {56, 57}, {57, 58}, {58, 59},
		},
		Faces: [][]int{
			{0, 1, 2, 3, 4}, {5, 6, 7, 8, 9}, {10, 11, 12, 13, 14}, {15, 16, 17, 18, 19},
			{20, 21, 22, 23, 24}, {25, 26, 27, 28, 29}, {30, 31, 32, 33, 34}, {35, 36, 37, 38, 39},
			{40, 41, 42, 43, 44}, {45, 46, 47, 48, 49}, {50, 51, 52, 53, 54}, {55, 56, 57, 58, 59},
			{0, 4, 11, 10, 9, 5}, {8, 9, 10, 14, 15, 16}, {7, 8, 16, 17, 37, 36},
			{6, 7, 36, 35, 33, 34}, {0, 1, 30, 34, 6, 5}, {3, 4, 11, 12, 20, 21},
			{12, 13, 42, 43, 24, 20}, {13, 14, 15, 19, 41, 42}, {1, 2, 25, 29, 31, 30},
			{18, 19, 41, 40, 46, 47}, {17, 18, 47, 48, 38, 37}, {38, 39, 51, 50, 49, 48},
			{32, 33, 35, 39, 51, 52}, {28, 29, 31, 32, 52, 53}, {27, 28, 53, 54, 59, 58},
			{2, 3, 21, 22, 26, 25}, {22, 23, 57, 58, 27, 26}, {23, 24, 43, 44, 56, 57},
			{40, 44, 56, 55, 45, 46}, {45, 49, 50, 54, 59, 55},
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
