package vector

import "math"

func Circumradius(a, b, c float64) float64 {

	return (a * b * c) / math.Sqrt(math.Abs((a+b+c)*(b+c-a)*(c+a-b)*(a+b-c)))

}
