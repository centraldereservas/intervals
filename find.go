package interval

func (intvls *intervals) FindIntervalsForValue(value int) []*Interval {
	intvls.Sort()
	var matches []*Interval
	for _, intvl := range intvls.Intervals {
		if intvl.Low > value {
			// due to the intervals are sorted, we can confirm that we will not find more matches
			break
		}
		if inBetweenInclusive(value, intvl.Low, intvl.High) {
			matches = append(matches, intvl)
		}
	}
	return matches
}
