package data

import "github.com/calummccain/coxeter/shared"

type MetricValues struct {
	E float64
	P float64
}

type CellData struct {
	Metric          byte
	Vertices        [][4]float64
	Edges           [][2]int
	Faces           [][]int
	NumVertices     int
	NumEdges        int
	NumFaces        int
	FaceReflections []string
	OuterReflection string
	V               [4]float64
	E               [4]float64
	F               [4]float64
	C               [4]float64
	CellType        string
	Vv              float64
	MetricValues    MetricValues
	Matrices        shared.Matrices
	Flip            func([4]float64) [4]float64
}
