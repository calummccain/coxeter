package vector

// func Cross3(vec1, vec2 [3]float64) [3]float64 {

// 	return [3]float64{vec1[1]*vec2[2] - vec1[2]*vec2[1], vec1[2]*vec2[0] - vec1[0]*vec2[2], vec1[0]*vec2[1] - vec1[1]*vec2[0]}

// }

// func Cross4(vec1, vec2, vec3 [4]float64) [4]float64 {

// 	m := [3][4]float64{vec1, vec2, vec3}

// 	return [4]float64{
// 		Determinant3([3][3]float64{
// 			{m[0][1], m[0][2], m[0][3]},
// 			{m[1][1], m[1][2], m[1][3]},
// 			{m[2][1], m[2][2], m[2][3]}}),
// 		-Determinant3([3][3]float64{
// 			{m[0][0], m[0][2], m[0][3]},
// 			{m[1][0], m[1][2], m[1][3]},
// 			{m[2][0], m[2][2], m[2][3]}}),
// 		Determinant3([3][3]float64{
// 			{m[0][0], m[0][1], m[0][3]},
// 			{m[1][0], m[1][1], m[1][3]},
// 			{m[2][0], m[2][1], m[2][3]}}),
// 		-Determinant3([3][3]float64{
// 			{m[0][0], m[0][1], m[0][2]},
// 			{m[1][0], m[1][1], m[1][2]},
// 			{m[2][0], m[2][1], m[2][2]}})}
// }
