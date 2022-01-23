package data

import (
	"fmt"
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

			//v := cellData.Fmat(cellData.Vertices[j])
			//v.Scale(1.0 / v.W)
			//fmt.Println(v)

			fmt.Println(cellData.Vertices[j], cellData.InnerProduct(cellData.Vertices[j], cellData.Vertices[j]), faces[i], cellData.InnerProduct(faces[i], faces[i]), cellData.InnerProduct(cellData.Vertices[j], faces[i]), cellData.DistanceSquared(cellData.Vertices[j], faces[i]))

			if math.Abs(cellData.DistanceSquared(cellData.Vertices[j], faces[i])-cellData.FV) < eps {

				nearestPoints = append(nearestPoints, j)
				k++

			}

		}

		faceData = append(faceData, nearestPoints)

	}

	cellData.Faces = faceData

}
