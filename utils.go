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
