package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8767 evaluates chebyshev collocation node order 8767.
func ComputeChebyshevColloc8767(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
