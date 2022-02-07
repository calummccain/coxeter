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

func (honeycomb *Honeycomb) OrderFaces() {

	newFaceData := [][]int{}

	var newFace []int
	var k int

	var l1 int

	for i := 0; i < len(honeycomb.Faces); i++ {

		if len(honeycomb.Faces[i]) < 3 {
			continue
		}

		newFace = []int{honeycomb.Faces[i][0]}
		k = 1

		for k < len(honeycomb.Faces[i]) {

			l1 = len(newFace)

			for j := 1; j < len(honeycomb.Faces[i]); j++ {

				if vector.IsInArray2(
					[2]int{
						min(newFace[len(newFace)-1], honeycomb.Faces[i][j]),
						max(newFace[len(newFace)-1], honeycomb.Faces[i][j]),
					},
					honeycomb.Edges) && !vector.IsInArray1(honeycomb.Faces[i][j], newFace) {

					newFace = append(newFace, honeycomb.Faces[i][j])
					k++

				}

			}

			if l1 == len(newFace) {
				break
			}

		}

		newFaceData = append(newFaceData, newFace)

	}

	honeycomb.Faces = newFaceData

}
