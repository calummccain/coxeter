package vector

func Diff(vec1, vec2 []float64) []float64 {

	d := make([]float64, 0, len(vec1))

	for i := 0; i < len(vec1); i++ {

		d = append(d, vec1[i]-vec2[i])

	}

	return d

}
