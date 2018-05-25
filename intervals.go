package interval

import (
	"math"
)

const (
	defaultMinLow  = 0
	defaultMaxHigh = math.MaxInt64
)

// Intervals is an interface to handle Interval structures discovering the existence of gaps or overlays
type Intervals interface {
	Add(itvl *Interval)

	// Sort sorts the intervals list by the Low property (ascending)
	Sort()

	// HasGaps returns true if exists gaps for the introduced intervals between MinLow and MaxHigh
	HasGaps() bool

	// Gaps first sorts (if necessary) and then returns the interval gaps
	Gaps() []*Interval

	// Merge first sorts (if necessary) and then fussion together overlapped intervals
	Merge() []*Interval

	// HasOverlapped returns true if exists overlapping for the introduced intervals
	HasOverlapped() bool

	// Overlapped first sorts (if necessary) and then returns the overlapped intervals
	Overlapped() []*Interval

	// FindIntervalsForValue returns all the intervals which contains the passed value
	FindIntervalsForValue(value int) []*Interval

	// Report first sorts (if necessary) and then creates a report of the interval sequence
	Report() string
}

// intervals implements Intervals interface
type intervals struct {
	Intervals      []*Interval
	GapsList       []*Interval
	OverlappedList []*Interval
	MergeList      []*Interval
	MinLow         int
	MaxHigh        int
	Sorted         bool
}

// NewIntervalsDefault is a constructor that returns an instance of the Intervals interface with default values
func NewIntervalsDefault() Intervals {
	return NewIntervals(defaultMinLow, defaultMaxHigh)
}

// NewIntervals is a constructor that returns an instance of the Intervals interface
func NewIntervals(minLow int, maxHigh int) Intervals {
	return &intervals{
		MinLow:    minLow,
		MaxHigh:   maxHigh,
		Intervals: []*Interval{},
		Sorted:    false,
	}
}
