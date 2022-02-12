package main

import (
	"fmt"
	"math"
)

func main() {
	const max = 10

	var cp, sp, cq, cr, sr, cs, ss float64
	var aa, ba, cb, dc float64
	for p := 3; p < max; p++ {
		for q := 3; q < max; q++ {
			for r := 3; r < max; r++ {
				for s := 3; s < max; s++ {
					cp = math.Pow(math.Cos(math.Pi/float64(p)), 2.0)
					cq = math.Pow(math.Cos(math.Pi/float64(q)), 2.0)
					cr = math.Pow(math.Cos(math.Pi/float64(r)), 2.0)
					cs = math.Pow(math.Cos(math.Pi/float64(s)), 2.0)
					sp = math.Pow(math.Sin(math.Pi/float64(p)), 2.0)
					//sq = math.Pow(math.Sin(math.Pi/float64(q)), 2.0)
					sr = math.Pow(math.Sin(math.Pi/float64(r)), 2.0)
					ss = math.Pow(math.Sin(math.Pi/float64(s)), 2.0)

					aa = (sp*(sr-cs) - cq*ss) / (cs * (sp - cq))
					ba = (sp*sr - cq) / (sp * cr)
					cb = (sp - cq) / cq
					dc = sp / cp

					if (aa > 0.0) && (ba > 0.0) && (cb > 0.0) && (dc > 0.0) {
						fmt.Printf("{%d,%d,%d,%d}\n", p, q, r, s)
					}

				}
			}
		}
	}

	for p := 3; p < max; p++ {
		for q := 3; q < max; q++ {
			for r := 3; r < max; r++ {
				cp = math.Pow(math.Cos(math.Pi/float64(p)), 2.0)
				cq = math.Pow(math.Cos(math.Pi/float64(q)), 2.0)
				cr = math.Pow(math.Cos(math.Pi/float64(r)), 2.0)
				sp = math.Pow(math.Sin(math.Pi/float64(p)), 2.0)
				//sq = math.Pow(math.Sin(math.Pi/float64(q)), 2.0)
				sr = math.Pow(math.Sin(math.Pi/float64(r)), 2.0)

				aa = (sr*sp - cq) / (sp * cr)
				ba = (sp - cq) / cq
				cb = sp / cp

				if (aa > 0.0) && (ba > 0.0) && (cb > 0.0) {
					fmt.Printf("{%d,%d,%d}\n", p, q, r)
				}
			}
		}
	}
}
