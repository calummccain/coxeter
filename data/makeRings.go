package data

import (
	"github.com/calummccain/coxeter/vector"
)

func (coxeter *Coxeter) MakeRing(initialRings []vector.Vec4) []vector.Vec4 {

	rings := initialRings

	newRing := []vector.Vec4{}

	var testRing []vector.Vec4

	for i := 0; i < len(coxeter.FaceReflections); i++ {

		testRing = vector.TransformVertices(rings, coxeter.FaceReflections[i], coxeter.A, coxeter.B, coxeter.C, coxeter.D)

		for j := 0; j < len(testRing); j++ {

			if !vector.IsInArray4(testRing[j], newRing) && !vector.IsInArray4(testRing[j], rings) {

				newRing = append(newRing, testRing[j])

			}

		}

	}

	rings = append(rings, newRing...)

	return rings

}
