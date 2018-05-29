package interval

func (intvls *intervals) getInclusives(intervalLow, intervalHigh int) (int, int) {
	low := intvls.getInclusiveLow(intervalLow)
	high := intvls.getInclusiveHigh(intervalHigh)
	if high < low {
		high = low
	}
	if low > high {
		low = high
	}
	return low, high
}

func (intvls *intervals) getInclusiveLow(value int) int {
	if intvls.LowInclusive {
		return value
	}
	// here the low is exclusive, we have to take the next value
	newLow := value + 1
	if newLow > intvls.MaxHigh {
		return intvls.MaxHigh
	}
	return newLow
}

func (intvls *intervals) getInclusiveHigh(value int) int {
	if intvls.HighInclusive {
		return value
	}
	// here the high is exclusive, we have to take the previous value
	newHigh := value - 1
	if newHigh < intvls.MinLow {
		return intvls.MinLow
	}
	return newHigh
}
