package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7462 evaluates chebyshev collocation node order 7462.
func ComputeChebyshevColloc7462(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
