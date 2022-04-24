package vector

import "math"

// a -> b
func DirectionalTangent(a, b Vec4) Vec4 {

	ab := a.Dot(b)

	return Scale4(Diff4(b, Scale4(a, ab)), 1.0/math.Sqrt(1-ab*ab))

}
