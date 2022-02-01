package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func TruncatedTetrahedronData(n float64) CellData {

	// Trig constants
	cos := math.Pow(math.Cos(math.Pi/n), 2.0)
	sin := 1.0 - cos
	cot := cos / sin

	third := 1.0 / 3.0

	// metric dividers
	eVal := math.Pi / math.Atan(Rt_2)
	pVal := math.Pi / math.Atan(Rt_11)
	space := Boundaries(n, eVal, pVal)

	// reflections
	d := func(v vector.Vec4) vector.Vec4 {
		return vector.Vec4{
			W: (3.0*sin-1.0)*(-v.W+v.X+v.Y-v.Z) + v.W,
			X: cos*(v.W-v.X-v.Y+v.Z) + v.X,
			Y: cos*(v.W-v.X-v.Y+v.Z) + v.Y,
			Z: cos*(-v.W+v.X+v.Y-v.Z) + v.Z,
		}
	}

	// metric
	var f func(vector.Vec4) vector.Vec4
	if space == 'p' {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: Rt11 * v.W, X: 3.0 * v.X, Y: 3.0 * v.Y, Z: 3.0 * v.Z}
		}
	} else if space == 'e' {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		f = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: math.Sqrt(math.Abs(cot/(22.0-2.0*cot))) * Rt3 * v.W,
				X: math.Sqrt(math.Abs((cot-2.0)/(22.0-2.0*cot))) * v.X,
				Y: math.Sqrt(math.Abs((cot-2.0)/(22.0-2.0*cot))) * v.Y,
				Z: math.Sqrt(math.Abs((cot-2.0)/(22.0-2.0*cot))) * v.Z,
			}
		}
	}

	// Inner product
	var innerProd func(vector.Vec4, vector.Vec4) float64
	if space == 'p' {
		innerProd = func(a, b vector.Vec4) float64 { return 11.0*a.W*b.W - 9.0*(a.X*b.X+a.Y*b.Y+a.Z*b.Z) }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (cot*a.W*b.W - (cot-2.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(22.0-2.0*cot)
		}
	}

	Vertices := []vector.Vec4{
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
	}

	return CellData{
		P:               3,
		Q:               3,
		R:               n,
		Space:           space,
		NumVertices:     12,
		NumEdges:        18,
		NumFaces:        8,
		FaceReflections: []string{"", "abc", "bc", "c"},
		OuterReflection: "d",
		CellType:        "spherical",
		EVal:            eVal,
		PVal:            pVal,
		Vertices:        Vertices,
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
		Amat:         func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: -v.Z, Z: -v.Y} },
		Bmat:         func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.Y, Y: v.X, Z: v.Z} },
		Cmat:         func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Z, Z: v.Y} },
		Dmat:         d,
		Fmat:         f,
		InnerProduct: innerProd,
	}
}
