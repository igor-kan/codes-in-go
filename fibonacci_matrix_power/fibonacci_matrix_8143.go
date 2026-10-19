package fibonacci_matrix_power

// ComputeFibonacciMatrix8143 evaluates fibonacci matrix power recurrence order 8143.
func ComputeFibonacciMatrix8143(x float64) float64 {
	f0, f1 := 1.0, 1.0
	for i := 0; i < 8; i++ {
		f0, f1 = f1, f0+f1*x*0.1
	}
	return f1
}
