package data

import (
	"math"

	"github.com/calummccain/coxeter/hyperbolic"
	"github.com/calummccain/coxeter/tesselations"
	"github.com/calummccain/coxeter/vector"
)

func HexagonalData(n float64, numberOfFaces int) CellData {

	eVal := 2.0
	pVal := 6.0

	cos := math.Pow(math.Cos(math.Pi/n), 2)

	den := math.Sqrt(math.Abs(1.0 - 4.0*cos/3.0))

	metric := Boundaries(n, eVal, pVal)

	cMat := func(v vector.Vec4) vector.Vec4 {

		return vector.Vec4{
			(1.0+2.0*cos)*v.W - 2.0*cos*cos*v.X - cos*v.Y - cos*v.Z,
			2.0*v.W + (1.0-2.0*cos)*v.X - v.Y - v.Z,
			3.0*v.W - 3.0*cos*v.X - 0.5*v.Y - 1.5*v.Z,
			v.W - cos*v.X - v.Y*0.5 + v.Z*0.5,
		}

	}

	var f func(vector.Vec4) vector.Vec4
	var a, b, c, d float64

	if metric == 'p' {

		a = 1.0
		b = 0.75
		c = 0.5
		d = 0.5 * Rt3

	} else {

		a = 1.0 / den
		b = cos / den
		c = math.Sqrt(cos/3.0) / den
		d = math.Sqrt(cos) / den

	}

	f = func(v vector.Vec4) vector.Vec4 {

		return vector.Vec4{a * v.W, b * v.X, c * v.Y, d * v.Z}

	}

	initialVerts := []vector.Vec4{
		{1, 0, 2, 0},
		{1, 0, -2, 0},
		{1, 0, 1, 1},
		{1, 0, 1, -1},
		{1, 0, -1, 1},
		{1, 0, -1, -1},
	}

	initialEdges := []vector.Vec4{
		vector.Scale4(vector.Vec4{2, 0, 3, 1}, 1.0/math.Sqrt(math.Abs(hyperbolic.HyperbolicNorm(f(vector.Vec4{2, 0, 3, 1}))))),
		vector.Scale4(vector.Vec4{2, 0, 0, 2}, 1.0/math.Sqrt(math.Abs(hyperbolic.HyperbolicNorm(f(vector.Vec4{2, 0, 0, 2}))))),
		vector.Scale4(vector.Vec4{2, 0, -3, 1}, 1.0/math.Sqrt(math.Abs(hyperbolic.HyperbolicNorm(f(vector.Vec4{2, 0, -3, 1}))))),
		vector.Scale4(vector.Vec4{2, 0, -3, -1}, 1.0/math.Sqrt(math.Abs(hyperbolic.HyperbolicNorm(f(vector.Vec4{2, 0, -3, -1}))))),
		vector.Scale4(vector.Vec4{2, 0, 0, -2}, 1.0/math.Sqrt(math.Abs(hyperbolic.HyperbolicNorm(f(vector.Vec4{2, 0, 0, -2}))))),
		vector.Scale4(vector.Vec4{2, 0, 3, -1}, 1.0/math.Sqrt(math.Abs(hyperbolic.HyperbolicNorm(f(vector.Vec4{2, 0, 3, -1}))))),
	}

	fVal := 0.0
	vv := 0.0
	ev := 0.0
	fv := 0.0

	if metric == 'p' {

		fVal = 1.0
		vv = 0.5
		ev = 0.25
		fv = 1.0

	} else {

		fVal = den
		vv = (1.0 - 2.0*cos/3.0) / math.Abs(1.0-4.0*cos/3.0)
		ev = (1.0 - cos) / (1.0 - 4.0*cos/3.0)
		fv = 1.0 / (1.0 - 4.0*cos/3.0)

	}

	fPoints, fNames := tesselations.MakeFaces(vector.Vec4{fVal, 0, 0, 0}, numberOfFaces, 6, matrices)

	v := tesselations.MakeRing(initialVerts, matrices, fNames)
	e := tesselations.MakeRing(initialEdges, matrices, fNames)

	faceData := tesselations.GenerateFaceData(math.Abs(fv), 6, fPoints, v, f)

	edgeData := tesselations.GenerateEdgeData(math.Abs(ev), e, v, f)

	faceData = tesselations.OrderFaces(6, faceData, edgeData)

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
		Amat: func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{v.W, v.X, 0.5*v.Y + 1.5*v.Z, 0.5 * (v.Y - v.Z)}
		},
		Bmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, v.X, v.Y, -v.Z} },
		Cmat: c,
		Dmat: func(v vector.Vec4) vector.Vec4 { return vector.Vec4{v.W, -v.X, v.Y, v.Z} },
		Emat: func(v vector.Vec4) vector.Vec4 { return v },
		Fmat: f,
	}
}
