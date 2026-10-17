package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7732 evaluates chebyshev collocation node order 7732.
func ComputeChebyshevColloc7732(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
