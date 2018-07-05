package intervals

import "fmt"

func (intvls *intervals) Merge() []*Interval {
	if intvls.MergeList == nil {
		intvls.MergeList = intvls.calculateMerged()
	}
	return intvls.MergeList
}

func (intvls *intervals) calculateMerged() []*Interval {
	list := []*Interval{}
	if len(intvls.Intervals) == 0 {
		return list
	}

	// sort intervals (if necessary)
	intvls.Sort()

	var currentMinLow int
	var currentMaxHigh int
	firstInitDone := false
	for _, intvl := range intvls.Intervals {
		// convert if necessary exclusive low/high values into inclusive ones
		low, high, err := intvls.getInclusives(intvl.Low, intvl.High)
		if err != nil {
			fmt.Printf("calculateMerged - unable to get inclusives: %v", err)
			continue
		}

		// in the first iteration it's for sure that don't exists merge, we just initialize variables
		if !firstInitDone {
			currentMinLow = low
			currentMaxHigh = high
			firstInitDone = true
			continue
		}

		// in case the current interval is overlapped or consecutive to the previous one, we need to update variables
		// and keep going and process the next interval
		if areSegmentsConsecutivesOrOverlapped(low, high, currentMinLow, currentMaxHigh) {
			// update control variables if necessary
			if low < currentMinLow {
				currentMinLow = low
			}
			if high > currentMaxHigh {
				currentMaxHigh = high
			}
			continue
		}

		// here the segments are not consecutive or overlapped so we close this merged segment (add to the list)
		list = append(list, &Interval{Low: currentMinLow, High: currentMaxHigh})

		// update control variables
		currentMinLow = low
		currentMaxHigh = high
	}
	// the last segment is pending to be added to the list
	list = append(list, &Interval{Low: currentMinLow, High: currentMaxHigh})
	return list
}

func areSegmentsConsecutivesOrOverlapped(low, high, lastLow, lastHigh int) bool {
	return areSegmentsOverlapped(low, high, lastLow, lastHigh) || areSegmentsConsecutives(low, high, lastLow, lastHigh)
}

func areSegmentsConsecutives(low, high, lastLow, lastHigh int) bool {
	return ((lastHigh + 1) == low) || ((high + 1) == lastLow)
}

func areSegmentsOverlapped(low, high, lastLow, lastHigh int) bool {
	if isLowInBetweenInclusive(low, high, lastLow, lastHigh) {
		return true
	}
	if isHighInBetweenInclusive(low, high, lastLow, lastHigh) {
		return true
	}
	return false
}
