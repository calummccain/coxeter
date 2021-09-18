package tesselations

import (
	"github.com/calummccain/coxeter/vector"

	"github.com/calummccain/coxeter/shared"
)

func MakeRing(initialRings [][4]float64, matrixDict shared.Matrices, fNames []string) [][4]float64 {

	rings := initialRings

	newRing := [][4]float64{}

	var testRing [][4]float64

	for i := 0; i < len(fNames); i++ {

		testRing = vector.TransformVertices(rings, fNames[i], matrixDict)

		for j := 0; j < len(testRing); j++ {

			if !vector.IsInArray(testRing[j], newRing) && !vector.IsInArray(testRing[j], rings) {

				newRing = append(newRing, testRing[j])

			}

		}

	}

	rings = append(rings, newRing...)

	return rings

}
