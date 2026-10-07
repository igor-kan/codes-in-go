package chebyshev_collocation

import "math"

// ComputeChebyshevColloc507 evaluates chebyshev collocation node order 507.
func ComputeChebyshevColloc507(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
