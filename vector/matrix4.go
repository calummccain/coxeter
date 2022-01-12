package vector

type Mat4 struct {
	WW float64
	WX float64
	WY float64
	WZ float64
	XW float64
	XX float64
	XY float64
	XZ float64
	YW float64
	YX float64
	YY float64
	YZ float64
	ZW float64
	ZX float64
	ZY float64
	ZZ float64
}

func (m *Mat4) Determinant() float64 {
	m1 := Determinant3(Mat3{m.XX, m.XY, m.XZ, m.YX, m.YY, m.YZ, m.ZX, m.ZY, m.ZZ})
	m2 := Determinant3(Mat3{m.XW, m.XY, m.XZ, m.YW, m.YY, m.YZ, m.ZW, m.ZY, m.ZZ})
	m3 := Determinant3(Mat3{m.XW, m.XX, m.XZ, m.YW, m.YX, m.YZ, m.ZW, m.ZX, m.ZZ})
	m4 := Determinant3(Mat3{m.XW, m.XX, m.XY, m.YW, m.YX, m.YY, m.ZW, m.ZX, m.ZY})
	return m.WW*m1 - m.WX*m2 + m.WY*m3 - m.WZ*m4
}

func Determinant4(m Mat4) float64 {
	m1 := Determinant3(Mat3{m.XX, m.XY, m.XZ, m.YX, m.YY, m.YZ, m.ZX, m.ZY, m.ZZ})
	m2 := Determinant3(Mat3{m.XW, m.XY, m.XZ, m.YW, m.YY, m.YZ, m.ZW, m.ZY, m.ZZ})
	m3 := Determinant3(Mat3{m.XW, m.XX, m.XZ, m.YW, m.YX, m.YZ, m.ZW, m.ZX, m.ZZ})
	m4 := Determinant3(Mat3{m.XW, m.XX, m.XY, m.YW, m.YX, m.YY, m.ZW, m.ZX, m.ZY})
	return m.WW*m1 - m.WX*m2 + m.WY*m3 - m.WZ*m4
}
