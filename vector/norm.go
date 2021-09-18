package vector

import "math"

func Norm2(vec [2]float64) float64 {

	return math.Sqrt(NormSquared2(vec))

}

func Norm3(vec [3]float64) float64 {

	return math.Sqrt(NormSquared3(vec))

}

func Norm4(vec [4]float64) float64 {

	return math.Sqrt(NormSquared4(vec))

}

func Norm(vec []float64) float64 {

	return math.Sqrt(NormSquared(vec))

}
