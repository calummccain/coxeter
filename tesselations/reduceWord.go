package tesselations

import (
	"strings"
)

func ReduceWord(word string, p, q, r int) string {

	done := false
	w1 := word
	w2 := word

	reduced := [16][2]string{
		{"aa", ""},
		{"bb", ""},
		{"cc", ""},
		{"dd", ""},
		{"aca", "c"},
		{"cac", "a"},
		{"ada", "d"},
		{"dad", "a"},
		{"bdb", "d"},
		{"dbd", "b"},
		{"", ""},
		{"", ""},
		{"", ""},
		{"", ""},
		{"", ""},
		{"", ""},
	}

	if p%2 == 0 {
		reduced[10] = [2]string{strings.Repeat("ab", (p+1)/2) + "a", strings.Repeat("ba", (p+1)/2-1) + "b"}
		reduced[11] = [2]string{strings.Repeat("ba", (p+1)/2) + "b", strings.Repeat("ab", (p+1)/2-1) + "a"}
	} else {
		reduced[10] = [2]string{strings.Repeat("ab", (p+1)/2), strings.Repeat("ba", (p+1)/2-1)}
		reduced[11] = [2]string{strings.Repeat("ba", (p+1)/2), strings.Repeat("ab", (p+1)/2-1)}
	}

	if q%2 == 0 {
		reduced[12] = [2]string{strings.Repeat("bc", (q+1)/2) + "b", strings.Repeat("cb", (q+1)/2-1) + "c"}
		reduced[13] = [2]string{strings.Repeat("cb", (q+1)/2) + "c", strings.Repeat("bc", (q+1)/2-1) + "b"}
	} else {
		reduced[12] = [2]string{strings.Repeat("bc", (q+1)/2), strings.Repeat("cb", (q+1)/2-1)}
		reduced[13] = [2]string{strings.Repeat("cb", (q+1)/2), strings.Repeat("bc", (q+1)/2-1)}
	}

	if r%2 == 0 {
		reduced[14] = [2]string{strings.Repeat("cd", (r+1)/2) + "c", strings.Repeat("dc", (r+1)/2-1) + "d"}
		reduced[15] = [2]string{strings.Repeat("dc", (r+1)/2) + "d", strings.Repeat("cd", (r+1)/2-1) + "c"}
	} else {
		reduced[14] = [2]string{strings.Repeat("cd", (r+1)/2), strings.Repeat("dc", (r+1)/2-1)}
		reduced[15] = [2]string{strings.Repeat("dc", (r+1)/2), strings.Repeat("cd", (r+1)/2-1)}
	}

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

func partialReduce(newWord string, reduced [16][2]string) string {

	for i := 0; i < 16; i++ {

		newWord = strings.Replace(newWord, reduced[i][0], reduced[i][1], -1)

	}

	return newWord

}
