package spherical

func InnerProduct(vec1, vec2 [4]float64) float64 {

	return vec1[0]*vec2[0] + vec1[1]*vec2[1] + vec1[2]*vec2[2] + vec1[3]*vec2[3]

}
