package data

import (
	"github.com/calummccain/coxeter/vector"
)

type MetricValues struct {
	E float64
	P float64
}

type CellData struct {
	Metric          byte
	Vertices        []vector.Vec4
	Edges           [][2]int
	Faces           [][]int
	NumVertices     int
	NumEdges        int
	NumFaces        int
	FaceReflections []string
	OuterReflection string
	V               vector.Vec4
	E               vector.Vec4
	F               vector.Vec4
	C               vector.Vec4
	CellType        string
	VV              float64
	CF              float64
	CE              float64
	CV              float64
	FE              float64
	FV              float64
	EV              float64
	MetricValues    MetricValues
	Amat            func(vector.Vec4) vector.Vec4
	Bmat            func(vector.Vec4) vector.Vec4
	Cmat            func(vector.Vec4) vector.Vec4
	Dmat            func(vector.Vec4) vector.Vec4
	Emat            func(vector.Vec4) vector.Vec4
	Fmat            func(vector.Vec4) vector.Vec4
}
