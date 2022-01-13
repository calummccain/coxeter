package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func TetrahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(Rt_2)
	pVal := 6.0

	cos := math.Pow(math.Cos(math.Pi/n), 2.0)
	sin := 1.0 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	cf := 1.5 * cot / (1 + cot)
	ce := 0.5 * cot
	fe := (1 + cot) / 3.0

	var cv, fv, ev, vv float64
	if metric == 'p' {
		cv = 3.0
		fv = 8.0
		ev = 2.0
		vv = 4.0
	} else {
		cv = 0.5 * cot / (3.0 - cot)
		fv = (1.0 + cot) / (3.0 * (3.0 - cot))
		ev = 1.0 / (3.0 - cot)
		vv = (cot - 1.0) / math.Abs(3.0-cot)
	}

	var d func(vector.Vec4) vector.Vec4

	if n == 3 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				-0.25*v.W + 1.25*v.X + 1.25*v.Y - 1.25*v.Z,
				0.25*v.W + 0.75*v.X - 0.25*v.Y + 0.25*v.Z,
				0.25*v.W - 0.25*v.X + 0.75*v.Y + 0.25*v.Z,
				-0.25*v.W + 0.25*v.X + 0.25*v.Y + 0.75*v.Z,
			}
		}
	} else if n == 4 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				0.5 * (v.W + v.X + v.Y - v.Z),
				0.5 * (v.W + v.X - v.Y + v.Z),
				0.5 * (v.W - v.X + v.Y + v.Z),
				0.5 * (-v.W + v.X + v.Y + v.Z),
			}
		}
	} else if n == 6 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				1.25*v.W - 0.25*v.X - 0.25*v.Y + 0.25*v.Z,
				0.75*v.W + 0.25*v.X - 0.75*v.Y + 0.75*v.Z,
				0.75*v.W - 0.75*v.X + 0.25*v.Y + 0.75*v.Z,
				-0.75*v.W + 0.75*v.X + 0.75*v.Y + 0.25*v.Z,
			}
		}
	} else {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				(3.0*sin-1)*(-v.W+v.X+v.Y-v.Z) + v.W,
				cos*(v.W-v.X-v.Y+v.Z) + v.X,
				cos*(v.W-v.X-v.Y+v.Z) + v.Y,
				cos*(-v.W+v.X+v.Y-v.Z) + v.Z,
			}
		}
	}

	var a, b float64
	if n == 3 {
		a = 0.25
		b = Rt5 * 0.25
	} else if n == 4 {
		a = 0.5
		b = 0.5
	} else if n == 5 {
		a = P2 * 0.5 * Rt_2
		b = P_1 * 0.5 * Rt_2
	} else if n == 6 {
		a = Rt3
		b = 1.0
	} else if metric == 'e' {
		a = 1.0
		b = 1.0
	} else {
		a = math.Sqrt(math.Abs(cot / (6.0 - 2.0*cot)))
		b = math.Sqrt(math.Abs((cot - 2.0) / (6.0 - 2.0*cot)))
	}
	f := func(v vector.Vec4) vector.Vec4 {
		return vector.Vec4{a * v.W, b * v.X, b * v.Y, b * v.Z}
	}

	return CellData{
		Metric:          metric,
		NumVertices:     4,
		NumEdges:        6,
		NumFaces:        4,
		FaceReflections: []string{"", "abc", "bc", "c"},
		OuterReflection: "d",
		V:               vector.Vec4{1, 1, 1, 1},
		E:               vector.Vec4{1, 1, 0, 0},
		F:               vector.Vec4{3, 1, 1, -1},
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
			vector.Vec4{1, 1, 1, 1},
			vector.Vec4{1, 1, -1, -1},
			vector.Vec4{1, -1, 1, -1},
			vector.Vec4{1, -1, -1, 1},
		},
		Edges: [][2]int{
			{0, 1}, {0, 2}, {0, 3},
			{1, 2}, {1, 3}, {2, 3},
		},
		Faces: [][]int{
			{0, 2, 1}, {1, 2, 3},
			{2, 0, 3}, {3, 0, 1},
		},
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, -v.Z, -v.Y} },
		Bmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.Y, v.X, v.Z} },
		Cmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Z, v.Y} },
		Dmat: d,
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}
}
