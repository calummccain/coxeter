package tesselations

import (
	"github.com/calummccain/coxeter/vector"
)

func MakeCells(cell vector.Vec4, faceReflections []string, numLayers int, A, B, C, D, E, F func(vector.Vec4) vector.Vec4) ([]vector.Vec4, []string) {

	cells := []vector.Vec4{}
	cellNames := []string{}

	outerLayer := []vector.Vec4{cell}
	outerNames := []string{""}

	var newLayer []vector.Vec4
	var newNames []string

	var testCenters []vector.Vec4

	i := 1

	for i <= numLayers {

		newLayer = []vector.Vec4{}
		newNames = []string{}

		for j := 0; j < len(faceReflections); j++ {

			testCenters = vector.TransformVertices(outerLayer, faceReflections[j]+"d", A, B, C, D, E, F)

			for k := 0; k < len(testCenters); k++ {

				if !(vector.IsInArray4(testCenters[k], cells) || vector.IsInArray4(testCenters[k], outerLayer) || vector.IsInArray4(testCenters[k], newLayer)) {

					newLayer = append(newLayer, testCenters[k])

					newNames = append(newNames, faceReflections[j]+"d"+outerNames[k])

				}

			}

		}

		cells = append(cells, outerLayer...)
		cellNames = append(cellNames, outerNames...)

		outerLayer = newLayer
		outerNames = newNames

		i++

	}

	cells = append(cells, outerLayer...)
	cellNames = append(cellNames, outerNames...)

	return cells, cellNames

}
