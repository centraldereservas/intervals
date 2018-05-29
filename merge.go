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
	if len(intvls.Intervals) == 0 {
		return list
	}

	var lastLow int
	var lastHigh int
	for i, intvl := range intvls.Intervals {
		if i == 0 {
			lastLow = intvl.Low
			lastHigh = intvl.High
			continue
		}
		if isLowInBetween(intvl.Low, intvl.High, lastLow, lastHigh) || isHighInBetween(intvl.Low, intvl.High, lastLow, lastHigh) {
			if intvl.Low < lastLow {
				lastLow = intvl.Low
			}
			if intvl.High > lastHigh {
				lastHigh = intvl.High
			}
			continue
		}
		list = append(list, &Interval{Low: lastLow, High: lastHigh})
		lastLow = intvl.Low
		lastHigh = intvl.High
	}
	list = append(list, &Interval{Low: lastLow, High: lastHigh})
	return list
}
