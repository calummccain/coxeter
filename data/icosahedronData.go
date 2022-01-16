package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func IcosahedronData(n float64) CellData {

	// Trig constants
	cos := math.Pow(math.Cos(math.Pi/n), 2.0)
	sin := 1.0 - cos
	cot := cos / sin

	// metric dividers
	eVal := math.Pi / math.Atan(P2)
	pVal := 10.0 / 3.0
	metric := Boundaries(n, eVal, pVal)

	// Goursat tetrahedron side lengths
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

	// reflections
	var d func(vector.Vec4) vector.Vec4
	if n == 3 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: 0.5 * ((P4-1.0)*v.W - (P3-3.0*P_1)*v.Y - (P-3.0*P_3)*v.Z),
				X: v.X,
				Y: 0.5 * (P5*v.W + (2.0-P4)*v.Y - P2*v.Z),
				Z: 0.5 * (P3*v.W - P2*v.Y + v.Z),
			}
		}
	} else {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: (6.0*P2*cos-1.0)*v.W + (2.0*P_1-6.0*P*cos)*v.Y + (2.0*P_3-6.0*cos*P_1)*v.Z,
				X: v.X,
				Y: 2.0*P5*cos*v.W + (1.0-2.0*P4*cos)*v.Y - 2.0*P2*cos*v.Z,
				Z: 2.0*P3*cos*v.W - 2.0*P2*cos*v.Y + (1.0-2.0*cos)*v.Z,
			}
		}
	}

	// metric
	var f func(vector.Vec4) vector.Vec4
	if n == 3 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: P3 * 0.5 * v.W, X: 0.5 * math.Sqrt(3.0*P-1.0) * v.X, Y: 0.5 * math.Sqrt(3.0*P-1.0) * v.Y, Z: 0.5 * math.Sqrt(3.0*P-1.0) * v.Z}
		}
	} else {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: P3 * math.Sqrt(math.Abs(cot/(P2+1.0-P4*cot))) * v.W,
				X: math.Sqrt(math.Abs((P4*cot-1.0)/(P2+1.0-P4*cot))) * v.X,
				Y: math.Sqrt(math.Abs((P4*cot-1.0)/(P2+1.0-P4*cot))) * v.Y,
				Z: math.Sqrt(math.Abs((P4*cot-1.0)/(P2+1.0-P4*cot))) * v.Z,
			}
		}
	}

	// Inner product
	var innerProd func(vector.Vec4, vector.Vec4) float64
	if n == 3 {
		innerProd = func(a, b vector.Vec4) float64 {
			return (P6*a.W*b.W - (3.0*P-1.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / 4.0
		}
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (P6*cot*a.W*b.W - (P4*cot-1.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(P2+1.0-P4*cot)
		}
	}

	V := vector.Vec4{W: 1, X: 1, Y: P, Z: 0}
	E := vector.Vec4{W: 1, X: 0, Y: P, Z: 0}
	F := vector.Vec4{W: 3, X: 0, Y: P3, Z: P}
	C := vector.Vec4{W: 1, X: 0, Y: 0, Z: 0}
	CFE := vector.Vec4{W: 0, X: 1, Y: 0, Z: 0}
	CFV := vector.Vec4{W: 0, X: P, Y: -1, Z: P_1}
	CEV := vector.Vec4{W: 0, X: 0, Y: 0, Z: 1}
	FEV := vector.Vec4{W: P*cot - P_3, X: 0, Y: P2, Z: 1}

	for _, vec := range []vector.Vec4{E, F, C, CFE, CFV, CEV, FEV} {
		vec.Scale(1.0 / math.Sqrt(math.Abs(innerProd(vec, vec))))
	}

	if metric != 'p' {
		V.Scale(1.0 / math.Sqrt(math.Abs(innerProd(V, V))))
	}
	E.Scale(1.0 / math.Sqrt(math.Abs(innerProd(E, E))))
	F.Scale(1.0 / math.Sqrt(math.Abs(innerProd(F, F))))
	C.Scale(1.0 / math.Sqrt(math.Abs(innerProd(C, C))))
	CFE.Scale(1.0 / math.Sqrt(math.Abs(innerProd(CFE, CFE))))
	CFV.Scale(1.0 / math.Sqrt(math.Abs(innerProd(CFV, CFV))))
	CEV.Scale(1.0 / math.Sqrt(math.Abs(innerProd(CEV, CEV))))
	FEV.Scale(1.0 / math.Sqrt(math.Abs(innerProd(FEV, FEV))))

	Vertices := []vector.Vec4{
		{W: 1, X: 1, Y: P, Z: 0}, {W: 1, X: 1, Y: -P, Z: 0}, {W: 1, X: -1, Y: P, Z: 0},
		{W: 1, X: -1, Y: -P, Z: 0}, {W: 1, X: P, Y: 0, Z: 1}, {W: 1, X: -P, Y: 0, Z: 1},
		{W: 1, X: P, Y: 0, Z: -1}, {W: 1, X: -P, Y: 0, Z: -1}, {W: 1, X: 0, Y: 1, Z: P},
		{W: 1, X: 0, Y: 1, Z: -P}, {W: 1, X: 0, Y: -1, Z: P}, {W: 1, X: 0, Y: -1, Z: -P},
	}

	return CellData{
		P:           3,
		Q:           5,
		R:           n,
		Metric:      metric,
		NumVertices: 12,
		NumEdges:    30,
		NumFaces:    20,
		FaceReflections: []string{
			"", "c", "bcbc", "bc", "cbc", "bacbcabacbc",
			"cbacbcabacbc", "bacabacbc", "bcabacbc",
			"cbcabacbc", "abcbc", "abc", "acbc",
			"abacabacbc", "abcabacbc", "cabcabacbc",
			"bacbc", "abacbc", "acabacbc", "cabacbc",
		},
		OuterReflection: "d",
		CellType:        "spherical",
		V:               V,
		E:               E,
		F:               F,
		C:               C,
		CFE:             CFE,
		CFV:             CFV,
		CEV:             CEV,
		FEV:             FEV,
		CF:              cf,
		CE:              ce,
		CV:              cv,
		FE:              fe,
		FV:              fv,
		EV:              ev,
		VV:              vv,
		EVal:            eVal,
		PVal:            pVal,
		Vertices:        Vertices,
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
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: -v.X, Y: v.Y, Z: v.Z} },
		Bmat: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: 0.5 * (v.X + P_1*v.Y - P*v.Z), Y: 0.5 * (P_1*v.X + P*v.Y + v.Z), Z: 0.5 * (-P*v.X + v.Y - P_1*v.Z)}
		},
		Cmat:         func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: -v.Z} },
		Dmat:         d,
		Fmat:         f,
		InnerProduct: innerProd,
	}
}
