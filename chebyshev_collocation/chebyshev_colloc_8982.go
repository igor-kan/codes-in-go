package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8982 evaluates chebyshev collocation node order 8982.
func ComputeChebyshevColloc8982(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
