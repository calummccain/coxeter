package vector

import "math"

func Norm2(vec [2]float64) float64 {

	return math.Sqrt(NormSquared2(vec))

}

func Norm3(vec Vec3) float64 {

	return math.Sqrt(vec.NormSquared())

}

func Norm4(vec Vec4) float64 {

	return math.Sqrt(vec.NormSquared())

}

func Norm(vec []float64) float64 {

	return math.Sqrt(NormSquared(vec))

}
