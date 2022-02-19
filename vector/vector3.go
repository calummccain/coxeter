package vector

import "math"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v Vec3) Norm() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) NormSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vec3) Scale(a float64) {
	v.X *= a
	v.Y *= a
	v.Z *= a
}

func (v *Vec3) Sum(w Vec3) {
	v.X += w.X
	v.Y += w.Y
	v.Z += w.Z
}

func (v *Vec3) Diff(w Vec3) {
	v.X -= w.X
	v.Y -= w.Y
	v.Z -= w.Z
}

func (v *Vec3) Normalise() {
	v.Scale(1 / math.Sqrt(v.NormSquared()))
}

func (v *Vec3) Dot(w Vec3) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

func (v *Vec3) Cross(w Vec3) Vec3 {
	return Vec3{v.Y*w.Z - v.Z*w.Y, v.Z*w.X - v.X*w.Z, v.X*w.Y - v.Y*w.X}
}

func Scale3(v Vec3, a float64) Vec3 {
	return Vec3{v.X * a, v.Y * a, v.Z * a}
}

func Sum3(v, w Vec3) Vec3 {
	return Vec3{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

func Diff3(v, w Vec3) Vec3 {
	return Vec3{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

func Dot3(v, w Vec3) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

func Midpoint3(v, w Vec3) Vec3 {
	return Scale3(Sum3(v, w), 0.5)
}

func Cross3(v, w Vec3) Vec3 {
	return Vec3{v.Y*w.Z - v.Z*w.Y, v.Z*w.X - v.X*w.Z, v.X*w.Y - v.Y*w.X}
}

func Norm3(v Vec3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func NormSquared3(v Vec3) float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func Distance3(v, w Vec3) float64 {
	return math.Sqrt(NormSquared3(Diff3(v, w)))
}

func DistanceSquared3(v, w Vec3) float64 {
	return NormSquared3(Diff3(v, w))
}
