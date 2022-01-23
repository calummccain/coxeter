package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func TruncatedOctahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(Rt2)
	pVal := math.Pi / math.Atan(Rt_5)

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1.0 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	third := 1.0 / 3.0
	twoThird := 2.0 / 3.0

	var vv float64

	if metric == 'p' {

		vv = 1.0

	} else {

		vv = cot / math.Abs(1.0-cot)

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
				2.0*v.W - v.X - v.Y - v.Z,
				v.W - v.Y - v.Z,
				v.W - v.X - v.Z,
				v.W - v.X - v.Y,
			}

		}

	} else {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				(6.0*cos-2.0)*(v.W-v.X-v.Y-v.Z) + v.W,
				2.0*cos*(v.W-v.X-v.Y-v.Z) + v.X,
				2.0*cos*(v.W-v.X-v.Y-v.Z) + v.Y,
				2.0*cos*(v.W-v.X-v.Y-v.Z) + v.Z,
			}

		}

	}

	var f func(vector.Vec4) vector.Vec4
	var a, b float64

	if metric == 'p' {

		a = 3.0 * Rt_5
		b = 3.0 * Rt_5

	} else if metric == 'e' {

		a = 1.0
		b = 1.0

	} else {

		a = 3.0 * math.Sqrt(math.Abs(cot/(5.0-cot)))
		b = 3.0 * math.Sqrt(math.Abs((1.0-2.0*cot)/(5.0-cot)))

	}

	f = func(v vector.Vec4) vector.Vec4 {

		return vector.Vec4{a * v.W, b * v.X, b * v.Y, b * v.Z}

	}

	return CellData{
		Metric:          metric,
		NumVertices:     24,
		NumEdges:        36,
		NumFaces:        14,
		FaceReflections: []string{"", "c", "bc", "cbc", "abc", "cabc", "bcabc", "cbcabc"},
		OuterReflection: "d",
		V:               vector.Vec4{0, 0, 0, 0},
		E:               vector.Vec4{0, 0, 0, 0},
		F:               vector.Vec4{0, 0, 0, 0},
		C:               vector.Vec4{0, 0, 0, 0},
		CellType:        "spherical",
		VV:              vv,
		EVal:            eVal,
		PVal:            pVal,
		Vertices: []vector.Vec4{
			{1, twoThird, third, 0}, {1, twoThird, 0, third}, {1, twoThird, -third, 0}, {1, twoThird, 0, -third},
			{1, 0, twoThird, third}, {1, third, twoThird, 0}, {1, 0, twoThird, -third}, {1, -third, twoThird, 0},
			{1, third, 0, twoThird}, {1, 0, third, twoThird}, {1, -third, 0, twoThird}, {1, 0, -third, twoThird},
			{1, -twoThird, third, 0}, {1, -twoThird, 0, third}, {1, -twoThird, -third, 0}, {1, -twoThird, 0, -third},
			{1, 0, -twoThird, third}, {1, third, -twoThird, 0}, {1, 0, -twoThird, -third}, {1, -third, -twoThird, 0},
			{1, third, 0, -twoThird}, {1, 0, third, -twoThird}, {1, -third, 0, -twoThird}, {1, 0, -third, -twoThird},
		},
		Edges: [][2]int{
			{0, 1}, {0, 3}, {0, 5}, {1, 2}, {1, 8}, {2, 3}, {2, 17}, {3, 20},
			{4, 5}, {4, 7}, {4, 9}, {5, 6}, {6, 7}, {6, 21}, {7, 12}, {8, 9},
			{8, 11}, {9, 10}, {10, 11}, {10, 13}, {11, 16}, {12, 13}, {12, 15},
			{13, 14}, {14, 15}, {14, 19}, {15, 22}, {16, 17}, {16, 19}, {17, 18},
			{18, 19}, {18, 23}, {20, 21}, {20, 23}, {21, 22}, {22, 23},
		},
		Faces: [][]int{
			{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11},
			{12, 13, 14, 15}, {16, 17, 18, 19}, {20, 21, 22, 23},
			{0, 1, 8, 9, 4, 5}, {0, 3, 20, 21, 6, 5},
			{4, 7, 12, 13, 10, 9}, {1, 2, 17, 16, 11, 8},
			{2, 3, 20, 23, 18, 17}, {6, 7, 12, 15, 22, 21},
			{10, 11, 16, 19, 14, 13}, {14, 15, 22, 23, 18, 19},
		},
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.Y, v.X, v.Z} },
		Bmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Z, v.Y} },
		Cmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Y, -v.Z} },
		Dmat: d,
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}
}
