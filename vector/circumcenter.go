// package vector

// func Circumcenter(u, v, w [2]float64) [2]float64 {

// 	a := -0.5 / Determinant3([3][3]float64{{u[0], u[1], 1}, {v[0], v[1], 1}, {w[0], w[1], 1}})
// 	ru := u[0]*u[0] + u[1]*u[1]
// 	rv := v[0]*v[0] + v[1]*v[1]
// 	rw := w[0]*w[0] + w[1]*w[1]
// 	bx := -Determinant3([3][3]float64{{ru, u[1], 1}, {rv, v[1], 1}, {rw, w[1], 1}})
// 	by := Determinant3([3][3]float64{{ru, u[0], 1}, {rv, v[0], 1}, {rw, w[0], 1}})

// 	return [2]float64{bx * a, by * a}

// }
