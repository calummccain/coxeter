package data

import (
	"math"

	"github.com/calummccain/coxeter/shared"
)

func TruncatedTetrahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(Rt_2)
	pVal := math.Pi / math.Atan(Rt_11)

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	third := 1.0 / 3.0

	var vv float64

	if metric == 'p' {

		vv = 4.0 / 11.0

	} else {

		vv = (cot + 7.0) / math.Abs(11.0-cot)

	}

	var f func([4]float64) [4]float64
	var a, b float64

	if metric == 'e' {

		a = 1
		b = 1

	} else if metric == 'p' {

		a = 1
		b = 3 * Rt_11

	} else {

		a = 3 * math.Sqrt(math.Abs(cot/(22.0-2.0*cot)))
		b = 3 * math.Sqrt(math.Abs((cot-2.0)/(22.0-2.0*cot)))

	}

	f = func(v [4]float64) [4]float64 {

		return [4]float64{a * v[0], b * v[1], b * v[2], b * v[3]}

	}

	return CellData{
		Metric:          metric,
		NumVertices:     12,
		NumEdges:        18,
		NumFaces:        8,
		FaceReflections: []string{"", "abc", "bc", "c"},
		OuterReflection: "d",
		V:               [4]float64{0, 0, 0, 0},
		E:               [4]float64{0, 0, 0, 0},
		F:               [4]float64{0, 0, 0, 0},
		C:               [4]float64{0, 0, 0, 0},
		CellType:        "spherical",
		Vv:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices: [][4]float64{
			{1, 1, third, third}, {1, third, 1, third}, {1, third, third, 1},
			{1, 1, -third, -third}, {1, third, -1, -third}, {1, third, -third, -1},
			{1, -1, third, -third}, {1, -third, 1, -third}, {1, -third, third, -1},
			{1, -1, -third, third}, {1, -third, -1, third}, {1, -third, -third, 1},
		},
		Edges: [][2]int{
			{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 7}, {2, 11},
			{3, 4}, {3, 5}, {4, 5}, {4, 10}, {5, 8}, {6, 7},
			{6, 8}, {6, 9}, {7, 8}, {9, 10}, {9, 11}, {10, 11},
		},
		Faces: [][]int{
			{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, 10, 11},
			{0, 2, 11, 10, 4, 3}, {0, 1, 7, 8, 5, 3},
			{1, 2, 11, 9, 6, 7}, {4, 5, 8, 6, 9, 10},
		},
		Matrices: shared.Matrices{
			A: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[1], -v[3], -v[2]} },
			B: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[2], v[1], v[3]} },
			C: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[1], v[3], v[2]} },
			D: func(v [4]float64) [4]float64 {
				return [4]float64{
					(3*sin-1)*(-v[0]+v[1]+v[2]-v[3]) + v[0],
					cos*(v[0]-v[1]-v[2]+v[3]) + v[1],
					cos*(v[0]-v[1]-v[2]+v[3]) + v[2],
					cos*(-v[0]+v[1]+v[2]-v[3]) + v[3],
				}
			},
			E: func(v [4]float64) [4]float64 { return v },
			F: f,
		},
		Flip: func(v [4]float64) [4]float64 { return [4]float64{-v[0], v[1], v[2], v[3]} },
	}
}
