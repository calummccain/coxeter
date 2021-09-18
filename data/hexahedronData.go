package data

import (
	"coxeter/shared"
	"math"
)

func HexahedronData(n float64) CellData {

	eVal := 4.0
	pVal := 6.0

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	var vv float64

	if metric == 'p' {

		vv = 2

	} else {

		vv = (1.0 + cot) / math.Abs(3.0-cot)

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

			cn := math.Cos(2.0 * math.Pi / n)

			return [4]float64{
				(1.0+2.0*cn)*v[0] - 2.0*cn*v[1],
				(2.0+2.0*cn)*v[0] - (1.0+2.0*cn)*v[1],
				v[2],
				v[3],
			}

		}

	}

	var f func([4]float64) [4]float64
	var a, b float64

	if n == 3 {

		a = 0.5
		b = 0.5

	} else if n == 4 {

		a = 1
		b = 1

	} else if n == 5 {

		a = P2 * Rt_2
		b = math.Sqrt(0.5 * P)

	} else if n == 6 {

		a = Rt3
		b = 1

	} else if metric == 'p' {

		a = 1
		b = 1

	} else {

		a = math.Sqrt(math.Abs(2.0 * cot / (3.0 - cot)))
		b = math.Sqrt(math.Abs((cot - 1.0) / (3.0 - cot)))

	}

	f = func(v [4]float64) [4]float64 {

		return [4]float64{a * v[0], b * v[1], b * v[2], b * v[3]}

	}

	return CellData{
		Metric:          metric,
		NumVertices:     8,
		NumEdges:        12,
		NumFaces:        6,
		FaceReflections: []string{"bc", "c", "cbabc", "abc", "", "babc"},
		OuterReflection: "d",
		V:               [4]float64{1, 1, 1, 1},
		E:               [4]float64{1, 1, 1, 0},
		F:               [4]float64{1, 1, 0, 0},
		C:               [4]float64{math.Sqrt(math.Abs(0.5 * (3.0 - cot) / cot)), 0, 0, 0},
		CellType:        "spherical",
		Vv:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices: [][4]float64{
			{1, 1, 1, 1}, {1, 1, -1, 1},
			{1, -1, -1, 1}, {1, -1, 1, 1},
			{1, 1, 1, -1}, {1, 1, -1, -1},
			{1, -1, -1, -1}, {1, -1, 1, -1},
		},
		Edges: [][2]int{
			{0, 3}, {3, 2}, {2, 1}, {1, 0},
			{7, 4}, {4, 5}, {5, 6}, {6, 7},
			{0, 4}, {1, 5}, {2, 6}, {3, 7},
		},
		Faces: [][]int{
			{0, 1, 2, 3}, {4, 7, 3, 0}, {7, 6, 2, 3},
			{4, 5, 6, 7}, {0, 1, 5, 4}, {1, 2, 6, 5},
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
