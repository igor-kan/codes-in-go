package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8507 evaluates chebyshev collocation node order 8507.
func ComputeChebyshevColloc8507(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
