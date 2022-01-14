package tesselations

import (
	"math"

	"github.com/calummccain/coxeter/hyperbolic"
	"github.com/calummccain/coxeter/vector"
)

// Only works for hyperbolic cells (no need to do for spherical or euclidean)
func GenerateEdgeData(ev float64, e, v []vector.Vec4, fmat func(vector.Vec4) vector.Vec4) [][2]int {

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

			if math.Abs(math.Pow(hyperbolic.HyperboloidInnerProduct(fmat(v[j]), fmat(e[i])), 2)-ev) < eps {

				nearestPoints[k] = j
				k++

			}

		}

		edgeData = append(edgeData, nearestPoints)

	}

	return edgeData

}
