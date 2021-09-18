package vector

func Determinant3(mat [3][3]float64) float64 {

	m1 := [2][2]float64{{mat[1][1], mat[1][2]}, {mat[2][1], mat[2][2]}}
	m2 := [2][2]float64{{mat[0][1], mat[0][2]}, {mat[2][1], mat[2][2]}}
	m3 := [2][2]float64{{mat[0][1], mat[0][2]}, {mat[1][1], mat[1][2]}}

	return mat[0][0]*Determinant2(m1) - mat[1][0]*Determinant2(m2) + mat[2][0]*Determinant2(m3)

}
