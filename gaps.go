package interval

func (intvls *intervals) HasGaps() bool {
	intvls.Gaps()
	if intvls.GapsList != nil && len(intvls.GapsList) > 0 {
		return true
	}
	return false
}

func (intvls *intervals) Gaps() []*Interval {
	if intvls.GapsList == nil {
		intvls.GapsList = intvls.calculateGaps()
	}
	return intvls.GapsList
}

func (intvls *intervals) calculateGaps() []*Interval {
	intvls.Sort()
	gaps := []*Interval{}
	lastMaxHigh := intvls.MinLow
	for _, intvl := range intvls.Intervals {
		if intvl.Low > lastMaxHigh {
			gaps = append(gaps, &Interval{Low: lastMaxHigh, High: intvl.Low - 1})
		}
		if intvl.High >= lastMaxHigh {
			lastMaxHigh = intvl.High + 1
		}
	}
	if lastMaxHigh < intvls.MaxHigh {
		gaps = append(gaps, &Interval{Low: lastMaxHigh, High: intvls.MaxHigh})
	}
	return gaps
}
