package interval

import "math"

func (intvls *intervals) HasOverlapped() bool {
	intvls.Overlapped()
	if intvls.OverlappedList != nil && len(intvls.OverlappedList) > 0 {
		return true
	}
	return false
}

func (intvls *intervals) Overlapped() []*Interval {
	if intvls.OverlappedList == nil {
		intvls.OverlappedList = intvls.calculateOverlapped()
	}
	return intvls.OverlappedList
}

func (intvls *intervals) calculateOverlapped() []*Interval {
	intvls.Sort()
	list := []*Interval{}
	lastMinLow := math.MaxInt64
	lastMaxHigh := math.MinInt64
	for i, intvl := range intvls.Intervals {
		if i > 0 {
			lowInBetween := isLowInBetween(lastMinLow, lastMaxHigh, intvl.Low, intvl.High)   //inBetweenInclusive(lastMinLow, intvl.Low, intvl.High) || inBetweenInclusive(intvl.Low, lastMinLow, lastMaxHigh)
			highInBetween := isHighInBetween(lastMinLow, lastMaxHigh, intvl.Low, intvl.High) //inBetweenInclusive(lastMaxHigh, intvl.Low, intvl.High) || inBetweenInclusive(intvl.High, lastMinLow, lastMaxHigh)
			if lowInBetween || highInBetween {
				greaterLow := max(intvl.Low, lastMinLow)
				lowerHigh := min(intvl.High, lastMaxHigh)
				list = append(list, &Interval{Low: greaterLow, High: lowerHigh})
			}
		}
		if intvl.Low < lastMinLow {
			lastMinLow = intvl.Low
		}
		if intvl.High > lastMaxHigh {
			lastMaxHigh = intvl.High
		}
	}
	return list
}

func (intvls *intervals) valueIsOverlapping(value int, overlapped []*Interval) bool {
	for _, ovrlp := range overlapped {
		if inBetweenInclusive(value, ovrlp.Low, ovrlp.High) {
			return true
		}
	}
	return false
}
