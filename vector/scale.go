package vector

func Scale1(vec [1]float64, a float64) [1]float64 {

	return [1]float64{a * vec[0]}

}

func Scale2(vec [2]float64, a float64) [2]float64 {

	return [2]float64{a * vec[0], a * vec[1]}

}

func Scale3(vec [3]float64, a float64) [3]float64 {

	return [3]float64{a * vec[0], a * vec[1], a * vec[2]}

}

func Scale4(vec [4]float64, a float64) [4]float64 {

	return [4]float64{a * vec[0], a * vec[1], a * vec[2], a * vec[3]}

}

func Scale(vec []float64, a float64) []float64 {

	d := make([]float64, 0, len(vec))

	for i := 0; i < len(vec); i++ {

		d = append(d, a*vec[i])

	}

	return d

}
