package vector

func DistanceSquared(vec1, vec2 []float64) float64 {

	return NormSquared(Diff(vec1, vec2))

}
