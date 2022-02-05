package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

const (
	eVal36n = 2.0 // Really?
	pVal36n = 3.0

	eVal36nTrunc = 2.0          // Really?
	pVal36nTrunc = 10.727915991 //math.Pi / math.Atan(Rt_11)

	eVal36nRect = 2.0   // Really?
	pVal36nRect = 1e100 //∞
)

func GoursatTetrahedron36n(n float64) GoursatTetrahedron {

	//cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)
	cos := math.Pow(math.Cos(math.Pi/n), 2.0)

	fe := 1.0 / (1.0 - cos)
	var cf, cv, fv, ev, ce float64
	if math.Abs(n-pVal36n) < BoundaryEps {
		ev = 12.0
		fv = 16.0
		cv = 16.0
		ce = 4.0 / 3.0
		cf = 1.0
	} else {
		ev = (1.0 - cos) / (1.0 - 4.0*cos)
		fv = 1.0 / (1.0 - 4.0*cos)
		cv = cos * cos / ((1.0 - 4.0*cos) * math.Abs(1.0-4.0*cos))
		cf = cos * cos / math.Abs(1.0-4.0*cos)
		ce = cos * cos / (math.Abs(1.0-4.0*cos) * (1.0 - cos))
	}

	return GoursatTetrahedron{
		V:   vector.Vec4{W: 1, X: 0, Y: 2, Z: 0},
		E:   vector.Vec4{W: 2, X: 0, Y: 1, Z: 1},
		F:   vector.Vec4{W: 1, X: 0, Y: 0, Z: 0},
		C:   vector.Vec4{W: cos, X: 1, Y: 0, Z: 0},
		CFE: vector.Vec4{W: 0, X: 0, Y: 3, Z: -1},
		CFV: vector.Vec4{W: 0, X: 0, Y: 0, Z: 1},
		CEV: vector.Vec4{W: 2 * cos, X: 2, Y: 1, Z: 1},
		FEV: vector.Vec4{W: 0, X: 1, Y: 0, Z: 0},
		CF:  cf,
		CE:  ce,
		CV:  cv,
		FE:  fe,
		FV:  fv,
		EV:  ev,
	}

}

func Coxeter36n(n float64, numberOfFaces int) (Coxeter, []vector.Vec4) {

	cos := math.Pow(math.Cos(math.Pi/n), 2.0)
	//sin := 1.0 - cos

	coxeter := Coxeter{
		P: 3.0,
		Q: 6.0,
		R: n,
		A: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: 0.5 * (-v.Y + 3.0*v.Z), Z: 0.5 * (v.Y + v.Z)}
		},
		B: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: -v.Z} },
		C: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: (1.0+2.0*cos)*v.W - 2.0*cos*cos*v.X - cos*v.Y - 3.0*cos*v.Z,
				X: 2.0*v.W + (1.0-2.0*cos)*v.X - v.Y - 3.0*v.Z,
				Y: v.W - cos*v.X + 0.5*v.Y - 1.5*v.Z,
				Z: v.W - cos*v.X - 0.5*v.Y - 0.5*v.Z,
			}
		},
		D:                  func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: -v.X, Y: v.Y, Z: v.Z} },
		FaceReflections:    []string{},
		GoursatTetrahedron: GoursatTetrahedron36n(n),
	}

	faceNames, facePoints := coxeter.MakeFaces(numberOfFaces)

	coxeter.FaceReflections = faceNames

	return coxeter, facePoints

}

func Honeycomb36n(n float64, numberOfFaces int) Honeycomb {

	coxeter, facePoints := Coxeter36n(n, numberOfFaces)

	cos := math.Pow(math.Cos(math.Pi/n), 2.0)

	den := math.Sqrt(math.Abs(1.0 - 4.0*cos))

	space := Boundaries(n, eVal36n, pVal36n)

	var scale func(vector.Vec4) vector.Vec4
	if space == 'p' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: 4.0 * v.W, X: v.X, Y: 2.0 * v.Y, Z: 2.0 * Rt3 * v.Z}
		}
	} else {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W / den, X: cos * v.X / den, Y: math.Sqrt(cos) * v.Y / den, Z: math.Sqrt(3.0*cos) * v.Z / den}
		}
	}

	var innerProd func(vector.Vec4, vector.Vec4) float64
	if space == 'p' {
		innerProd = func(a, b vector.Vec4) float64 { return 16.0*a.W*b.W - a.X*b.X - 4.0*a.Y*b.Y - 12.0*a.Z*b.Z }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (a.W*b.W - cos*cos*a.X*b.X - cos*a.Y*b.Y - 3.0*cos*a.Z*b.Z) / math.Abs(1.0-4.0*cos)
		}
	}

	initialVerts := []vector.Vec4{
		{W: 1, X: 0, Y: 2, Z: 0},
		{W: 1, X: 0, Y: -1, Z: 1},
		{W: 1, X: 0, Y: -1, Z: -1},
	}

	initialEdges := []vector.Vec4{
		{W: 2, X: 0, Y: 1, Z: 1},
		{W: 2, X: 0, Y: -2, Z: 0},
		{W: 2, X: 0, Y: 1, Z: -1},
	}

	honeycomb := Honeycomb{
		Coxeter:      coxeter,
		CellType:     "euclidean",
		Vertices:     []vector.Vec4{},
		Edges:        [][2]int{},
		Faces:        [][]int{},
		EVal:         eVal36n,
		PVal:         pVal36n,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}

	honeycomb.Vertices = honeycomb.Coxeter.MakeRing(initialVerts)

	edgePoints := honeycomb.Coxeter.MakeRing(initialEdges)

	honeycomb.GenerateFaceData(facePoints)

	honeycomb.GenerateEdgeData(edgePoints)

	honeycomb.OrderFaces()

	return honeycomb
}

