package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

const (
	eVal34n = 3.2885355431 //math.Pi / math.Atan(Rt2)
	pVal34n = 4.0

	eVal34nTrunc = 3.2885355431 //math.Pi / math.Atan(Rt2)
	pVal34nTrunc = 7.4704783652 //math.Pi / math.Atan(Rt_5)

	eVal34nRect = 3.2885355431 //math.Pi / math.Atan(Rt2)
	pVal34nRect = 1e100        //∞
)

func GoursatTetrahedron34n(n float64) GoursatTetrahedron {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	cf := 3.0 * cot / (1 + cot)
	ce := 2.0 * cot
	fe := 2.0 * (1 + cot) / 3.0

	var cv, fv, ev float64
	if math.Abs(n-pVal34n) < BoundaryEps {
		cv = 1.0
		fv = 2.0
		ev = 1.0
	} else {
		cv = cot / (1.0 - cot)
		fv = (1.0 + cot) / (3.0 * (1.0 - cot))
		ev = 0.5 / (1.0 - cot)
	}

	return GoursatTetrahedron{
		V:   vector.Vec4{W: 1, X: 1, Y: 0, Z: 0},
		E:   vector.Vec4{W: 2, X: 1, Y: 1, Z: 0},
		F:   vector.Vec4{W: 3, X: 1, Y: 1, Z: 1},
		C:   vector.Vec4{W: 1, X: 0, Y: 0, Z: 0},
		CFE: vector.Vec4{W: 0, X: 1, Y: -1, Z: 0},
		CFV: vector.Vec4{W: 0, X: 0, Y: 1, Z: -1},
		CEV: vector.Vec4{W: 0, X: 0, Y: 0, Z: 1},
		FEV: vector.Vec4{W: 2.0*cot - 1.0, X: cot, Y: cot, Z: cot},
		CF:  cf,
		CE:  ce,
		CV:  cv,
		FE:  fe,
		FV:  fv,
		EV:  ev,
	}

}

func Coxeter34n(n float64) Coxeter {

	cos := math.Pow(math.Cos(math.Pi/n), 2.0)
	// sin := 1.0 - cos

	return Coxeter{
		P: 3.0,
		Q: 4.0,
		R: n,
		A: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.Y, Y: v.X, Z: v.Z} },
		B: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Z, Z: v.Y} },
		C: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: -v.Z} },
		D: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: (6*cos-2)*(v.W-v.X-v.Y-v.Z) + v.W,
				X: 2*cos*(v.W-v.X-v.Y-v.Z) + v.X,
				Y: 2*cos*(v.W-v.X-v.Y-v.Z) + v.Y,
				Z: 2*cos*(v.W-v.X-v.Y-v.Z) + v.Z,
			}
		},
		FaceReflections:    []string{"", "c", "bc", "cbc", "abc", "cabc", "bcabc", "cbcabc"},
		GoursatTetrahedron: GoursatTetrahedron34n(n),
	}

}

