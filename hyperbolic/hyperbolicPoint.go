package hyperbolic

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

type HPoint struct {
	H    vector.Vec4
	K    vector.Vec3
	P    vector.Vec3
	U    vector.Vec3
	Norm float64
}

func InitHPoint(w, x, y, z float64) HPoint {
	return HPoint{H: vector.Vec4{W: w, X: x, Y: y, Z: z}}
}

func InitHPointVec4(v vector.Vec4) HPoint {
	return HPoint{H: v}
}

func Scale3(p vector.Vec3, a float64) vector.Vec3 {
	return vector.Vec3{X: p.X * a, Y: p.Y * a, Z: p.Z * a}
}

func Scale4(p vector.Vec4, a float64) vector.Vec4 {
	return vector.Vec4{W: p.W * a, X: p.X * a, Y: p.Y * a, Z: p.Z * a}
}

func Sum4(p, q vector.Vec4) vector.Vec4 {
	return vector.Vec4{W: p.W + q.W, X: p.X + q.X, Y: p.Y + q.Y, Z: p.Z + q.Z}
}

func Diff4(p, q vector.Vec4) vector.Vec4 {
	return vector.Vec4{W: p.W - q.W, X: p.X - q.X, Y: p.Y - q.Y, Z: p.Z - q.Z}
}

func (p *HPoint) HyperboloidInnerProduct(q HPoint) float64 {
	return p.H.W*q.H.W - p.H.X*q.H.X - p.H.Y*q.H.Y - p.H.Z*q.H.Z
}

func (p *HPoint) HyperbolicNorm() {
	p.Norm = p.H.W*p.H.W - p.H.X*p.H.X - p.H.Y*p.H.Y - p.H.Z*p.H.Z
}

func (p *HPoint) Normalise() {
	p.H.Scale(1.0 / math.Sqrt(math.Abs(p.HyperboloidInnerProduct(*p))))
}

func (p *HPoint) HyperboloidToKlein() {
	inv := 1.0 / p.H.W
	p.K.X = p.H.X * inv
	p.K.Y = p.H.Y * inv
	p.K.Z = p.H.Z * inv
}

func (p *HPoint) HyperboloidToPoincare() {
	eps := HyperboloidToPoincareEps
	var inv float64
	if math.Abs(p.Norm) < eps {
		inv = 1.0 / p.H.W
	} else if p.Norm > eps {
		inv = 1.0 / (1.0 + p.H.W)
	} else {
		inv = 1.0 / (p.H.X*p.H.X + p.H.Y*p.H.Y + p.H.Z*p.H.Z)
	}
	p.P.X = p.H.X * inv
	p.P.Y = p.H.Y * inv
	p.P.Z = p.H.Z * inv
}

func (p *HPoint) HyperboloidToUHP() {
	eps := HyperboloidToUHPEps
	if math.Abs(p.H.W-p.H.Z) < eps {
		p.U.X = p.H.X / p.H.W
		p.U.Y = p.H.Y / p.H.W
		p.U.Z = math.Inf(1)
	} else if p.Norm < eps {
		p.U.X = p.H.X / (p.H.W - p.H.Z)
		p.U.Y = p.H.Y / (p.H.W - p.H.Z)
		p.U.Z = 0.0
	} else {
		p.U.X = p.H.X / (p.H.W - p.H.Z)
		p.U.Y = p.H.Y / (p.H.W - p.H.Z)
		p.U.Z = 1.0 / (p.H.W - p.H.Z)
	}
}

func (p *HPoint) KleinToHyperboloid() {
	inv := 1.0 / (1.0 - p.K.NormSquared())
	p.H.W = inv
	p.H.X = p.K.X * inv
	p.H.Y = p.K.Y * inv
	p.H.Z = p.K.Z * inv
}

func (p *HPoint) KleinToPoincare() {
	eps := KleinToPoincareEps
	if p.K.NormSquared() < 1-eps {
		p.P = p.K
	} else {
		p.P = Scale3(p.K, 1.0/(1.0+math.Sqrt(1.0-p.K.NormSquared())))
	}
}

func (p *HPoint) KleinToUHP() {
	p.U = Scale3(vector.Vec3{X: p.K.X, Y: p.K.Y, Z: math.Sqrt(1.0 - p.K.NormSquared())}, 1.0/(1.0-p.K.Z))
}

func (p *HPoint) PoincareToHyperboloid() {
	eps := PoincareToHyperboloidEps
	r := p.P.NormSquared()
	if math.Abs(r-1) < eps {
		p.H = vector.Vec4{W: 1, X: p.P.X, Y: p.P.Y, Z: p.P.Z}
	} else {
		p.H = Scale4(vector.Vec4{W: (1.0 + r) * 0.5, X: p.P.X, Y: p.P.Y, Z: p.P.Z}, 2.0/(1.0-r))
	}
}

func (p *HPoint) PoincareToKlein() {
	p.K = Scale3(p.P, 2.0/(1.0+p.P.NormSquared()))
}

func (p *HPoint) PoincareToUHP() {
	eps := PoincareToUHPEps
	r := p.P.NormSquared()
	s := 1 / (r + 1.0 - 2.0*p.P.Z)

	if s < eps {
		p.U = vector.Vec3{X: p.P.X, Y: p.P.Y, Z: math.Inf(1)}
	} else if r > 1-eps {
		p.U = Scale3(vector.Vec3{X: p.P.X, Y: p.P.Y, Z: 0}, 2.0*s)
	} else {
		p.U = Scale3(vector.Vec3{X: p.P.X, Y: p.P.Y, Z: (1.0 - r) * 0.5}, 2.0*s)
	}
}

func (p *HPoint) UHPToHyperboloid() {
	eps := UHPToHyperboloidEps
	r := p.U.NormSquared()
	if p.U.Z < eps {
		p.H = vector.Vec4{W: (r + 1.0) * 0.5, X: p.U.X, Y: p.U.Y, Z: (r - 1.0) * 0.5}
	} else {
		p.H = Scale4(vector.Vec4{W: (r + 1.0) * 0.5, X: p.U.X, Y: p.U.Y, Z: (r - 1.0) * 0.5}, 1.0/p.U.Z)
	}
}

func (p *HPoint) UHPToKlein() {
	r := p.U.NormSquared()
	p.K = Scale3(vector.Vec3{X: p.U.X, Y: p.U.Y, Z: (r - 1.0) * 0.5}, 2.0/(r+1.0))
}

func (p *HPoint) UHPToPoincare() {
	r := p.U.NormSquared()
	p.P = Scale3(vector.Vec3{X: p.U.X, Y: p.U.Y, Z: (r - 1.0) * 0.5}, 2.0/(r+1.0+2.0*p.U.Z))
}
