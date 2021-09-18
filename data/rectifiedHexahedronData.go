package data

import (
	"coxeter/shared"
	"math"
)

func RectifiedHexahedronData(n float64) CellData {

	eVal := 4.0
	pVal := math.Inf(1)

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	vv := 0.5*cot + 0.5

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

			c := math.Cos(math.Pi / n)

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

		a = 1
		b = 1

	} else {

		a = math.Sqrt(math.Abs(cot))
		b = math.Sqrt(math.Abs(0.5*cot - 0.5))

	}

	f = func(v [4]float64) [4]float64 {

		return [4]float64{a * v[0], b * v[1], b * v[2], b * v[3]}

	}

	return CellData{
		Metric:          metric,
		NumVertices:     12,
		NumEdges:        24,
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
			{1, 1, 1, 0}, {1, 1, 0, 1}, {1, 0, 1, 1},
			{1, 1, -1, 0}, {1, -1, 0, 1}, {1, 0, 1, -1},
			{1, -1, 1, 0}, {1, 1, 0, -1}, {1, 0, -1, 1},
			{1, -1, -1, 0}, {1, -1, 0, -1}, {1, 0, -1, -1},
		},
		Edges: [][2]int{
			{0, 1}, {1, 2}, {2, 0}, {0, 5}, {5, 7}, {7, 0},
			{3, 7}, {7, 11}, {11, 3}, {1, 3}, {3, 8}, {8, 1},
			{4, 8}, {4, 9}, {8, 9}, {2, 4}, {2, 6}, {4, 6},
			{5, 6}, {5, 10}, {6, 10}, {9, 10}, {9, 11}, {10, 11},
		},
		Faces: [][]int{
			{0, 1, 2}, {0, 5, 7}, {3, 7, 11}, {1, 3, 8},
			{4, 8, 9}, {2, 4, 6}, {5, 6, 10}, {9, 10, 11},
			{0, 5, 6, 2}, {1, 2, 4, 8}, {0, 1, 3, 7},
			{3, 8, 9, 11}, {5, 7, 11, 10}, {4, 6, 10, 9},
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
