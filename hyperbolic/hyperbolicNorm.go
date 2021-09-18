package hyperbolic

func HyperbolicNorm(vec [4]float64) float64 {

	return vec[0]*vec[0] - vec[1]*vec[1] - vec[2]*vec[2] - vec[3]*vec[3]

}
