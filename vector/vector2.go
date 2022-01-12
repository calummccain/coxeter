package vector

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (v Vec2) NormSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vec2) Scale(a float64) {
	v.X *= a
	v.Y *= a
}

func (v *Vec2) Sum(w Vec2) {
	v.X += w.X
	v.Y += w.Y
}

func (v *Vec2) Diff(w Vec2) {
	v.X -= w.X
	v.Y -= w.Y
}

func (v *Vec2) Normalise() {
	v.Scale(1 / math.Sqrt(v.NormSquared()))
}

func (v *Vec2) Dot(w Vec2) float64 {
	return v.X*w.X + v.Y*w.Y
}

func (v *Vec2) Cross(w Vec2) float64 {
	return v.X*w.Y - v.Y*w.X
}

func Scale2(v Vec2, a float64) Vec2 {
	return Vec2{v.X * a, v.Y * a}
}

func Sum2(v, w Vec2) Vec2 {
	return Vec2{v.X + w.X, v.Y + w.Y}
}

func Diff2(v, w Vec2) Vec2 {
	return Vec2{v.X - w.X, v.Y - w.Y}
}

func Dot2(v, w Vec2) float64 {
	return v.X*w.X + v.Y*w.Y
}

func Midpoint2(v, w Vec2) Vec2 {
	return Scale2(Sum2(v, w), 0.5)
}

func Cross2(v, w Vec2) float64 {
	return v.X*w.Y - v.Y*w.X
}
