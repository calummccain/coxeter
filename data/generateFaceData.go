package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func (honeycomb *Honeycomb) GenerateFaceData(faces []vector.Vec4) {

	faceData := [][]int{}
	var nearestPoints []int
	var k int

	eps := GenerateFaceDataEps

	for i := 0; i < len(faces); i++ {

		nearestPoints = []int{}
		k = 0

		for j := 0; j < len(honeycomb.Vertices); j++ {

			if k == int(honeycomb.Coxeter.P) {

				break

			}

			if math.Abs(honeycomb.DistanceSquared(honeycomb.Vertices[j], faces[i])-honeycomb.Coxeter.GoursatTetrahedron.FV) < eps {

				nearestPoints = append(nearestPoints, j)
				k++

			}

		}

		faceData = append(faceData, nearestPoints)

	}

	honeycomb.Faces = faceData

}
