package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8742 evaluates chebyshev collocation node order 8742.
func ComputeChebyshevColloc8742(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
