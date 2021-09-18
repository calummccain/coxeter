package data

import (
	"coxeter/shared"
	"math"
)

func DodecahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(P)
	pVal := 6.0

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	var vv float64

	if metric == 'p' {

		vv = 2.0 * P_2

	} else {

		vv = (cot + P + P_1) / math.Abs(cot-3.0)

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

	if n == 3 {

		a = P2 * 0.5 * Rt_2
		b = P_1 * 0.5 * Rt_2

	} else if n == 4 {

		a = P2 * Rt_2
		b = math.Sqrt(P * 0.5)

	} else if n == 5 {

		a = P4 * 0.5
		b = P * 0.5 * math.Sqrt(4*P-1)

	} else if n == 6 {

		a = Rt3
		b = 1

	} else if metric == 'e' {

		a = 1
		b = 1

	} else {

		a = P2 * math.Sqrt(math.Abs(cot/(cot-3.0)))
		b = math.Sqrt(math.Abs((P2*cot - 1.0) / (cot - 3.0)))

	}

	f = func(v [4]float64) [4]float64 {

		return [4]float64{a * v[0], b * v[1], b * v[2], b * v[3]}

	}

	return CellData{
		Metric:      metric,
		NumVertices: 20,
		NumEdges:    30,
		NumFaces:    12,
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
			{1, 1, 1, 1}, {1, 1, 1, -1}, {1, 1, -1, 1}, {1, 1, -1, -1},
			{1, -1, 1, 1}, {1, -1, 1, -1}, {1, -1, -1, 1}, {1, -1, -1, -1},
			{1, 0, P, P_1}, {1, 0, P, -P_1}, {1, 0, -P, P_1}, {1, 0, -P, -P_1},
			{1, P, P_1, 0}, {1, P, -P_1, 0}, {1, -P, P_1, 0}, {1, -P, -P_1, 0},
			{1, P_1, 0, P}, {1, -P_1, 0, P}, {1, P_1, 0, -P}, {1, -P_1, 0, -P},
		},
		Edges: [][2]int{
			{0, 8}, {0, 12}, {0, 16}, {1, 9},
			{1, 12}, {1, 18}, {2, 10}, {2, 13},
			{2, 16}, {3, 11}, {3, 13}, {3, 18},
			{4, 8}, {4, 14}, {4, 17}, {5, 9},
			{5, 14}, {5, 19}, {6, 10}, {6, 15},
			{6, 17}, {7, 11}, {7, 15}, {7, 19},
			{8, 9}, {10, 11}, {12, 13}, {14, 15},
			{16, 17}, {18, 19},
		},
		Faces: [][]int{
			{0, 16, 2, 13, 12}, {1, 12, 13, 3, 18},
			{0, 12, 1, 9, 8}, {0, 8, 4, 17, 16},
			{2, 16, 17, 6, 10}, {1, 18, 19, 5, 9},
			{4, 8, 9, 5, 14}, {5, 19, 7, 15, 14},
			{6, 17, 4, 14, 15}, {3, 13, 2, 10, 11},
			{3, 11, 7, 19, 18}, {11, 10, 6, 15, 7},
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
