package vector

type Mat3 struct {
	XX float64
	XY float64
	XZ float64
	YX float64
	YY float64
	YZ float64
	ZX float64
	ZY float64
	ZZ float64
}

func (m *Mat3) Determinant() float64 {
	return m.XX*Determinant2(Mat2{m.YY, m.YZ, m.ZY, m.ZZ}) - m.XY*Determinant2(Mat2{m.XY, m.XZ, m.ZY, m.ZZ}) + m.XZ*Determinant2(Mat2{m.XY, m.XZ, m.YY, m.YZ})
}

func Determinant3(m Mat3) float64 {
	return m.XX*Determinant2(Mat2{m.YY, m.YZ, m.ZY, m.ZZ}) - m.XY*Determinant2(Mat2{m.XY, m.XZ, m.ZY, m.ZZ}) + m.XZ*Determinant2(Mat2{m.XY, m.XZ, m.YY, m.YZ})
}
