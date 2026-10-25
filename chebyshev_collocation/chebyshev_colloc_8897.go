package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8897 evaluates chebyshev collocation node order 8897.
func ComputeChebyshevColloc8897(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
