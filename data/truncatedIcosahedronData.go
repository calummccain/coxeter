package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
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

	var d func(vector.Vec4) vector.Vec4

	if n == 3 {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				0.5 * ((P4-1.0)*v.W - (P3-3.0*P_1)*v.Y - (P-3.0*P_3)*v.Z),
				v.X,
				0.5 * (P5*v.W + (2.0-P4)*v.Y - P2*v.Z),
				0.5 * (P3*v.W - P2*v.Y + v.Z),
			}

		}

	} else {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				(6.0*P2*cos-1.0)*v.W + (2.0*P_1-6.0*P*cos)*v.Y + (2.0*P_3-6.0*cos*P_1)*v.Z,
				v.X,
				2.0*P5*cos*v.W + (1.0-2.0*P4*cos)*v.Y - 2.0*P2*cos*v.Z,
				2.0*P3*cos*v.W - 2.0*P2*cos*v.Y + (1.0-2.0*cos)*v.Z,
			}

		}

	}

	var f func(vector.Vec4) vector.Vec4
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

	f = func(v vector.Vec4) vector.Vec4 {

		return vector.Vec4{a * v.W, b * v.X, b * v.Y, b * v.Z}

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
		V:               vector.Vec4{0, 0, 0, 0},
		E:               vector.Vec4{0, 0, 0, 0},
		F:               vector.Vec4{0, 0, 0, 0},
		C:               vector.Vec4{0, 0, 0, 0},
		CellType:        "spherical",
		VV:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices: []vector.Vec4{
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
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, -v.X, v.Y, v.Z} },
		Bmat: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{v.W, 0.5 * (v.X + P_1*v.Y - P*v.Z), 0.5 * (P_1*v.X + P*v.Y + v.Z), 0.5 * (-P*v.X + v.Y - P_1*v.Z)}
		},
		Cmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Y, -v.Z} },
		Dmat: d,
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}
}
