package check

// Contains check contain
func Contains(x string, A []string) bool {
	for _, elem := range A {
		if x == elem {
			return true
		}
	}
	return false
}

// ContainsFloat check Contain float
func ContainsFloat(x float64, A []float64) bool {
	for _, elem := range A {
		if x == elem {
			return true
		}
	}
	return false
}

// ContainsInt check Contain float
func ContainsInt(x int, A []int) bool {
	for _, elem := range A {
		if x == elem {
			return true
		}
	}
	return false
}
