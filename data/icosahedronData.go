package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func IcosahedronData(n float64) CellData {

	eVal := math.Pi / math.Atan(P2)
	pVal := 10.0 / 3.0

	cos := math.Pow(math.Cos(math.Pi/n), 2)
	sin := 1 - cos
	cot := cos / sin

	metric := Boundaries(n, eVal, pVal)

	cf := 3.0 * P2 * cot / (1 + cot)
	ce := P4 * cot
	fe := P2 * (1 + cot) / 3.0

	var cv, fv, ev, vv float64

	if metric == 'p' {
		cv = 0.0
		fv = 0.0
		ev = 0.0
		vv = 2.0 / (P + 2.0)
	} else {
		cv = P6 * cot / (1.0 + P2 - P4*cot)
		fv = P4 * (1.0 + cot) / (3.0 * (1.0 + P2 - P4*cot))
		ev = P2 / (1.0 + P2 - P4*cot)
		vv = (P2 - 1.0 + P4*cot) / math.Abs(P2+1.0-P4*cot)
	}

	var d func(vector.Vec4) vector.Vec4

	if n == 3 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				0.5 * ((P4-1.0)*v.W - (P3-3.0*P_1)*v.Y - (P-3.0*P_3)*v.Z),
				v.X,
				0.5 * (P5*v.W + (2.0-P4)*v.Y - P2*v.Z),
				0.5 * (P3*v.W - P2*v.Y + v.Z),
			}
		}
	} else {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				(6.0*P2*cos-1.0)*v.W + (2.0*P_1-6.0*P*cos)*v.Y + (2.0*P_3-6.0*cos*P_1)*v.Z,
				v.X,
				2.0*P5*cos*v.W + (1.0-2.0*P4*cos)*v.Y - 2.0*P2*cos*v.Z,
				2.0*P3*cos*v.W - 2.0*P2*cos*v.Y + (1.0-2.0*cos)*v.Z,
			}
		}
	}

	var f func(vector.Vec4) vector.Vec4

	f3_1 := P3 / 2.0
	f3_2 := math.Sqrt(3.0*P-1.0) / 2.0

	fp := 1.0 / math.Sqrt(P+2.0)

	fn_1 := P3 * math.Sqrt(math.Abs(cot/(P2+1.0-P4*cot)))
	fn_2 := math.Sqrt(math.Abs((P4*cot - 1.0) / (P2 + 1.0 - P4*cot)))

	if n == 3 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{f3_1 * v.W, f3_2 * v.X, f3_2 * v.Y, f3_2 * v.Z}
		}
	} else if metric == 'e' {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{v.W, v.X, v.Y, v.Z}
		}
	} else if metric == 'p' {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{v.W, fp * v.X, fp * v.Y, fp * v.Z}
		}
	} else {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{fn_1 * v.W, fn_2 * v.X, fn_2 * v.Y, fn_2 * v.Z}
		}
	}

	return CellData{
		Metric:      metric,
		NumVertices: 12,
		NumEdges:    30,
		NumFaces:    20,
		FaceReflections: []string{
			"", "c", "bcbc", "bc", "cbc", "bacbcabacbc",
			"cbacbcabacbc", "bacabacbc", "bcabacbc",
			"cbcabacbc", "abcbc", "abc", "acbc",
			"abacabacbc", "abcabacbc", "cabcabacbc",
			"bacbc", "abacbc", "acabacbc", "cabacbc"},
		OuterReflection: "d",
		V:               vector.Vec4{1, 1, P, 0},
		E:               vector.Vec4{1, 0, P, 0},
		F:               vector.Vec4{3, 0, P3, P},
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
			{1, 1, P, 0}, {1, 1, -P, 0}, {1, -1, P, 0},
			{1, -1, -P, 0}, {1, P, 0, 1}, {1, -P, 0, 1},
			{1, P, 0, -1}, {1, -P, 0, -1}, {1, 0, 1, P},
			{1, 0, 1, -P}, {1, 0, -1, P}, {1, 0, -1, -P},
		},
		Edges: [][2]int{
			{0, 2}, {0, 4}, {0, 6}, {0, 8},
			{0, 9}, {1, 3}, {1, 4}, {1, 6},
			{1, 10}, {1, 11}, {2, 5}, {2, 7},
			{2, 8}, {2, 9}, {3, 5}, {3, 7},
			{3, 10}, {3, 11}, {4, 6}, {4, 8},
			{4, 10}, {5, 7}, {5, 8}, {5, 10},
			{6, 9}, {6, 11}, {7, 9}, {7, 11},
			{8, 10}, {9, 11},
		},
		Faces: [][]int{
			{0, 8, 2}, {0, 2, 9}, {0, 6, 4}, {0, 4, 8},
			{0, 9, 6}, {1, 3, 10}, {1, 11, 3}, {1, 4, 6},
			{1, 10, 4}, {1, 6, 11}, {2, 5, 7}, {2, 8, 5},
			{2, 7, 9}, {3, 7, 5}, {3, 5, 10}, {3, 11, 7},
			{4, 10, 8}, {5, 8, 10}, {6, 9, 11}, {7, 11, 9},
		},
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, -v.X, v.Y, v.Z} },
		Bmat: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{v.W, 0.5 * (v.X + P_1*v.Y - P*v.Z), 0.5 * (P_1*v.X + P*v.Y + v.Z), 0.5 * (-P*v.X + v.Y - P_1*v.Z)}
		},
		Cmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Y, -v.Z} },
		Dmat: d,
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}
}
