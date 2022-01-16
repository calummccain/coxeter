package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func TetrahedronData(n float64) CellData {

	// Trig constants
	cos := math.Pow(math.Cos(math.Pi/n), 2.0)
	sin := 1.0 - cos
	cot := cos / sin

	// metric dividers
	eVal := math.Pi / math.Atan(Rt_2)
	pVal := 6.0
	metric := Boundaries(n, eVal, pVal)

	// Goursat tetrahedron side lengths
	cf := 1.5 * cot / (1 + cot)
	ce := 0.5 * cot
	fe := (1.0 + cot) / 3.0

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
		vv = (cot - 1.0) / (3.0 - cot)
	}

	// reflections
	var d func(vector.Vec4) vector.Vec4
	if n == 3 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: -0.25*v.W + 1.25*v.X + 1.25*v.Y - 1.25*v.Z,
				X: 0.25*v.W + 0.75*v.X - 0.25*v.Y + 0.25*v.Z,
				Y: 0.25*v.W - 0.25*v.X + 0.75*v.Y + 0.25*v.Z,
				Z: -0.25*v.W + 0.25*v.X + 0.25*v.Y + 0.75*v.Z,
			}
		}
	} else if n == 4 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: 0.5 * (v.W + v.X + v.Y - v.Z),
				X: 0.5 * (v.W + v.X - v.Y + v.Z),
				Y: 0.5 * (v.W - v.X + v.Y + v.Z),
				Z: 0.5 * (-v.W + v.X + v.Y + v.Z),
			}
		}
	} else if n == 6 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: 1.25*v.W - 0.25*v.X - 0.25*v.Y + 0.25*v.Z,
				X: 0.75*v.W + 0.25*v.X - 0.75*v.Y + 0.75*v.Z,
				Y: 0.75*v.W - 0.75*v.X + 0.25*v.Y + 0.75*v.Z,
				Z: -0.75*v.W + 0.75*v.X + 0.75*v.Y + 0.25*v.Z,
			}
		}
	} else {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: (3.0*sin-1.0)*(-v.W+v.X+v.Y-v.Z) + v.W,
				X: cos*(v.W-v.X-v.Y+v.Z) + v.X,
				Y: cos*(v.W-v.X-v.Y+v.Z) + v.Y,
				Z: cos*(-v.W+v.X+v.Y-v.Z) + v.Z,
			}
		}
	}

	// metric
	var f func(vector.Vec4) vector.Vec4
	if n == 3 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: 0.25 * v.W, X: Rt5 * 0.25 * v.X, Y: Rt5 * 0.25 * v.Y, Z: Rt5 * 0.25 * v.Z}
		}
	} else if n == 4 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: 0.5 * v.W, X: 0.5 * v.X, Y: 0.5 * v.Y, Z: 0.5 * v.Z}
		}
	} else if n == 5 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: P2 * 0.5 * Rt_2 * v.W, X: P_1 * 0.5 * Rt_2 * v.X, Y: P_1 * 0.5 * Rt_2 * v.Y, Z: P_1 * 0.5 * Rt_2 * v.Z}
		}
	} else if n == 6 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: Rt3 * v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else if metric == 'e' {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: math.Sqrt(math.Abs(cot/(6.0-2.0*cot))) * Rt3 * v.W,
				X: math.Sqrt(math.Abs((cot-2.0)/(6.0-2.0*cot))) * v.X,
				Y: math.Sqrt(math.Abs((cot-2.0)/(6.0-2.0*cot))) * v.Y,
				Z: math.Sqrt(math.Abs((cot-2.0)/(6.0-2.0*cot))) * v.Z,
			}
		}
	}

	// Inner product
	var innerProd func(vector.Vec4, vector.Vec4) float64
	if n == 3 {
		innerProd = func(a, b vector.Vec4) float64 { return (a.W*b.W + 5.0*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / 16.0 }
	} else if n == 4 {
		innerProd = func(a, b vector.Vec4) float64 { return (a.W*b.W + a.X*b.X + a.Y*b.Y + a.Z*b.Z) / 4.0 }
	} else if n == 5 {
		innerProd = func(a, b vector.Vec4) float64 { return (P6*a.W*b.W + a.X*b.X + a.Y*b.Y + a.Z*b.Z) / (8.0 * P2) }
	} else if n == 6 {
		innerProd = func(a, b vector.Vec4) float64 { return 3.0*a.W*b.W - a.X*b.X - a.Y*b.Y - a.Z*b.Z }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (cot*a.W*b.W - (cot-2.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(6.0-2.0*cot)
		}
	}

	V := vector.Vec4{W: 1, X: 1, Y: 1, Z: 1}
	E := vector.Vec4{W: 1, X: 1, Y: 0, Z: 0}
	F := vector.Vec4{W: 3, X: 1, Y: 1, Z: -1}
	C := vector.Vec4{W: 1, X: 0, Y: 0, Z: 0}
	CFE := vector.Vec4{W: 0, X: 0, Y: 1, Z: 1}
	CFV := vector.Vec4{W: 0, X: 1, Y: -1, Z: 0}
	CEV := vector.Vec4{W: 0, X: 0, Y: 1, Z: -1}
	FEV := vector.Vec4{W: cot - 2.0, X: cot, Y: cot, Z: -cot}

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
		{W: 1, X: 1, Y: 1, Z: 1},
		{W: 1, X: 1, Y: -1, Z: -1},
		{W: 1, X: -1, Y: 1, Z: -1},
		{W: 1, X: -1, Y: -1, Z: 1},
	}

	return CellData{
		P:               3,
		Q:               3,
		R:               n,
		Metric:          metric,
		NumVertices:     4,
		NumEdges:        6,
		NumFaces:        4,
		FaceReflections: []string{"", "abc", "bc", "c"},
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
			{0, 1}, {0, 2}, {0, 3},
			{1, 2}, {1, 3}, {2, 3},
		},
		Faces: [][]int{
			{0, 2, 1}, {1, 2, 3},
			{2, 0, 3}, {3, 0, 1},
		},
		Amat:         func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: -v.Z, Z: -v.Y} },
		Bmat:         func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.Y, Y: v.X, Z: v.Z} },
		Cmat:         func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Z, Z: v.Y} },
		Dmat:         d,
		Fmat:         f,
		InnerProduct: innerProd,
	}
}
