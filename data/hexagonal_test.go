package data

import (
	"testing"
)

func TestHexagonal(t *testing.T) {

	for i := 6; i <= 6; i++ {
		data := HexagonalData(float64(i), 7)
		DataTest(data, t)

	}
}
