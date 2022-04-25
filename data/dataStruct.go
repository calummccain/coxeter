package data

import (
	"math"

	"github.com/calummccain/coxeter/vector"
)

const IPEpsilon = 0.00001

type VectorIP struct {
	Vector vector.Vec4
	IP     float64
}

type Cell struct {
	Vertices []vector.Vec4
	Edges    [][]int
	Faces    [][]int
}

var vect = []vector.Vec4{
	{W: 1, X: 0, Y: 0, Z: 0},
	{W: 0, X: 1, Y: 0, Z: 0},
	{W: 0, X: 0, Y: 1, Z: 0},
	{W: 0, X: 0, Y: 0, Z: 1},
}

type UniformHoneycomb struct {
	GoursatTetrahedron GoursatTetrahedron
	G                  VectorIP
	G_VE               VectorIP
	G_VF               VectorIP
	G_VC               VectorIP
	G_EF               VectorIP
	G_EC               VectorIP
	G_FC               VectorIP
	G_VEF              VectorIP
	G_VEC              VectorIP
	G_VFC              VectorIP
	G_EFC              VectorIP
}

type ReflectionMatrices struct {
	A func(v vector.Vec4) vector.Vec4
	B func(v vector.Vec4) vector.Vec4
	C func(v vector.Vec4) vector.Vec4
	D func(v vector.Vec4) vector.Vec4
}

type Distances struct {
	VE float64
	VF float64
	VC float64
	EF float64
	EC float64
	FC float64
}

type Normals struct {
	EFC vector.Vec4
	VFC vector.Vec4
	VEC vector.Vec4
	VEF vector.Vec4
}

type Heights struct {
	V float64
	E float64
	F float64
	C float64
}

type BaseReflections struct {
	V func(vector.Vec4) vector.Vec4
	E func(vector.Vec4) vector.Vec4
	F func(vector.Vec4) vector.Vec4
	C func(vector.Vec4) vector.Vec4
}

type Reflections struct {
	V []Reflect
	E []Reflect
	F []Reflect
	C []Reflect
}

type GoursatTetrahedron struct {
	P               float64
	Q               float64
	R               float64
	V               vector.Vec4
	E               vector.Vec4
	F               vector.Vec4
	C               vector.Vec4
	Distances       Distances
	Normals         Normals
	Heights         Heights
	BaseReflections BaseReflections
	Reflections     Reflections
	Scale           func(vector.Vec4) vector.Vec4
	IP              func(vector.Vec4, vector.Vec4) float64
	Metric          string
}

func (gt *GoursatTetrahedron) DistanceSquared(u, v vector.Vec4) float64 {

	return gt.IP(u, v) * gt.IP(u, v) / (gt.IP(u, u) * gt.IP(v, v))

}

// ProjectToPlane removes the n component of g and returns the result using the inner product from gt
//
// i.e. projects g to the hyperplane with normal n.
//
//  ProjectToPlane(g, n) = g - <g,n> n
func (gt *GoursatTetrahedron) ProjectToPlane(g, n vector.Vec4) vector.Vec4 {

	return vector.Diff4(g, vector.Scale4(n, gt.IP(g, n)))

}

func (gt *GoursatTetrahedron) ReflectionGenerator(n vector.Vec4) func(vector.Vec4) vector.Vec4 {

	return func(v vector.Vec4) vector.Vec4 {
		return vector.Diff4(v, vector.Scale4(n, 2.0*gt.IP(v, n)/gt.IP(n, n)))
	}

}

func (gt *GoursatTetrahedron) ProjectToLine(d, n, m vector.Vec4) vector.Vec4 {

	//n = gt.Normalise(n)
	//m = gt.Normalise(m)

	nm := gt.IP(n, m)
	dn := gt.IP(d, n)
	dm := gt.IP(d, m)

	return vector.Diff4(d, vector.Sum4(vector.Scale4(n, (dn-nm*dm)/(1-nm*nm)), vector.Scale4(m, (dm-nm*dn)/(1-nm*nm))))

}

