package data

import (
	"github.com/calummccain/coxeter/vector"
)

func (coxeter *Coxeter) MakeFaces(maxNumber int) ([]string, []vector.Vec4) {

	faces := []vector.Vec4{coxeter.GoursatTetrahedron.F}
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

		for j < int(coxeter.P) {

			testCenters = vector.TransformVertices(faces, spec, coxeter.A, coxeter.B, coxeter.C, coxeter.D)

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

	return faceNames, faces

}
