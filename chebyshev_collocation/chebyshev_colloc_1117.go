package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1117 evaluates chebyshev collocation node order 1117.
func ComputeChebyshevColloc1117(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
