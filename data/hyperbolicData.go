package data

import (
	"coxeter/hyperbolic"
	"coxeter/shared"
	"coxeter/tesselations"
	"coxeter/vector"
	"math"
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

	cMat := func(v [4]float64) [4]float64 {

		r := sr*sr*sp*sp - cq*cq

		return [4]float64{
			(1.0-2.0*r/(sp*sp))*v[0] + (2.0*r*cr/(sp*cp*cq))*v[1] + (2.0*r/(cp*sp*sp))*v[2],
			(2.0*cp*cq*cr/sp)*v[0] + (1.0-2.0*cr*cr)*v[1] - (2.0*cr*cq/sp)*v[2],
			(2.0*cp*cq*cq/(sp*sp))*v[0] - (2.0*cr*cq/sp)*v[1] + (1.0-2.0*cq*cq/(sp*sp))*v[2],
			v[3],
		}

	}

	var f func([4]float64) [4]float64
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

	f = func(v [4]float64) [4]float64 {

		return [4]float64{a * v[0], b * v[1], b * v[2], b * v[3]}

	}

	matrices := shared.Matrices{
		A: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[1], v[2], -v[3]} },
		B: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[1], cp2*v[2] + sp2*v[3], sp2*v[2] - cp2*v[3]} },
		C: cMat,
		D: func(v [4]float64) [4]float64 { return [4]float64{v[0], -v[1], v[2], v[3]} },
		E: func(v [4]float64) [4]float64 { return v },
		F: f,
	}

	initialVerts := [][4]float64{}
	initialEdges := [][4]float64{}
	for i := 0; i < p; i++ {
		initialVerts = append(initialVerts, [4]float64{1.0, 0.0, math.Cos(math.Pi * float64(2*i+1) / float64(p)), math.Sin(math.Pi * float64(2*i+1) / float64(p))})
		initialEdges = append(initialEdges, [4]float64{1.0, 0.0, cp * math.Cos(math.Pi*float64(2*i)/float64(p)), cp * math.Sin(math.Pi*float64(2*i)/float64(p))})
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

	fPoints, fNames := tesselations.MakeFaces([4]float64{fVal, 0, 0, 0}, numberOfFaces, p, matrices)

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
		V:               [4]float64{0, 0, 0, 0},
		E:               [4]float64{0, 0, 0, 0},
		F:               [4]float64{0, 0, 0, 0},
		C:               [4]float64{0, 0, 0, 0},
		CellType:        "hyperbolic",
		Vv:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices:        v,
		Edges:           edgeData,
		Faces:           faceData,
		Matrices:        matrices,
		Flip:            func(v [4]float64) [4]float64 { return [4]float64{v[0], v[2], v[3], v[1]} },
	}
}
