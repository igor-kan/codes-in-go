package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8837 evaluates chebyshev collocation node order 8837.
func ComputeChebyshevColloc8837(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
