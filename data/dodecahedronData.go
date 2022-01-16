package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func DodecahedronData(n float64) CellData {

	// Trig constants
	cos := math.Pow(math.Cos(math.Pi/n), 2.0)
	sin := 1.0 - cos
	cot := cos / sin

	// metric dividers
	eVal := math.Pi / math.Atan(P)
	pVal := 6.0
	metric := Boundaries(n, eVal, pVal)

	// Goursat tetrahedron side lengths
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

	// reflections
	var d func(vector.Vec4) vector.Vec4
	if n == 3 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: (P*v.W + v.X*P_3 + v.Z*P_4) * 0.5,
				X: (P3*v.W - v.X*P_1 - P*v.Z) * 0.5,
				Y: v.Y,
				Z: (P2*v.W - P*v.X + v.Z) * 0.5,
			}
		}
	} else if n == 4 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: P2*v.W - v.X - v.Z*P_1,
				X: P3*v.W - P*v.X - P*v.Z,
				Y: v.Y,
				Z: P2*v.W - P*v.X,
			}
		}
	} else if n == 5 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: ((4*P+1)*v.W - (5-P)*v.X - (5*P-6)*v.Z) * 0.5,
				X: (P5*v.W + (2-P4)*v.X - P3*v.Z) * 0.5,
				Y: v.Y,
				Z: (P4*v.W - P3*v.X - v.Z*P_1) * 0.5,
			}
		}
	} else if n == 6 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: ((2+P4)*v.W - P3*v.X - P2*v.Z) * 0.5,
				X: (3*P3*v.W + (2-3*P2)*v.X - 3*P*v.Z) * 0.5,
				Y: v.Y,
				Z: (3*P2*v.W - 3*P*v.X - v.Z) * 0.5,
			}
		}
	} else {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: (2*P*Rt5*cos-1)*v.W - (2*Rt5*cos-2*P_1)*v.X - (2*Rt5*cos*P_1-2*P_2)*v.Z,
				X: 2*P3*cos*v.W + (1-2*P2*cos)*v.X - 2*P*cos*v.Z,
				Y: v.Y,
				Z: 2*P2*cos*v.W - 2*P*cos*v.X + (1-2*cos)*v.Z,
			}
		}
	}

	// metric
	var f func(vector.Vec4) vector.Vec4
	if n == 3 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: P2 * 0.5 * Rt_2 * v.W, X: P_1 * 0.5 * Rt_2 * v.X, Y: P_1 * 0.5 * Rt_2 * v.Y, Z: P_1 * 0.5 * Rt_2 * v.Z}
		}
	} else if n == 4 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: P2 * Rt_2 * v.W, X: math.Sqrt(P*0.5) * v.X, Y: math.Sqrt(P*0.5) * v.Y, Z: math.Sqrt(P*0.5) * v.Z}
		}
	} else if n == 5 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: P4 * 0.5 * v.W, X: 0.5 * math.Sqrt(4*P-1) * v.X, Y: 0.5 * math.Sqrt(4*P-1) * v.Y, Z: 0.5 * math.Sqrt(4*P-1) * v.Z}
		}
	} else if n == 6 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: Rt3 * v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: P2 * math.Sqrt(math.Abs(cot/(cot-3.0))) * v.W,
				X: math.Sqrt(math.Abs((P2*cot-1.0)/(cot-3.0))) * v.X,
				Y: math.Sqrt(math.Abs((P2*cot-1.0)/(cot-3.0))) * v.Y,
				Z: math.Sqrt(math.Abs((P2*cot-1.0)/(cot-3.0))) * v.Z,
			}
		}
	}

	// Inner product
	var innerProd func(vector.Vec4, vector.Vec4) float64
	if n == 3 {
		innerProd = func(a, b vector.Vec4) float64 { return (P6*a.W*b.W + a.X*b.X + a.Y*b.Y + a.Z*b.Z) / (8.0 * P2) }
	} else if n == 4 {
		innerProd = func(a, b vector.Vec4) float64 { return (P4*a.W*b.W - P*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / 2.0 }
	} else if n == 5 {
		innerProd = func(a, b vector.Vec4) float64 {
			return (P8*a.W*b.W - P2*(4.0*P-1.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / 4
		}
	} else if n == 6 {
		innerProd = func(a, b vector.Vec4) float64 { return 3.0*a.W*b.W - a.X*b.X - a.Y*b.Y - a.Z*b.Z }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (P4*cot*a.W*b.W - (P2*cot-1.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(3.0-cot)
		}
	}

	V := vector.Vec4{W: 1, X: P, Y: P_1, Z: 0}
	E := vector.Vec4{W: 1, X: P, Y: 0, Z: 0}
	F := vector.Vec4{W: 3 - P, X: P, Y: 0, Z: 1}
	C := vector.Vec4{W: 1, X: 0, Y: 0, Z: 0}
	CFE := vector.Vec4{W: 0, X: 0, Y: 1, Z: 0}
	CFV := vector.Vec4{W: 0, X: 1, Y: -P2, Z: -P}
	CEV := vector.Vec4{W: 0, X: 0, Y: 0, Z: 1}
	FEV := vector.Vec4{W: P2*cot - 1.0, X: P3 * cot, Y: 0, Z: P2 * cot}

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
		{W: 1, X: 1, Y: 1, Z: 1}, {W: 1, X: 1, Y: 1, Z: -1}, {W: 1, X: 1, Y: -1, Z: 1}, {W: 1, X: 1, Y: -1, Z: -1},
		{W: 1, X: -1, Y: 1, Z: 1}, {W: 1, X: -1, Y: 1, Z: -1}, {W: 1, X: -1, Y: -1, Z: 1}, {W: 1, X: -1, Y: -1, Z: -1},
		{W: 1, X: 0, Y: P, Z: P_1}, {W: 1, X: 0, Y: P, Z: -P_1}, {W: 1, X: 0, Y: -P, Z: P_1}, {W: 1, X: 0, Y: -P, Z: -P_1},
		{W: 1, X: P, Y: P_1, Z: 0}, {W: 1, X: P, Y: -P_1, Z: 0}, {W: 1, X: -P, Y: P_1, Z: 0}, {W: 1, X: -P, Y: -P_1, Z: 0},
		{W: 1, X: P_1, Y: 0, Z: P}, {W: 1, X: -P_1, Y: 0, Z: P}, {W: 1, X: P_1, Y: 0, Z: -P}, {W: 1, X: -P_1, Y: 0, Z: -P},
	}

	return CellData{
		P:           5,
		Q:           3,
		R:           n,
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
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: -v.Y, Z: v.Z} },
		Bmat: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: 0.5 * (P*v.X + v.Y + v.Z*P_1), Y: 0.5 * (v.X - v.Y*P_1 - P*v.Z), Z: 0.5 * (v.X*P_1 - P*v.Y + v.Z)}
		},
		Cmat:         func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: -v.Z} },
		Dmat:         d,
		Fmat:         f,
		InnerProduct: innerProd,
	}
}
