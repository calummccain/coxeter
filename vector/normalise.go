package vector

func Normalise2(vec [2]float64) [2]float64 {

	return Scale2(vec, 1.0/Norm2(vec))

}

func Normalise3(vec [3]float64) [3]float64 {

	return Scale3(vec, 1.0/Norm3(vec))

}

func Normalise4(vec [4]float64) [4]float64 {

	return Scale4(vec, 1.0/Norm4(vec))

}

func Normalise(vec []float64) []float64 {

	return Scale(vec, 1.0/Norm(vec))

}
