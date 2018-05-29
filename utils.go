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

func inBetween(i, min, max int, lowInclusive, highInclusive bool) bool {
	if lowInclusive && highInclusive {
		return inBetweenInclusive(i, min, max)
	}
	if !lowInclusive && !highInclusive {
		return inBetweenExclusive(i, min, max)
	}
	if lowInclusive && !highInclusive {
		return inBetweenLowInclusive(i, min, max)
	}
	if !lowInclusive && highInclusive {
		return inBetweenHighInclusive(i, min, max)
	}
	return false
}

// both inclusive
func inBetweenInclusive(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	}
	return false
}

// only low is inclusive
func inBetweenLowInclusive(i, min, max int) bool {
	if (i >= min) && (i < max) {
		return true
	}
	return false
}

// only high is inclusive
func inBetweenHighInclusive(i, min, max int) bool {
	if (i > min) && (i <= max) {
		return true
	}
	return false
}

// both exclusive
func inBetweenExclusive(i, min, max int) bool {
	if (i > min) && (i < max) {
		return true
	}
	return false
}

func isLowInBetween(low1, high1, low2, high2 int, lowInclusive, highInclusive bool) bool {
	if inBetween(low1, low2, high2, lowInclusive, highInclusive) || inBetween(low2, low1, high1, lowInclusive, highInclusive) {
		return true
	}
	return false
}

func isLowInBetweenInclusive(low1, high1, low2, high2 int) bool {
	if inBetweenInclusive(low1, low2, high2) || inBetweenInclusive(low2, low1, high1) {
		return true
	}
	return false
}

func isHighInBetween(low1, high1, low2, high2 int, lowInclusive, highInclusive bool) bool {
	if inBetween(high1, low2, high2, lowInclusive, highInclusive) || inBetween(high2, low1, high1, lowInclusive, highInclusive) {
		return true
	}
	return false
}

func isHighInBetweenInclusive(low1, high1, low2, high2 int) bool {
	if inBetweenInclusive(high1, low2, high2) || inBetweenInclusive(high2, low1, high1) {
		return true
	}
	return false
}
