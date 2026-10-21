package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8547 evaluates chebyshev collocation node order 8547.
func ComputeChebyshevColloc8547(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
