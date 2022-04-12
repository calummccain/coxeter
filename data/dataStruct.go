package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

type UniformHoneycomb struct {
	G     vector.Vec4
	G_VE  vector.Vec4
	G_VF  vector.Vec4
	G_VC  vector.Vec4
	G_EF  vector.Vec4
	G_EC  vector.Vec4
	G_FC  vector.Vec4
	G_VEF vector.Vec4
	G_VEC vector.Vec4
	G_VFC vector.Vec4
	G_EFC vector.Vec4
}

type ReflectionMatrices struct {
	A func(v vector.Vec4) vector.Vec4
	B func(v vector.Vec4) vector.Vec4
	C func(v vector.Vec4) vector.Vec4
	D func(v vector.Vec4) vector.Vec4
}

type GoursatTetrahedron struct {
	P               float64
	Q               float64
	R               float64
	V               vector.Vec4
	E               vector.Vec4
	F               vector.Vec4
	C               vector.Vec4
	Scale           func(vector.Vec4) vector.Vec4
	IP              func(vector.Vec4, vector.Vec4) float64
	Matrices        ReflectionMatrices
	FaceReflections []string
}

// ProjectToPlane removes the n component of g and returns the result
// i.e. projects g to the hyperplane with normal n.
//  ProjectToPlane(g, n) = g - <g,n>/<n,n>n
func (gt *GoursatTetrahedron) ProjectToPlane(g, n vector.Vec4) vector.Vec4 {

	return vector.Diff4(g, vector.Scale4(n, gt.IP(g, n)/gt.IP(n, n)))

}

// ScaleVEFC normalises the Goursat Tetrahedron's vertices with respect to the inner product.
func (gt *GoursatTetrahedron) ScaleVEFC() {

	gt.V = gt.Normalise(gt.V)
	gt.E = gt.Normalise(gt.E)
	gt.F = gt.Normalise(gt.F)
	gt.C = gt.Normalise(gt.C)

}

// Normalise normalises the vector using the inner product from the Goursat Tetrahedron.
//
//  Normalise(n) = n/sqrt(|<n,n>|)
func (gt *GoursatTetrahedron) Normalise(v vector.Vec4) vector.Vec4 {

	return vector.Scale4(v, 1.0/math.Sqrt(math.Abs(gt.IP(v, v))))

}

func (gt *GoursatTetrahedron) UniformHoneycombGenerator(v, e, f, c float64) UniformHoneycomb {

	gt.ScaleVEFC()

	uf := UniformHoneycomb{}

	sum := v + e + f + c
	// general point G = (vV + eE + fF + cC)/(v + e + f + c)
	G := vector.Scale4(gt.V, v/sum)
	G.Sum(vector.Scale4(gt.E, e/sum))
	G.Sum(vector.Scale4(gt.F, f/sum))
	G.Sum(vector.Scale4(gt.C, c/sum))

	uf.G = gt.Normalise(G)

	projV, projE, projF, projC := gt.GramSchmidt(gt.V, gt.E, gt.F, gt.C)

	uf.G_VEF = gt.Normalise(gt.ProjectToPlane(G, projC))
	uf.G_VE = gt.Normalise(gt.ProjectToPlane(uf.G_VEF, projF))

	projV, projF, projC, projE = gt.GramSchmidt(gt.V, gt.F, gt.C, gt.E)

	uf.G_VFC = gt.Normalise(gt.ProjectToPlane(G, projE))
	uf.G_VF = gt.Normalise(gt.ProjectToPlane(uf.G_VFC, projC))

	projV, projC, projE, projF = gt.GramSchmidt(gt.V, gt.C, gt.E, gt.F)

	uf.G_VEC = gt.Normalise(gt.ProjectToPlane(G, projF))
	uf.G_VC = gt.Normalise(gt.ProjectToPlane(uf.G_VEC, projE))

	projE, projF, projC, projV = gt.GramSchmidt(gt.E, gt.F, gt.C, gt.V)

	uf.G_EFC = gt.Normalise(gt.ProjectToPlane(G, projV))
	uf.G_EF = gt.Normalise(gt.ProjectToPlane(uf.G_EFC, projC))

	projE, projC, projF, projV = gt.GramSchmidt(gt.E, gt.C, gt.F, gt.V)

	uf.G_EC = gt.Normalise(gt.ProjectToPlane(gt.ProjectToPlane(G, projV), projF))

	projF, projC, projE, projV = gt.GramSchmidt(gt.F, gt.C, gt.E, gt.V)

	uf.G_FC = gt.Normalise(gt.ProjectToPlane(gt.ProjectToPlane(G, projV), projE))

	return uf

}

// GramSchmidt runs the Gram-Schmidt algorithm on the four vectors w, x, y, z with the inner product given by gt.
//
// Returns:
//  a = w
//  b = x - <x,a>a
//  c = y - <y,a>a - <y,b>b
//  d = z - <z,a>a - <z,b>b - <z,c>c
func (gt *GoursatTetrahedron) GramSchmidt(w, x, y, z vector.Vec4) (vector.Vec4, vector.Vec4, vector.Vec4, vector.Vec4) {

	a := w
	b := gt.ProjectToPlane(x, a)
	c := gt.ProjectToPlane(gt.ProjectToPlane(y, a), b)
	d := gt.ProjectToPlane(gt.ProjectToPlane(gt.ProjectToPlane(z, a), b), c)

	return a, b, c, d
}
