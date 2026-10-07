package chebyshev_collocation

import "math"

// ComputeChebyshevColloc532 evaluates chebyshev collocation node order 532.
func ComputeChebyshevColloc532(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
