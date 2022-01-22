package data

import (
	"testing"
)

func TestIcosahedron(t *testing.T) {

	for i := 3; i <= 20; i++ {

		DataTest(IcosahedronData(float64(i)), t)

	}
}
