package vector

func PointInPolygon(point [3]float64, polygon [][3]float64) bool {

	var v0, v1, v2 [3]float64
	var dot01, dot12, dot20, dot11, dot22 float64

	for i := 1; i < len(polygon)-1; i++ {

		v0 = Diff3(polygon[0], point)
		v1 = Diff3(polygon[i], point)
		v2 = Diff3(polygon[i+1], point)

		dot01 = Dot3(v0, v1)
		dot12 = Dot3(v1, v2)
		dot20 = Dot3(v2, v0)
		dot11 = Dot3(v1, v1)
		dot22 = Dot3(v2, v2)

		if (dot12*dot20 > dot22*dot01) && (dot01*dot12 > dot11*dot20) {

			return true

		}

	}

	return false

}
