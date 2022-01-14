package tesselations

import (
	"math"

	"github.com/calummccain/coxeter/hyperbolic"
	"github.com/calummccain/coxeter/vector"
)

func GenerateFaceData(fv float64, numEdges int, f []vector.Vec4, v []vector.Vec4, fmat func(vector.Vec4) vector.Vec4) [][]int {

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

			if math.Abs(math.Pow(hyperbolic.HyperboloidInnerProduct(fmat(v[j]), fmat(f[i])), 2)-fv) < eps {

				nearestPoints = append(nearestPoints, j)
				k++

			}

		}

		faceData = append(faceData, nearestPoints)

	}

	return faceData

}
