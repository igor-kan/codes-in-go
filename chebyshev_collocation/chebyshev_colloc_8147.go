package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8147 evaluates chebyshev collocation node order 8147.
func ComputeChebyshevColloc8147(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
