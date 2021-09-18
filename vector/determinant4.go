package vector

func Determinant4(mat [4][4]float64) float64 {

	m1 := [3][3]float64{{mat[1][1], mat[1][2], mat[1][3]}, {mat[2][1], mat[2][2], mat[2][3]}, {mat[3][1], mat[3][2], mat[3][3]}}
	m2 := [3][3]float64{{mat[1][0], mat[1][2], mat[1][3]}, {mat[2][0], mat[2][2], mat[2][3]}, {mat[3][0], mat[3][2], mat[3][3]}}
	m3 := [3][3]float64{{mat[1][0], mat[1][1], mat[1][3]}, {mat[2][0], mat[2][1], mat[2][3]}, {mat[3][0], mat[3][1], mat[3][3]}}
	m4 := [3][3]float64{{mat[1][0], mat[1][1], mat[1][2]}, {mat[2][0], mat[2][1], mat[2][2]}, {mat[3][0], mat[3][1], mat[3][2]}}

	return mat[0][0]*Determinant3(m1) - mat[0][1]*Determinant3(m2) + mat[0][2]*Determinant3(m3) - mat[0][3]*Determinant3(m4)

}
