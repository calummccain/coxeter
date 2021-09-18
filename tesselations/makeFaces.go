package tesselations

import (
	"coxeter/shared"
	"coxeter/vector"
)

func MakeFaces(face [4]float64, maxNumber int, numEdges int, matrixDict shared.Matrices) ([][4]float64, []string) {

	faces := [][4]float64{face}
	faceNames := []string{""}
	i := 1

	var j int
	var spec string
	var newFaces [][4]float64
	var newNames []string
	var testCenters [][4]float64

	for i < maxNumber {

		j = 0
		spec = "c"
		newFaces = [][4]float64{}
		newNames = []string{}

		for j < numEdges {

			testCenters = vector.TransformVertices(faces, spec, matrixDict)

			for k := 0; k < len(testCenters); k++ {

				if !(vector.IsInArray(testCenters[k], faces) || vector.IsInArray(testCenters[k], newFaces)) {

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
