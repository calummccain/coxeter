package vector

//
// | XX XY |
// | YX YY |
//

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

func InitialiseFromVectors(v []Vec4) Mat4 {
	m := Mat4{}

	m.WW = v[0].W
	m.WX = v[1].W
	m.WY = v[2].W
	m.WZ = v[3].W

	m.XW = v[0].X
	m.XX = v[1].X
	m.XY = v[2].X
	m.XZ = v[3].X

	m.YW = v[0].Y
	m.YX = v[1].Y
	m.YY = v[2].Y
	m.YZ = v[3].Y

	m.ZW = v[0].Z
	m.ZX = v[1].Z
	m.ZY = v[2].Z
	m.ZZ = v[3].Z

	return m

}

func (m *Mat4) Transpose() Mat4 {
	n := Mat4{}

	n.WW = m.WW
	n.WX = m.XW
	n.WY = m.YW
	n.WZ = m.ZW

	n.XW = m.WX
	n.XX = m.XX
	n.XY = m.YX
	n.XZ = m.ZX

	n.YW = m.WY
	n.YX = m.XY
	n.YY = m.YY
	n.YZ = m.ZY

	n.ZW = m.WZ
	n.ZX = m.XZ
	n.ZY = m.YZ
	n.ZZ = m.ZZ

	return n

}

func (m *Mat4) MatrixByVector(v Vec4) Vec4 {
	w := Vec4{
		W: m.WW*v.W + m.WX*v.X + m.WY*v.Y + m.WZ*v.Z,
		X: m.XW*v.W + m.XX*v.X + m.XY*v.Y + m.XZ*v.Z,
		Y: m.YW*v.W + m.YX*v.X + m.YY*v.Y + m.YZ*v.Z,
		Z: m.ZW*v.W + m.ZX*v.X + m.ZY*v.Y + m.ZZ*v.Z,
	}
	return w
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
