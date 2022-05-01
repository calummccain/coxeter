package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

func eValpqr(p, q, r float64) float64 {

	sp := math.Sin(math.Pi / p)
	cq := math.Cos(math.Pi / q)

	eVal := 0.0

	if (p-2.0)*(q-2.0) < 4.0 {

		eVal = math.Pi / math.Asin(cq/sp)

	} else {

		eVal = 2.0

	}

	return eVal

}

func pValpqr(p, q, r float64) float64 {

	return 1.0 / (0.5 - 1.0/q)

}

func GoursatTetrahedronpqr(p, q, r float64) GoursatTetrahedron {

	cp := math.Cos(math.Pi / p)
	sp := math.Sin(math.Pi / p)

	cq := math.Cos(math.Pi / q)
	//sq := math.Sin(math.Pi / float64(q))

	cr := math.Cos(math.Pi / r)
	sr := math.Sin(math.Pi / r)

	gt := GoursatTetrahedron{
		P:    p,
		Q:    q,
		R:    r,
		V:    vector.Vec4{W: 1, X: 0, Y: cp, Z: sp},
		E:    vector.Vec4{W: 1, X: 0, Y: cp, Z: 0},
		F:    vector.Vec4{W: 1, X: 0, Y: 0, Z: 0},
		C:    vector.Vec4{W: 1, X: cp * cq / (sp * cr), Y: 0, Z: 0},
		EVal: eValpqr(p, q, r),
		PVal: pValpqr(p, q, r),
	}

	gt.Metric = Boundaries(r, gt.EVal, gt.PVal)

	if gt.Metric == "p" {
		gt.Scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{W: v.W, X: v.X, Y: v.Y, Z: v.Z}
		}
	} else {
		den := sp * math.Sqrt(math.Abs(sr*sr-cq*cq))
		gt.Scale = func(v vector.Vec4) vector.Vec4 {
			return vector.Vec4{
				W: cp * cq * v.W / den,
				X: math.Sqrt(math.Abs(sp*sp*sr*sr-cq*cq)) * v.X / den,
				Y: math.Sqrt(math.Abs(sp*sp*sr*sr-cq*cq)) * v.Y / den,
				Z: math.Sqrt(math.Abs(sp*sp*sr*sr-cq*cq)) * v.Z / den,
			}
		}
	}

	gt.Populate()

	return gt

}
