package interval

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func inBetweenInclusive(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	}
	return false
}

func inBetweenExclusive(i, min, max int) bool {
	if (i > min) && (i < max) {
		return true
	}
	return false
}

func isLowInBetween(low1, high1, low2, high2 int) bool {
	if inBetweenInclusive(low1, low2, high2) || inBetweenInclusive(low2, low1, high1) {
		return true
	}
	return false
}

func isHighInBetween(low1, high1, low2, high2 int) bool {
	if inBetweenInclusive(high1, low2, high2) || inBetweenInclusive(high2, low1, high1) {
		return true
	}
	return false
}
