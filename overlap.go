package intervals

import (
	"fmt"
	"math"
)

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
	list := []*Interval{}
	if len(intvls.Intervals) == 0 {
		return list
	}

	// sort intervals (if necessary)
	intvls.Sort()

	lastMinLow := math.MaxInt64
	lastMaxHigh := math.MinInt64
	firstInitDone := false
	for _, intvl := range intvls.Intervals {
		// convert if necessary exclusive low/high values into inclusive ones
		low, high, err := intvls.getInclusives(intvl.Low, intvl.High)
		if err != nil {
			fmt.Printf("calculateOverlapped - unable to get inclusives: %v", err)
			continue
		}

		// for the first iteration make no sense those operations
		if firstInitDone {
			// check if the front or back side of the current segment overlaps with the previous one
			lowInBetween := isLowInBetweenInclusive(lastMinLow, lastMaxHigh, low, high)
			highInBetween := isHighInBetweenInclusive(lastMinLow, lastMaxHigh, low, high)
			if lowInBetween || highInBetween {
				// extract which part is overlapped, create a new interval and add it to the list
				biggestLow := max(low, lastMinLow)
				smallestHigh := min(high, lastMaxHigh)
				list = append(list, &Interval{Low: biggestLow, High: smallestHigh})
			}
		}
		// update control variables (if necessary)
		if low < lastMinLow {
			lastMinLow = low
		}
		if high > lastMaxHigh {
			lastMaxHigh = high
		}
		firstInitDone = true
	}
	return list
}

func (intvls *intervals) valueIsOverlapping(value int, overlapped []*Interval) bool {
	for _, ovrlp := range overlapped {
		if inBetween(value, ovrlp.Low, ovrlp.High, intvls.LowInclusive, intvls.HighInclusive) {
			return true
		}
	}
	return false
}
