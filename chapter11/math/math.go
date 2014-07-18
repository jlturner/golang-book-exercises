package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Finds the minimum and maximum values of a series of numbers
func MinMax(xs []float64) (min, max float64) {
	for i, v := range xs {
		if i == 0 {
			min = v
			max = v
		} else {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
	}
	return
}

