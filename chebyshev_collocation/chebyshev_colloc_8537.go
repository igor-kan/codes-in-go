package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8537 evaluates chebyshev collocation node order 8537.
func ComputeChebyshevColloc8537(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
