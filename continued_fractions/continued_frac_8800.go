package continued_fractions

// ComputeContinuedFrac8800 evaluates continued fraction approximant order 8800.
func ComputeContinuedFrac8800(x float64) float64 {
	a := 1.0
	for k := 5; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
