package tesselations

import (
	"math"

	"github.com/calummccain/coxeter/hyperbolic"
)

func GenerateEdgeData(evDist float64, e, v [][4]float64, fmat func([4]float64) [4]float64) [][2]int {

	edgeData := [][2]int{}
	var nearestPoints [2]int
	var k int

	eps := GenerateEdgeDataEps

	for i := 0; i < len(e); i++ {

		nearestPoints = [2]int{}
		k = 0

		for j := 0; j < len(v); j++ {

			if k == 2 {

				break

			}

			if math.Abs(math.Pow(hyperbolic.HyperboloidInnerProduct(fmat(v[j]), fmat(e[i])), 2)-evDist) < eps {

				nearestPoints[k] = j
				k++

			}

		}

		edgeData = append(edgeData, nearestPoints)

	}

	return edgeData

}
