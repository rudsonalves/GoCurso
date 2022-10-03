package statistics

func Max(e ...float64) float64 {
	max := e[0]
	for _, v := range e {
		if v > max {
			max = v
		}
	}

	return max
}

func Min(e ...float64) float64 {
	min := e[0]
	for _, v := range e {
		if min > v {
			min = v
		}
	}

	return min
}

func MidRange(e ...float64) float64 {
	max := Max(e...)
	min := Min(e...)

	return (max + min) / 2
}

func ArithmeticMean(e ...float64) float64 {
	sum := 0.
	for _, v := range e {
		sum += v
	}

	return sum / float64(len(e))
}
