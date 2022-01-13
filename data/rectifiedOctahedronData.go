package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func RectifiedOctahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(Rt2)
	pVal := math.Inf(1)

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	var vv float64

	if metric == 'p' {

		vv = 1

	} else {

		vv = cot / math.Abs(1-cot)

	}

	var d func(vector.Vec4) vector.Vec4

	if n == 3 {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				0.5 * (v.W + v.X + v.Y + v.Z),
				0.5 * (v.W + v.X - v.Y - v.Z),
				0.5 * (v.W - v.X + v.Y - v.Z),
				0.5 * (v.W - v.X - v.Y + v.Z),
			}

		}

	} else if n == 4 {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				2*v.W - v.X - v.Y - v.Z,
				v.W - v.Y - v.Z,
				v.W - v.X - v.Z,
				v.W - v.X - v.Y,
			}

		}

	} else {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				(6*cos-2)*(v.W-v.X-v.Y-v.Z) + v.W,
				2*cos*(v.W-v.X-v.Y-v.Z) + v.X,
				2*cos*(v.W-v.X-v.Y-v.Z) + v.Y,
				2*cos*(v.W-v.X-v.Y-v.Z) + v.Z,
			}

		}

	}

	var f func(vector.Vec4) vector.Vec4
	var a, b float64

	if n == 3 {

		a = Rt2 * Rt_3
		b = Rt_3 * Rt_2

	} else if n == 4 {

		a = Rt2
		b = Rt_2

	} else if metric == 'e' {

		a = 1
		b = 1

	} else {

		a = math.Sqrt(math.Abs(2.0 * cot))
		b = math.Sqrt(math.Abs((0.5 - cot)))

	}

	f = func(v vector.Vec4) vector.Vec4 {

		return vector.Vec4{a * v.W, b * v.X, b * v.Y, b * v.Z}

	}

	return CellData{
		Metric:          metric,
		NumVertices:     12,
		NumEdges:        24,
		NumFaces:        14,
		FaceReflections: []string{"", "c", "bc", "cbc", "abc", "cabc", "bcabc", "cbcabc"},
		OuterReflection: "d",
		V:               vector.Vec4{0, 0, 0, 0},
		E:               vector.Vec4{0, 0, 0, 0},
		F:               vector.Vec4{0, 0, 0, 0},
		C:               vector.Vec4{0, 0, 0, 0},
		CellType:        "spherical",
		VV:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices: []vector.Vec4{
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
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.Y, v.X, v.Z} },
		Bmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Z, v.Y} },
		Cmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Y, -v.Z} },
		Dmat: d,
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}
}
