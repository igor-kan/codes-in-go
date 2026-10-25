package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8892 evaluates chebyshev collocation node order 8892.
func ComputeChebyshevColloc8892(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
