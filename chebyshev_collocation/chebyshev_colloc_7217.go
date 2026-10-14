package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7217 evaluates chebyshev collocation node order 7217.
func ComputeChebyshevColloc7217(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
