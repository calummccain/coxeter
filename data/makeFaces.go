package data

import (
	"github.com/calummccain/coxeter/vector"
)

func (cellData *CellData) MakeFaces(maxNumber int) []vector.Vec4 {

	faces := []vector.Vec4{cellData.F}
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

		for j < int(cellData.P) {

			testCenters = vector.TransformVertices(faces, spec, cellData.Amat, cellData.Bmat, cellData.Cmat, cellData.Dmat)

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

	cellData.FaceReflections = faceNames

	return faces

}