func Honeycomb34n(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	space := Boundaries(n, eVal34n, pVal34n)

	var scale func(vector.Vec4) vector.Vec4
	if space == 'p' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else if space == 'e' {
		// TODO
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: math.Sqrt(math.Abs(cot/(1.0-cot))) * v.W,
				X: math.Sqrt(math.Abs((1.0-2.0*cot)/(1.0-cot))) * v.X,
				Y: math.Sqrt(math.Abs((1.0-2.0*cot)/(1.0-cot))) * v.Y,
				Z: math.Sqrt(math.Abs((1.0-2.0*cot)/(1.0-cot))) * v.Z,
			}
		}
	}

	var innerProd func(vector.Vec4, vector.Vec4) float64
	if space == 'p' {
		innerProd = func(a, b vector.Vec4) float64 { return a.W*b.W - a.X*b.X - a.Y*b.Y - a.Z*b.Z }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (cot*a.W*b.W - (2.0*cot-1.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(1.0-cot)
		}
	}

	return Honeycomb{
		Coxeter:  Coxeter34n(n),
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
		EVal:         eVal34n,
		PVal:         pVal34n,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}

func Honeycomb34nTrunc(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	third := 1.0 / 3.0
	twoThird := 2.0 / 3.0

	space := Boundaries(n, eVal34nTrunc, pVal34nTrunc)

	var scale func(vector.Vec4) vector.Vec4
	if space == 'p' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: Rt5 * v.W, X: 3.0 * v.X, Y: 3.0 * v.Y, Z: 3.0 * v.Z}
		}
	} else if space == 'e' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: 3.0 * math.Sqrt(math.Abs(cot/(5.0-cot))) * v.W,
				X: 3.0 * math.Sqrt(math.Abs((2.0*cot-1.0)/(5.0-cot))) * v.X,
				Y: 3.0 * math.Sqrt(math.Abs((2.0*cot-1.0)/(5.0-cot))) * v.Y,
				Z: 3.0 * math.Sqrt(math.Abs((2.0*cot-1.0)/(5.0-cot))) * v.Z,
			}
		}
	}

	var innerProd func(vector.Vec4, vector.Vec4) float64
	if space == 'p' {
		innerProd = func(a, b vector.Vec4) float64 { return 5.0*a.W*b.W - 9.0*(a.X*b.X+a.Y*b.Y+a.Z*b.Z) }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (9.0*cot*a.W*b.W - 9.0*(2.0*cot-1.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(5.0-cot)
		}
	}

	return Honeycomb{
		Coxeter:  Coxeter34n(n),
		CellType: "spherical",
		Vertices: []vector.Vec4{
			{W: 1, X: twoThird, Y: third, Z: 0},
			{W: 1, X: twoThird, Y: 0, Z: third},
			{W: 1, X: twoThird, Y: -third, Z: 0},
			{W: 1, X: twoThird, Y: 0, Z: -third},
			{W: 1, X: 0, Y: twoThird, Z: third},
			{W: 1, X: third, Y: twoThird, Z: 0},
			{W: 1, X: 0, Y: twoThird, Z: -third},
			{W: 1, X: -third, Y: twoThird, Z: 0},
			{W: 1, X: third, Y: 0, Z: twoThird},
			{W: 1, X: 0, Y: third, Z: twoThird},
			{W: 1, X: -third, Y: 0, Z: twoThird},
			{W: 1, X: 0, Y: -third, Z: twoThird},
			{W: 1, X: -twoThird, Y: third, Z: 0},
			{W: 1, X: -twoThird, Y: 0, Z: third},
			{W: 1, X: -twoThird, Y: -third, Z: 0},
			{W: 1, X: -twoThird, Y: 0, Z: -third},
			{W: 1, X: 0, Y: -twoThird, Z: third},
			{W: 1, X: third, Y: -twoThird, Z: 0},
			{W: 1, X: 0, Y: -twoThird, Z: -third},
			{W: 1, X: -third, Y: -twoThird, Z: 0},
			{W: 1, X: third, Y: 0, Z: -twoThird},
			{W: 1, X: 0, Y: third, Z: -twoThird},
			{W: 1, X: -third, Y: 0, Z: -twoThird},
			{W: 1, X: 0, Y: -third, Z: -twoThird},
		},
		Edges: [][2]int{
			{0, 1}, {0, 3}, {0, 5}, {1, 2}, {1, 8}, {2, 3}, {2, 17}, {3, 20},
			{4, 5}, {4, 7}, {4, 9}, {5, 6}, {6, 7}, {6, 21}, {7, 12}, {8, 9},
			{8, 11}, {9, 10}, {10, 11}, {10, 13}, {11, 16}, {12, 13}, {12, 15},
			{13, 14}, {14, 15}, {14, 19}, {15, 22}, {16, 17}, {16, 19}, {17, 18},
			{18, 19}, {18, 23}, {20, 21}, {20, 23}, {21, 22}, {22, 23},
		},
		Faces: [][]int{
			{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11},
			{12, 13, 14, 15}, {16, 17, 18, 19}, {20, 21, 22, 23},
			{0, 1, 8, 9, 4, 5}, {0, 3, 20, 21, 6, 5},
			{4, 7, 12, 13, 10, 9}, {1, 2, 17, 16, 11, 8},
			{2, 3, 20, 23, 18, 17}, {6, 7, 12, 15, 22, 21},
			{10, 11, 16, 19, 14, 13}, {14, 15, 22, 23, 18, 19},
		},
		EVal:         eVal34nTrunc,
		PVal:         pVal34nTrunc,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}

func Honeycomb34nRect(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	space := Boundaries(n, eVal34nRect, pVal34nRect)

	var scale func(vector.Vec4) vector.Vec4
	if space == 'e' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: math.Sqrt(math.Abs(2.0*cot)) * v.W,
				X: math.Sqrt(math.Abs((cot-1.0)/2.0)) * v.X,
				Y: math.Sqrt(math.Abs((cot-1.0)/2.0)) * v.Y,
				Z: math.Sqrt(math.Abs((cot-1.0)/2.0)) * v.Z,
			}
		}
	}

	innerProd := func(a, b vector.Vec4) float64 {
		return 2.0*cot*a.W*b.W - (cot-0.5)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)
	}

	return Honeycomb{
		Coxeter:  Coxeter34n(n),
		CellType: "spherical",
		Vertices: []vector.Vec4{
			{W: 1, X: 1, Y: 1, Z: 0},
			{W: 1, X: 1, Y: 0, Z: 1},
			{W: 1, X: 0, Y: 1, Z: 1},
			{W: 1, X: 1, Y: -1, Z: 0},
			{W: 1, X: -1, Y: 0, Z: 1},
			{W: 1, X: 0, Y: 1, Z: -1},
			{W: 1, X: -1, Y: 1, Z: 0},
			{W: 1, X: 1, Y: 0, Z: -1},
			{W: 1, X: 0, Y: -1, Z: 1},
			{W: 1, X: -1, Y: -1, Z: 0},
			{W: 1, X: -1, Y: 0, Z: -1},
			{W: 1, X: 0, Y: -1, Z: -1},
		},
		Edges: [][2]int{
			{0, 1}, {1, 2}, {2, 0}, {0, 5}, {5, 7}, {7, 0},
			{3, 7}, {7, 11}, {11, 3}, {1, 3}, {3, 8}, {8, 1},
			{4, 8}, {4, 9}, {8, 9}, {2, 4}, {2, 6}, {4, 6},
			{5, 6}, {5, 10}, {6, 10}, {9, 10}, {9, 11}, {10, 11},
		},
		Faces: [][]int{
			{0, 1, 2}, {0, 5, 7}, {3, 7, 11}, {1, 3, 8},
			{4, 8, 9}, {2, 4, 6}, {5, 6, 10}, {9, 10, 11},
			{0, 5, 6, 2}, {1, 2, 4, 8}, {0, 1, 3, 7},
			{3, 8, 9, 11}, {5, 7, 11, 10}, {4, 6, 10, 9},
		},
		EVal:         eVal34nRect,
		PVal:         pVal34nRect,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}
