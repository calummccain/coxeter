package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

// Only works for hyperbolic cells (no need to do for spherical or euclidean)
func (honeycomb *Honeycomb) GenerateEdgeData(edges []vector.Vec4) {

	edgeData := [][2]int{}
	var nearestPoints [2]int
	var k int

	eps := GenerateEdgeDataEps

	for i := 0; i < len(edges); i++ {

		nearestPoints = [2]int{}
		k = 0

		for j := 0; j < len(honeycomb.Vertices); j++ {

			if k == 2 {

				break

			}

			if math.Abs(honeycomb.DistanceSquared(honeycomb.Vertices[j], edges[i])-honeycomb.Coxeter.GoursatTetrahedron.EV) < eps {

				nearestPoints[k] = j
				k++

			}

		}

		edgeData = append(edgeData, nearestPoints)

	}

	honeycomb.Edges = edgeData

}
