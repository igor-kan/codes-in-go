package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8977 evaluates chebyshev collocation node order 8977.
func ComputeChebyshevColloc8977(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
