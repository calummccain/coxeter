package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

const (
	eVal35n = 2.6051148443 //math.Pi / math.Atan(phi^2)
	pVal35n = 3.3333333333

	eVal35nTrunc = 2.6051148443 //math.Pi / math.Atan(phi^2)
	pVal35nTrunc = 6.4642820231 //math.Pi / math.Atan(phi^2/sqrt(9phi^2+1))

	eVal35nRect = 2.6051148443 //math.Pi / math.Atan(phi^2)
	pVal35nRect = 1e100        //∞
)

func GoursatTetrahedron35n(n float64) GoursatTetrahedron {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	cf := 3.0 * P2 * cot / (1 + cot)
	ce := P4 * cot
	fe := P2 * (1 + cot) / 3.0

	var cv, fv, ev float64
	if math.Abs(n-pVal35n) < BoundaryEps {
		cv = 0.0
		fv = 0.0
		ev = 0.0
	} else {
		cv = P6 * cot / (1.0 + P2 - P4*cot)
		fv = P4 * (1.0 + cot) / (3.0 * (1.0 + P2 - P4*cot))
		ev = P2 / (1.0 + P2 - P4*cot)
	}

	return GoursatTetrahedron{
		V:   vector.Vec4{W: 1, X: 1, Y: P, Z: 0},
		E:   vector.Vec4{W: 1, X: 0, Y: P, Z: 0},
		F:   vector.Vec4{W: 3, X: 0, Y: P3, Z: P},
		C:   vector.Vec4{W: 1, X: 0, Y: 0, Z: 0},
		CFE: vector.Vec4{W: 0, X: 1, Y: 0, Z: 0},
		CFV: vector.Vec4{W: 0, X: P, Y: -1, Z: P2},
		CEV: vector.Vec4{W: 0, X: 0, Y: 0, Z: 1},
		FEV: vector.Vec4{W: P*cot - P_3, X: 0, Y: P2 * cot, Z: cot},
		CF:  cf,
		CE:  ce,
		CV:  cv,
		FE:  fe,
		FV:  fv,
		EV:  ev,
	}

}

func Coxeter35n(n float64) Coxeter {

	cos := math.Pow(math.Cos(math.Pi/n), 2.0)
	// sin := 1.0 - cos

	return Coxeter{
		P: 3.0,
		Q: 5.0,
		R: n,
		A: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: -v.X, Y: v.Y, Z: v.Z} },
		B: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: 0.5 * (v.X + P_1*v.Y - P*v.Z), Y: 0.5 * (P_1*v.X + P*v.Y + v.Z), Z: 0.5 * (-P*v.X + v.Y - P_1*v.Z)}
		},
		C: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: -v.Z} },
		D: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: (6.0*P2*cos-1.0)*v.W + (2.0*P_1-6.0*P*cos)*v.Y + (2.0*P_3-6.0*cos*P_1)*v.Z,
				X: v.X,
				Y: 2.0*P5*cos*v.W + (1.0-2.0*P4*cos)*v.Y - 2.0*P2*cos*v.Z,
				Z: 2.0*P3*cos*v.W - 2.0*P2*cos*v.Y + (1.0-2.0*cos)*v.Z,
			}
		},
		FaceReflections: []string{
			"", "c", "bcbc", "bc", "cbc", "bacbcabacbc",
			"cbacbcabacbc", "bacabacbc", "bcabacbc",
			"cbcabacbc", "abcbc", "abc", "acbc",
			"abacabacbc", "abcabacbc", "cabcabacbc",
			"bacbc", "abacbc", "acabacbc", "cabacbc",
		},
		GoursatTetrahedron: GoursatTetrahedron35n(n),
	}

}

