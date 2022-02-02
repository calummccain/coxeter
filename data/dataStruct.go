package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

type GoursatTetrahedron struct {
	V   vector.Vec4
	E   vector.Vec4
	F   vector.Vec4
	C   vector.Vec4
	CFE vector.Vec4
	CFV vector.Vec4
	CEV vector.Vec4
	FEV vector.Vec4
	CF  float64
	CE  float64
	CV  float64
	FE  float64
	FV  float64
	EV  float64
}

type Coxeter struct {
	P                  float64
	Q                  float64
	R                  float64
	A                  func(vector.Vec4) vector.Vec4
	B                  func(vector.Vec4) vector.Vec4
	C                  func(vector.Vec4) vector.Vec4
	D                  func(vector.Vec4) vector.Vec4
	FaceReflections    []string
	GoursatTetrahedron GoursatTetrahedron
}

type Honeycomb struct {
	Coxeter      Coxeter
	CellType     string
	Vertices     []vector.Vec4
	Edges        [][2]int
	Faces        [][]int
	EVal         float64
	PVal         float64
	Space        byte
	Scale        func(vector.Vec4) vector.Vec4
	InnerProduct func(vector.Vec4, vector.Vec4) float64
}

type Cell struct {
	Vertices    []vector.Vec4
	Edges       [][2]int
	Faces       [][]int
	NumVertices int
	NumEdges    int
	NumFaces    int
}

func (honeycomb *Honeycomb) DistanceSquared(a, b vector.Vec4) float64 {

	if honeycomb.Space == 'e' {
		return honeycomb.InnerProduct(vector.Diff4(a, b), vector.Diff4(a, b))
	}

	den := 1.0

	if math.Abs(honeycomb.InnerProduct(a, a)) > DistanceSquaredEps {
		den *= honeycomb.InnerProduct(a, a)
	}

	if math.Abs(honeycomb.InnerProduct(b, b)) > DistanceSquaredEps {
		den *= honeycomb.InnerProduct(b, b)
	}

	return honeycomb.InnerProduct(a, b) * honeycomb.InnerProduct(a, b) / den

}
