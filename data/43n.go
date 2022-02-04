package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

const (
	eVal43n = 4.0
	pVal43n = 6.0

	eVal43nTrunc = 4.0
	pVal43nTrunc = 11.465072284 // math.Pi / math.Atan(math.Sqrt(1.0/(7.0+4.0*Rt2)))

	eVal43nRect = 4.0
	pVal43nRect = 1e100 //∞
)

func GoursatTetrahedron43n(n float64) GoursatTetrahedron {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	cf := 2.0 * cot / (1 + cot)
	ce := cot
	fe := 0.5 * (1 + cot)

	var cv, fv, ev float64
	if math.Abs(n-pVal43n) < BoundaryEps {
		cv = 3.0
		fv = 2.0
		ev = 1.0
	} else {
		cv = 2.0 * cot / (3.0 - cot)
		fv = (1.0 + cot) / (3.0 - cot)
		ev = 2.0 / (3.0 - cot)
	}

	return GoursatTetrahedron{
		V:   vector.Vec4{W: 1, X: 1, Y: 1, Z: 1},
		E:   vector.Vec4{W: 1, X: 1, Y: 1, Z: 0},
		F:   vector.Vec4{W: 1, X: 1, Y: 0, Z: 0},
		C:   vector.Vec4{W: 1, X: 0, Y: 0, Z: 0},
		CFE: vector.Vec4{W: 0, X: 0, Y: 0, Z: 1},
		CFV: vector.Vec4{W: 0, X: 0, Y: 1, Z: -1},
		CEV: vector.Vec4{W: 0, X: 1, Y: -1, Z: 0},
		FEV: vector.Vec4{W: cot - 1.0, X: 2.0 * cot, Y: 0, Z: 0},
		CF:  cf,
		CE:  ce,
		CV:  cv,
		FE:  fe,
		FV:  fv,
		EV:  ev,
	}

}

func Coxeter43n(n float64) Coxeter {

	//cos := math.Pow(math.Cos(math.Pi/n), 2.0)
	cn := math.Cos(2.0 * math.Pi / n)
	// sin := 1.0 - cos

	return Coxeter{
		P: 4.0,
		Q: 3.0,
		R: n,
		A: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: -v.Z} },
		B: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Z, Z: v.Y} },
		C: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.Y, Y: v.X, Z: v.Z} },
		D: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: (1.0+2.0*cn)*v.W - 2.0*cn*v.X,
				X: (2.0+2.0*cn)*v.W - (1.0+2.0*cn)*v.X,
				Y: v.Y,
				Z: v.Z,
			}
		},
		FaceReflections:    []string{"bc", "c", "cbabc", "abc", "", "babc"},
		GoursatTetrahedron: GoursatTetrahedron43n(n),
	}

}

