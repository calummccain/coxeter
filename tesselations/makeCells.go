package tesselations

import (
	"coxeter/shared"
	"coxeter/vector"
)

func MakeCells(cell [4]float64, faceReflections []string, numLayers int, matrixDict shared.Matrices) ([][4]float64, []string) {

	cells := [][4]float64{}
	cellNames := []string{}

	outerLayer := [][4]float64{cell}
	outerNames := []string{""}

	var newLayer [][4]float64
	var newNames []string

	var testCenters [][4]float64

	i := 1

	for i <= numLayers {

		newLayer = [][4]float64{}
		newNames = []string{}

		for j := 0; j < len(faceReflections); j++ {

			testCenters = vector.TransformVertices(outerLayer, faceReflections[j]+"d", matrixDict)

			for k := 0; k < len(testCenters); k++ {

				if !(vector.IsInArray(testCenters[k], cells) || vector.IsInArray(testCenters[k], outerLayer) || vector.IsInArray(testCenters[k], newLayer)) {

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
