package vector

func Distance(vec1, vec2 []float64) float64 {

	return Norm(Diff(vec1, vec2))

}
