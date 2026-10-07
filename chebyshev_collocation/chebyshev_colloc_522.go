package chebyshev_collocation

import "math"

// ComputeChebyshevColloc522 evaluates chebyshev collocation node order 522.
func ComputeChebyshevColloc522(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