func Honeycomb36nTrunc(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	third := 1.0 / 3.0

	space := Boundaries(n, eVal36nTrunc, pVal36nTrunc)

	var scale func(vector.Vec4) vector.Vec4
	if space == 'p' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: Rt11 * v.W, X: 3.0 * v.X, Y: 3.0 * v.Y, Z: 3.0 * v.Z}
		}
	} else if space == 'e' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: math.Sqrt(math.Abs(cot/(22.0-2.0*cot))) * Rt3 * v.W,
				X: math.Sqrt(math.Abs((cot-2.0)/(22.0-2.0*cot))) * v.X,
				Y: math.Sqrt(math.Abs((cot-2.0)/(22.0-2.0*cot))) * v.Y,
				Z: math.Sqrt(math.Abs((cot-2.0)/(22.0-2.0*cot))) * v.Z,
			}
		}
	}

	var innerProd func(vector.Vec4, vector.Vec4) float64
	if space == 'p' {
		innerProd = func(a, b vector.Vec4) float64 { return 11.0*a.W*b.W - 9.0*(a.X*b.X+a.Y*b.Y+a.Z*b.Z) }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (cot*a.W*b.W - (cot-2.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(22.0-2.0*cot)
		}
	}

	return Honeycomb{
		Coxeter:  Coxeter36n(n),
		CellType: "spherical",
		Vertices: []vector.Vec4{
			{1, 1, third, third},
			{W: 1, X: third, Y: 1, Z: third},
			{W: 1, X: third, Y: third, Z: 1},
			{W: 1, X: 1, Y: -third, Z: -third},
			{W: 1, X: third, Y: -1, Z: -third},
			{W: 1, X: third, Y: -third, Z: -1},
			{W: 1, X: -1, Y: third, Z: -third},
			{W: 1, X: -third, Y: 1, Z: -third},
			{W: 1, X: -third, Y: third, Z: -1},
			{W: 1, X: -1, Y: -third, Z: third},
			{W: 1, X: -third, Y: -1, Z: third},
			{W: 1, X: -third, Y: -third, Z: 1},
		},
		Edges: [][2]int{
			{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 7}, {2, 11},
			{3, 4}, {3, 5}, {4, 5}, {4, 10}, {5, 8}, {6, 7},
			{6, 8}, {6, 9}, {7, 8}, {9, 10}, {9, 11}, {10, 11},
		},
		Faces: [][]int{
			{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, 10, 11},
			{0, 2, 11, 10, 4, 3}, {0, 1, 7, 8, 5, 3},
			{1, 2, 11, 9, 6, 7}, {4, 5, 8, 6, 9, 10},
		},
		EVal:         eVal36nTrunc,
		PVal:         pVal36nTrunc,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}

func Honeycomb36nRect(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	space := Boundaries(n, eVal36nRect, pVal36nRect)

	var scale func(vector.Vec4) vector.Vec4
	if space == 'e' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: math.Sqrt(math.Abs(cot/2.0)) * Rt3 * v.W,
				X: math.Sqrt(math.Abs((cot-2.0)/2.0)) * v.X,
				Y: math.Sqrt(math.Abs((cot-2.0)/2.0)) * v.Y,
				Z: math.Sqrt(math.Abs((cot-2.0)/2.0)) * v.Z,
			}
		}
	}

	innerProd := func(a, b vector.Vec4) float64 {
		return (cot*a.W*b.W - (cot-2.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / 2.0
	}

	return Honeycomb{
		Coxeter:  Coxeter36n(n),
		CellType: "spherical",
		Vertices: []vector.Vec4{
			{W: 1, X: 1, Y: 0, Z: 0},
			{W: 1, X: -1, Y: 0, Z: 0},
			{W: 1, X: 0, Y: 1, Z: 0},
			{W: 1, X: 0, Y: -1, Z: 0},
			{W: 1, X: 0, Y: 0, Z: 1},
			{W: 1, X: 0, Y: 0, Z: -1},
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
		EVal:         eVal36nRect,
		PVal:         pVal36nRect,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}
