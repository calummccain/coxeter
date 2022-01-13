package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func RectifiedDodecahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(P)
	pVal := math.Inf(1)

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1.0 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	vv := 0.5 * (cot + P)

	var d func(vector.Vec4) vector.Vec4

	if n == 3 {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				(P*v.W + v.X*P_3 + v.Z*P_4) * 0.5,
				(P3*v.W - v.X*P_1 - P*v.Z) * 0.5,
				v.Y,
				(P2*v.W - P*v.X + v.Z) * 0.5,
			}

		}

	} else if n == 4 {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				P2*v.W - v.X - v.Z*P_1,
				P3*v.W - P*v.X - P*v.Z,
				v.Y,
				P2*v.W - P*v.X,
			}

		}

	} else if n == 5 {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				((4*P+1)*v.W - (5-P)*v.X - (5*P-6)*v.Z) * 0.5,
				(P5*v.W + (2-P4)*v.X - P3*v.Z) * 0.5,
				v.Y,
				(P4*v.W - P3*v.X - v.Z*P_1) * 0.5,
			}

		}

	} else if n == 6 {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				((2+P4)*v.W - P3*v.X - P2*v.Z) * 0.5,
				(3*P3*v.W + (2-3*P2)*v.X - 3*P*v.Z) * 0.5,
				v.Y,
				(3*P2*v.W - 3*P*v.X - v.Z) * 0.5,
			}

		}

	} else {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				(2*P*Rt5*cos-1)*v.W - (2*Rt5*cos-2*P_1)*v.X - (2*Rt5*cos*P_1-2*P_2)*v.Z,
				2*P3*cos*v.W + (1-2*P2*cos)*v.X - 2*P*cos*v.Z,
				v.Y,
				2*P2*cos*v.W - 2*P*cos*v.X + (1-2*cos)*v.Z,
			}

		}

	}

	var f func(vector.Vec4) vector.Vec4
	var a, b float64

	if metric == 'e' {

		a = 1.0
		b = 1.0

	} else {

		a = P * math.Sqrt(math.Abs(cot))
		b = math.Sqrt(math.Abs(cot - P_2))

	}

	f = func(v vector.Vec4) vector.Vec4 {

		return vector.Vec4{a * v.W, b * v.X, b * v.Y, b * v.Z}

	}

	return CellData{
		Metric:      metric,
		NumVertices: 30,
		NumEdges:    60,
		NumFaces:    32,
		FaceReflections: []string{
			"", "c", "bc", "acacbabc", "cacbabc",
			"cbabc", "bacbabc", "cbabacbabc", "babacbabc",
			"abc", "acbabc", "abacbabc",
		},
		OuterReflection: "d",
		V:               vector.Vec4{1, P, P_1, 0},
		E:               vector.Vec4{1, P, 0, 0},
		F:               vector.Vec4{3 - P, P, 0, 1},
		C:               vector.Vec4{1, 0, 0, 0},
		CellType:        "spherical",
		VV:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices: []vector.Vec4{
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
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, -v.Y, v.Z} },
		Bmat: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{v.W, 0.5 * (P*v.X + v.Y + v.Z*P_1), 0.5 * (v.X - v.Y*P_1 - P*v.Z), 0.5 * (v.X*P_1 - P*v.Y + v.Z)}
		},
		Cmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Y, -v.Z} },
		Dmat: d,
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}
}
