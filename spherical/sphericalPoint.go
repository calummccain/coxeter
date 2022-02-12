package spherical

import (
	"github.com/calummccain/coxeter/vector"
)

type sPoint struct {
	H vector.Vec4
	S vector.Vec3
}

func InitSPoint(w, x, y, z float64) sPoint {
	return sPoint{H: vector.Vec4{W: w, X: x, Y: y, Z: z}}
}

func (p *sPoint) HyperToStereo() {

	p.S = vector.Vec3{X: p.H.X, Y: p.H.Y, Z: p.H.Z}
	p.S.Scale(1.0 / (1.0 - p.H.W))

}

func (p *sPoint) StereoToHyper() {

	r := p.S.NormSquared()

	p.H = vector.Vec4{W: (r - 1.0) * 0.5, X: p.S.X, Y: p.S.Y, Z: p.S.Z}
	p.H.Scale(2.0 / (r + 1.0))

}

func (p sPoint) HyperboloidInnerProduct(q sPoint) float64 {
	return p.H.W*q.H.W + p.H.X*q.H.X + p.H.Y*q.H.Y + p.H.Z*q.H.Z
}
