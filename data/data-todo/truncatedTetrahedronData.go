package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
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

	var f func(vector.Vec4) vector.Vec4
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

	f = func(v vector.Vec4) vector.Vec4 {

		return vector.Vec4{a * v.W, b * v.X, b * v.Y, b * v.Z}

	}

	return CellData{
		Metric:          metric,
		NumVertices:     12,
		NumEdges:        18,
		NumFaces:        8,
		FaceReflections: []string{"", "abc", "bc", "c"},
		OuterReflection: "d",
		V:               vector.Vec4{0, 0, 0, 0},
		E:               vector.Vec4{0, 0, 0, 0},
		F:               vector.Vec4{0, 0, 0, 0},
		C:               vector.Vec4{0, 0, 0, 0},
		CellType:        "spherical",
		VV:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices: []vector.Vec4{
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
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, -v.Z, -v.Y} },
		Bmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.Y, v.X, v.Z} },
		Cmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Z, v.Y} },
		Dmat: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				(3*sin-1)*(-v.W+v.X+v.Y-v.Z) + v.W,
				cos*(v.W-v.X-v.Y+v.Z) + v.X,
				cos*(v.W-v.X-v.Y+v.Z) + v.Y,
				cos*(-v.W+v.X+v.Y-v.Z) + v.Z,
			}
		},
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}
}
