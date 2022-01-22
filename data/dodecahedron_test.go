package data

import (
	"testing"
)

func TestDodecahedron(t *testing.T) {

	for i := 3; i <= 20; i++ {

		DataTest(DodecahedronData(float64(i)), t)

	}
}
