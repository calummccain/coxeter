package data

import (
	"strings"

	"github.com/calummccain/coxeter/vector"
)

const MAXLENGTH = 1000

func ReduceWord(word string, p, q, r int) string {

	done := false
	w1 := word
	w2 := word

	reduced := [10][2]string{
		{"aa", ""},
		{"bb", ""},
		{"cc", ""},
		{"dd", ""},
		{"ca", "ac"},
		{"da", "ad"},
		{"db", "bd"},
		{"", ""},
		{"", ""},
		{"", ""},
	}

	if p%2 == 0 {
		reduced[7] = [2]string{strings.Repeat("ab", p/2), strings.Repeat("ba", p/2)}
	} else {
		reduced[7] = [2]string{strings.Repeat("ab", p/2) + "a", strings.Repeat("ba", p/2) + "b"}
	}

	if q%2 == 0 {
		reduced[8] = [2]string{strings.Repeat("bc", q/2), strings.Repeat("cb", q/2)}
	} else {
		reduced[8] = [2]string{strings.Repeat("bc", q/2) + "b", strings.Repeat("cb", q/2) + "c"}
	}

	if r%2 == 0 {
		reduced[9] = [2]string{strings.Repeat("cd", r/2), strings.Repeat("dc", r/2)}
	} else {
		reduced[9] = [2]string{strings.Repeat("cd", r/2) + "c", strings.Repeat("dc", r/2) + "d"}
	}

	//fmt.Println(reduced)

	for !done {

		w2 = partialReduce(w1, reduced)

		if w1 == w2 {

			done = true

		}

		w1 = w2

		continue

	}

	return w1

}

func partialReduce(newWord string, reduced [10][2]string) string {

	for i := 0; i < 10; i++ {

		newWord = strings.Replace(newWord, reduced[i][0], reduced[i][1], -1)

	}

	return newWord

}

type Reflect struct {
	Word   string
	Matrix []vector.Vec4
}

func InStringSlice(word string, slice []string) bool {

	for _, w := range slice {
		if word == w {
			return true
		}
	}
	return false
}

/*
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
*/

func sameMatrix(u, v []vector.Vec4) bool {
	sum := 0.0
	for i, w := range u {
		sum += vector.DistanceSquared4(w, v[i])
	}
	return sum < 0.001
}

func sameMatrixArray(u Reflect, arr []Reflect) bool {
	for _, m := range arr {
		if sameMatrix(u.Matrix, m.Matrix) {
			return true
		}
	}
	return false
}

func (gt *GoursatTetrahedron) EnumerateReflections(reflections []Reflect, faceReflections []string) []Reflect {

	var new Reflect
	var newRef []Reflect

	for _, w := range faceReflections {
		for _, v := range reflections {
			new = Reflect{
				Word: w + v.Word,
				Matrix: vector.TransformVertices(
					v.Matrix,
					w,
					gt.BaseReflections.V,
					gt.BaseReflections.E,
					gt.BaseReflections.F,
					gt.BaseReflections.C,
				),
			}
			if !sameMatrixArray(new, reflections) && !sameMatrixArray(new, newRef) {
				newRef = append(newRef, new)
			}
		}
	}

	// TODO Refactor MAXLENGTH to appear at the beginning of loop
	if len(newRef) > 0 && len(reflections) < MAXLENGTH {
		return gt.EnumerateReflections(append(reflections, newRef...), faceReflections)
	} else {
		return reflections
	}

}
