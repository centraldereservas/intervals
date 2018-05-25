package interval

func (intvls *intervals) Add(itvl *Interval) {
	intvls.Intervals = append(intvls.Intervals, itvl)
	intvls.reset()
}

func (intvls *intervals) reset() {
	intvls.Sorted = false
	intvls.GapsList = nil
	intvls.OverlappedList = nil
	intvls.MergeList = nil
}
