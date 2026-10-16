package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7627 evaluates chebyshev collocation node order 7627.
func ComputeChebyshevColloc7627(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
