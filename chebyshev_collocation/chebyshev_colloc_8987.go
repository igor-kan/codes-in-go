package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8987 evaluates chebyshev collocation node order 8987.
func ComputeChebyshevColloc8987(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
