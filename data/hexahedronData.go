package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func HexahedronData(n float64) CellData {

	// Trig constants
	cos := math.Pow(math.Cos(math.Pi/n), 2.0)
	sin := 1.0 - cos
	cot := cos / sin

	// metric dividers
	eVal := 4.0
	pVal := 6.0
	metric := Boundaries(n, eVal, pVal)

	// Goursat tetrahedron side lengths
	cf := 2.0 * cot / (1 + cot)
	ce := cot
	fe := 0.5 * (1 + cot)

	var cv, fv, ev, vv float64
	if metric == 'p' {
		cv = 0.0
		fv = 0.0
		ev = 0.0
		vv = 2.0
	} else if metric == 'e' {
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

	// reflections
	var d func(vector.Vec4) vector.Vec4
	if n == 3 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: v.X,
				X: v.W,
				Y: v.Y,
				Z: v.Z,
			}
		}
	} else if n == 4 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: v.W,
				X: 2.0*v.W - v.X,
				Y: v.Y,
				Z: v.Z,
			}
		}
	} else if n == 5 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: P*v.W - P_1*v.X,
				X: P2*v.W - P*v.X,
				Y: v.Y,
				Z: v.Z,
			}
		}
	} else if n == 6 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: 2.0*v.W - v.X,
				X: 3.0*v.W - 2.0*v.X,
				Y: v.Y,
				Z: v.Z,
			}
		}
	} else {
		cn := math.Cos(2.0 * math.Pi / n)
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: (1.0+2.0*cn)*v.W - 2.0*cn*v.X,
				X: (2.0+2.0*cn)*v.W - (1.0+2.0*cn)*v.X,
				Y: v.Y,
				Z: v.Z,
			}
		}
	}

	// metric
	var f func(vector.Vec4) vector.Vec4
	if n == 3 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: 0.5 * v.W, X: 0.5 * v.X, Y: 0.5 * v.Y, Z: 0.5 * v.Z}
		}
	} else if n == 4 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else if n == 5 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: P2 * Rt_2 * v.W, X: math.Sqrt(0.5*P) * v.X, Y: math.Sqrt(0.5*P) * v.Y, Z: math.Sqrt(0.5*P) * v.Z}
		}
	} else if n == 6 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: Rt3 * v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: math.Sqrt(math.Abs(2.0*cot/(3.0-cot))) * v.W,
				X: math.Sqrt(math.Abs((cot-1.0)/(3.0-cot))) * v.X,
				Y: math.Sqrt(math.Abs((cot-1.0)/(3.0-cot))) * v.Y,
				Z: math.Sqrt(math.Abs((cot-1.0)/(3.0-cot))) * v.Z,
			}
		}
	}

	// Inner product
	var innerProd func(vector.Vec4, vector.Vec4) float64
	if n == 3 {
		innerProd = func(a, b vector.Vec4) float64 { return (a.W*b.W + a.X*b.X + a.Y*b.Y + a.Z*b.Z) / 4.0 }
	} else if n == 4 {
		innerProd = func(a, b vector.Vec4) float64 { return a.X*b.X + a.Y*b.Y + a.Z*b.Z }
	} else if n == 5 {
		innerProd = func(a, b vector.Vec4) float64 { return (P4*a.W*b.W - P*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / 2.0 }
	} else if n == 6 {
		innerProd = func(a, b vector.Vec4) float64 { return 3.0*a.W*b.W - a.X*b.X - a.Y*b.Y - a.Z*b.Z }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (2.0*cot*a.W*b.W - (cot-1.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(3.0-cot)
		}
	}

	V := vector.Vec4{W: 1, X: 1, Y: 1, Z: 1}
	E := vector.Vec4{W: 1, X: 1, Y: 1, Z: 0}
	F := vector.Vec4{W: 1, X: 1, Y: 0, Z: 0}
	C := vector.Vec4{W: 1, X: 0, Y: 0, Z: 0}
	CFE := vector.Vec4{W: 0, X: 0, Y: 0, Z: 1}
	CFV := vector.Vec4{W: 0, X: 0, Y: 1, Z: -1}
	CEV := vector.Vec4{W: 0, X: 1, Y: -1, Z: 0}
	FEV := vector.Vec4{W: cot - 1.0, X: 2.0 * cot, Y: 0, Z: 0}

	Vertices := []vector.Vec4{
		{W: 1, X: 1, Y: 1, Z: 1}, {W: 1, X: 1, Y: -1, Z: 1},
		{W: 1, X: -1, Y: -1, Z: 1}, {W: 1, X: -1, Y: 1, Z: 1},
		{W: 1, X: 1, Y: 1, Z: -1}, {W: 1, X: 1, Y: -1, Z: -1},
		{W: 1, X: -1, Y: -1, Z: -1}, {W: 1, X: -1, Y: 1, Z: -1},
	}

	return CellData{
		P:               4,
		Q:               3,
		R:               n,
		Metric:          metric,
		NumVertices:     8,
		NumEdges:        12,
		NumFaces:        6,
		FaceReflections: []string{"bc", "c", "cbabc", "abc", "", "babc"},
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
			{0, 3}, {3, 2}, {2, 1}, {1, 0},
			{7, 4}, {4, 5}, {5, 6}, {6, 7},
			{0, 4}, {1, 5}, {2, 6}, {3, 7},
		},
		Faces: [][]int{
			{0, 1, 2, 3}, {4, 7, 3, 0}, {7, 6, 2, 3},
			{4, 5, 6, 7}, {0, 1, 5, 4}, {1, 2, 6, 5},
		},
		Amat:         func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: -v.Z} },
		Bmat:         func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Z, Z: v.Y} },
		Cmat:         func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.Y, Y: v.X, Z: v.Z} },
		Dmat:         d,
		Fmat:         f,
		InnerProduct: innerProd,
	}
}
