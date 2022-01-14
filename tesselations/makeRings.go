package tesselations

import (
	"github.com/calummccain/coxeter/vector"
)

func MakeRing(initialRings []vector.Vec4, A, B, C, D, E, F func(vector.Vec4) vector.Vec4, fNames []string) []vector.Vec4 {

	rings := initialRings

	newRing := []vector.Vec4{}

	var testRing []vector.Vec4

	for i := 0; i < len(fNames); i++ {

		testRing = vector.TransformVertices(rings, fNames[i], A, B, C, D, E, F)

		for j := 0; j < len(testRing); j++ {

			if !vector.IsInArray4(testRing[j], newRing) && !vector.IsInArray4(testRing[j], rings) {

				newRing = append(newRing, testRing[j])

			}

		}

	}

	rings = append(rings, newRing...)

	return rings

}
