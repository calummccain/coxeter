package vector

// XX XY XZ
// YX YY YZ
// ZX ZY ZZ
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

func InitialiseMat3(u, v, w Vec3) Mat3 {
	return Mat3{
		XX: u.X,
		XY: u.Y,
		XZ: u.Z,
		YX: v.X,
		YY: v.Y,
		YZ: v.Z,
		ZX: w.X,
		ZY: w.Y,
		ZZ: w.Z,
	}
}

func (m *Mat3) Determinant() float64 {
	return m.XX*Determinant2(Mat2{m.YY, m.YZ, m.ZY, m.ZZ}) - m.XY*Determinant2(Mat2{m.YX, m.YZ, m.ZX, m.ZZ}) + m.XZ*Determinant2(Mat2{m.YX, m.YY, m.ZX, m.ZY})
}

func Determinant3(m Mat3) float64 {
	return m.XX*Determinant2(Mat2{m.YY, m.YZ, m.ZY, m.ZZ}) - m.XY*Determinant2(Mat2{m.YX, m.YZ, m.ZX, m.ZZ}) + m.XZ*Determinant2(Mat2{m.YX, m.YY, m.ZX, m.ZY})
}
