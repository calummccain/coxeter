package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func eValpqr(p, q, r float64) float64 {

	sp := math.Sin(math.Pi / p)
	cq := math.Cos(math.Pi / q)

	eVal := 0.0

	if (p-2.0)*(q-2.0) < 4.0 {

		eVal = math.Pi / math.Asin(cq/sp)

	} else {

		eVal = 2.0

	}

	return eVal

}

func pValpqr(p, q, r float64) float64 {

	return 1.0 / (0.5 - 1.0/q)

}

func GoursatTetrahedronpqr(p, q, r float64) GoursatTetrahedron {

	cp := math.Cos(math.Pi / p)
	sp := math.Sin(math.Pi / p)

	cq := math.Cos(math.Pi / q)
	//sq := math.Sin(math.Pi / float64(q))

	cr := math.Cos(math.Pi / r)
	sr := math.Sin(math.Pi / r)

	//cot := 1.0 / math.Pow(math.Tan(math.Pi/n), 2.0)
	//cos := math.Pow(math.Cos(math.Pi/n), 2.0)

	//fe := 1.0 / (1.0 - cos)
	var fe, cf, cv, fv, ev, ce float64
	if math.Abs(r-pValpqr(p, q, r)) < BoundaryEps {
		ev = sp * sp
		fv = 1.0
	} else {
		ev = sr * sr * cp * cp / (sr*sr - cq*cq)
		fv = cp * cp * cq * cq / (sp * sp * (sr*sr - cq*cq))
	}

	return GoursatTetrahedron{
		V:   vector.Vec4{W: 1, X: 0, Y: cp, Z: sp},
		E:   vector.Vec4{W: 1, X: 0, Y: cp, Z: 0},
		F:   vector.Vec4{W: 1, X: 0, Y: 0, Z: 0},
		C:   vector.Vec4{W: 1, X: cp * cq / (sp * cr), Y: 0, Z: 0},
		CFE: vector.Vec4{W: 0, X: 0, Y: 0, Z: 0},
		CFV: vector.Vec4{W: 0, X: 0, Y: 0, Z: 0},
		CEV: vector.Vec4{W: 0, X: 0, Y: 0, Z: 0},
		FEV: vector.Vec4{W: 0, X: 0, Y: 0, Z: 0},
		CF:  cf,
		CE:  ce,
		CV:  cv,
		FE:  fe,
		FV:  fv,
		EV:  ev,
	}

}

func Coxeterpqr(p, q, r float64, numberOfFaces int) (Coxeter, []vector.Vec4) {

	//cos := math.Pow(math.Cos(math.Pi/n), 2.0)
	//sin := 1.0 - cos

	cp := math.Cos(math.Pi / p)
	sp := math.Sin(math.Pi / p)

	cq := math.Cos(math.Pi / q)
	//sq := math.Sin(math.Pi / float64(q))

	cr := math.Cos(math.Pi / r)
	sr := math.Sin(math.Pi / r)

	cp2 := math.Cos(2.0 * math.Pi / p)
	sp2 := math.Sin(2.0 * math.Pi / p)

	m := sr*sr*sp*sp - cq*cq

	coxeter := Coxeter{
		P: p,
		Q: q,
		R: r,
		A: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: -v.Z}
		},
		B: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: cp2*v.Y + sp2*v.Z, Z: sp2*v.Y - cp2*v.Z}
		},
		C: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: (1.0-2.0*m/(sp*sp))*v.W + (2.0*m*cr/(sp*cp*cq))*v.X + (2.0*m/(cp*sp*sp))*v.Y,
				X: (2.0*cp*cq*cr/sp)*v.W + (1.0-2.0*cr*cr)*v.X - (2.0*cr*cq/sp)*v.Y,
				Y: (2.0*cp*cq*cq/(sp*sp))*v.W - (2.0*cr*cq/sp)*v.X + (1.0-2.0*cq*cq/(sp*sp))*v.Y,
				Z: v.Z,
			}
		},
		D: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: -v.X, Y: v.Y, Z: v.Z}
		},
		FaceReflections:    []string{},
		GoursatTetrahedron: GoursatTetrahedronpqr(p, q, r),
	}

	faceNames, facePoints := coxeter.MakeFaces(numberOfFaces)

	coxeter.FaceReflections = faceNames

	return coxeter, facePoints

}

func Honeycombpqr(p, q, r float64, numberOfFaces int) Honeycomb {

	coxeter, facePoints := Coxeterpqr(p, q, r, numberOfFaces)

	cp := math.Cos(math.Pi / float64(p))
	sp := math.Sin(math.Pi / float64(p))

	cq := math.Cos(math.Pi / float64(q))
	//sq := math.Sin(math.Pi / float64(q))

	//cr := math.Cos(math.Pi / r)
	sr := math.Sin(math.Pi / r)

	den := sp * math.Sqrt(math.Abs(sr*sr-cq*cq))

	space := Boundaries(r, eValpqr(p, q, r), pValpqr(p, q, r))

	var scale func(vector.Vec4) vector.Vec4
	if space == 'p' {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: cp * cq * v.W / den,
				X: math.Sqrt(math.Abs(sp*sp*sr*sr-cq*cq)) * v.X / den,
				Y: math.Sqrt(math.Abs(sp*sp*sr*sr-cq*cq)) * v.Y / den,
				Z: math.Sqrt(math.Abs(sp*sp*sr*sr-cq*cq)) * v.Z / den,
			}
		}
	}

	var innerProd func(vector.Vec4, vector.Vec4) float64
	if space == 'p' {
		innerProd = func(a, b vector.Vec4) float64 { return a.W*b.W - a.X*b.X - a.Y*b.Y - a.Z*b.Z }
	} else {
		innerProd = func(a, b vector.Vec4) float64 {
			return (cp*cp*cq*cq*a.W*b.W - (sp*sp*sr*sr-cq*cq)*(a.X*b.X-a.Y*b.Y-a.Z*b.Z)) / math.Abs(sr*sr-cq*cq)
		}
	}

	initialVerts := []vector.Vec4{}
	initialEdges := []vector.Vec4{}
	for i := 0; i < int(p); i++ {
		initialVerts = append(initialVerts, vector.Vec4{
			W: 1.0,
			X: 0.0,
			Y: math.Cos(math.Pi * float64(2*i+1) / float64(p)),
			Z: math.Sin(math.Pi * float64(2*i+1) / float64(p)),
		})
		initialEdges = append(initialEdges, vector.Vec4{
			W: 1.0,
			X: 0.0,
			Y: cp * math.Cos(math.Pi*float64(2*i)/float64(p)),
			Z: cp * math.Sin(math.Pi*float64(2*i)/float64(p)),
		})
		initialEdges[i] = vector.Scale4(initialEdges[i], 1.0/(sp*cp*sr))
	}

	honeycomb := Honeycomb{
		Coxeter:      coxeter,
		CellType:     "hyperbolic",
		Vertices:     []vector.Vec4{},
		Edges:        [][2]int{},
		Faces:        [][]int{},
		EVal:         eValpqr(p, q, r),
		PVal:         pValpqr(p, q, r),
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
func HoneycombpqrTrunc(n float64, numberOfFaces int) Honeycomb {

	return Honeycomb{}

}

// TODO
func HoneycombpqrRect(n float64, numberOfFaces int) Honeycomb {

	return Honeycomb{}

}
