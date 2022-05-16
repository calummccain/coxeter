package vector

import "math"

// a -> b
func SDirectionalTangent(a, b Vec4) Vec4 {

	ab := a.Dot(b)

	return Scale4(Diff4(b, Scale4(a, ab)), 1.0/math.Sqrt(1-ab*ab))

}

// a -> b
func EDirectionalTangent(a, b Vec4) Vec4 {

	v := Diff4(b, a)
	v.Normalise()

	return v

}

// a -> b
func HDirectionalTangent(a, b Vec4) Vec4 {

	ab := a.HDot(b)

	return Scale4(Diff4(b, Scale4(a, ab)), 1.0/math.Sqrt(ab*ab-1))

}
