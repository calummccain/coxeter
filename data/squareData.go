package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func SquareData(n float64, numberOfFaces int) CellData {

	// Trig constants
	cos := math.Pow(math.Cos(math.Pi/n), 2.0)

	den := math.Sqrt(math.Abs(1.0 - 2.0*cos))

	// metric dividers
	eVal := 2.0
	pVal := 4.0
	metric := Boundaries(n, eVal, pVal)

	// Goursat tetrahedron side lengths
	fe := 1.0 / (1.0 - cos)
	var cv, fv, ev, vv, ce, cf float64
	if metric == 'p' {
		vv = 4.0
		ev = 2.0
		fv = 4.0
		cv = 4.0
		ce = 2.0
		cf = 1.0
	} else {
		vv = 1.0 / (1.0 - 2.0*cos)
		ev = (1.0 - cos) / (1.0 - 2.0*cos)
		fv = 1.0 / (1.0 - 2.0*cos)
		cv = cos * cos / ((1.0 - 2.0*cos) * math.Abs(1.0-2.0*cos))
		ce = cos * cos / (math.Abs(1.0-2.0*cos) * (1.0 - cos))
		cf = cos * cos / math.Abs(1.0-2.0*cos)
	}

	initialVerts := []vector.Vec4{
		{W: 1, X: 0, Y: 1, Z: 0},
		{W: 1, X: 0, Y: -1, Z: 0},
		{W: 1, X: 0, Y: 0, Z: 1},
		{W: 1, X: 0, Y: 0, Z: -1},
	}

	initialEdges := []vector.Vec4{
		{W: 2, X: 0, Y: 1, Z: 1},
		{W: 2, X: 0, Y: 1, Z: -1},
		{W: 2, X: 0, Y: -1, Z: 1},
		{W: 2, X: 0, Y: -1, Z: -1},
	}

	//metric
	var f func(vector.Vec4) vector.Vec4
	if n == 3 {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: Rt_2 * v.W, X: 0.5 * Rt_2 * v.X, Y: v.Y, Z: v.Z}
		}
	} else if metric == 'p' {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: 2.0 * v.W, X: v.X, Y: 2.0 * v.Y, Z: 2.0 * v.Z}
		}
	} else {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W / den, X: cos * v.X / den, Y: math.Sqrt(cos) * v.Y / den, Z: math.Sqrt(cos) * v.Z / den}
		}
	}

	// reflections
	Amat := func(v vector.Vec4) vector.Vec4 {
		return vector.Vec4{W: v.W, X: v.X, Y: v.Z, Z: v.Y}
	}
	Bmat := func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: -v.Z} }
	Cmat := func(v vector.Vec4) vector.Vec4 {
		return vector.Vec4{
			W: (1.0+2.0*cos)*v.W - 2.0*cos*cos*v.X - 2.0*cos*v.Y - 2.0*cos*v.Z,
			X: 2.0*v.W + (1.0-2.0*cos)*v.X - 2.0*v.Y - 2.0*v.Z,
			Y: v.W - cos*v.X - v.Z,
			Z: v.W - cos*v.X - v.Y,
		}
	}
	Dmat := func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: -v.X, Y: v.Y, Z: v.Z} }

	// Inner product
	var innerProd func(vector.Vec4, vector.Vec4) float64
	if n == 3 {
		innerProd = func(a, b vector.Vec4) float64 { return (16.0*a.W*b.W - a.X*b.X - 8.0*a.Y*b.Y - 8.0*a.Z*b.Z) / 8.0 }
	} else if metric == 'p' {
		innerProd = func(a, b vector.Vec4) float64 { return 4.0*a.W*b.W - a.X*b.X - 4.0*a.Y*b.Y - 4.0*a.Z*b.Z }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (a.W*b.W - cos*cos*a.X*b.X - 2.0*cos*a.Y*b.Y - 2.0*cos*a.Z*b.Z) / math.Abs(1.0-2.0*cos)
		}
	}

	V := vector.Vec4{W: 1, X: 0, Y: 1, Z: 0}
	E := vector.Vec4{W: 2, X: 0, Y: 1, Z: 1}
	F := vector.Vec4{W: 1, X: 0, Y: 0, Z: 0}
	C := vector.Vec4{W: cos, X: 1, Y: 0, Z: 0}
	CFE := vector.Vec4{W: 0, X: 0, Y: 1, Z: -1}
	CFV := vector.Vec4{W: 0, X: 0, Y: 0, Z: 1}
	CEV := vector.Vec4{W: 2 * cos, X: 2, Y: 1, Z: 1}
	FEV := vector.Vec4{W: 0, X: 1, Y: 0, Z: 0}

	initialData := CellData{
		P:               4,
		Q:               4,
		R:               n,
		Metric:          metric,
		NumVertices:     0,
		NumEdges:        0,
		NumFaces:        0,
		FaceReflections: []string{},
		OuterReflection: "d",
		CellType:        "euclidean",
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
		Vertices:        []vector.Vec4{},
		Edges:           [][2]int{},
		Faces:           [][]int{},
		Amat:            Amat,
		Bmat:            Bmat,
		Cmat:            Cmat,
		Dmat:            Dmat,
		Fmat:            f,
		InnerProduct:    innerProd,
	}

	fPoints := initialData.MakeFaces(numberOfFaces)

	initialData.Vertices = initialData.MakeRing(initialVerts)

	edges := initialData.MakeRing(initialEdges)

	initialData.GenerateFaceData(fPoints)

	initialData.GenerateEdgeData(edges)

	initialData.OrderFaces()

	return initialData
}
