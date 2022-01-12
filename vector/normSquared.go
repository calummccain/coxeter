package vector

// func NormSquared2(vec [2]float64) float64 {

// 	return vec[0]*vec[0] + vec[1]*vec[1]

// }

// func NormSquared3(vec [3]float64) float64 {

// 	return vec[0]*vec[0] + vec[1]*vec[1] + vec[2]*vec[2]

// }

// func NormSquared4(vec [4]float64) float64 {

// 	return vec[0]*vec[0] + vec[1]*vec[1] + vec[2]*vec[2] + vec[3]*vec[3]

// }

func NormSquared(vec []float64) float64 {

	d := 0.0

	for i := 0; i < len(vec); i++ {

		d += vec[i] * vec[i]

	}

	return d

}
