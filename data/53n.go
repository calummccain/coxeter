package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

const (
	eVal53n = 3.0884042547 // math.Pi / math.Atan(P)
	pVal53n = 6.0

	eVal53nTrunc = 3.0884042547 // math.Pi / math.Atan(P)
	pVal53nTrunc = 11.465072284 // math.Pi / math.Atan(math.Sqrt(1.0/(7.0+4.0*Rt2)))

	eVal53nRect = 3.0884042547 // math.Pi / math.Atan(P)
	pVal53nRect = 1e100        //∞
)

func GoursatTetrahedron53n(n float64) GoursatTetrahedron {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	cf := (P + 2.0) * cot / (1 + cot)
	ce := P2 * cot
	fe := P2 * (1 + cot) / (P + 2.0)

	var cv, fv, ev float64
	if math.Abs(n-pVal53n) < BoundaryEps {
		cv = 3.0
		fv = 2.0 * (1.0 - Rt_5)
		ev = P_2
	} else {
		cv = P4 * cot / (3.0 - cot)
		fv = P4 * (1.0 + cot) / ((P + 2) * (3.0 - cot))
		ev = P2 / (3.0 - cot)
	}

	return GoursatTetrahedron{
		V:   vector.Vec4{W: 1, X: P, Y: P_1, Z: 0},
		E:   vector.Vec4{W: 1, X: P, Y: 0, Z: 0},
		F:   vector.Vec4{W: 3 - P, X: P, Y: 0, Z: 1},
		C:   vector.Vec4{W: 1, X: 0, Y: 0, Z: 0},
		CFE: vector.Vec4{W: 0, X: 0, Y: 1, Z: 0},
		CFV: vector.Vec4{W: 0, X: 1, Y: -P2, Z: -P},
		CEV: vector.Vec4{W: 0, X: 0, Y: 0, Z: 1},
		FEV: vector.Vec4{W: P2*cot - 1.0, X: P3 * cot, Y: 0, Z: P2 * cot},
		CF:  cf,
		CE:  ce,
		CV:  cv,
		FE:  fe,
		FV:  fv,
		EV:  ev,
	}

}

func Coxeter53n(n float64) Coxeter {

	cos := math.Pow(math.Cos(math.Pi/n), 2.0)

	return Coxeter{
		P: 5.0,
		Q: 3.0,
		R: n,
		A: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: -v.Y, Z: v.Z} },
		B: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: 0.5 * (P*v.X + v.Y + v.Z*P_1), Y: 0.5 * (v.X - v.Y*P_1 - P*v.Z), Z: 0.5 * (v.X*P_1 - P*v.Y + v.Z)}
		},
		C: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: -v.Z} },
		D: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: (2*P*Rt5*cos-1)*v.W - (2*Rt5*cos-2*P_1)*v.X - (2*Rt5*cos*P_1-2*P_2)*v.Z,
				X: 2*P3*cos*v.W + (1-2*P2*cos)*v.X - 2*P*cos*v.Z,
				Y: v.Y,
				Z: 2*P2*cos*v.W - 2*P*cos*v.X + (1-2*cos)*v.Z,
			}
		},
		FaceReflections: []string{
			"", "c", "bc", "acacbabc", "cacbabc",
			"cbabc", "bacbabc", "cbabacbabc", "babacbabc",
			"abc", "acbabc", "abacbabc",
		},
		GoursatTetrahedron: GoursatTetrahedron53n(n),
	}

}

