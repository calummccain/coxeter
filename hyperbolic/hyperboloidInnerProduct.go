package hyperbolic

import "github.com/calummccain/coxeter/vector"

func HyperboloidInnerProduct(vec1, vec2 vector.Vec4) float64 {

	return vec1.W*vec2.W - vec1.X*vec2.X - vec1.Y*vec2.Y - vec1.Z*vec2.Z

}
