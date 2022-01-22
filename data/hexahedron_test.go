package data

import (
	"testing"
)

func TestHexahedron(t *testing.T) {

	for i := 3; i <= 20; i++ {

		DataTest(HexahedronData(float64(i)), t)

	}
}
