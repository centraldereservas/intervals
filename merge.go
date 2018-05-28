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
	pendingToAdd := false
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
			pendingToAdd = true
			continue
		}
		list = append(list, &Interval{Low: lastLow, High: lastHigh})
		lastLow = intvl.Low
		lastHigh = intvl.High
		pendingToAdd = false
	}
	if pendingToAdd == true {
		list = append(list, &Interval{Low: lastLow, High: lastHigh})
	}
	return list
}
