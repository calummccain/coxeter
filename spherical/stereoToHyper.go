package spherical

import (
	"coxeter/vector"
)

func StereoToHyper(vec [3]float64) [4]float64 {

	r := vector.NormSquared(vec[:])

	return vector.Scale4([4]float64{(r - 1.0) * 0.5, vec[0], vec[1], vec[2]}, 2.0/(r+1.0))

}
