package vector

type Mat2 struct {
	XX float64
	XY float64
	YX float64
	YY float64
}

func (m *Mat2) Determinant() float64 {
	return m.XX*m.YY - m.XY*m.YX
}

func (m *Mat2) Scale(a float64) {
	m.XX *= a
	m.XY *= a
	m.YX *= a
	m.YY *= a
}

func (m *Mat2) Sum(n Mat2) {
	m.XX += n.XX
	m.XY += n.XY
	m.YX += n.YX
	m.YY += n.YY
}

func Determinant2(m Mat2) float64 {
	return m.XX*m.YY - m.XY*m.YX
}