func Honeycomb53n(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	space := Boundaries(n, eVal53n, pVal53n)

	var scale func(vector.Vec4) vector.Vec4
	if space == 'p' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: Rt3 * v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else if space == 'e' {
		// TODO
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: P2 * math.Sqrt(math.Abs(cot/(cot-3.0))) * v.W,
				X: math.Sqrt(math.Abs((P2*cot-1.0)/(cot-3.0))) * v.X,
				Y: math.Sqrt(math.Abs((P2*cot-1.0)/(cot-3.0))) * v.Y,
				Z: math.Sqrt(math.Abs((P2*cot-1.0)/(cot-3.0))) * v.Z,
			}
		}
	}

	var innerProd func(vector.Vec4, vector.Vec4) float64
	if space == 'p' {
		innerProd = func(a, b vector.Vec4) float64 { return 3.0*a.W*b.W - a.X*b.X - a.Y*b.Y - a.Z*b.Z }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (P4*cot*a.W*b.W - (P2*cot-1.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(3.0-cot)
		}
	}

	return Honeycomb{
		Coxeter:  Coxeter53n(n),
		CellType: "spherical",
		Vertices: []vector.Vec4{
			{W: 1, X: 1, Y: 1, Z: 1},
			{W: 1, X: 1, Y: 1, Z: -1},
			{W: 1, X: 1, Y: -1, Z: 1},
			{W: 1, X: 1, Y: -1, Z: -1},
			{W: 1, X: -1, Y: 1, Z: 1},
			{W: 1, X: -1, Y: 1, Z: -1},
			{W: 1, X: -1, Y: -1, Z: 1},
			{W: 1, X: -1, Y: -1, Z: -1},
			{W: 1, X: 0, Y: P, Z: P_1},
			{W: 1, X: 0, Y: P, Z: -P_1},
			{W: 1, X: 0, Y: -P, Z: P_1},
			{W: 1, X: 0, Y: -P, Z: -P_1},
			{W: 1, X: P, Y: P_1, Z: 0},
			{W: 1, X: P, Y: -P_1, Z: 0},
			{W: 1, X: -P, Y: P_1, Z: 0},
			{W: 1, X: -P, Y: -P_1, Z: 0},
			{W: 1, X: P_1, Y: 0, Z: P},
			{W: 1, X: -P_1, Y: 0, Z: P},
			{W: 1, X: P_1, Y: 0, Z: -P},
			{W: 1, X: -P_1, Y: 0, Z: -P},
		},
		Edges: [][2]int{
			{0, 8}, {0, 12}, {0, 16}, {1, 9},
			{1, 12}, {1, 18}, {2, 10}, {2, 13},
			{2, 16}, {3, 11}, {3, 13}, {3, 18},
			{4, 8}, {4, 14}, {4, 17}, {5, 9},
			{5, 14}, {5, 19}, {6, 10}, {6, 15},
			{6, 17}, {7, 11}, {7, 15}, {7, 19},
			{8, 9}, {10, 11}, {12, 13}, {14, 15},
			{16, 17}, {18, 19},
		},
		Faces: [][]int{
			{0, 16, 2, 13, 12}, {1, 12, 13, 3, 18},
			{0, 12, 1, 9, 8}, {0, 8, 4, 17, 16},
			{2, 16, 17, 6, 10}, {1, 18, 19, 5, 9},
			{4, 8, 9, 5, 14}, {5, 19, 7, 15, 14},
			{6, 17, 4, 14, 15}, {3, 13, 2, 10, 11},
			{3, 11, 7, 19, 18}, {11, 10, 6, 15, 7},
		},
		EVal:         eVal53n,
		PVal:         pVal53n,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}

func Honeycomb53nTrunc(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	space := Boundaries(n, eVal53nTrunc, pVal53nTrunc)

	A := (P4 - 1.0) * 0.2
	B := 1.0 + Rt_5
	C := A * P_1
	D := B * P_1
	E := (1 - Rt_5) * 0.5
	F := (1 + Rt_5) * 0.5
	G := P_1 * Rt_5

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
		Coxeter:  Coxeter53n(n),
		CellType: "spherical",
		Vertices: []vector.Vec4{
			{W: 1, X: B, Y: C, Z: -E},
			{W: 1, X: B, Y: C, Z: E},
			{W: 1, X: P, Y: G, Z: 0},
			{W: 1, X: P, Y: -G, Z: 0},
			{W: 1, X: B, Y: -C, Z: -E},
			{W: 1, X: B, Y: -C, Z: E},
			{W: 1, X: F, Y: -A, Z: D},
			{W: 1, X: A, Y: -D, Z: F},
			{W: 1, X: D, Y: -F, Z: A},
			{W: 1, X: C, Y: -E, Z: B},
			{W: 1, X: C, Y: E, Z: B},
			{W: 1, X: G, Y: 0, Z: P},
			{W: 1, X: D, Y: F, Z: A},
			{W: 1, X: A, Y: D, Z: F},
			{W: 1, X: F, Y: A, Z: D},
			{W: 1, X: E, Y: B, Z: C},
			{W: 1, X: 0, Y: P, Z: G},
			{W: 1, X: -E, Y: B, Z: C},
			{W: 1, X: -F, Y: A, Z: D},
			{W: 1, X: -A, Y: D, Z: F},
			{W: 1, X: -D, Y: F, Z: A},
			{W: 1, X: -C, Y: E, Z: B},
			{W: 1, X: -G, Y: 0, Z: P},
			{W: 1, X: -C, Y: -E, Z: B},
			{W: 1, X: -D, Y: -F, Z: A},
			{W: 1, X: -A, Y: -D, Z: F},
			{W: 1, X: -F, Y: -A, Z: D},
			{W: 1, X: -E, Y: -B, Z: C},
			{W: 1, X: E, Y: -B, Z: C},
			{W: 1, X: 0, Y: -P, Z: G},
			{W: 1, X: 0, Y: -P, Z: -G},
			{W: 1, X: E, Y: -B, Z: -C},
			{W: 1, X: -E, Y: -B, Z: -C},
			{W: 1, X: F, Y: -A, Z: -D},
			{W: 1, X: A, Y: -D, Z: -F},
			{W: 1, X: D, Y: -F, Z: -A},
			{W: 1, X: C, Y: -E, Z: -B},
			{W: 1, X: G, Y: 0, Z: -P},
			{W: 1, X: C, Y: E, Z: -B},
			{W: 1, X: D, Y: F, Z: -A},
			{W: 1, X: F, Y: A, Z: -D},
			{W: 1, X: A, Y: D, Z: -F},
			{W: 1, X: E, Y: B, Z: -C},
			{W: 1, X: -E, Y: B, Z: -C},
			{W: 1, X: 0, Y: P, Z: -G},
			{W: 1, X: -B, Y: C, Z: -E},
			{W: 1, X: -B, Y: C, Z: E},
			{W: 1, X: -P, Y: G, Z: 0},
			{W: 1, X: -P, Y: -G, Z: 0},
			{W: 1, X: -B, Y: -C, Z: E},
			{W: 1, X: -B, Y: -C, Z: -E},
			{W: 1, X: -A, Y: -D, Z: -F},
			{W: 1, X: -F, Y: -A, Z: -D},
			{W: 1, X: -D, Y: -F, Z: -A},
			{W: 1, X: -C, Y: -E, Z: -B},
			{W: 1, X: -G, Y: 0, Z: -P},
			{W: 1, X: -C, Y: E, Z: -B},
			{W: 1, X: -D, Y: F, Z: -A},
			{W: 1, X: -F, Y: A, Z: -D},
			{W: 1, X: -A, Y: D, Z: -F},
		},
		Edges: [][2]int{
			{0, 1}, {0, 2}, {0, 41}, {1, 2}, {1, 13},
			{2, 3}, {3, 4}, {3, 5}, {4, 5}, {4, 34},
			{5, 7}, {6, 7}, {6, 8}, {6, 28}, {7, 8},
			{8, 9}, {9, 10}, {9, 11}, {10, 11}, {10, 12},
			{11, 22}, {12, 13}, {12, 14}, {13, 14}, {14, 15},
			{15, 16}, {15, 17}, {16, 17}, {16, 44}, {17, 18},
			{18, 19}, {18, 20}, {19, 20}, {19, 46}, {20, 21},
			{21, 22}, {21, 23}, {22, 23}, {23, 24}, {24, 25},
			{24, 26}, {25, 26}, {25, 49}, {26, 27}, {27, 28},
			{27, 29}, {28, 29}, {29, 30}, {30, 31}, {30, 32},
			{31, 32}, {31, 33}, {32, 52}, {33, 34}, {33, 35},
			{34, 35}, {35, 36}, {36, 37}, {36, 38}, {37, 38},
			{37, 55}, {38, 39}, {39, 40}, {39, 41}, {40, 41},
			{40, 42}, {42, 43}, {42, 44}, {43, 44}, {43, 58},
			{45, 46}, {45, 47}, {45, 59}, {46, 47}, {47, 48},
			{48, 49}, {48, 50}, {49, 50}, {50, 51}, {51, 52},
			{51, 53}, {52, 53}, {53, 54}, {54, 55}, {54, 56},
			{55, 56}, {56, 57}, {57, 58}, {57, 59}, {58, 59},
		},
		Faces: [][]int{
			{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, 10, 11}, {12, 13, 14},
			{15, 16, 17}, {18, 19, 20}, {21, 22, 23}, {24, 25, 26}, {27, 28, 29},
			{30, 31, 32}, {33, 34, 35}, {36, 37, 38}, {39, 40, 41}, {42, 43, 44},
			{45, 46, 47}, {48, 49, 50}, {51, 52, 53}, {54, 55, 56}, {57, 58, 59},
			{0, 1, 13, 14, 15, 16, 44, 42, 40, 41}, {1, 2, 3, 5, 7, 8, 9, 10, 12, 13},
			{0, 2, 3, 4, 34, 35, 36, 38, 39, 41}, {10, 11, 22, 21, 20, 18, 17, 15, 14, 12},
			{4, 5, 7, 6, 28, 29, 30, 31, 33, 34}, {6, 8, 9, 11, 22, 23, 24, 26, 27, 28},
			{25, 26, 27, 29, 30, 32, 52, 51, 50, 49}, {31, 32, 52, 53, 54, 55, 37, 36, 35, 33},
			{37, 38, 39, 40, 42, 43, 58, 57, 56, 55}, {16, 17, 18, 19, 46, 45, 59, 58, 43, 44},
			{19, 20, 21, 23, 24, 25, 49, 48, 47, 46}, {45, 47, 48, 50, 51, 53, 54, 56, 57, 59},
		},
		EVal:         eVal53nTrunc,
		PVal:         pVal53nTrunc,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}

func Honeycomb53nRect(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	space := Boundaries(n, eVal53nRect, pVal53nRect)

	var scale func(vector.Vec4) vector.Vec4
	if space == 'e' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: P2 * math.Sqrt(cot) * v.W,
				X: math.Sqrt(math.Abs(P2*cot-P_2)) * v.X,
				Y: math.Sqrt(math.Abs(P2*cot-P_2)) * v.Y,
				Z: math.Sqrt(math.Abs(P2*cot-P_2)) * v.Z,
			}
		}
	}

	innerProd := func(a, b vector.Vec4) float64 {
		return P4*cot*a.W*b.W - (P2*cot-P_2)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)
	}

	return Honeycomb{
		Coxeter:  Coxeter53n(n),
		CellType: "spherical",
		Vertices: []vector.Vec4{
			{W: 1, X: P * 0.5, Y: 0.5, Z: P2 * 0.5},
			{W: 1, X: P * 0.5, Y: -0.5, Z: P2 * 0.5},
			{W: 1, X: P2 * 0.5, Y: -P * 0.5, Z: 0.5},
			{W: 1, X: P, Y: 0, Z: 0},
			{W: 1, X: P2 * 0.5, Y: P * 0.5, Z: 0.5},
			{W: 1, X: 0, Y: 0, Z: P},
			{W: 1, X: 0.5, Y: -P2 * 0.5, Z: P * 0.5},
			{W: 1, X: P2 * 0.5, Y: -P * 0.5, Z: -0.5},
			{W: 1, X: P2 * 0.5, Y: P * 0.5, Z: -0.5},
			{W: 1, X: 0.5, Y: P2 * 0.5, Z: P * 0.5},
			{W: 1, X: -P * 0.5, Y: -0.5, Z: P2 * 0.5},
			{W: 1, X: -0.5, Y: -P2 * 0.5, Z: P * 0.5},
			{W: 1, X: 0, Y: -P, Z: 0},
			{W: 1, X: 0.5, Y: -P2 * 0.5, Z: -P * 0.5},
			{W: 1, X: P * 0.5, Y: -0.5, Z: -P2 * 0.5},
			{W: 1, X: P * 0.5, Y: 0.5, Z: -P2 * 0.5},
			{W: 1, X: 0.5, Y: P2 * 0.5, Z: -P * 0.5},
			{W: 1, X: 0, Y: P, Z: 0},
			{W: 1, X: -0.5, Y: P2 * 0.5, Z: P * 0.5},
			{W: 1, X: -P * 0.5, Y: 0.5, Z: P2 * 0.5},
			{W: 1, X: -P2 * 0.5, Y: -P * 0.5, Z: 0.5},
			{W: 1, X: -0.5, Y: -P2 * 0.5, Z: -P * 0.5},
			{W: 1, X: 0, Y: 0, Z: -P},
			{W: 1, X: -0.5, Y: P2 * 0.5, Z: -P * 0.5},
			{W: 1, X: -P2 * 0.5, Y: P * 0.5, Z: 0.5},
			{W: 1, X: -P2 * 0.5, Y: -P * 0.5, Z: -0.5},
			{W: 1, X: -P * 0.5, Y: -0.5, Z: -P2 * 0.5},
			{W: 1, X: -P * 0.5, Y: 0.5, Z: -P2 * 0.5},
			{W: 1, X: -P2 * 0.5, Y: P * 0.5, Z: -0.5},
			{W: 1, X: -P, Y: 0, Z: 0},
		},
		Edges: [][2]int{
			{0, 1}, {0, 4}, {0, 5}, {0, 9},
			{1, 5}, {1, 2}, {1, 6}, {2, 6},
			{2, 7}, {2, 3}, {3, 7}, {3, 8},
			{3, 4}, {4, 8}, {4, 9}, {5, 10},
			{5, 19}, {6, 11}, {6, 12}, {7, 13},
			{7, 14}, {8, 15}, {8, 16}, {9, 17},
			{9, 18}, {10, 11}, {10, 20}, {10, 19},
			{11, 20}, {11, 12}, {12, 13}, {12, 21},
			{13, 21}, {13, 14}, {14, 15}, {14, 22},
			{15, 16}, {15, 22}, {16, 17}, {16, 23},
			{17, 18}, {17, 23}, {18, 19}, {18, 24},
			{19, 24}, {20, 29}, {20, 25}, {21, 25},
			{21, 26}, {22, 26}, {22, 27}, {23, 27},
			{23, 28}, {24, 28}, {24, 29}, {25, 26},
			{26, 27}, {27, 28}, {28, 29}, {29, 25},
		},
		Faces: [][]int{
			{0, 1, 2, 3, 4}, {0, 5, 19, 18, 9}, {1, 6, 11, 10, 5}, {2, 7, 13, 12, 6},
			{3, 8, 15, 14, 7}, {4, 9, 17, 16, 8}, {11, 12, 21, 25, 20}, {13, 14, 22, 26, 21},
			{15, 16, 23, 27, 22}, {17, 18, 24, 28, 23}, {19, 10, 20, 29, 24}, {25, 26, 27, 28, 29},
			{0, 1, 5}, {1, 2, 6}, {2, 3, 7}, {3, 4, 8}, {4, 0, 9},
			{5, 10, 19}, {6, 11, 12}, {7, 13, 14}, {8, 15, 16}, {9, 17, 18},
			{10, 11, 20}, {12, 13, 21}, {14, 15, 22}, {16, 17, 23}, {18, 19, 24},
			{20, 25, 29}, {21, 25, 26}, {22, 26, 27}, {23, 27, 28}, {24, 28, 29},
		},
		EVal:         eVal53nRect,
		PVal:         pVal53nRect,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}
