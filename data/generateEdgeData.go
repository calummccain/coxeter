package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

// Only works for hyperbolic cells (no need to do for spherical or euclidean)
func (cellData *CellData) GenerateEdgeData(edges []vector.Vec4) {

	edgeData := [][2]int{}
	var nearestPoints [2]int
	var k int

	eps := GenerateEdgeDataEps

	for i := 0; i < len(edges); i++ {

		nearestPoints = [2]int{}
		k = 0

		for j := 0; j < len(cellData.Vertices); j++ {

			if k == 2 {

				break

			}

			if math.Abs(cellData.DistanceSquared(cellData.Vertices[j], edges[i])-cellData.EV) < eps {

				nearestPoints[k] = j
				k++

			}

		}

		edgeData = append(edgeData, nearestPoints)

	}

	cellData.Edges = edgeData

}
