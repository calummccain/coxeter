package data

import (
	"math"

	"github.com/calummccain/coxeter/shared"
)

func TruncatedDodecahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(P)
	pVal := math.Pi / math.Atan(P/math.Sqrt(5.0*P4+1.0))

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1.0 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	A := (P4 - 1.0) * 0.2
	B := 1.0 + Rt_5
	C := A * P_1
	D := B * P_1
	E := (1 - Rt_5) * 0.5
	F := (1 + Rt_5) * 0.5
	G := P_1 * Rt_5

	var vv float64

	if metric == 'p' {

		vv = 2.0 / (5.0*P4 + 1.0)

	} else {

		vv = (P_2*cot*0.2 + 0.6*P) / math.Abs(1.0+P_4*0.2-P_2*cot*0.2)

	}

	var d func([4]float64) [4]float64

	if n == 3 {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				(P*v[0] + v[1]*P_3 + v[3]*P_4) * 0.5,
				(P3*v[0] - v[1]*P_1 - P*v[3]) * 0.5,
				v[2],
				(P2*v[0] - P*v[1] + v[3]) * 0.5,
			}

		}

	} else if n == 4 {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				P2*v[0] - v[1] - v[3]*P_1,
				P3*v[0] - P*v[1] - P*v[3],
				v[2],
				P2*v[0] - P*v[1],
			}

		}

	} else if n == 5 {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				((4*P+1)*v[0] - (5-P)*v[1] - (5*P-6)*v[3]) * 0.5,
				(P5*v[0] + (2-P4)*v[1] - P3*v[3]) * 0.5,
				v[2],
				(P4*v[0] - P3*v[1] - v[3]*P_1) * 0.5,
			}

		}

	} else if n == 6 {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				((2+P4)*v[0] - P3*v[1] - P2*v[3]) * 0.5,
				(3*P3*v[0] + (2-3*P2)*v[1] - 3*P*v[3]) * 0.5,
				v[2],
				(3*P2*v[0] - 3*P*v[1] - v[3]) * 0.5,
			}

		}

	} else {

		d = func(v [4]float64) [4]float64 {

			return [4]float64{
				(2*P*Rt5*cos-1)*v[0] - (2*Rt5*cos-2*P_1)*v[1] - (2*Rt5*cos*P_1-2*P_2)*v[3],
				2*P3*cos*v[0] + (1-2*P2*cos)*v[1] - 2*P*cos*v[3],
				v[2],
				2*P2*cos*v[0] - 2*P*cos*v[1] + (1-2*cos)*v[3],
			}

		}

	}

	var f func([4]float64) [4]float64
	var a, b float64

	if metric == 'p' {

		a = 1.0
		b = 1.0 / math.Sqrt(P2+0.2*P_2)

	} else if metric == 'e' {

		a = 1.0
		b = 1.0

	} else {

		a = P * math.Sqrt(math.Abs(cot/(1.0+P_4*0.2-P_2*cot*0.2)))
		b = math.Sqrt(math.Abs((cot - P_2) / (1.0 + P_4*0.2 - P_2*cot*0.2)))

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
			"", "c", "bc", "acacbabc", "cacbabc",
			"cbabc", "bacbabc", "cbabacbabc", "babacbabc",
			"abc", "acbabc", "abacbabc",
		},
		OuterReflection: "d",
		V:               [4]float64{1, P, P_1, 0},
		E:               [4]float64{1, P, 0, 0},
		F:               [4]float64{3 - P, P, 0, 1},
		C:               [4]float64{1, 0, 0, 0},
		CellType:        "spherical",
		Vv:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices: [][4]float64{
			{1, B, C, -E}, {1, B, C, E}, {1, P, G, 0},
			{1, P, -G, 0}, {1, B, -C, -E}, {1, B, -C, E},
			{1, F, -A, D}, {1, A, -D, F}, {1, D, -F, A},
			{1, C, -E, B}, {1, C, E, B}, {1, G, 0, P},
			{1, D, F, A}, {1, A, D, F}, {1, F, A, D},
			{1, E, B, C}, {1, 0, P, G}, {1, -E, B, C},
			{1, -F, A, D}, {1, -A, D, F}, {1, -D, F, A},
			{1, -C, E, B}, {1, -G, 0, P}, {1, -C, -E, B},
			{1, -D, -F, A}, {1, -A, -D, F}, {1, -F, -A, D},
			{1, -E, -B, C}, {1, E, -B, C}, {1, 0, -P, G},
			{1, 0, -P, -G}, {1, E, -B, -C}, {1, -E, -B, -C},
			{1, F, -A, -D}, {1, A, -D, -F}, {1, D, -F, -A},
			{1, C, -E, -B}, {1, G, 0, -P}, {1, C, E, -B},
			{1, D, F, -A}, {1, F, A, -D}, {1, A, D, -F},
			{1, E, B, -C}, {1, -E, B, -C}, {1, 0, P, -G},
			{1, -B, C, -E}, {1, -B, C, E}, {1, -P, G, 0},
			{1, -P, -G, 0}, {1, -B, -C, E}, {1, -B, -C, -E},
			{1, -A, -D, -F}, {1, -F, -A, -D}, {1, -D, -F, -A},
			{1, -C, -E, -B}, {1, -G, 0, -P}, {1, -C, E, -B},
			{1, -D, F, -A}, {1, -F, A, -D}, {1, -A, D, -F},
		},
		Edges: [][2]int{
			{0, 1}, {0, 2}, {0, 41}, {1, 2}, {1, 13},
			{2, 3}, {3, 4}, {3, 5}, {4, 5}, {4, 34},
			{5, 7}, {6, 7}, {6, 8}, {6, 28}, {7, 8},
			{8, 9}, {9, 10}, {9, 11}, {10, 11}, {10, 12},
			{11, 22}, {12, 13}, {12, 14}, {13, 14}, {14, 15},
			{15, 16}, {15, 17}, {16, 17}, {16, 44}, {17, 18},
			{18, 19}, {18, 20}, {19, 20}, {19, 46}, {20, 21},
			{21, 22}, {21, 23}, {22, 23}, {23, 24}, {24, 25},
			{24, 26}, {25, 26}, {25, 49}, {26, 27}, {27, 28},
			{27, 29}, {28, 29}, {29, 30}, {30, 31}, {30, 32},
			{31, 32}, {31, 33}, {32, 52}, {33, 34}, {33, 35},
			{34, 35}, {35, 36}, {36, 37}, {36, 38}, {37, 38},
			{37, 55}, {38, 39}, {39, 40}, {39, 41}, {40, 41},
			{40, 42}, {42, 43}, {42, 44}, {43, 44}, {43, 58},
			{45, 46}, {45, 47}, {45, 59}, {46, 47}, {47, 48},
			{48, 49}, {48, 50}, {49, 50}, {50, 51}, {51, 52},
			{51, 53}, {52, 53}, {53, 54}, {54, 55}, {54, 56},
			{55, 56}, {56, 57}, {57, 58}, {57, 59}, {58, 59},
		},
		Faces: [][]int{
			{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, 10, 11}, {12, 13, 14},
			{15, 16, 17}, {18, 19, 20}, {21, 22, 23}, {24, 25, 26}, {27, 28, 29},
			{30, 31, 32}, {33, 34, 35}, {36, 37, 38}, {39, 40, 41}, {42, 43, 44},
			{45, 46, 47}, {48, 49, 50}, {51, 52, 53}, {54, 55, 56}, {57, 58, 59},
			{0, 1, 13, 14, 15, 16, 44, 42, 40, 41}, {1, 2, 3, 5, 7, 8, 9, 10, 12, 13},
			{0, 2, 3, 4, 34, 35, 36, 38, 39, 41}, {10, 11, 22, 21, 20, 18, 17, 15, 14, 12},
			{4, 5, 7, 6, 28, 29, 30, 31, 33, 34}, {6, 8, 9, 11, 22, 23, 24, 26, 27, 28},
			{25, 26, 27, 29, 30, 32, 52, 51, 50, 49}, {31, 32, 52, 53, 54, 55, 37, 36, 35, 33},
			{37, 38, 39, 40, 42, 43, 58, 57, 56, 55}, {16, 17, 18, 19, 46, 45, 59, 58, 43, 44},
			{19, 20, 21, 23, 24, 25, 49, 48, 47, 46}, {45, 47, 48, 50, 51, 53, 54, 56, 57, 59},
		},
		Matrices: shared.Matrices{
			A: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[1], -v[2], v[3]} },
			B: func(v [4]float64) [4]float64 {
				return [4]float64{v[0], 0.5 * (P*v[1] + v[2] + v[3]*P_1), 0.5 * (v[1] - v[2]*P_1 - P*v[3]), 0.5 * (v[1]*P_1 - P*v[2] + v[3])}
			},
			C: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[1], v[2], -v[3]} },
			D: d,
			E: func(v [4]float64) [4]float64 { return v },
			F: f,
		},
		Flip: func(v [4]float64) [4]float64 { return [4]float64{-v[0], v[1], v[2], v[3]} },
	}
}
