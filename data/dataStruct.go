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

	if cellData.Metric == 'e' {
		return cellData.InnerProduct(vector.Diff4(a, b), vector.Diff4(a, b))
	}

	den := 1.0

	if math.Abs(cellData.InnerProduct(a, a)) < DistanceSquaredEps {
		//den *= cellData.InnerProduct(vector.Vec4{W: a.W, X: 0, Y: 0, Z: 0}, vector.Vec4{W: a.W, X: 0, Y: 0, Z: 0})
	} else {
		den *= cellData.InnerProduct(a, a)
	}

	if math.Abs(cellData.InnerProduct(b, b)) < DistanceSquaredEps {
		//den *= cellData.InnerProduct(vector.Vec4{W: b.W, X: 0, Y: 0, Z: 0}, vector.Vec4{W: b.W, X: 0, Y: 0, Z: 0})
	} else {
		den *= cellData.InnerProduct(b, b)
	}

	return cellData.InnerProduct(a, b) * cellData.InnerProduct(a, b) / den

}
