package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8832 evaluates chebyshev collocation node order 8832.
func ComputeChebyshevColloc8832(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
