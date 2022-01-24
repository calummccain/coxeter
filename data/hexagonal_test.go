package data

import (
	"testing"
)

func TestHexagonal(t *testing.T) {

	for i := 3; i <= 20; i++ {

		DataTest(HexagonalData(float64(i), 100), t)

	}
}
