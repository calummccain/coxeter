package vector

import "math"

func EllipseFinder(v0, v1, v2, v3, v4 [2]float64) (a, b, c, d, e, f float64) {

	c0 := [5]float64{v0[0] * v0[0], v1[0] * v1[0], v2[0] * v2[0], v3[0] * v3[0], v4[0] * v4[0]}
	c1 := [5]float64{v0[0] * v0[1], v1[0] * v1[1], v2[0] * v2[1], v3[0] * v3[1], v4[0] * v4[1]}
	c2 := [5]float64{v0[1] * v0[1], v1[1] * v1[1], v2[1] * v2[1], v3[1] * v3[1], v4[1] * v4[1]}
	c3 := [5]float64{v0[0], v1[0], v2[0], v3[0], v4[0]}
	c4 := [5]float64{v0[1], v1[1], v2[1], v3[1], v4[1]}
	c5 := [5]float64{1, 1, 1, 1, 1}

	a = Determinant5([5][5]float64{c1, c2, c3, c4, c5})
	b = -Determinant5([5][5]float64{c0, c2, c3, c4, c5})
	c = Determinant5([5][5]float64{c0, c1, c3, c4, c5})
	d = -Determinant5([5][5]float64{c0, c1, c2, c4, c5})
	e = Determinant5([5][5]float64{c0, c1, c2, c3, c5})
	f = -Determinant5([5][5]float64{c0, c1, c2, c3, c4})

	if math.Signbit(a) {

		a = -a
		b = -b
		c = -c
		d = -d
		e = -e
		f = -f

	}

	return a, b, c, d, e, f

}
