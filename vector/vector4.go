package vector

import "math"

type Vec4 struct {
	W float64
	X float64
	Y float64
	Z float64
}

func (v Vec4) NormSquared() float64 {
	return v.W*v.W + v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vec4) Scale(a float64) {
	v.W *= a
	v.X *= a
	v.Y *= a
	v.Z *= a
}

func (v *Vec4) Sum(w Vec4) {
	v.W += w.W
	v.X += w.X
	v.Y += w.Y
	v.Z += w.Z
}

func (v *Vec4) Diff(w Vec4) {
	v.W -= w.W
	v.X -= w.X
	v.Y -= w.Y
	v.Z -= w.Z
}

func (v *Vec4) Normalise() {
	v.Scale(1 / math.Sqrt(v.NormSquared()))
}

func (v *Vec4) Dot(w Vec4) float64 {
	return v.W*w.W + v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

func Scale4(v Vec4, a float64) Vec4 {
	return Vec4{v.W * a, v.X * a, v.Y * a, v.Z * a}
}

func Sum4(v, w Vec4) Vec4 {
	return Vec4{v.W + w.W, v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

func Diff4(v, w Vec4) Vec4 {
	return Vec4{v.W - w.W, v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

func Dot4(v, w Vec4) float64 {
	return v.W*w.W + v.X*w.X + v.Y*w.Y + v.Z*w.Z
}
