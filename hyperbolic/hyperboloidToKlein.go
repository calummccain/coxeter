package hyperbolic

func HyperboloidToKlein(vec [4]float64) [3]float64 {

	inv := 1.0 / vec[0]

	return [3]float64{vec[1] * inv, vec[2] * inv, vec[3] * inv}

}
