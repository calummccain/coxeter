package vector

func IsInArray(testVector Vec4, groupVectors []Vec4) bool {

	eps := IsInArrayEps

	for i := 0; i < len(groupVectors); i++ {

		if DistanceSquared4(testVector, groupVectors[i]) < eps {

			return true

		}
	}

	return false

}

func IsInArray2(testVector Vec2, groupVectors []Vec2) bool {

	for i := 0; i < len(groupVectors); i++ {

		if groupVectors[i].X == testVector.X && groupVectors[i].Y == testVector.Y {

			return true

		}
	}

	return false

}

func IsInArray1(value int, list []int) bool {

	for i := 0; i < len(list); i++ {

		if list[i] == value {

			return true

		}
	}

	return false

}
