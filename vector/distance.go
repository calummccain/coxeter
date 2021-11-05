package vector

func Distance3(vec1, vec2 [3]float64) float64 {

	return Norm3(Diff3(vec1, vec2))

}

func Distance(vec1, vec2 []float64) float64 {

	return Norm(Diff(vec1, vec2))

}
