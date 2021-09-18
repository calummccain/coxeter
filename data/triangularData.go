package data

import (
	"coxeter/hyperbolic"
	"coxeter/shared"
	"coxeter/tesselations"
	"coxeter/vector"
	"math"
)

func TriangularData(n float64, numberOfFaces int) CellData {

	eVal := 2.0
	pVal := 3.0

	cos := math.Pow(math.Cos(math.Pi/n), 2)

	den := math.Sqrt(math.Abs(1.0 - 4.0*cos))

	metric := Boundaries(n, eVal, pVal)

	var vv float64

	if metric == 'p' {

		vv = 1.5

	} else {

		vv = (1.0 + 2.0*cos) / math.Abs(1.0-4.0*cos)

	}

	cMat := func(v [4]float64) [4]float64 {

		return [4]float64{
			(1.0+2.0*cos)*v[0] - 2.0*(cos*cos)*v[1] - cos*v[2] - 3.0*cos*v[3],
			2.0*v[0] + (1.0-2.0*cos)*v[1] - v[2] - 3.0*v[3],
			v[0] - cos*v[1] + 0.5*v[2] - 1.5*v[3],
			v[0] - cos*v[1] - 0.5*v[2] - 0.5*v[3],
		}

	}

	var f func([4]float64) [4]float64
	var a, b, c, d float64

	if metric == 'p' {

		a = 1.0
		b = 0.25
		c = 0.5
		d = 0.5 * Rt3

	} else {

		a = 1.0 / den
		b = cos / den
		c = math.Sqrt(cos) / den
		d = Rt3 * c

	}

	f = func(v [4]float64) [4]float64 {

		return [4]float64{a * v[0], b * v[1], c * v[2], d * v[3]}

	}

	matrices := shared.Matrices{
		A: func(v [4]float64) [4]float64 {
			return [4]float64{v[0], v[1], -0.5*v[2] + 1.5*v[3], 0.5 * (v[2] + v[3])}
		},
		B: func(v [4]float64) [4]float64 { return [4]float64{v[0], v[1], v[2], -v[3]} },
		C: cMat,
		D: func(v [4]float64) [4]float64 { return [4]float64{v[0], -v[1], v[2], v[3]} },
		E: func(v [4]float64) [4]float64 { return v },
		F: f,
	}

	initialVerts := [][4]float64{
		{1, 0, 2, 0},
		{1, 0, -1, -1},
		{1, 0, -1, 1},
	}

	initialEdges := [][4]float64{
		vector.Scale4([4]float64{2, 0, 1, -1}, 1.0/math.Sqrt(math.Abs(hyperbolic.HyperbolicNorm(f([4]float64{2, 0, 1, -1}))))),
		vector.Scale4([4]float64{2, 0, 1, 1}, 1.0/math.Sqrt(math.Abs(hyperbolic.HyperbolicNorm(f([4]float64{2, 0, 1, 1}))))),
		vector.Scale4([4]float64{1, 0, -1, 0}, 1.0/math.Sqrt(math.Abs(hyperbolic.HyperbolicNorm(f([4]float64{1, 0, -1, 0}))))),
	}

	fVal := 0.0
	ev := 0.0
	fv := 0.0
	if metric == 'p' {

		fVal = 1.0
		ev = 0.75
		fv = 1.0

	} else {

		fVal = den
		ev = (1.0 - cos) / (1.0 - 4.0*cos)
		fv = 1.0 / (1.0 - 4.0*cos)

	}

	fPoints, fNames := tesselations.MakeFaces([4]float64{fVal, 0, 0, 0}, numberOfFaces, 3, matrices)

	v := tesselations.MakeRing(initialVerts, matrices, fNames)
	e := tesselations.MakeRing(initialEdges, matrices, fNames)

	faceData := tesselations.GenerateFaceData(math.Abs(fv), 3, fPoints, v, f)

	edgeData := tesselations.GenerateEdgeData(math.Abs(ev), e, v, f)
	faceData = tesselations.OrderFaces(3, faceData, edgeData)

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
		CellType:        "euclidean",
		Vv:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices:        v,
		Edges:           edgeData,
		Faces:           faceData,
		Matrices:        matrices,
		Flip:            func(v [4]float64) [4]float64 { return [4]float64{v[0], v[2], v[3], v[1]} },
	}
}
