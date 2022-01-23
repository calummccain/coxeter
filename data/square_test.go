package data

import (
	"testing"
)

func TestSquare(t *testing.T) {

	for i := 3; i <= 20; i++ {
		DataTest(SquareData(float64(i), 7), t)
	}
}
