package data

import (
	"github.com/calummccain/coxeter/vector"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (cellData *CellData) OrderFaces() {

	newFaceData := [][]int{}

	var newFace []int
	var k int

	for i := 0; i < len(cellData.Faces); i++ {

		newFace = []int{cellData.Faces[i][0]}
		k = 1

		for k < int(cellData.P) {

			for j := 1; j < int(cellData.P); j++ {

				if vector.IsInArray2([2]int{min(newFace[len(newFace)-1], cellData.Faces[i][j]), max(newFace[len(newFace)-1], cellData.Faces[i][j])}, cellData.Edges) && !vector.IsInArray1(cellData.Faces[i][j], newFace) {

					newFace = append(newFace, cellData.Faces[i][j])
					k++

				}

			}

		}

		newFaceData = append(newFaceData, newFace)

	}

	cellData.Faces = newFaceData

}
