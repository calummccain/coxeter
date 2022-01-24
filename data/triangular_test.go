package data

import (
	"testing"
)

func TestTriangular(t *testing.T) {

	for i := 3; i <= 20; i++ {
		DataTest(TriangularData(float64(i), 7), t)
	}
}
