package tesselations

import (
	"coxeter/vector"
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

func OrderFaces(numEdges int, faceData [][]int, edgeData [][2]int) [][]int {

	newFaceData := [][]int{}

	var newFace []int
	var k int

	for i := 0; i < len(faceData); i++ {

		newFace = []int{faceData[i][0]}
		k = 1

		for k < numEdges {

			for j := 1; j < numEdges; j++ {

				if vector.IsInArray2([2]int{min(newFace[len(newFace)-1], faceData[i][j]), max(newFace[len(newFace)-1], faceData[i][j])}, edgeData) && !vector.IsInArray1(faceData[i][j], newFace) {

					newFace = append(newFace, faceData[i][j])
					k++

				}

			}

		}

		newFaceData = append(newFaceData, newFace)

	}

	return newFaceData

}
