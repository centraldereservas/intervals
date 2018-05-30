package interval

import "fmt"

func (intvls *intervals) getInclusives(intervalLow, intervalHigh int) (int, int, error) {
	low := intvls.getInclusiveLow(intervalLow)
	high := intvls.getInclusiveHigh(intervalHigh)

	err := intvls.checkValidInterval(low, high)
	if err != nil {
		return 0, 0, err
	}
	return low, high, nil
}

func (intvls *intervals) checkValidInterval(low, high int) error {
	// if we get an incoherent result it's because it's invalid, for example in this case:
	//		(low,high)=(0,1) --> if we add 1 to the Low and substract 1 to the High ends with (1,0) which is not valid
	if low > high {
		return fmt.Errorf("low (%v) can not be bigger than high (%v)", low, high)
	}
	if high < low {
		return fmt.Errorf("high (%v) can not be smaller than low (%v)", high, low)
	}
	return nil
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
