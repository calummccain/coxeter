package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

const (
	eVal33n = 5.1042993121 //math.Pi / math.Atan(Rt_2)
	pVal33n = 6.0

	eVal33nTrunc = 5.1042993121 //math.Pi / math.Atan(Rt_2)
	pVal33nTrunc = 10.727915991 //math.Pi / math.Atan(Rt_11)

	eVal33nRect = 5.1042993121 //math.Pi / math.Atan(Rt_2)
	pVal33nRect = 1e100        //∞
)

func GoursatTetrahedron33n(n float64) GoursatTetrahedron {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	cf := 1.5 * cot / (1 + cot)
	ce := 0.5 * cot
	fe := (1.0 + cot) / 3.0

	var cv, fv, ev float64
	if n == pVal33n {
		cv = 3.0
		fv = 8.0 / 3.0
		ev = 2.0
	} else {
		cv = 0.5 * cot / (3.0 - cot)
		fv = (1.0 + cot) / (3.0 * (3.0 - cot))
		ev = 1.0 / (3.0 - cot)
	}

	return GoursatTetrahedron{
		V:   vector.Vec4{W: 1, X: 1, Y: 1, Z: 1},
		E:   vector.Vec4{W: 1, X: 1, Y: 0, Z: 0},
		F:   vector.Vec4{W: 3, X: 1, Y: 1, Z: -1},
		C:   vector.Vec4{W: 1, X: 0, Y: 0, Z: 0},
		CFE: vector.Vec4{W: 0, X: 0, Y: 1, Z: 1},
		CFV: vector.Vec4{W: 0, X: 1, Y: -1, Z: 0},
		CEV: vector.Vec4{W: 0, X: 0, Y: 1, Z: -1},
		FEV: vector.Vec4{W: cot - 2.0, X: cot, Y: cot, Z: -cot},
		CF:  cf,
		CE:  ce,
		CV:  cv,
		FE:  fe,
		FV:  fv,
		EV:  ev,
	}

}

func Coxeter33n(n float64) Coxeter {

	cos := math.Pow(math.Cos(math.Pi/n), 2.0)
	sin := 1.0 - cos

	return Coxeter{
		P: 3.0,
		Q: 3.0,
		R: n,
		A: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: -v.Z, Z: -v.Y} },
		B: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.Y, Y: v.X, Z: v.Z} },
		C: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Z, Z: v.Y} },
		D: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: (3.0*sin-1.0)*(-v.W+v.X+v.Y-v.Z) + v.W,
				X: cos*(v.W-v.X-v.Y+v.Z) + v.X,
				Y: cos*(v.W-v.X-v.Y+v.Z) + v.Y,
				Z: cos*(-v.W+v.X+v.Y-v.Z) + v.Z,
			}
		},
		FaceReflections:    []string{"", "abc", "bc", "c"},
		GoursatTetrahedron: GoursatTetrahedron33n(n),
	}

}

func Honeycomb33n(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	space := Boundaries(n, eVal33n, pVal33n)

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
				W: math.Sqrt(math.Abs(cot/(6.0-2.0*cot))) * v.W,
				X: math.Sqrt(math.Abs((cot-2.0)/(6.0-2.0*cot))) * v.X,
				Y: math.Sqrt(math.Abs((cot-2.0)/(6.0-2.0*cot))) * v.Y,
				Z: math.Sqrt(math.Abs((cot-2.0)/(6.0-2.0*cot))) * v.Z,
			}
		}
	}

	var innerProd func(vector.Vec4, vector.Vec4) float64
	if space == 'p' {
		innerProd = func(a, b vector.Vec4) float64 { return 3.0*a.W*b.W - a.X*b.X - a.Y*b.Y - a.Z*b.Z }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (cot*a.W*b.W - (cot-2.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(6.0-2.0*cot)
		}
	}

	return Honeycomb{
		Coxeter:  Coxeter33n(n),
		CellType: "spherical",
		Vertices: []vector.Vec4{
			{W: 1, X: 1, Y: 1, Z: 1},
			{W: 1, X: 1, Y: -1, Z: -1},
			{W: 1, X: -1, Y: 1, Z: -1},
			{W: 1, X: -1, Y: -1, Z: 1},
		},
		Edges: [][2]int{
			{0, 1}, {0, 2}, {0, 3},
			{1, 2}, {1, 3}, {2, 3},
		},
		Faces: [][]int{
			{0, 2, 1}, {1, 2, 3},
			{2, 0, 3}, {3, 0, 1},
		},
		EVal:         eVal33n,
		PVal:         pVal33n,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}

func Honeycomb33nTrunc(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	third := 1.0 / 3.0

	space := Boundaries(n, eVal33nTrunc, pVal33nTrunc)

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
		Coxeter:  Coxeter33n(n),
		CellType: "spherical",
		Vertices: []vector.Vec4{
			{W: 1, X: 1, Y: third, Z: third},
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
		EVal:         eVal33nTrunc,
		PVal:         pVal33nTrunc,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}

func Honeycomb33nRect(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	space := Boundaries(n, eVal33nRect, pVal33nRect)

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
		Coxeter:  Coxeter33n(n),
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
		EVal:         eVal33nRect,
		PVal:         pVal33nRect,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}
