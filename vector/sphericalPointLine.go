package vector

// NearestPointOnLineToPoint returns the point on the line through a and b to the point c in S^3

func NearestPointOnLineToPoint(a, b, c Vec4) Vec4 {

	ac := a.Dot(c)
	ab := a.Dot(b)
	bc := b.Dot(c)

	return Sum4(Scale4(a, (ac-ab*bc)/(1-ab*ab)), Scale4(b, (bc-ab*ac)/(1-ab*ab)))

}
