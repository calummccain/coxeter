package data

import (
	"testing"
)

func TextOctahedron(t *testing.T) {

	for i := 3; i <= 20; i++ {

		DataTest(OctahedronData(float64(i)), t)

	}
}
