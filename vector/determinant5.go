package vector

func Determinant5(mat [5][5]float64) float64 {

	m0 := Determinant4([4][4]float64{
		{mat[1][1], mat[1][2], mat[1][3], mat[1][4]},
		{mat[2][1], mat[2][2], mat[2][3], mat[2][4]},
		{mat[3][1], mat[3][2], mat[3][3], mat[3][4]},
		{mat[4][1], mat[4][2], mat[4][3], mat[4][4]},
	})

	m1 := Determinant4([4][4]float64{
		{mat[1][0], mat[1][2], mat[1][3], mat[1][4]},
		{mat[2][0], mat[2][2], mat[2][3], mat[2][4]},
		{mat[3][0], mat[3][2], mat[3][3], mat[3][4]},
		{mat[4][0], mat[4][2], mat[4][3], mat[4][4]},
	})

	m2 := Determinant4([4][4]float64{
		{mat[1][0], mat[1][1], mat[1][3], mat[1][4]},
		{mat[2][0], mat[2][1], mat[2][3], mat[2][4]},
		{mat[3][0], mat[3][1], mat[3][3], mat[3][4]},
		{mat[4][0], mat[4][1], mat[4][3], mat[4][4]},
	})

	m3 := Determinant4([4][4]float64{
		{mat[1][0], mat[1][1], mat[1][2], mat[1][4]},
		{mat[2][0], mat[2][1], mat[2][2], mat[2][4]},
		{mat[3][0], mat[3][1], mat[3][2], mat[3][4]},
		{mat[4][0], mat[4][1], mat[4][2], mat[4][4]},
	})

	m4 := Determinant4([4][4]float64{
		{mat[1][0], mat[1][1], mat[1][2], mat[1][3]},
		{mat[2][0], mat[2][1], mat[2][2], mat[2][3]},
		{mat[3][0], mat[3][1], mat[3][2], mat[3][3]},
		{mat[4][0], mat[4][1], mat[4][2], mat[4][3]},
	})

	return mat[0][0]*m0 - mat[0][1]*m1 + mat[0][2]*m2 - mat[0][3]*m3 + mat[0][4]*m4

}