func Honeycomb35n(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	space := Boundaries(n, eVal35n, pVal35n)

	var scale func(vector.Vec4) vector.Vec4
	if space == 'p' {
		//TODO
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
				W: P3 * math.Sqrt(math.Abs(cot/(P2+1.0-P4*cot))) * v.W,
				X: math.Sqrt(math.Abs((P4*cot-1.0)/(P2+1.0-P4*cot))) * v.X,
				Y: math.Sqrt(math.Abs((P4*cot-1.0)/(P2+1.0-P4*cot))) * v.Y,
				Z: math.Sqrt(math.Abs((P4*cot-1.0)/(P2+1.0-P4*cot))) * v.Z,
			}
		}
	}

	var innerProd func(vector.Vec4, vector.Vec4) float64
	if space == 'p' {
		// TODO
		innerProd = func(a, b vector.Vec4) float64 { return a.W*b.W - a.X*b.X - a.Y*b.Y - a.Z*b.Z }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (P6*cot*a.W*b.W - (P4*cot-1.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(P2+1.0-P4*cot)
		}
	}

	return Honeycomb{
		Coxeter:  Coxeter35n(n),
		CellType: "spherical",
		Vertices: []vector.Vec4{
			{W: 1, X: 1, Y: P, Z: 0},
			{W: 1, X: 1, Y: -P, Z: 0},
			{W: 1, X: -1, Y: P, Z: 0},
			{W: 1, X: -1, Y: -P, Z: 0},
			{W: 1, X: P, Y: 0, Z: 1},
			{W: 1, X: -P, Y: 0, Z: 1},
			{W: 1, X: P, Y: 0, Z: -1},
			{W: 1, X: -P, Y: 0, Z: -1},
			{W: 1, X: 0, Y: 1, Z: P},
			{W: 1, X: 0, Y: 1, Z: -P},
			{W: 1, X: 0, Y: -1, Z: P},
			{W: 1, X: 0, Y: -1, Z: -P},
		},
		Edges: [][2]int{
			{0, 2}, {0, 4}, {0, 6}, {0, 8},
			{0, 9}, {1, 3}, {1, 4}, {1, 6},
			{1, 10}, {1, 11}, {2, 5}, {2, 7},
			{2, 8}, {2, 9}, {3, 5}, {3, 7},
			{3, 10}, {3, 11}, {4, 6}, {4, 8},
			{4, 10}, {5, 7}, {5, 8}, {5, 10},
			{6, 9}, {6, 11}, {7, 9}, {7, 11},
			{8, 10}, {9, 11},
		},
		Faces: [][]int{
			{0, 8, 2}, {0, 2, 9}, {0, 6, 4}, {0, 4, 8},
			{0, 9, 6}, {1, 3, 10}, {1, 11, 3}, {1, 4, 6},
			{1, 10, 4}, {1, 6, 11}, {2, 5, 7}, {2, 8, 5},
			{2, 7, 9}, {3, 7, 5}, {3, 5, 10}, {3, 11, 7},
			{4, 10, 8}, {5, 8, 10}, {6, 9, 11}, {7, 11, 9},
		},
		EVal:         eVal35n,
		PVal:         pVal35n,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}

func Honeycomb35nTrunc(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	A := P / 3.0
	B := (P + 2.0) / 3.0
	C := 1.0 / 3.0
	D := P3 / 3.0
	E := 2.0 * P / 3.0
	F := 2.0 / 3.0

	space := Boundaries(n, eVal35nTrunc, pVal35nTrunc)

	var scale func(vector.Vec4) vector.Vec4
	if space == 'p' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: math.Sqrt(9.0*P2+1.0) * v.W, X: 3.0 * v.X, Y: 3.0 * v.Y, Z: 3.0 * v.Z}
		}
	} else if space == 'e' {
		// TODO
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: 3.0 * P3 * math.Sqrt(math.Abs(cot/(9*P2+1.0-P4*cot))) * v.W,
				X: 3.0 * math.Sqrt(math.Abs((P4*cot-1.0)/(9*P2+1.0-P4*cot))) * v.X,
				Y: 3.0 * math.Sqrt(math.Abs((P4*cot-1.0)/(9*P2+1.0-P4*cot))) * v.Y,
				Z: 3.0 * math.Sqrt(math.Abs((P4*cot-1.0)/(9*P2+1.0-P4*cot))) * v.Z,
			}
		}
	}

	var innerProd func(vector.Vec4, vector.Vec4) float64
	if space == 'p' {
		innerProd = func(a, b vector.Vec4) float64 { return (9.0*P2+1.0)*a.W*b.W - 9.0*(a.X*b.X+a.Y*b.Y+a.Z*b.Z) }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (9.0*P6*cot*a.W*b.W - 9.0*(P4*cot-1.0)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)) / math.Abs(9*P2+1.0-P4*cot)
		}
	}

	return Honeycomb{
		Coxeter:  Coxeter35n(n),
		CellType: "spherical",
		Vertices: []vector.Vec4{
			{W: 1, X: F, Y: D, Z: A},
			{W: 1, X: B, Y: E, Z: C},
			{W: 1, X: B, Y: E, Z: -C},
			{W: 1, X: F, Y: D, Z: -A},
			{W: 1, X: C, Y: P, Z: 0},
			{W: 1, X: C, Y: B, Z: E},
			{W: 1, X: A, Y: F, Z: D},
			{W: 1, X: 0, Y: C, Z: P},
			{W: 1, X: -A, Y: F, Z: D},
			{W: 1, X: -C, Y: B, Z: E},
			{W: 1, X: -F, Y: D, Z: A},
			{W: 1, X: -C, Y: P, Z: 0},
			{W: 1, X: -F, Y: D, Z: -A},
			{W: 1, X: -B, Y: E, Z: -C},
			{W: 1, X: -B, Y: E, Z: C},
			{W: 1, X: -D, Y: A, Z: F},
			{W: 1, X: -E, Y: C, Z: B},
			{W: 1, X: -E, Y: -C, Z: B},
			{W: 1, X: -D, Y: -A, Z: F},
			{W: 1, X: -P, Y: 0, Z: C},
			{W: 1, X: -C, Y: B, Z: -E},
			{W: 1, X: C, Y: B, Z: -E},
			{W: 1, X: A, Y: F, Z: -D},
			{W: 1, X: 0, Y: C, Z: -P},
			{W: 1, X: -A, Y: F, Z: -D},
			{W: 1, X: D, Y: A, Z: -F},
			{W: 1, X: E, Y: C, Z: -B},
			{W: 1, X: E, Y: -C, Z: -B},
			{W: 1, X: D, Y: -A, Z: -F},
			{W: 1, X: P, Y: 0, Z: -C},
			{W: 1, X: D, Y: A, Z: F},
			{W: 1, X: P, Y: 0, Z: C},
			{W: 1, X: D, Y: -A, Z: F},
			{W: 1, X: E, Y: -C, Z: B},
			{W: 1, X: E, Y: C, Z: B},
			{W: 1, X: A, Y: -F, Z: D},
			{W: 1, X: 0, Y: -C, Z: P},
			{W: 1, X: -A, Y: -F, Z: D},
			{W: 1, X: -C, Y: -B, Z: E},
			{W: 1, X: C, Y: -B, Z: E},
			{W: 1, X: -D, Y: -A, Z: -F},
			{W: 1, X: -P, Y: 0, Z: -C},
			{W: 1, X: -D, Y: A, Z: -F},
			{W: 1, X: -E, Y: C, Z: -B},
			{W: 1, X: -E, Y: -C, Z: -B},
			{W: 1, X: -F, Y: -D, Z: -A},
			{W: 1, X: -B, Y: -E, Z: -C},
			{W: 1, X: -B, Y: -E, Z: C},
			{W: 1, X: -F, Y: -D, Z: A},
			{W: 1, X: -C, Y: -P, Z: 0},
			{W: 1, X: C, Y: -P, Z: 0},
			{W: 1, X: F, Y: -D, Z: A},
			{W: 1, X: B, Y: -E, Z: C},
			{W: 1, X: B, Y: -E, Z: -C},
			{W: 1, X: F, Y: -D, Z: -A},
			{W: 1, X: -C, Y: -B, Z: -E},
			{W: 1, X: -A, Y: -F, Z: -D},
			{W: 1, X: 0, Y: -C, Z: -P},
			{W: 1, X: A, Y: -F, Z: -D},
			{W: 1, X: C, Y: -B, Z: -E},
		},
		Edges: [][2]int{
			{0, 1}, {0, 4}, {0, 5}, {1, 2}, {1, 30},
			{2, 3}, {2, 25}, {3, 4}, {3, 21}, {4, 11},
			{5, 6}, {5, 9}, {6, 7}, {6, 34}, {7, 8},
			{7, 36}, {8, 9}, {8, 16}, {9, 10}, {10, 11},
			{10, 14}, {11, 12}, {12, 13}, {12, 20}, {13, 14},
			{13, 42}, {14, 15}, {15, 16}, {15, 19}, {16, 17},
			{17, 18}, {17, 37}, {18, 19}, {18, 47}, {19, 41},
			{20, 21}, {20, 24}, {21, 22}, {22, 23}, {22, 26},
			{23, 24}, {23, 57}, {24, 43}, {25, 26}, {25, 29},
			{26, 27}, {27, 28}, {27, 58}, {28, 29}, {28, 53},
			{29, 31}, {30, 31}, {30, 34}, {31, 32}, {32, 33},
			{32, 52}, {33, 34}, {33, 35}, {35, 36}, {35, 39},
			{36, 37}, {37, 38}, {38, 39}, {38, 48}, {39, 51},
			{40, 41}, {40, 44}, {40, 46}, {41, 42}, {42, 43},
			{43, 44}, {44, 56}, {45, 46}, {45, 49}, {45, 55},
			{46, 47}, {47, 48}, {48, 49}, {49, 50}, {50, 51},
			{50, 54}, {51, 52}, {52, 53}, {53, 54}, {54, 59},
			{55, 56}, {55, 59}, {56, 57}, {57, 58}, {58, 59},
		},
		Faces: [][]int{
			{0, 1, 2, 3, 4}, {5, 6, 7, 8, 9}, {10, 11, 12, 13, 14}, {15, 16, 17, 18, 19},
			{20, 21, 22, 23, 24}, {25, 26, 27, 28, 29}, {30, 31, 32, 33, 34}, {35, 36, 37, 38, 39},
			{40, 41, 42, 43, 44}, {45, 46, 47, 48, 49}, {50, 51, 52, 53, 54}, {55, 56, 57, 58, 59},
			{0, 4, 11, 10, 9, 5}, {8, 9, 10, 14, 15, 16}, {7, 8, 16, 17, 37, 36},
			{6, 7, 36, 35, 33, 34}, {0, 1, 30, 34, 6, 5}, {3, 4, 11, 12, 20, 21},
			{12, 13, 42, 43, 24, 20}, {13, 14, 15, 19, 41, 42}, {1, 2, 25, 29, 31, 30},
			{18, 19, 41, 40, 46, 47}, {17, 18, 47, 48, 38, 37}, {38, 39, 51, 50, 49, 48},
			{32, 33, 35, 39, 51, 52}, {28, 29, 31, 32, 52, 53}, {27, 28, 53, 54, 59, 58},
			{2, 3, 21, 22, 26, 25}, {22, 23, 57, 58, 27, 26}, {23, 24, 43, 44, 56, 57},
			{40, 44, 56, 55, 45, 46}, {45, 49, 50, 54, 59, 55},
		},
		EVal:         eVal35nTrunc,
		PVal:         pVal35nTrunc,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}

func Honeycomb35nRect(n float64) Honeycomb {

	cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)

	space := Boundaries(n, eVal35nRect, pVal35nRect)

	var scale func(vector.Vec4) vector.Vec4
	if space == 'e' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: P2 * math.Sqrt(math.Abs(cot)) * v.W,
				X: math.Sqrt(math.Abs((P4*cot-1.0)/P2)) * v.X,
				Y: math.Sqrt(math.Abs((P4*cot-1.0)/P2)) * v.Y,
				Z: math.Sqrt(math.Abs((P4*cot-1.0)/P2)) * v.Z,
			}
		}
	}

	innerProd := func(a, b vector.Vec4) float64 {
		return P4*cot*a.W*b.W - (P2*cot-P_2)*(a.X*b.X+a.Y*b.Y+a.Z*b.Z)
	}

	return Honeycomb{
		Coxeter:  Coxeter35n(n),
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
		EVal:         eVal35nRect,
		PVal:         pVal35nRect,
		Space:        space,
		Scale:        scale,
		InnerProduct: innerProd,
	}
}
