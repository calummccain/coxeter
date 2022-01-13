package data

import (
	"testing"

	"github.com/calummccain/coxeter/vector"
)

const eps = 0.0001

func TestTetrahedron(t *testing.T) {
	data := TetrahedronData(3)
	if vector.DistanceSquared4(data.Amat(data.C), data.C) > eps {
		t.Errorf("C is not fixed by As")
	}
}
