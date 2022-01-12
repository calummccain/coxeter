package vector

import "math"

type Vec4 struct {
	W float64
	X float64
	Y float64
	Z float64
}

func (v *Vec4) NormSquared() float64 {
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

func (v *Vec4) ToSlice() [4]float64 {
	return [4]float64{v.W, v.X, v.Y, v.Z}
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

func Midpoint4(v, w Vec4) Vec4 {
	return Scale4(Sum4(v, w), 0.5)
}

func NormSquared4(v Vec4) float64 {
	return v.W*v.W + v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func DistanceSquared4(v, w Vec4) float64 {
	return NormSquared4(Diff4(v, w))
}

func Cross4(vec1, vec2, vec3 Vec4) Vec4 {
	m := [3][4]float64{vec1.ToSlice(), vec2.ToSlice(), vec3.ToSlice()}
	return Vec4{
		Determinant3(Mat3{
			m[0][1], m[0][2], m[0][3],
			m[1][1], m[1][2], m[1][3],
			m[2][1], m[2][2], m[2][3]}),
		-Determinant3(Mat3{
			m[0][0], m[0][2], m[0][3],
			m[1][0], m[1][2], m[1][3],
			m[2][0], m[2][2], m[2][3]}),
		Determinant3(Mat3{
			m[0][0], m[0][1], m[0][3],
			m[1][0], m[1][1], m[1][3],
			m[2][0], m[2][1], m[2][3]}),
		-Determinant3(Mat3{
			m[0][0], m[0][1], m[0][2],
			m[1][0], m[1][1], m[1][2],
			m[2][0], m[2][1], m[2][2]}),
	}
}
