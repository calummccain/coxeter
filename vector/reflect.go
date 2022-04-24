package vector

func Reflect(x, n Vec4) Vec4 {
	return Diff4(x, Scale4(n, 2.0*x.Dot(n)))
}
