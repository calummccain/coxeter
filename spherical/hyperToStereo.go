package spherical

import (
	"coxeter/vector"
)

func HyperToStereo(vec [4]float64) [3]float64 {

	return vector.Scale3([3]float64{vec[1], vec[2], vec[3]}, 1.0/(1.0-vec[0]))

}
