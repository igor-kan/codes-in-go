package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7622 evaluates chebyshev collocation node order 7622.
func ComputeChebyshevColloc7622(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
