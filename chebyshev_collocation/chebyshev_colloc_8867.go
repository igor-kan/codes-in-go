package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8867 evaluates chebyshev collocation node order 8867.
func ComputeChebyshevColloc8867(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
