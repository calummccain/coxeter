package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func (cellData *CellData) GenerateFaceData(faces []vector.Vec4) {

	faceData := [][]int{}
	var nearestPoints []int
	var k int

	eps := GenerateFaceDataEps

	for i := 0; i < len(faces); i++ {

		nearestPoints = []int{}
		k = 0

		for j := 0; j < len(cellData.Vertices); j++ {

			if k == int(cellData.P) {

				break

			}

			if math.Abs(cellData.DistanceSquared(cellData.Vertices[j], faces[i])-cellData.FV) < eps {

				nearestPoints = append(nearestPoints, j)
				k++

			}

		}

		faceData = append(faceData, nearestPoints)

	}

	cellData.Faces = faceData

}
