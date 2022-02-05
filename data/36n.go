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

// TODO
func Honeycomb36nTrunc(n float64, numberOfFaces int) Honeycomb {

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

// TODO
func Honeycomb36nRect(n float64, numberOfFaces int) Honeycomb {

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
