package data

import (
	"math"
	"testing"

	"github.com/calummccain/coxeter/vector"
)

func DataTest(data CellData, t *testing.T) {

	// Test reflection matrices fix correct vertices of Goursat tetrahedron
	if vector.DistanceSquared4(data.Amat(data.C), data.C) > TestEps {
		t.Errorf("%f cell is not fixed by A \n A(C) = %v \n C = %v", data.R, data.Amat(data.C), data.C)
	}

	if vector.DistanceSquared4(data.Amat(data.F), data.F) > TestEps {
		t.Errorf("%f face is not fixed by A \n A(F) = %v \n F = %v", data.R, data.Amat(data.F), data.F)
	}

	if vector.DistanceSquared4(data.Amat(data.E), data.E) > TestEps {
		t.Errorf("%f edge is not fixed by A \n A(E) = %v \n E = %v", data.R, data.Amat(data.E), data.E)
	}

	if vector.DistanceSquared4(data.Bmat(data.C), data.C) > TestEps {
		t.Errorf("%f cell is not fixed by B \n B(C) = %v \n C = %v", data.R, data.Bmat(data.C), data.C)
	}

	if vector.DistanceSquared4(data.Bmat(data.F), data.F) > TestEps {
		t.Errorf("%f face is not fixed by B \n B(F) = %v \n F = %v", data.R, data.Bmat(data.F), data.F)
	}

	if vector.DistanceSquared4(data.Bmat(data.V), data.V) > TestEps {
		t.Errorf("%f vertex is not fixed by B \n B(V) = %v \n V = %v", data.R, data.Bmat(data.V), data.V)
	}

	if vector.DistanceSquared4(data.Cmat(data.C), data.C) > TestEps {
		t.Errorf("%f cell is not fixed by C \n C(C) = %v \n C = %v", data.R, data.Cmat(data.C), data.C)
	}

	if vector.DistanceSquared4(data.Cmat(data.E), data.E) > TestEps {
		t.Errorf("%f edge is not fixed by C \n C(E) = %v \n E = %v", data.R, data.Cmat(data.E), data.E)
	}

	if vector.DistanceSquared4(data.Cmat(data.V), data.V) > TestEps {
		t.Errorf("%f vertex is not fixed by C \n C(V) = %v \n V = %v", data.R, data.Cmat(data.V), data.V)
	}

	if vector.DistanceSquared4(data.Dmat(data.F), data.F) > TestEps {
		t.Errorf("%f face is not fixed by D \n D(F) = %v \n F = %v", data.R, data.Dmat(data.F), data.F)
	}

	if vector.DistanceSquared4(data.Dmat(data.E), data.E) > TestEps {
		t.Errorf("%f edge is not fixed by D \n D(E) = %v \n E = %v", data.R, data.Dmat(data.E), data.E)
	}

	if vector.DistanceSquared4(data.Dmat(data.V), data.V) > TestEps {
		t.Errorf("%f vertex is not fixed by D \n D(V) = %v \n V = %v", data.R, data.Dmat(data.V), data.V)
	}

	// Check that the pythagorean identities are satisfied
	if data.Space != 'e' {
		if math.Abs(data.CF*data.FE-data.CE) > TestEps {
			t.Errorf("%f CF * FE != CE \n CF * FE = %f \n CE = %f", data.R, data.CF*data.FE, data.CE)
		}

		if math.Abs(data.CF*data.FV-data.CV) > TestEps {
			t.Errorf("%f CF * FV != CV \n CF * FV = %f \n CV = %f", data.R, data.CF*data.FV, data.CV)
		}

		if math.Abs(data.CE*data.EV-data.CV) > TestEps {
			t.Errorf("%f CE * EV != CV \n CE * EV = %f \n CV = %f", data.R, data.CE*data.EV, data.CV)
		}

		if math.Abs(data.FE*data.EV-data.FV) > TestEps {
			t.Errorf("%f FE * EV != FV \n FE * EV = %f \n FV = %f", data.R, data.FE*data.EV, data.FV)
		}
	}

	// Check that the vertices have the correct norm
	for _, vertex := range data.Vertices {
		if data.Space == 's' {
			if math.Abs(data.InnerProduct(vertex, vertex)-1.0) > TestEps {
				t.Errorf("%f Magnitude of %v != 1 \n <v,v> = %f", data.R, vertex, data.InnerProduct(vertex, vertex))
			}
		} else if data.Space == 'h' {
			if math.Abs(data.InnerProduct(vertex, vertex)-1.0) > TestEps {
				t.Errorf("%f Magnitude of %v != 1 \n <v,v> = %f", data.R, vertex, data.InnerProduct(vertex, vertex))
			}
		} else if data.Space == 'p' {
			if math.Abs(data.InnerProduct(vertex, vertex)) > TestEps {
				t.Errorf("%f Magnitude of %v != 0 \n <v,v> = %f", data.R, vertex, data.InnerProduct(vertex, vertex))
			}
		} else if data.Space == 'u' {
			if math.Abs(data.InnerProduct(vertex, vertex)+1.0) > TestEps {
				t.Errorf("%f Magnitude of %v != -1 \n <v,v> = %f", data.R, vertex, data.InnerProduct(vertex, vertex))
			}
		}
	}

	// Check that the distances between vertices bounding an edge are correct
	for _, edge := range data.Edges {
		if math.Abs(data.DistanceSquared(data.Vertices[edge[0]], data.Vertices[edge[1]])-math.Pow(data.VV, 2.0)) > TestEps {
			t.Errorf("%f Distance Squared between %v and %v doesn't equal %f \n <a,b> = %f", data.R, data.Vertices[edge[0]], data.Vertices[edge[1]], math.Pow(data.VV, 2.0), data.DistanceSquared(data.Vertices[edge[0]], data.Vertices[edge[1]]))
		}
	}

	// Check that the Goursat tetrahedron has the correct edge lengths
	if math.Abs(data.DistanceSquared(data.C, data.F)-data.CF) > TestEps {
		t.Errorf("%f Distance squared between C and F equals %f not %f", data.R, data.DistanceSquared(data.C, data.F), data.CF)
	}

	if math.Abs(data.DistanceSquared(data.C, data.E)-data.CE) > TestEps {
		t.Errorf("%f Distance squared between C and E equals %f not %f", data.R, data.DistanceSquared(data.C, data.E), data.CE)
	}

	if math.Abs(data.DistanceSquared(data.C, data.V)-data.CV) > TestEps {
		t.Errorf("%f Distance squared between C and V equals %f not %f", data.R, data.DistanceSquared(data.C, data.V), data.CV)
	}

	if math.Abs(data.DistanceSquared(data.F, data.E)-data.FE) > TestEps {
		t.Errorf("%f Distance squared between F and E equals %f not %f", data.R, data.DistanceSquared(data.F, data.E), data.FE)
	}

	if math.Abs(data.DistanceSquared(data.F, data.V)-data.FV) > TestEps {
		t.Errorf("%f Distance squared between F and V equals %f not %f", data.R, data.DistanceSquared(data.F, data.V), data.FV)
	}

	if math.Abs(data.DistanceSquared(data.E, data.V)-data.EV) > TestEps {
		t.Errorf("%f Distance squared between E and V equals %f not %f", data.R, data.DistanceSquared(data.E, data.V), data.EV)
	}

	// Correct Goursat tetrahedron angles
	if data.Space != 'e' {
		if math.Abs(data.DistanceSquared(data.CFE, data.CFV)-math.Pow(math.Cos(math.Pi/data.P), 2.0)) > TestEps {
			t.Errorf("%f Angle between CFE and CFV is not pi/%f", data.R, data.R)
		}

		if math.Abs(data.DistanceSquared(data.CFE, data.FEV)) > TestEps {
			t.Errorf("%f Angle between CFE and FEV is not pi/2", data.R)
		}

		if math.Abs(data.DistanceSquared(data.CFE, data.CEV)) > TestEps {
			t.Errorf("%f Angle between CFE and CEV is not pi/2", data.R)
		}

		if math.Abs(data.DistanceSquared(data.CFV, data.CEV)-math.Pow(math.Cos(math.Pi/data.Q), 2.0)) > TestEps {
			t.Errorf("%f Angle between CFV and CEV is not pi/%f", data.R, data.Q)
		}

		if math.Abs(data.DistanceSquared(data.CFV, data.FEV)) > TestEps {
			t.Errorf("%f Angle between CFV and FEV is not pi/2", data.R)
		}

		if math.Abs(data.DistanceSquared(data.CEV, data.FEV)-math.Pow(math.Cos(math.Pi/data.R), 2.0)) > TestEps {
			t.Errorf("%f Angle between CEV and FEV is not pi/%f, n = %f", data.R, data.R, math.Pi/(math.Acos(math.Sqrt(data.DistanceSquared(data.CEV, data.FEV)))))
		}
	}

}
