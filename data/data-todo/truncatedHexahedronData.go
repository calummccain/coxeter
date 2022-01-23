package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func TruncatedHexahedronData(n float64) CellData {

	eVal := 4.0
	pVal := math.Pi / math.Atan(math.Sqrt(1.0/(7.0+4.0*Rt2)))

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1.0 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	factor := Rt2 - 1.0

	var vv float64

	if metric == 'p' {

		vv = 2.0 / (7.0 + 4.0*Rt2)

	} else {

		vv = ((3.0-2.0*Rt2)*cot + 2.0*Rt2 - 1.0) / math.Abs(-(3.0-2.0*Rt2)*cot+5.0-2.0*Rt2)

	}

	var d func(vector.Vec4) vector.Vec4

	if n == 3 {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				v.X,
				v.W,
				v.Y,
				v.Z,
			}

		}

	} else if n == 4 {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				v.W,
				2.0*v.W - v.X,
				v.Y,
				v.Z,
			}

		}

	} else if n == 5 {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				P*v.W - P_1*v.X,
				P2*v.W - P*v.X,
				v.Y,
				v.Z,
			}

		}

	} else if n == 6 {

		d = func(v vector.Vec4) vector.Vec4 {

			return vector.Vec4{
				2.0*v.W - v.X,
				3.0*v.W - 2.0*v.X,
				v.Y,
				v.Z,
			}

		}

	} else {

		d = func(v vector.Vec4) vector.Vec4 {

			c := math.Cos(2.0 * math.Pi / n)

			return vector.Vec4{
				(1.0+2.0*c)*v.W - 2.0*c*v.X,
				2.0 + 2.0*c*v.W - (1.0+2.0*c)*v.X,
				v.Y,
				v.Z,
			}

		}

	}

	var f func(vector.Vec4) vector.Vec4
	var a, b float64

	if n == 4 {

		a = 1.0
		b = 1.0

	} else if metric == 'p' {

		a = 1.0
		b = 1.0 / math.Sqrt(5.0-2.0*Rt2)

	} else {

		a = math.Sqrt(math.Abs(2.0 * cot / (5.0 - 2.0*Rt2 - (3.0-2.0*Rt2)*cot)))
		b = math.Sqrt(math.Abs((cot - 1.0) / (5.0 - 2.0*Rt2 - (3.0-2.0*Rt2)*cot)))

	}

	f = func(v vector.Vec4) vector.Vec4 {

		return vector.Vec4{a * v.W, b * v.X, b * v.Y, b * v.Z}

	}

	return CellData{
		Metric:          metric,
		NumVertices:     24,
		NumEdges:        36,
		NumFaces:        14,
		FaceReflections: []string{"bc", "c", "cbabc", "abc", "", "babc"},
		OuterReflection: "d",
		V:               vector.Vec4{1, 1, 1, 1},
		E:               vector.Vec4{1, 1, 1, 0},
		F:               vector.Vec4{1, 1, 0, 0},
		C:               vector.Vec4{1, 0, 0, 0},
		CellType:        "spherical",
		VV:              vv,
		EVal:            eVal,
		PVal:            pVal,
		Vertices: []vector.Vec4{
			{1, 1, 1, factor}, {1, 1, factor, 1}, {1, factor, 1, 1},
			{1, 1, 1, -factor}, {1, 1, factor, -1}, {1, factor, 1, -1},
			{1, 1, -1, factor}, {1, 1, -factor, 1}, {1, factor, -1, 1},
			{1, 1, -1, -factor}, {1, 1, -factor, -1}, {1, factor, -1, -1},
			{1, -1, 1, factor}, {1, -1, factor, 1}, {1, -factor, 1, 1},
			{1, -1, 1, -factor}, {1, -1, factor, -1}, {1, -factor, 1, -1},
			{1, -1, -1, factor}, {1, -1, -factor, 1}, {1, -factor, -1, 1},
			{1, -1, -1, -factor}, {1, -1, -factor, -1}, {1, -factor, -1, -1},
		},
		Edges: [][2]int{
			{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 7}, {2, 14}, {3, 4}, {3, 5},
			{4, 5}, {4, 10}, {5, 17}, {6, 7}, {6, 8}, {6, 9}, {7, 8}, {8, 20},
			{9, 10}, {9, 11}, {10, 11}, {11, 23}, {12, 13}, {12, 14}, {12, 15}, {13, 14},
			{13, 19}, {15, 16}, {15, 17}, {16, 17}, {16, 22}, {18, 19}, {18, 20}, {18, 21},
			{19, 20}, {21, 22}, {21, 23}, {22, 23},
		},
		Faces: [][]int{
			{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, 10, 11},
			{12, 13, 14}, {15, 16, 17}, {18, 19, 20}, {21, 22, 23},
			{0, 1, 7, 6, 9, 10, 4, 3}, {0, 2, 14, 12, 15, 17, 5, 3},
			{1, 2, 14, 13, 19, 20, 8, 7}, {6, 8, 20, 18, 21, 23, 11, 9},
			{4, 5, 17, 16, 22, 23, 11, 10}, {12, 13, 19, 18, 21, 22, 16, 15},
		},
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Y, -v.Z} },
		Bmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Z, v.Y} },
		Cmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.Y, v.X, v.Z} },
		Dmat: d,
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}
}
