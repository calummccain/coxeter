package vector

func NearestPointOnPlaneToPoint(a, n Vec4) Vec4 {
	an := a.Dot(n)
	return Diff4(a, Scale4(n, an))
}
