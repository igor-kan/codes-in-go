package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8777 evaluates chebyshev collocation node order 8777.
func ComputeChebyshevColloc8777(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
