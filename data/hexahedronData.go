package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func HexahedronData(n float64) CellData {

	eVal := 4.0
	pVal := 6.0

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	cf := 2.0 * cot / (1 + cot)
	ce := cot
	fe := 0.5 * (1 + cot)

	var cv, fv, ev, vv float64

	if metric == 'p' {
		cv = 0.0
		fv = 0.0
		ev = 0.0
		vv = 2.0
	} else {
		cv = 2.0 * cot / (3.0 - cot)
		fv = (1.0 + cot) / (3.0 - cot)
		ev = 2.0 / (3.0 - cot)
		vv = (1 + cot) / math.Abs(cot-3.0)
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

			cn := math.Cos(2.0 * math.Pi / n)

			return vector.Vec4{
				(1.0+2.0*cn)*v.W - 2.0*cn*v.X,
				(2.0+2.0*cn)*v.W - (1.0+2.0*cn)*v.X,
				v.Y,
				v.Z,
			}

		}

	}

	var f func(vector.Vec4) vector.Vec4
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

	f = func(v vector.Vec4) vector.Vec4 {

		return vector.Vec4{a * v.W, b * v.X, b * v.Y, b * v.Z}

	}

	return CellData{
		Metric:          metric,
		NumVertices:     8,
		NumEdges:        12,
		NumFaces:        6,
		FaceReflections: []string{"bc", "c", "cbabc", "abc", "", "babc"},
		OuterReflection: "d",
		V:               vector.Vec4{1, 1, 1, 1},
		E:               vector.Vec4{1, 1, 1, 0},
		F:               vector.Vec4{1, 1, 0, 0},
		C:               vector.Vec4{1, 0, 0, 0},
		CellType:        "spherical",
		CF:              cf,
		CE:              ce,
		CV:              cv,
		FE:              fe,
		FV:              fv,
		EV:              ev,
		VV:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices: []vector.Vec4{
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
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Y, -v.Z} },
		Bmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Z, v.Y} },
		Cmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.Y, v.X, v.Z} },
		Dmat: d,
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}
}
