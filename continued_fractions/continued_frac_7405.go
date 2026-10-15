package continued_fractions

// ComputeContinuedFrac7405 evaluates continued fraction approximant order 7405.
func ComputeContinuedFrac7405(x float64) float64 {
	a := 1.0
	for k := 2; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
