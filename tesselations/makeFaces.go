package tesselations

import (
	"github.com/calummccain/coxeter/vector"
)

func MakeFaces(face vector.Vec4, maxNumber int, numEdges int, A, B, C, D, E, F func(vector.Vec4) vector.Vec4) ([]vector.Vec4, []string) {

	faces := []vector.Vec4{face}
	faceNames := []string{""}
	i := 1

	var j int
	var spec string
	var newFaces []vector.Vec4
	var newNames []string
	var testCenters []vector.Vec4

	for i < maxNumber {

		j = 0
		spec = "c"
		newFaces = []vector.Vec4{}
		newNames = []string{}

		for j < numEdges {

			testCenters = vector.TransformVertices(faces, spec, A, B, C, D, E, F)

			for k := 0; k < len(testCenters); k++ {

				if !(vector.IsInArray4(testCenters[k], faces) || vector.IsInArray4(testCenters[k], newFaces)) {

					newFaces = append(newFaces, testCenters[k])

					newNames = append(newNames, spec+faceNames[k])

				}

			}

			spec = "ab" + spec
			j++

		}

		faces = append(faces, newFaces...)
		faceNames = append(faceNames, newNames...)

		i = len(faces)

	}

	return faces, faceNames

}
