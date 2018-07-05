package intervals_test

import (
	"testing"

	"github.com/centraldereservas/intervals"
)

func TestOverlapped(t *testing.T) {
	// tests for low/high inclusive
	demo001 := buildIntervalsDemo001()
	demo002 := buildIntervalsDemo002()
	demo003 := buildIntervalsDemo003()
	demo004 := buildIntervalsDemo004()
	demo005 := buildIntervalsDemo005()
	demo006 := buildIntervalsDemo006()
	demo007 := buildIntervalsDemo007()
	demo008 := buildIntervalsDemo008()
	demo009 := buildIntervalsDemo009()
	demo010 := buildIntervalsDemo010()
	demo011 := buildIntervalsDemo011()
	demo012 := buildIntervalsDemo012()
	demo013 := buildIntervalsDemo013()
	demo014 := buildIntervalsDemo014()
	demo015 := buildIntervalsDemo015()

	// tests for low/high exclusive
	demo101 := buildIntervalsDemo101()
	demo102 := buildIntervalsDemo102()
	demo103 := buildIntervalsDemo103()
	demo104 := buildIntervalsDemo104()
	demo105 := buildIntervalsDemo105()
	demo106 := buildIntervalsDemo106()
	demo107 := buildIntervalsDemo107()
	demo108 := buildIntervalsDemo108()
	demo109 := buildIntervalsDemo109()
	demo110 := buildIntervalsDemo110()
	demo111 := buildIntervalsDemo111()
	demo112 := buildIntervalsDemo112()

	// tests for low inclusive and high exclusive
	demo201 := buildIntervalsDemo201()
	demo202 := buildIntervalsDemo202()
	demo203 := buildIntervalsDemo203()
	demo204 := buildIntervalsDemo204()
	demo205 := buildIntervalsDemo205()
	demo206 := buildIntervalsDemo206()
	demo207 := buildIntervalsDemo207()
	demo208 := buildIntervalsDemo208()
	demo209 := buildIntervalsDemo209()
	demo210 := buildIntervalsDemo210()
	demo211 := buildIntervalsDemo211()
	demo212 := buildIntervalsDemo212()

	// tests for low exclusive and high inclusive
	demo301 := buildIntervalsDemo301()
	demo302 := buildIntervalsDemo302()
	demo303 := buildIntervalsDemo303()
	demo304 := buildIntervalsDemo304()
	demo305 := buildIntervalsDemo305()
	demo306 := buildIntervalsDemo306()
	demo307 := buildIntervalsDemo307()
	demo308 := buildIntervalsDemo308()
	demo309 := buildIntervalsDemo309()
	demo310 := buildIntervalsDemo310()
	demo311 := buildIntervalsDemo311()
	demo312 := buildIntervalsDemo312()
	tt := []struct {
		name             string
		intvls           intervals.Intervals
		expectedOverlaps []intervals.Interval
	}{
		{name: "demo001", intvls: demo001.Intervals, expectedOverlaps: demo001.ExpectedOverlaps},
		{name: "demo002", intvls: demo002.Intervals, expectedOverlaps: demo002.ExpectedOverlaps},
		{name: "demo003", intvls: demo003.Intervals, expectedOverlaps: demo003.ExpectedOverlaps},
		{name: "demo004", intvls: demo004.Intervals, expectedOverlaps: demo004.ExpectedOverlaps},
		{name: "demo005", intvls: demo005.Intervals, expectedOverlaps: demo005.ExpectedOverlaps},
		{name: "demo006", intvls: demo006.Intervals, expectedOverlaps: demo006.ExpectedOverlaps},
		{name: "demo007", intvls: demo007.Intervals, expectedOverlaps: demo007.ExpectedOverlaps},
		{name: "demo008", intvls: demo008.Intervals, expectedOverlaps: demo008.ExpectedOverlaps},
		{name: "demo009", intvls: demo009.Intervals, expectedOverlaps: demo009.ExpectedOverlaps},
		{name: "demo010", intvls: demo010.Intervals, expectedOverlaps: demo010.ExpectedOverlaps},
		{name: "demo011", intvls: demo011.Intervals, expectedOverlaps: demo011.ExpectedOverlaps},
		{name: "demo012", intvls: demo012.Intervals, expectedOverlaps: demo012.ExpectedOverlaps},
		{name: "demo013", intvls: demo013.Intervals, expectedOverlaps: demo013.ExpectedOverlaps},
		{name: "demo014", intvls: demo014.Intervals, expectedOverlaps: demo014.ExpectedOverlaps},
		{name: "demo015", intvls: demo015.Intervals, expectedOverlaps: demo015.ExpectedOverlaps},

		{name: "demo101", intvls: demo101.Intervals, expectedOverlaps: demo101.ExpectedOverlaps},
		{name: "demo102", intvls: demo102.Intervals, expectedOverlaps: demo102.ExpectedOverlaps},
		{name: "demo103", intvls: demo103.Intervals, expectedOverlaps: demo103.ExpectedOverlaps},
		{name: "demo104", intvls: demo104.Intervals, expectedOverlaps: demo104.ExpectedOverlaps},
		{name: "demo105", intvls: demo105.Intervals, expectedOverlaps: demo105.ExpectedOverlaps},
		{name: "demo106", intvls: demo106.Intervals, expectedOverlaps: demo106.ExpectedOverlaps},
		{name: "demo107", intvls: demo107.Intervals, expectedOverlaps: demo107.ExpectedOverlaps},
		{name: "demo108", intvls: demo108.Intervals, expectedOverlaps: demo108.ExpectedOverlaps},
		{name: "demo109", intvls: demo109.Intervals, expectedOverlaps: demo109.ExpectedOverlaps},
		{name: "demo110", intvls: demo110.Intervals, expectedOverlaps: demo110.ExpectedOverlaps},
		{name: "demo111", intvls: demo111.Intervals, expectedOverlaps: demo111.ExpectedOverlaps},
		{name: "demo112", intvls: demo112.Intervals, expectedOverlaps: demo112.ExpectedOverlaps},

		{name: "demo201", intvls: demo201.Intervals, expectedOverlaps: demo201.ExpectedOverlaps},
		{name: "demo202", intvls: demo202.Intervals, expectedOverlaps: demo202.ExpectedOverlaps},
		{name: "demo203", intvls: demo203.Intervals, expectedOverlaps: demo203.ExpectedOverlaps},
		{name: "demo204", intvls: demo204.Intervals, expectedOverlaps: demo204.ExpectedOverlaps},
		{name: "demo205", intvls: demo205.Intervals, expectedOverlaps: demo205.ExpectedOverlaps},
		{name: "demo206", intvls: demo206.Intervals, expectedOverlaps: demo206.ExpectedOverlaps},
		{name: "demo207", intvls: demo207.Intervals, expectedOverlaps: demo207.ExpectedOverlaps},
		{name: "demo208", intvls: demo208.Intervals, expectedOverlaps: demo208.ExpectedOverlaps},
		{name: "demo209", intvls: demo209.Intervals, expectedOverlaps: demo209.ExpectedOverlaps},
		{name: "demo210", intvls: demo210.Intervals, expectedOverlaps: demo210.ExpectedOverlaps},
		{name: "demo211", intvls: demo211.Intervals, expectedOverlaps: demo211.ExpectedOverlaps},
		{name: "demo212", intvls: demo212.Intervals, expectedOverlaps: demo212.ExpectedOverlaps},

		{name: "demo301", intvls: demo301.Intervals, expectedOverlaps: demo301.ExpectedOverlaps},
		{name: "demo302", intvls: demo302.Intervals, expectedOverlaps: demo302.ExpectedOverlaps},
		{name: "demo303", intvls: demo303.Intervals, expectedOverlaps: demo303.ExpectedOverlaps},
		{name: "demo304", intvls: demo304.Intervals, expectedOverlaps: demo304.ExpectedOverlaps},
		{name: "demo305", intvls: demo305.Intervals, expectedOverlaps: demo305.ExpectedOverlaps},
		{name: "demo306", intvls: demo306.Intervals, expectedOverlaps: demo306.ExpectedOverlaps},
		{name: "demo307", intvls: demo307.Intervals, expectedOverlaps: demo307.ExpectedOverlaps},
		{name: "demo308", intvls: demo308.Intervals, expectedOverlaps: demo308.ExpectedOverlaps},
		{name: "demo309", intvls: demo309.Intervals, expectedOverlaps: demo309.ExpectedOverlaps},
		{name: "demo310", intvls: demo310.Intervals, expectedOverlaps: demo310.ExpectedOverlaps},
		{name: "demo311", intvls: demo311.Intervals, expectedOverlaps: demo311.ExpectedOverlaps},
		{name: "demo312", intvls: demo312.Intervals, expectedOverlaps: demo312.ExpectedOverlaps},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// if tc.name == "demo105" {
			// 	t.Log("break here")
			// }
			overlaps := tc.intvls.Overlapped()
			unexpectedLength := false
			if len(overlaps) != len(tc.expectedOverlaps) {
				t.Errorf("unexpected overlaps length: expected %d; got %d", len(tc.expectedOverlaps), len(overlaps))
				unexpectedLength = true
			}
			if !unexpectedLength {
				for i, overlap := range overlaps {
					if overlap.Low != tc.expectedOverlaps[i].Low || overlap.High != tc.expectedOverlaps[i].High {
						t.Errorf("unexpected overlap[%d]: expected %v; got %v", i, tc.expectedOverlaps[i], overlap)
					}
				}
			}
		})
	}
}
