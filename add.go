package interval

func (intvls *intervals) Add(low, high int, obj interface{}) error {
	itvl := &Interval{
		Low:    low,
		High:   high,
		Object: obj,
	}
	return intvls.AddInterval(itvl)
}

func (intvls *intervals) AddInterval(itvl *Interval) error {
	low := intvls.getInclusiveLow(itvl.Low)
	high := intvls.getInclusiveHigh(itvl.High)

	err := intvls.checkValidInterval(low, high)
	if err != nil {
		return err
	}

	intvls.Intervals = append(intvls.Intervals, itvl)
	intvls.reset()
	return nil
}

func (intvls *intervals) reset() {
	intvls.Sorted = false
	intvls.GapsList = nil
	intvls.OverlappedList = nil
	intvls.MergeList = nil
}
