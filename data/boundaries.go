package data

func Boundaries(n, e, p float64) string {

	eps := BoundaryEps

	if n <= 2.0 {

		return "x"

	} else if n <= e-eps {

		return "s"

	} else if n <= e+eps {

		return "e"

	} else if n <= p-eps {

		return "h"

	} else if n <= p+eps {

		return "p"

	} else {

		return "u"

	}

}
