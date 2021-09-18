package tesselations

import (
	"coxeter/hyperbolic"
	"math"
)

func GenerateFaceData(fvDist float64, numEdges int, f [][4]float64, v [][4]float64, fmat func([4]float64) [4]float64) [][]int {

	faceData := [][]int{}
	var nearestPoints []int
	var k int

	eps := GenerateFaceDataEps

	for i := 0; i < len(f); i++ {

		nearestPoints = []int{}
		k = 0

		for j := 0; j < len(v); j++ {

			if k == numEdges {

				break

			}

			if math.Abs(math.Pow(hyperbolic.HyperboloidInnerProduct(fmat(v[j]), fmat(f[i])), 2)-fvDist) < eps {

				nearestPoints = append(nearestPoints, j)
				k++

			}

		}

		faceData = append(faceData, nearestPoints)

	}

	return faceData

}
