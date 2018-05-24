package interval

import "fmt"

type Interval struct {
	Low    int
	High   int
	Object interface{}
}

func (itvl Interval) String() string {
	return fmt.Sprintf("(%v, %v)", itvl.Low, itvl.High)
}

// ByLow implements sort.Interface for []Interval based on the Low field.
type ByLow []*Interval

func (itvls ByLow) Len() int           { return len(itvls) }
func (itvls ByLow) Swap(i, j int)      { itvls[i], itvls[j] = itvls[j], itvls[i] }
func (itvls ByLow) Less(i, j int) bool { return itvls[i].Low < itvls[j].Low }

// ByHigh implements sort.Interface for []Interval based on the High field.
type ByHigh []*Interval

func (itvls ByHigh) Len() int           { return len(itvls) }
func (itvls ByHigh) Swap(i, j int)      { itvls[i], itvls[j] = itvls[j], itvls[i] }
func (itvls ByHigh) Less(i, j int) bool { return itvls[i].High > itvls[j].High }
