package vector

import (
	"math"
)

func Circum4(u, v, w, x [3]float64) ([3]float64, float64) {

	nu := Dot3(u, u)
	nv := Dot3(v, v)
	nw := Dot3(w, w)
	nx := Dot3(x, x)

	a := 2 * Determinant4([4][4]float64{{u[0], u[1], u[2], 1}, {v[0], v[1], v[2], 1}, {w[0], w[1], w[2], 1}, {x[0], x[1], x[2], 1}})
	g := Determinant4([4][4]float64{{nu, u[0], u[1], u[2]}, {nv, v[0], v[1], v[2]}, {nw, w[0], w[1], w[2]}, {nx, x[0], x[1], x[2]}})

	dx := Determinant4([4][4]float64{{nu, u[1], u[2], 1}, {nv, v[1], v[2], 1}, {nw, w[1], w[2], 1}, {nx, x[1], x[2], 1}})
	dy := -Determinant4([4][4]float64{{nu, u[0], u[2], 1}, {nv, v[0], v[2], 1}, {nw, w[0], w[2], 1}, {nx, x[0], x[2], 1}})
	dz := Determinant4([4][4]float64{{nu, u[0], u[1], 1}, {nv, v[0], v[1], 1}, {nw, w[0], w[1], 1}, {nx, x[0], x[1], 1}})

	return [3]float64{dx / a, dy / a, dz / a}, math.Sqrt((dx*dx + dy*dy + dz*dz - 2*a*g)) / math.Abs(a)

}