func Honeycomb43n(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	space := Boundaries(n, eVal43n, pVal43n)

	var scale func(vector.Vec4) vector.Vec4
	if space == 'p' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: Rt3 * v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else if space == 'e' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: math.Sqrt(math.Abs(2.0*cot/(3.0-cot))) * v.W,
				X: math.Sqrt(math.Abs((cot-1.0)/(3.0-cot))) * v.X,
				Y: math.Sqrt(math.Abs((cot-1.0)/(3.0-cot))) * v.Y,
				Z: math.Sqrt(math.Abs((cot-1.0)/(3.0-cot))) * v.Z,
			}
		}
	}

	var innerProd func(vector.Vec4, vector.Vec4) float64
	if space == 'p' {
		innerProd = func(a, b vector.Vec4) float64 { return 3.0*a.W*b.W - a.X*b.X - a.Y*b.Y - a.Z*b.Z }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (2.0*cot*a.W*b.W - (cot-1.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(3.0-cot)
		}
	}

	return Honeycomb{
		Coxeter:  Coxeter43n(n),
		CellType: "spherical",
		Vertices: []vector.Vec4{
			{W: 1, X: 1, Y: 1, Z: 1},
			{W: 1, X: 1, Y: -1, Z: 1},
			{W: 1, X: -1, Y: -1, Z: 1},
			{W: 1, X: -1, Y: 1, Z: 1},
			{W: 1, X: 1, Y: 1, Z: -1},
			{W: 1, X: 1, Y: -1, Z: -1},
			{W: 1, X: -1, Y: -1, Z: -1},
			{W: 1, X: -1, Y: 1, Z: -1},
		},
		Edges: [][2]int{
			{0, 3}, {3, 2}, {2, 1}, {1, 0},
			{7, 4}, {4, 5}, {5, 6}, {6, 7},
			{0, 4}, {1, 5}, {2, 6}, {3, 7},
		},
		Faces: [][]int{
			{0, 1, 2, 3}, {4, 7, 3, 0}, {7, 6, 2, 3},
			{4, 5, 6, 7}, {0, 1, 5, 4}, {1, 2, 6, 5},
		},
		EVal:         eVal43n,
		PVal:         pVal43n,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}

func Honeycomb43nTrunc(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	factor := Rt2 - 1.0

	space := Boundaries(n, eVal43nTrunc, pVal43nTrunc)

	var scale func(vector.Vec4) vector.Vec4
	if space == 'p' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: math.Sqrt(5.0-2.0*Rt2) * v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else if space == 'e' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: math.Sqrt(math.Abs(2.0*cot/(5.0-2.0*Rt2-(3.0-2.0*Rt2)*cot))) * v.W,
				X: math.Sqrt(math.Abs((cot-1.0)/(5.0-2.0*Rt2-(3.0-2.0*Rt2)*cot))) * v.X,
				Y: math.Sqrt(math.Abs((cot-1.0)/(5.0-2.0*Rt2-(3.0-2.0*Rt2)*cot))) * v.Y,
				Z: math.Sqrt(math.Abs((cot-1.0)/(5.0-2.0*Rt2-(3.0-2.0*Rt2)*cot))) * v.Z,
			}
		}
	}

	var innerProd func(vector.Vec4, vector.Vec4) float64
	if space == 'p' {
		innerProd = func(a, b vector.Vec4) float64 { return (5.0-2.0*Rt2)*a.W*b.W - (a.X*b.X + a.Y*b.Y + a.Z*b.Z) }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (2.0*cot*a.W*b.W - (cot-1.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(5.0-2.0*Rt2-(3.0-2.0*Rt2)*cot)
		}
	}

	return Honeycomb{
		Coxeter:  Coxeter43n(n),
		CellType: "spherical",
		Vertices: []vector.Vec4{
			{W: 1, X: 1, Y: 1, Z: factor},
			{W: 1, X: 1, Y: factor, Z: 1},
			{W: 1, X: factor, Y: 1, Z: 1},
			{W: 1, X: 1, Y: 1, Z: -factor},
			{W: 1, X: 1, Y: factor, Z: -1},
			{W: 1, X: factor, Y: 1, Z: -1},
			{W: 1, X: 1, Y: -1, Z: factor},
			{W: 1, X: 1, Y: -factor, Z: 1},
			{W: 1, X: factor, Y: -1, Z: 1},
			{W: 1, X: 1, Y: -1, Z: -factor},
			{W: 1, X: 1, Y: -factor, Z: -1},
			{W: 1, X: factor, Y: -1, Z: -1},
			{W: 1, X: -1, Y: 1, Z: factor},
			{W: 1, X: -1, Y: factor, Z: 1},
			{W: 1, X: -factor, Y: 1, Z: 1},
			{W: 1, X: -1, Y: 1, Z: -factor},
			{W: 1, X: -1, Y: factor, Z: -1},
			{W: 1, X: -factor, Y: 1, Z: -1},
			{W: 1, X: -1, Y: -1, Z: factor},
			{W: 1, X: -1, Y: -factor, Z: 1},
			{W: 1, X: -factor, Y: -1, Z: 1},
			{W: 1, X: -1, Y: -1, Z: -factor},
			{W: 1, X: -1, Y: -factor, Z: -1},
			{W: 1, X: -factor, Y: -1, Z: -1},
		},
		Edges: [][2]int{
			{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 7}, {2, 14}, {3, 4}, {3, 5},
			{4, 5}, {4, 10}, {5, 17}, {6, 7}, {6, 8}, {6, 9}, {7, 8}, {8, 20},
			{9, 10}, {9, 11}, {10, 11}, {11, 23}, {12, 13}, {12, 14}, {12, 15}, {13, 14},
			{13, 19}, {15, 16}, {15, 17}, {16, 17}, {16, 22}, {18, 19}, {18, 20}, {18, 21},
			{19, 20}, {21, 22}, {21, 23}, {22, 23},
		},
		Faces: [][]int{
			{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, 10, 11},
			{12, 13, 14}, {15, 16, 17}, {18, 19, 20}, {21, 22, 23},
			{0, 1, 7, 6, 9, 10, 4, 3}, {0, 2, 14, 12, 15, 17, 5, 3},
			{1, 2, 14, 13, 19, 20, 8, 7}, {6, 8, 20, 18, 21, 23, 11, 9},
			{4, 5, 17, 16, 22, 23, 11, 10}, {12, 13, 19, 18, 21, 22, 16, 15},
		},
		EVal:         eVal43nTrunc,
		PVal:         pVal43nTrunc,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}

func Honeycomb43nRect(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	space := Boundaries(n, eVal43nRect, pVal43nRect)

	var scale func(vector.Vec4) vector.Vec4
	if space == 'e' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: math.Sqrt(math.Abs(cot)) * v.W,
				X: math.Sqrt(math.Abs((cot-1.0)/2.0)) * v.X,
				Y: math.Sqrt(math.Abs((cot-1.0)/2.0)) * v.Y,
				Z: math.Sqrt(math.Abs((cot-1.0)/2.0)) * v.Z,
			}
		}
	}

	innerProd := func(a, b vector.Vec4) float64 {
		return cot*a.W*b.W - (0.5*cot-0.5)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)
	}

	return Honeycomb{
		Coxeter:  Coxeter43n(n),
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
		EVal:         eVal43nRect,
		PVal:         pVal43nRect,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}
