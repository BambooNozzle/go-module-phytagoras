package go_module_phytagoras

import "math"

const (
	powerOfTwo = 2
)

func FindC(a float64, b float64) float64 {
	aSqr := math.Pow(a, powerOfTwo)
	bSqr := math.Pow(b, powerOfTwo)
	return math.Sqrt(aSqr * bSqr)
}
