package vector

func IsInArray(testVector [4]float64, groupVectors [][4]float64) bool {

	eps := IsInArrayEps

	for i := 0; i < len(groupVectors); i++ {

		if DistanceSquared(testVector[:], groupVectors[i][:]) < eps {

			return true

		}
	}

	return false

}

func IsInArray2(testVector [2]int, groupVectors [][2]int) bool {

	for i := 0; i < len(groupVectors); i++ {

		if groupVectors[i][0] == testVector[0] && groupVectors[i][1] == testVector[1] {

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
