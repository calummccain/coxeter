package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

type CellData struct {
	P               float64
	Q               float64
	R               float64
	Metric          byte
	Vertices        []vector.Vec4
	Edges           [][2]int
	Faces           [][]int
	NumVertices     int
	NumEdges        int
	NumFaces        int
	FaceReflections []string
	OuterReflection string
	CellType        string
	V               vector.Vec4
	E               vector.Vec4
	F               vector.Vec4
	C               vector.Vec4
	CFE             vector.Vec4
	CFV             vector.Vec4
	CEV             vector.Vec4
	FEV             vector.Vec4
	VV              float64
	CF              float64
	CE              float64
	CV              float64
	FE              float64
	FV              float64
	EV              float64
	EVal            float64
	PVal            float64
	Amat            func(vector.Vec4) vector.Vec4
	Bmat            func(vector.Vec4) vector.Vec4
	Cmat            func(vector.Vec4) vector.Vec4
	Dmat            func(vector.Vec4) vector.Vec4
	Fmat            func(vector.Vec4) vector.Vec4
	InnerProduct    func(vector.Vec4, vector.Vec4) float64
}

func (cellData *CellData) DistanceSquared(a, b vector.Vec4) float64 {

	if cellData.Metric == 'p' {
		return math.Pow(cellData.InnerProduct(a, b), 2.0)
	} else {
		return math.Pow(cellData.InnerProduct(a, b), 2.0) / (cellData.InnerProduct(a, a) * cellData.InnerProduct(b, b))
	}

}
