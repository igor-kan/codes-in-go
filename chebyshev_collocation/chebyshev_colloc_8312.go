package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8312 evaluates chebyshev collocation node order 8312.
func ComputeChebyshevColloc8312(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
