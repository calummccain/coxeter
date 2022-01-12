package vector

func Midpoint1(vec1, vec2 [1]float64) [1]float64 {

	return Scale1(Sum1(vec1, vec2), 0.5)

}

// func Midpoint2(vec1, vec2 [2]float64) [2]float64 {

// 	return Scale2(Sum2(vec1, vec2), 0.5)

// }

// func Midpoint3(vec1, vec2 [3]float64) [3]float64 {

// 	return Scale3(Sum3(vec1, vec2), 0.5)

// }

// func Midpoint4(vec1, vec2 [4]float64) [4]float64 {

// 	return Scale4(Sum4(vec1, vec2), 0.5)

// }

func Midpoint(vec1, vec2 []float64) []float64 {

	return Scale(Sum(vec1, vec2), 0.5)

}
