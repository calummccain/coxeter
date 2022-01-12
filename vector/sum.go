package vector

func Sum1(vec1, vec2 [1]float64) [1]float64 {

	return [1]float64{vec1[0] + vec2[0]}

}

// func Sum2(vec1, vec2 [2]float64) [2]float64 {

// 	return [2]float64{vec1[0] + vec2[0], vec1[1] + vec2[1]}

// }

// func Sum3(vec1, vec2 [3]float64) [3]float64 {

// 	return [3]float64{vec1[0] + vec2[0], vec1[1] + vec2[1], vec1[2] + vec2[2]}

// }

// func Sum4(vec1, vec2 [4]float64) [4]float64 {

// 	return [4]float64{vec1[0] + vec2[0], vec1[1] + vec2[1], vec1[2] + vec2[2], vec1[3] + vec2[3]}

// }

func Sum(vec1, vec2 []float64) []float64 {

	d := make([]float64, 0, len(vec1))

	for i := 0; i < len(vec1); i++ {

		d = append(d, vec1[i]+vec2[i])

	}

	return d

}
