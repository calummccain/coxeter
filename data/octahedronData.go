package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func OctahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(Rt2)
	pVal := 4.0

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1.0 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	cf := 3.0 * cot / (1 + cot)
	ce := 2.0 * cot
	fe := 2.0 * (1 + cot) / 3.0

	var cv, fv, ev, vv float64

	if metric == 'p' {
		cv = 1.0
		fv = 2.0
		ev = 1.0
		vv = 1.0
	} else {
		cv = cot / (1.0 - cot)
		fv = (1.0 + cot) / (3.0 * (1.0 - cot))
		ev = 0.5 / (1.0 - cot)
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
		a = Rt_2
		b = Rt_2
	} else if n == 4 {
		a = 1.0
		b = 1.0
	} else if metric == 'e' {
		a = 1.0
		b = 1.0
	} else {
		a = math.Sqrt(math.Abs(cot / (1.0 - cot)))
		b = math.Sqrt(math.Abs((1.0 - 2.0*cot) / (1.0 - cot)))
	}
	f = func(v vector.Vec4) vector.Vec4 {
		return vector.Vec4{a * v.W, b * v.X, b * v.Y, b * v.Z}
	}

	return CellData{
		Metric:          metric,
		NumVertices:     6,
		NumEdges:        12,
		NumFaces:        8,
		FaceReflections: []string{"", "c", "bc", "cbc", "abc", "cabc", "bcabc", "cbcabc"},
		OuterReflection: "d",
		V:               vector.Vec4{1, 1, 0, 0},
		E:               vector.Vec4{2, 1, 1, 0},
		F:               vector.Vec4{3, 1, 1, 1},
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
			vector.Vec4{1, 1, 0, 0},
			vector.Vec4{1, -1, 0, 0},
			vector.Vec4{1, 0, 1, 0},
			vector.Vec4{1, 0, -1, 0},
			vector.Vec4{1, 0, 0, 1},
			vector.Vec4{1, 0, 0, -1},
		},
		Edges: [][2]int{
			{0, 2}, {0, 3}, {0, 4}, {0, 5},
			{1, 2}, {1, 3}, {1, 4}, {1, 5},
			{2, 4}, {4, 3}, {3, 5}, {5, 2},
		},
		Faces: [][]int{
			{0, 2, 4}, {0, 5, 2},
			{0, 4, 3}, {0, 3, 5},
			{1, 4, 2}, {1, 2, 5},
			{1, 3, 4}, {1, 5, 3},
		},
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.Y, v.X, v.Z} },
		Bmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Z, v.Y} },
		Cmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Y, -v.Z} },
		Dmat: d,
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}
}
