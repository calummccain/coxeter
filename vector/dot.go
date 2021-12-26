package vector

func Dot2(vec1, vec2 [2]float64) float64 {

	return vec1[0]*vec2[0] + vec1[1]*vec2[1]

}

// func Dot3(vec1, vec2 [3]float64) float64 {

// 	return vec1[0]*vec2[0] + vec1[1]*vec2[1] + vec1[2]*vec2[2]

// }

// func Dot4(vec1, vec2 [4]float64) float64 {

// 	return vec1[0]*vec2[0] + vec1[1]*vec2[1] + vec1[2]*vec2[2] + vec1[3]*vec2[3]

// }

func Dot(vec1, vec2 []float64) float64 {

	d := 0.0

	for i := 0; i < len(vec1); i++ {

		d += vec1[i] * vec2[i]

	}

	return d
}
