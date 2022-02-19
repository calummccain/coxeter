package vector

func Coplanar(v1, v2, v3 Vec3) Plane {
	mat := IniitialiseMat3(v1, v2, v3)
	normal := Cross3(Diff3(v2, v1), Diff3(v3, v1))
	normal.Normalise()

	return Plane{
		Normal: normal,
		D:      mat.Determinant(),
	}
}
