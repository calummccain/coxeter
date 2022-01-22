package data

import (
	"testing"
)

func TestTetrahedron(t *testing.T) {

	for i := 3; i <= 20; i++ {

		DataTest(TetrahedronData(float64(i)), t)

	}
}
