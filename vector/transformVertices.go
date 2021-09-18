package vector

import "coxeter/shared"

func TransformVertices(baseVertices [][4]float64, transformation string, matrices shared.Matrices) [][4]float64 {

	newVertices := [][4]float64{}
	e1 := [4]float64{1, 0, 0, 0}
	e2 := [4]float64{0, 1, 0, 0}
	e3 := [4]float64{0, 0, 1, 0}
	e4 := [4]float64{0, 0, 0, 1}

	m := map[byte]func([4]float64) [4]float64{
		'a': matrices.A,
		'b': matrices.B,
		'c': matrices.C,
		'd': matrices.D,
		'e': matrices.E,
		'f': matrices.F,
	}

	if transformation != "" {

		i := len(transformation) - 1
		for i > -1 {

			e1 = m[transformation[i]](e1)
			e2 = m[transformation[i]](e2)
			e3 = m[transformation[i]](e3)
			e4 = m[transformation[i]](e4)
			i--

		}

	}

	newVertex := [4]float64{0, 0, 0, 0}

	for j := 0; j < len(baseVertices); j++ {

		newVertex = Sum4(Scale4(e1, baseVertices[j][0]), Scale4(e2, baseVertices[j][1]))
		newVertex = Sum4(newVertex, Scale4(e3, baseVertices[j][2]))
		newVertex = Sum4(newVertex, Scale4(e4, baseVertices[j][3]))

		newVertices = append(newVertices, newVertex)

	}

	return newVertices

}
