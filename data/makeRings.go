package data

import (
	"github.com/calummccain/coxeter/vector"
)

func (cellData *CellData) MakeRing(initialRings []vector.Vec4) []vector.Vec4 {

	rings := initialRings

	newRing := []vector.Vec4{}

	var testRing []vector.Vec4

	for i := 0; i < len(cellData.FaceReflections); i++ {

		testRing = vector.TransformVertices(rings, cellData.FaceReflections[i], cellData.Amat, cellData.Bmat, cellData.Cmat, cellData.Dmat)

		for j := 0; j < len(testRing); j++ {

			if !vector.IsInArray4(testRing[j], newRing) && !vector.IsInArray4(testRing[j], rings) {

				newRing = append(newRing, testRing[j])

			}

		}

	}

	rings = append(rings, newRing...)

	return rings

}
