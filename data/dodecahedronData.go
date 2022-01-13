package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func DodecahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(P)
	pVal := 6.0

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	cf := (P + 2.0) * cot / (1 + cot)
	ce := P2 * cot
	fe := P2 * (1 + cot) / (P + 2.0)

	var cv, fv, ev, vv float64

	if metric == 'p' {
		cv = 0.0
		fv = 0.0
		ev = 0.0
		vv = 2.0 * P_2
	} else {
		cv = P4 * cot / (3.0 - cot)
		fv = P4 * (1.0 + cot) / ((P + 2) * (3.0 - cot))
		ev = P2 / (3.0 - cot)
		vv = (cot + P + P_1) / math.Abs(cot-3.0)
	}

	var d func(vector.Vec4) vector.Vec4
	if n == 3 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				(P*v.W + v.X*P_3 + v.Z*P_4) * 0.5,
				(P3*v.W - v.X*P_1 - P*v.Z) * 0.5,
				v.Y,
				(P2*v.W - P*v.X + v.Z) * 0.5,
			}
		}
	} else if n == 4 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				P2*v.W - v.X - v.Z*P_1,
				P3*v.W - P*v.X - P*v.Z,
				v.Y,
				P2*v.W - P*v.X,
			}
		}
	} else if n == 5 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				((4*P+1)*v.W - (5-P)*v.X - (5*P-6)*v.Z) * 0.5,
				(P5*v.W + (2-P4)*v.X - P3*v.Z) * 0.5,
				v.Y,
				(P4*v.W - P3*v.X - v.Z*P_1) * 0.5,
			}
		}
	} else if n == 6 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				((2+P4)*v.W - P3*v.X - P2*v.Z) * 0.5,
				(3*P3*v.W + (2-3*P2)*v.X - 3*P*v.Z) * 0.5,
				v.Y,
				(3*P2*v.W - 3*P*v.X - v.Z) * 0.5,
			}
		}
	} else {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				(2*P*Rt5*cos-1)*v.W - (2*Rt5*cos-2*P_1)*v.X - (2*Rt5*cos*P_1-2*P_2)*v.Z,
				2*P3*cos*v.W + (1-2*P2*cos)*v.X - 2*P*cos*v.Z,
				v.Y,
				2*P2*cos*v.W - 2*P*cos*v.X + (1-2*cos)*v.Z,
			}
		}
	}

	var f func(vector.Vec4) vector.Vec4
	var a, b float64

	if n == 3 {
		a = P2 * 0.5 * Rt_2
		b = P_1 * 0.5 * Rt_2
	} else if n == 4 {
		a = P2 * Rt_2
		b = math.Sqrt(P * 0.5)
	} else if n == 5 {
		a = P4 * 0.5
		b = P * 0.5 * math.Sqrt(4*P-1)
	} else if n == 6 {
		a = Rt3
		b = 1
	} else if metric == 'e' {
		a = 1
		b = 1
	} else {
		a = P2 * math.Sqrt(math.Abs(cot/(cot-3.0)))
		b = math.Sqrt(math.Abs((P2*cot - 1.0) / (cot - 3.0)))
	}
	f = func(v vector.Vec4) vector.Vec4 {
		return vector.Vec4{a * v.W, b * v.X, b * v.Y, b * v.Z}
	}

	return CellData{
		Metric:      metric,
		NumVertices: 20,
		NumEdges:    30,
		NumFaces:    12,
		FaceReflections: []string{
			"", "c", "bc", "acacbabc", "cacbabc",
			"cbabc", "bacbabc", "cbabacbabc", "babacbabc",
			"abc", "acbabc", "abacbabc",
		},
		OuterReflection: "d",
		V:               vector.Vec4{1, P, P_1, 0},
		E:               vector.Vec4{1, P, 0, 0},
		F:               vector.Vec4{3 - P, P, 0, 1},
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
			{1, 1, 1, 1}, {1, 1, 1, -1}, {1, 1, -1, 1}, {1, 1, -1, -1},
			{1, -1, 1, 1}, {1, -1, 1, -1}, {1, -1, -1, 1}, {1, -1, -1, -1},
			{1, 0, P, P_1}, {1, 0, P, -P_1}, {1, 0, -P, P_1}, {1, 0, -P, -P_1},
			{1, P, P_1, 0}, {1, P, -P_1, 0}, {1, -P, P_1, 0}, {1, -P, -P_1, 0},
			{1, P_1, 0, P}, {1, -P_1, 0, P}, {1, P_1, 0, -P}, {1, -P_1, 0, -P},
		},
		Edges: [][2]int{
			{0, 8}, {0, 12}, {0, 16}, {1, 9},
			{1, 12}, {1, 18}, {2, 10}, {2, 13},
			{2, 16}, {3, 11}, {3, 13}, {3, 18},
			{4, 8}, {4, 14}, {4, 17}, {5, 9},
			{5, 14}, {5, 19}, {6, 10}, {6, 15},
			{6, 17}, {7, 11}, {7, 15}, {7, 19},
			{8, 9}, {10, 11}, {12, 13}, {14, 15},
			{16, 17}, {18, 19},
		},
		Faces: [][]int{
			{0, 16, 2, 13, 12}, {1, 12, 13, 3, 18},
			{0, 12, 1, 9, 8}, {0, 8, 4, 17, 16},
			{2, 16, 17, 6, 10}, {1, 18, 19, 5, 9},
			{4, 8, 9, 5, 14}, {5, 19, 7, 15, 14},
			{6, 17, 4, 14, 15}, {3, 13, 2, 10, 11},
			{3, 11, 7, 19, 18}, {11, 10, 6, 15, 7},
		},
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, -v.Y, v.Z} },
		Bmat: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{v.W, 0.5 * (P*v.X + v.Y + v.Z*P_1), 0.5 * (v.X - v.Y*P_1 - P*v.Z), 0.5 * (v.X*P_1 - P*v.Y + v.Z)}
		},
		Cmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Y, -v.Z} },
		Dmat: d,
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}
}
