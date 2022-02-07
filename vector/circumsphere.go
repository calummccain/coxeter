package vector

import (
	"math"
)

func Circumsphere2(u, v, w Vec2) Sphere2 {

	nu := u.NormSquared()
	nv := v.NormSquared()
	nw := w.NormSquared()

	a := 2 * Determinant3(Mat3{u.X, u.Y, 1, v.X, v.Y, 1, w.X, w.Y, 1})
	a = 1.0 / a
	g := Determinant3(Mat3{nu, u.X, u.Y, nv, v.X, v.Y, nw, w.X, w.Y}) * a

	dx := Determinant3(Mat3{nu, u.Y, 1, nv, v.Y, 1, nw, w.Y, 1}) * a
	dy := -Determinant3(Mat3{nu, u.X, 1, nv, v.X, 1, nw, w.X, 1}) * a

	return Sphere2{Vec2{dx, dy}, math.Sqrt((dx*dx + dy*dy - 2*g))}

}

func Circumsphere3(u, v, w, x Vec3) Sphere3 {

	nu := u.NormSquared()
	nv := v.NormSquared()
	nw := w.NormSquared()
	nx := x.NormSquared()

	a := 2 * Determinant4(Mat4{u.X, u.Y, u.Z, 1, v.X, v.Y, v.Z, 1, w.X, w.Y, w.Z, 1, x.X, x.Y, x.Z, 1})
	a = 1.0 / a
	g := Determinant4(Mat4{nu, u.X, u.Y, u.Z, nv, v.X, v.Y, v.Z, nw, w.X, w.Y, w.Z, nx, x.X, x.Y, x.Z}) * a

	dx := Determinant4(Mat4{nu, u.Y, u.Z, 1, nv, v.Y, v.Z, 1, nw, w.Y, w.Z, 1, nx, x.Y, x.Z, 1}) * a
	dy := -Determinant4(Mat4{nu, u.X, u.Z, 1, nv, v.X, v.Z, 1, nw, w.X, w.Z, 1, nx, x.X, x.Z, 1}) * a
	dz := Determinant4(Mat4{nu, u.X, u.Y, 1, nv, v.X, v.Y, 1, nw, w.X, w.Y, 1, nx, x.X, x.Y, 1}) * a

	return Sphere3{Vec3{dx, dy, dz}, math.Sqrt((dx*dx + dy*dy + dz*dz - 2*g))}

}
