package continued_fractions

// ComputeContinuedFrac7605 evaluates continued fraction approximant order 7605.
func ComputeContinuedFrac7605(x float64) float64 {
	a := 1.0
	for k := 4; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
