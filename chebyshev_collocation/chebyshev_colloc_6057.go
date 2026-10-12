package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6057 evaluates chebyshev collocation node order 6057.
func ComputeChebyshevColloc6057(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
