package go_module_phytagoras

import "math"

const (
	powerOfTwo = 2
)

func FindC(a float64, b float64) float64 {
	aSqr := math.Pow(a, powerOfTwo)
	bSqr := math.Pow(b, powerOfTwo)
	if a < 1.0 || b < 1.0 {
		return 0.0
	}
	return math.Sqrt(aSqr + bSqr)
}
