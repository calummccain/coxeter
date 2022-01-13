package data

import (
	"math"

	"github.com/calummccain/coxeter/hyperbolic"
	"github.com/calummccain/coxeter/shared"
	"github.com/calummccain/coxeter/tesselations"
	"github.com/calummccain/coxeter/vector"
)

func HyperbolicData(p, q int, r float64, numberOfFaces int) CellData {

	cp := math.Cos(math.Pi / float64(p))
	sp := math.Sin(math.Pi / float64(p))

	cq := math.Cos(math.Pi / float64(q))
	// sq := math.Sin(math.Pi / float64(q))

	cr := math.Cos(math.Pi / r)
	sr := math.Sin(math.Pi / r)

	cp2 := math.Cos(2.0 * math.Pi / float64(p))
	sp2 := math.Sin(2.0 * math.Pi / float64(p))

	den := sp * math.Sqrt(math.Abs(sr*sr-cq*cq))

	var eVal float64
	if (p-2)*(q-2) < 4 {
		eVal = math.Pi / math.Asin(cq/sp)
	} else {
		eVal = 2.0
	}

	pVal := 2.0 * float64(q) / (float64(q) - 2.0)

	metric := Boundaries(r, eVal, pVal)

	cMat := func(v vector.Vec4) vector.Vec4 {

		r := sr*sr*sp*sp - cq*cq

		return vector.Vec4{
			(1.0-2.0*r/(sp*sp))*v.W + (2.0*r*cr/(sp*cp*cq))*v.X + (2.0*r/(cp*sp*sp))*v.Y,
			(2.0*cp*cq*cr/sp)*v.W + (1.0-2.0*cr*cr)*v.X - (2.0*cr*cq/sp)*v.Y,
			(2.0*cp*cq*cq/(sp*sp))*v.W - (2.0*cr*cq/sp)*v.X + (1.0-2.0*cq*cq/(sp*sp))*v.Y,
			v.Z,
		}

	}

	var f func(vector.Vec4) vector.Vec4
	var a, b float64

	if metric == 'e' {

		a = 1.0
		b = 1.0

	} else if metric == 'p' {

		a = 1.0
		b = 1.0

	} else {

		a = cp * cq / den
		b = math.Sqrt(math.Abs(sp*sp*sr*sr-cq*cq)) / den
	}

	f = func(v vector.Vec4) vector.Vec4 {

		return vector.Vec4{a * v.W, b * v.X, b * v.Y, b * v.Z}

	}

	matrices := shared.Matrices{
		Amat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Y, -v.Z} },
		Bmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, cp2*v.Y + sp2*v.Z, sp2*v.Y - cp2*v.Z} },
		C: cMat,
		Dmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, -v.X, v.Y, v.Z} },
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}

	initialVerts := []vector.Vec4{}
	initialEdges := []vector.Vec4{}
	for i := 0; i < p; i++ {
		initialVerts = append(initialVerts, vector.Vec4{1.0, 0.0, math.Cos(math.Pi * float64(2*i+1) / float64(p)), math.Sin(math.Pi * float64(2*i+1) / float64(p))})
		initialEdges = append(initialEdges, vector.Vec4{1.0, 0.0, cp * math.Cos(math.Pi*float64(2*i)/float64(p)), cp * math.Sin(math.Pi*float64(2*i)/float64(p))})
		initialEdges[i] = vector.Scale4(initialEdges[i], 1.0/math.Sqrt(math.Abs(hyperbolic.HyperbolicNorm(f(initialEdges[i])))))
	}

	fVal := 0.0
	vv := 0.0
	ev := 0.0
	fv := 0.0

	if metric == 'p' {

		fVal = 1.0
		vv = 2.0 * sp * sp
		ev = sp * sp
		fv = 1.0

	} else {

		fVal = sp * math.Sqrt(math.Abs(sr*sr-cq*cq)) / (cp * cq)
		vv = ((2.0*cp*cp-1)*sr*sr + cq*cq) / math.Abs(sr*sr-cq*cq)
		ev = sr * sr * cp * cp / (sr*sr - cq*cq)
		fv = cp * cp * cq * cq / (sp * sp * (sr*sr - cq*cq))

	}

	fPoints, fNames := tesselations.MakeFaces(vector.Vec4{fVal, 0, 0, 0}, numberOfFaces, p, matrices)

	v := tesselations.MakeRing(initialVerts, matrices, fNames)

	e := tesselations.MakeRing(initialEdges, matrices, fNames)

	faceData := tesselations.GenerateFaceData(math.Abs(fv), p, fPoints, v, f)

	edgeData := tesselations.GenerateEdgeData(math.Abs(ev), e, v, f)

	faceData = tesselations.OrderFaces(p, faceData, edgeData)

	return CellData{
		Metric:          metric,
		NumVertices:     len(v),
		NumEdges:        len(edgeData),
		NumFaces:        len(faceData),
		FaceReflections: fNames,
		OuterReflection: "d",
		V:               vector.Vec4{0, 0, 0, 0},
		E:               vector.Vec4{0, 0, 0, 0},
		F:               vector.Vec4{0, 0, 0, 0},
		C:               vector.Vec4{0, 0, 0, 0},
		CellType:        "hyperbolic",
		VV:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices:        v,
		Edges:           edgeData,
		Faces:           faceData,
		Matrices:        matrices,
		Flip:            func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.Y, v.Z, v.X} },
	}
}
