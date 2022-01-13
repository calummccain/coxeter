package data

import (
	"math"

	"github.com/calummccain/coxeter/hyperbolic"
	"github.com/calummccain/coxeter/shared"
	"github.com/calummccain/coxeter/tesselations"
	"github.com/calummccain/coxeter/vector"
)

func SquareData(n float64, numberOfFaces int) CellData {

	eVal := 2.0
	pVal := 4.0

	cos := math.Pow(math.Cos(math.Pi/n), 2)

	den := math.Sqrt(math.Abs(1.0 - 2.0*cos))

	metric := Boundaries(n, eVal, pVal)

	var vv float64

	if metric == 'p' {

		vv = 1.0

	} else {

		vv = 1.0 / math.Abs(1.0-2.0*cos)

	}

	cMat := func(v vector.Vec4) vector.Vec4 {

		return vector.Vec4{
			(1.0+2.0*cos)*v.W - 2.0*(cos*cos)*v.X - 2.0*cos*v.Y - 2.0*cos*v.Z,
			2.0*v.W + (1.0-2.0*cos)*v.X - 2.0*v.Y - 2.0*v.Z,
			v.W - cos*v.X - v.Z,
			v.W - cos*v.X - v.Y,
		}

	}

	var f func(vector.Vec4) vector.Vec4
	var a, b, c, d float64

	if metric == 'p' {

		a = 1.0
		b = 0.5
		c = 1.0
		d = 1.0

	} else {

		a = 1.0 / den
		b = cos / den
		c = Rt2 * math.Sqrt(cos) / den
		d = c

	}

	f = func(v vector.Vec4) vector.Vec4 {

		return vector.Vec4{a * v.W, b * v.X, c * v.Y, d * v.Z}

	}

	matrices := shared.Matrices{
		Amat: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{v.W, v.X, v.Z, v.Y}
		},
		Bmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Y, -v.Z} },
		C: cMat,
		Dmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, -v.X, v.Y, v.Z} },
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}

	initialVerts := []vector.Vec4{
		{1, 0, 1, 0},
		{1, 0, -1, 0},
		{1, 0, 0, 1},
		{1, 0, 0, -1},
	}

	initialEdges := []vector.Vec4{
		vector.Scale4(vector.Vec4{2, 0, 1, 1}, 1.0/math.Sqrt(math.Abs(hyperbolic.HyperbolicNorm(f(vector.Vec4{2, 0, 1, 1}))))),
		vector.Scale4(vector.Vec4{2, 0, 1, -1}, 1.0/math.Sqrt(math.Abs(hyperbolic.HyperbolicNorm(f(vector.Vec4{2, 0, 1, -1}))))),
		vector.Scale4(vector.Vec4{2, 0, -1, 1}, 1.0/math.Sqrt(math.Abs(hyperbolic.HyperbolicNorm(f(vector.Vec4{2, 0, -1, 1}))))),
		vector.Scale4(vector.Vec4{2, 0, -1, -1}, 1.0/math.Sqrt(math.Abs(hyperbolic.HyperbolicNorm(f(vector.Vec4{2, 0, -1, -1}))))),
	}

	fVal := 0.0
	ev := 0.0
	fv := 0.0
	if metric == 'p' {

		fVal = 1.0
		ev = 0.5
		fv = 1.0

	} else {

		fVal = den
		ev = math.Abs((1.0 - cos) / (1.0 - 2.0*cos))
		fv = 1.0 / (1.0 - 2.0*cos)

	}

	fPoints, fNames := tesselations.MakeFaces(vector.Vec4{fVal, 0, 0, 0}, numberOfFaces, 4, matrices)

	v := tesselations.MakeRing(initialVerts, matrices, fNames)
	e := tesselations.MakeRing(initialEdges, matrices, fNames)

	faceData := tesselations.GenerateFaceData(math.Abs(fv), 4, fPoints, v, f)

	edgeData := tesselations.GenerateEdgeData(math.Abs(ev), e, v, f)
	faceData = tesselations.OrderFaces(4, faceData, edgeData)

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
		CellType:        "euclidean",
		VV:              vv,
		MetricValues:    MetricValues{E: eVal, P: pVal},
		Vertices:        v,
		Edges:           edgeData,
		Faces:           faceData,
		Matrices:        matrices,
		Flip:            func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.Y, v.Z, v.X} },
	}
}
