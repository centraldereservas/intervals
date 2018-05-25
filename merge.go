package interval

func (intvls *intervals) Merge() []*Interval {
	if intvls.MergeList == nil {
		intvls.MergeList = intvls.calculateMerged()
	}
	return intvls.MergeList
}

func (intvls *intervals) calculateMerged() []*Interval {
	intvls.Sort()
	list := []*Interval{}
	var lastLow int
	var lastHigh int
	for i, intvl := range intvls.Intervals {
		if i == 0 {
			lastLow = intvl.Low
			lastHigh = intvl.High
			continue
		}
		if inBetweenInclusive(intvl.Low, lastLow, lastHigh) {
			// because the intervals are previously sorted, we just need to take care of the High value
			if intvl.High > lastHigh {
				lastHigh = intvl.High
			}
			continue
		}
		list = append(list, &Interval{Low: lastLow, High: lastHigh})
	}
	list = append(list, &Interval{Low: lastLow, High: lastHigh})
	return list
}
