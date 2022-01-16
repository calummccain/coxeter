package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func OctahedronData(n float64) CellData {

	// Trig constants
	cos := math.Pow(math.Cos(math.Pi/n), 2.0)
	sin := 1.0 - cos
	cot := cos / sin

	// metric dividers
	eVal := math.Pi / math.Atan(Rt2)
	pVal := 4.0
	metric := Boundaries(n, eVal, pVal)

	// Goursat tetrahedron side lengths
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

	// reflections
	var d func(vector.Vec4) vector.Vec4
	if n == 3 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: 0.5 * (v.W + v.X + v.Y + v.Z),
				X: 0.5 * (v.W + v.X - v.Y - v.Z),
				Y: 0.5 * (v.W - v.X + v.Y - v.Z),
				Z: 0.5 * (v.W - v.X - v.Y + v.Z),
			}
		}
	} else if n == 4 {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: 2*v.W - v.X - v.Y - v.Z,
				X: v.W - v.Y - v.Z,
				Y: v.W - v.X - v.Z,
				Z: v.W - v.X - v.Y,
			}
		}
	} else {
		d = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: (6*cos-2)*(v.W-v.X-v.Y-v.Z) + v.W,
				X: 2*cos*(v.W-v.X-v.Y-v.Z) + v.X,
				Y: 2*cos*(v.W-v.X-v.Y-v.Z) + v.Y,
				Z: 2*cos*(v.W-v.X-v.Y-v.Z) + v.Z,
			}
		}
	}

	// metric
	var f func(vector.Vec4) vector.Vec4
	if n == 3 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: Rt_2 * v.W, X: Rt_2 * v.X, Y: Rt_2 * v.Y, Z: Rt_2 * v.Z}
		}
	} else if n == 4 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: math.Sqrt(math.Abs(cot/(1.0-cot))) * v.W,
				X: math.Sqrt(math.Abs((1.0-2.0*cot)/(1.0-cot))) * v.X,
				Y: math.Sqrt(math.Abs((1.0-2.0*cot)/(1.0-cot))) * v.Y,
				Z: math.Sqrt(math.Abs((1.0-2.0*cot)/(1.0-cot))) * v.Z,
			}
		}
	}

	// Inner product
	var innerProd func(vector.Vec4, vector.Vec4) float64
	if n == 3 {
		innerProd = func(a, b vector.Vec4) float64 { return (a.W*b.W + a.X*b.X + a.Y*b.Y + a.Z*b.Z) / 2.0 }
	} else if n == 4 {
		innerProd = func(a, b vector.Vec4) float64 { return a.W*b.W - a.X*b.X - a.Y*b.Y - a.Z*b.Z }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (cot*a.W*b.W - (2.0*cot-1.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(1.0-cot)
		}
	}

	V := vector.Vec4{W: 1, X: 1, Y: 0, Z: 0}
	E := vector.Vec4{W: 2, X: 1, Y: 1, Z: 0}
	F := vector.Vec4{W: 3, X: 1, Y: 1, Z: 1}
	C := vector.Vec4{W: 1, X: 0, Y: 0, Z: 0}
	CFE := vector.Vec4{W: 0, X: 1, Y: -1, Z: 0}
	CFV := vector.Vec4{W: 0, X: 0, Y: 1, Z: -1}
	CEV := vector.Vec4{W: 0, X: 0, Y: 0, Z: 1}
	FEV := vector.Vec4{W: 2.0*cot - 1.0, X: cot, Y: cot, Z: cot}

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
		{W: 1, X: 1, Y: 0, Z: 0},
		{W: 1, X: -1, Y: 0, Z: 0},
		{W: 1, X: 0, Y: 1, Z: 0},
		{W: 1, X: 0, Y: -1, Z: 0},
		{W: 1, X: 0, Y: 0, Z: 1},
		{W: 1, X: 0, Y: 0, Z: -1},
	}

	return CellData{
		P:               3,
		Q:               4,
		R:               n,
		Metric:          metric,
		NumVertices:     6,
		NumEdges:        12,
		NumFaces:        8,
		FaceReflections: []string{"", "c", "bc", "cbc", "abc", "cabc", "bcabc", "cbcabc"},
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
		Amat:         func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.Y, Y: v.X, Z: v.Z} },
		Bmat:         func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Z, Z: v.Y} },
		Cmat:         func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: -v.Z} },
		Dmat:         d,
		Fmat:         f,
		InnerProduct: innerProd,
	}
}