// ScaleVEFC normalises the Goursat Tetrahedron's vertices with respect to the inner product.
func (gt *GoursatTetrahedron) ScaleVEFC() {

	gt.V = gt.Scale(gt.V)
	gt.E = gt.Scale(gt.E)
	gt.F = gt.Scale(gt.F)
	gt.C = gt.Scale(gt.C)

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

func (gt *GoursatTetrahedron) GenerateInnerProduct() {

	if gt.Metric == "s" {
		gt.IP = vector.Dot4
	}
	// TODO Other metrics
}

func (gt *GoursatTetrahedron) Populate() {

	gt.GenerateInnerProduct()

	gt.ScaleVEFC()

	gt.Distances = Distances{
		VE: gt.IP(gt.V, gt.E),
		VF: gt.IP(gt.V, gt.F),
		VC: gt.IP(gt.V, gt.C),
		EF: gt.IP(gt.E, gt.F),
		EC: gt.IP(gt.E, gt.C),
		FC: gt.IP(gt.F, gt.C),
	}

	// TODO Directional tangent only valid for spherical
	gt.Normals.EFC = vector.DirectionalTangent(gt.E, gt.V)
	gt.Normals.VFC = vector.Scale4(
		vector.Diff4(
			vector.Scale4(gt.E, 1.0-gt.Distances.VF*gt.Distances.VF),
			vector.Sum4(
				vector.Scale4(gt.V, gt.Distances.VE*(1.0-gt.Distances.EF*gt.Distances.EF)),
				vector.Scale4(gt.F, gt.Distances.EF*(1.0-gt.Distances.VE*gt.Distances.VE)),
			),
		), 1.0/math.Sqrt((1.0-gt.Distances.VF*gt.Distances.VF)*(1.0-gt.Distances.EF*gt.Distances.EF)*(1.0-gt.Distances.VE*gt.Distances.VE)),
	)
	gt.Normals.VEC = vector.Scale4(
		vector.Diff4(
			vector.Scale4(gt.F, 1.0-gt.Distances.EC*gt.Distances.EC),
			vector.Sum4(
				vector.Scale4(gt.C, gt.Distances.FC*(1.0-gt.Distances.EF*gt.Distances.EF)),
				vector.Scale4(gt.E, gt.Distances.EF*(1.0-gt.Distances.FC*gt.Distances.FC)),
			),
		), 1.0/math.Sqrt((1.0-gt.Distances.EC*gt.Distances.EC)*(1.0-gt.Distances.EF*gt.Distances.EF)*(1.0-gt.Distances.FC*gt.Distances.FC)),
	)
	gt.Normals.VEF = vector.DirectionalTangent(gt.F, gt.C)

	gt.Heights.V = math.Sqrt(1.0 - gt.Distances.VE*gt.Distances.VE)
	gt.Heights.E = math.Sqrt(1.0-gt.Distances.EF*gt.Distances.EF) * math.Sqrt(1.0-gt.Distances.VE*gt.Distances.VE) / math.Sqrt(1.0-gt.Distances.VF*gt.Distances.VF)
	gt.Heights.F = math.Sqrt(1.0-gt.Distances.EF*gt.Distances.EF) * math.Sqrt(1.0-gt.Distances.FC*gt.Distances.FC) / math.Sqrt(1.0-gt.Distances.EC*gt.Distances.EC)
	gt.Heights.C = math.Sqrt(1.0 - gt.Distances.FC*gt.Distances.FC)

	gt.BaseReflections.V = gt.ReflectionGenerator(gt.Normals.EFC)
	gt.BaseReflections.E = gt.ReflectionGenerator(gt.Normals.VFC)
	gt.BaseReflections.F = gt.ReflectionGenerator(gt.Normals.VEC)
	gt.BaseReflections.C = gt.ReflectionGenerator(gt.Normals.VEF)

	gt.Reflections.V = gt.EnumerateReflections([]Reflect{{Word: "", Matrix: vect}}, []string{"e", "f", "c"})
	gt.Reflections.E = gt.EnumerateReflections([]Reflect{{Word: "", Matrix: vect}}, []string{"v", "f", "c"})
	gt.Reflections.F = gt.EnumerateReflections([]Reflect{{Word: "", Matrix: vect}}, []string{"v", "e", "c"})
	gt.Reflections.C = gt.EnumerateReflections([]Reflect{{Word: "", Matrix: vect}}, []string{"v", "e", "f"})

}

func (gt *GoursatTetrahedron) UniformHoneycombGenerator(v, e, f, c float64) UniformHoneycomb {

	uf := UniformHoneycomb{
		GoursatTetrahedron: *gt,
	}

	G := vector.Scale4(gt.V, v/gt.Heights.V)
	G.Sum(vector.Scale4(gt.E, e/gt.Heights.E))
	G.Sum(vector.Scale4(gt.F, f/gt.Heights.F))
	G.Sum(vector.Scale4(gt.C, c/gt.Heights.C))

	uf.G = VectorIP{
		Vector: gt.Normalise(G),
	}

	uf.G.IP = gt.IP(uf.G.Vector, uf.G.Vector)

	uf.G_VEF.Vector = gt.Normalise(gt.ProjectToPlane(G, gt.Normals.VEF))
	uf.G_VEF.IP = gt.IP(uf.G_VEF.Vector, uf.G.Vector)

	uf.G_VE.Vector = gt.Normalise(gt.ProjectToLine(G, gt.Normals.VEF, gt.Normals.VEC))
	uf.G_VE.IP = gt.IP(uf.G_VE.Vector, uf.G.Vector)

	uf.G_VFC.Vector = gt.Normalise(gt.ProjectToPlane(G, gt.Normals.VFC))
	uf.G_VFC.IP = gt.IP(uf.G_VFC.Vector, uf.G.Vector)

	uf.G_VF.Vector = gt.Normalise(gt.ProjectToLine(G, gt.Normals.VEF, gt.Normals.VFC))
	uf.G_VF.IP = gt.IP(uf.G_VF.Vector, uf.G.Vector)

	uf.G_VEC.Vector = gt.Normalise(gt.ProjectToPlane(G, gt.Normals.VEC))
	uf.G_VEC.IP = gt.IP(uf.G_VEC.Vector, uf.G.Vector)

	uf.G_VC.Vector = gt.Normalise(gt.ProjectToLine(G, gt.Normals.VEC, gt.Normals.VFC))
	uf.G_VC.IP = gt.IP(uf.G_VC.Vector, uf.G.Vector)

	uf.G_EFC.Vector = gt.Normalise(gt.ProjectToPlane(G, gt.Normals.EFC))
	uf.G_EFC.IP = gt.IP(uf.G_EFC.Vector, uf.G.Vector)

	uf.G_EF.Vector = gt.Normalise(gt.ProjectToLine(G, gt.Normals.VEF, gt.Normals.EFC))
	uf.G_EF.IP = gt.IP(uf.G_EF.Vector, uf.G.Vector)

	uf.G_EC.Vector = gt.Normalise(gt.ProjectToLine(G, gt.Normals.VEC, gt.Normals.EFC))
	uf.G_EC.IP = gt.IP(uf.G_EC.Vector, uf.G.Vector)

	uf.G_FC.Vector = gt.Normalise(gt.ProjectToLine(G, gt.Normals.VFC, gt.Normals.EFC))
	uf.G_FC.IP = gt.IP(uf.G_FC.Vector, uf.G.Vector)

	return uf

}

func (uh *UniformHoneycomb) GenerateCells() []Cell {

	cells := []Cell{}

	viableCells := [][]Reflect{}
	viableFaces := []VectorIP{}
	viableEdges := []VectorIP{}

	if math.Abs(uh.GoursatTetrahedron.IP(uh.G.Vector, uh.GoursatTetrahedron.V)-1.0) > IPEpsilon {
		viableCells = append(viableCells, uh.GoursatTetrahedron.Reflections.V)
	}

	if math.Abs(uh.GoursatTetrahedron.IP(uh.G.Vector, uh.GoursatTetrahedron.E)-1.0) > IPEpsilon {
		viableCells = append(viableCells, uh.GoursatTetrahedron.Reflections.E)
	}

	if math.Abs(uh.GoursatTetrahedron.IP(uh.G.Vector, uh.GoursatTetrahedron.F)-1.0) > IPEpsilon {
		viableCells = append(viableCells, uh.GoursatTetrahedron.Reflections.F)
	}

	if math.Abs(uh.GoursatTetrahedron.IP(uh.G.Vector, uh.GoursatTetrahedron.C)-1.0) > IPEpsilon {
		viableCells = append(viableCells, uh.GoursatTetrahedron.Reflections.C)
	}

	for _, f := range []VectorIP{uh.G_VE, uh.G_VF, uh.G_VC, uh.G_EF, uh.G_EC, uh.G_FC} {
		if math.Abs(f.IP-1.0) > IPEpsilon {
			viableFaces = append(viableFaces, f)
		}
	}

	for _, e := range []VectorIP{uh.G_VEF, uh.G_VEC, uh.G_VFC, uh.G_EFC} {
		if math.Abs(e.IP-1.0) > IPEpsilon {
			viableEdges = append(viableEdges, e)
		}
	}

	var testVector vector.Vec4

	for _, reflections := range viableCells {

		verts := []vector.Vec4{}
		faces := []VectorIP{}
		edges := []VectorIP{}

		faceVectors := []vector.Vec4{}
		edgeVectors := []vector.Vec4{}

		for _, m := range reflections {

			mat := vector.InitialiseFromVectors(m.Matrix)
			mat = mat.Transpose()

			testVector = mat.MatrixByVector(uh.G.Vector)
			if !vector.IsInArray4(testVector, verts) {
				verts = append(verts, testVector)
			}

			for _, f := range viableFaces {
				testVector = mat.MatrixByVector(f.Vector)
				if !vector.IsInArray4(testVector, faceVectors) {
					faceVectors = append(faceVectors, testVector)
					faces = append(faces, VectorIP{Vector: testVector, IP: f.IP})
				}
			}

			for _, e := range viableEdges {
				testVector = mat.MatrixByVector(e.Vector)
				if !vector.IsInArray4(testVector, edgeVectors) {
					edgeVectors = append(edgeVectors, testVector)
					edges = append(edges, VectorIP{Vector: testVector, IP: e.IP})
				}
			}

		}

		newCell := Cell{
			Vertices: verts,
		}

		var subFace []int
		for _, f := range faces {
			subFace = []int{}
			for j, v := range newCell.Vertices {
				if math.Abs(uh.GoursatTetrahedron.IP(f.Vector, v)-f.IP) < IPEpsilon {
					subFace = append(subFace, j)
				}
			}
			if len(subFace) > 2 {
				newCell.Faces = append(newCell.Faces, subFace)
			}
		}

		var subEdge []int
		for _, e := range edges {
			subEdge = []int{}
			for j, v := range newCell.Vertices {
				if math.Abs(uh.GoursatTetrahedron.IP(e.Vector, v)-e.IP) < IPEpsilon {
					subEdge = append(subEdge, j)
				}
			}
			if len(subEdge) > 1 {
				newCell.Edges = append(newCell.Edges, subEdge)
			}
		}

		if len(newCell.Faces) > 3 {
			cells = append(cells, newCell)
		}
	}

	return cells

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
