package vector

func TransformVertices(baseVertices []Vec4, transformation string, A, B, C, D, E, F func(Vec4) Vec4) []Vec4 {

	newVertices := []Vec4{}
	e1 := Vec4{1, 0, 0, 0}
	e2 := Vec4{0, 1, 0, 0}
	e3 := Vec4{0, 0, 1, 0}
	e4 := Vec4{0, 0, 0, 1}

	m := map[byte]func(Vec4) Vec4{
		'a': A,
		'b': B,
		'c': C,
		'd': D,
		'e': E,
		'f': F,
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

	newVertex := Vec4{0, 0, 0, 0}

	for j := 0; j < len(baseVertices); j++ {

		newVertex = Sum4(Scale4(e1, baseVertices[j].W), Scale4(e2, baseVertices[j].X))
		newVertex = Sum4(newVertex, Scale4(e3, baseVertices[j].Y))
		newVertex = Sum4(newVertex, Scale4(e4, baseVertices[j].Z))

		newVertices = append(newVertices, newVertex)

	}

	return newVertices

}
