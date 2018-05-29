package interval

func (intvls *intervals) FindIntervalsForValue(value int) []*Interval {
	// sort intervals (if necessary)
	intvls.Sort()

	var matches []*Interval
	for _, intvl := range intvls.Intervals {
		// convert if necessary exclusive low/high values into inclusive ones
		low, high := intvls.getInclusives(intvl.Low, intvl.High)

		// check if we have to stop searching
		if low > value {
			// due to the intervals are sorted byLow, we can confirm that we will not find more matches
			break
		}
		if inBetweenInclusive(value, low, high) {
			matches = append(matches, intvl)
		}
	}
	return matches
}
