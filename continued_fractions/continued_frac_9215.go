package continued_fractions

// ComputeContinuedFrac9215 evaluates continued fraction approximant order 9215.
func ComputeContinuedFrac9215(x float64) float64 {
	a := 1.0
	for k := 6; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
