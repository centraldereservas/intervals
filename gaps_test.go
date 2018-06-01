package interval_test

import (
	"testing"

	"bitbucket.org/differenttravel/interval"
)

func TestGaps(t *testing.T) {
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
		name         string
		intvls       interval.Intervals
		expectedGaps []interval.Interval
	}{
		{name: "demo001", intvls: demo001.Intervals, expectedGaps: demo001.ExpectedGaps},
		{name: "demo002", intvls: demo002.Intervals, expectedGaps: demo002.ExpectedGaps},
		{name: "demo003", intvls: demo003.Intervals, expectedGaps: demo003.ExpectedGaps},
		{name: "demo004", intvls: demo004.Intervals, expectedGaps: demo004.ExpectedGaps},
		{name: "demo005", intvls: demo005.Intervals, expectedGaps: demo005.ExpectedGaps},
		{name: "demo006", intvls: demo006.Intervals, expectedGaps: demo006.ExpectedGaps},
		{name: "demo007", intvls: demo007.Intervals, expectedGaps: demo007.ExpectedGaps},
		{name: "demo008", intvls: demo008.Intervals, expectedGaps: demo008.ExpectedGaps},
		{name: "demo009", intvls: demo009.Intervals, expectedGaps: demo009.ExpectedGaps},
		{name: "demo010", intvls: demo010.Intervals, expectedGaps: demo010.ExpectedGaps},
		{name: "demo011", intvls: demo011.Intervals, expectedGaps: demo011.ExpectedGaps},
		{name: "demo012", intvls: demo012.Intervals, expectedGaps: demo012.ExpectedGaps},
		{name: "demo013", intvls: demo013.Intervals, expectedGaps: demo013.ExpectedGaps},
		{name: "demo014", intvls: demo014.Intervals, expectedGaps: demo014.ExpectedGaps},
		{name: "demo015", intvls: demo015.Intervals, expectedGaps: demo015.ExpectedGaps},

		{name: "demo101", intvls: demo101.Intervals, expectedGaps: demo101.ExpectedGaps},
		{name: "demo102", intvls: demo102.Intervals, expectedGaps: demo102.ExpectedGaps},
		{name: "demo103", intvls: demo103.Intervals, expectedGaps: demo103.ExpectedGaps},
		{name: "demo104", intvls: demo104.Intervals, expectedGaps: demo104.ExpectedGaps},
		{name: "demo105", intvls: demo105.Intervals, expectedGaps: demo105.ExpectedGaps},
		{name: "demo106", intvls: demo106.Intervals, expectedGaps: demo106.ExpectedGaps},
		{name: "demo107", intvls: demo107.Intervals, expectedGaps: demo107.ExpectedGaps},
		{name: "demo108", intvls: demo108.Intervals, expectedGaps: demo108.ExpectedGaps},
		{name: "demo109", intvls: demo109.Intervals, expectedGaps: demo109.ExpectedGaps},
		{name: "demo110", intvls: demo110.Intervals, expectedGaps: demo110.ExpectedGaps},
		{name: "demo111", intvls: demo111.Intervals, expectedGaps: demo111.ExpectedGaps},
		{name: "demo112", intvls: demo112.Intervals, expectedGaps: demo112.ExpectedGaps},

		{name: "demo201", intvls: demo201.Intervals, expectedGaps: demo201.ExpectedGaps},
		{name: "demo202", intvls: demo202.Intervals, expectedGaps: demo202.ExpectedGaps},
		{name: "demo203", intvls: demo203.Intervals, expectedGaps: demo203.ExpectedGaps},
		{name: "demo204", intvls: demo204.Intervals, expectedGaps: demo204.ExpectedGaps},
		{name: "demo205", intvls: demo205.Intervals, expectedGaps: demo205.ExpectedGaps},
		{name: "demo206", intvls: demo206.Intervals, expectedGaps: demo206.ExpectedGaps},
		{name: "demo207", intvls: demo207.Intervals, expectedGaps: demo207.ExpectedGaps},
		{name: "demo208", intvls: demo208.Intervals, expectedGaps: demo208.ExpectedGaps},
		{name: "demo209", intvls: demo209.Intervals, expectedGaps: demo209.ExpectedGaps},
		{name: "demo210", intvls: demo210.Intervals, expectedGaps: demo210.ExpectedGaps},
		{name: "demo211", intvls: demo211.Intervals, expectedGaps: demo211.ExpectedGaps},
		{name: "demo212", intvls: demo212.Intervals, expectedGaps: demo212.ExpectedGaps},

		{name: "demo301", intvls: demo301.Intervals, expectedGaps: demo301.ExpectedGaps},
		{name: "demo302", intvls: demo302.Intervals, expectedGaps: demo302.ExpectedGaps},
		{name: "demo303", intvls: demo303.Intervals, expectedGaps: demo303.ExpectedGaps},
		{name: "demo304", intvls: demo304.Intervals, expectedGaps: demo304.ExpectedGaps},
		{name: "demo305", intvls: demo305.Intervals, expectedGaps: demo305.ExpectedGaps},
		{name: "demo306", intvls: demo306.Intervals, expectedGaps: demo306.ExpectedGaps},
		{name: "demo307", intvls: demo307.Intervals, expectedGaps: demo307.ExpectedGaps},
		{name: "demo308", intvls: demo308.Intervals, expectedGaps: demo308.ExpectedGaps},
		{name: "demo309", intvls: demo309.Intervals, expectedGaps: demo309.ExpectedGaps},
		{name: "demo310", intvls: demo310.Intervals, expectedGaps: demo310.ExpectedGaps},
		{name: "demo311", intvls: demo311.Intervals, expectedGaps: demo311.ExpectedGaps},
		{name: "demo312", intvls: demo312.Intervals, expectedGaps: demo312.ExpectedGaps},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// if tc.name == "demo014" {
			// 	fmt.Printf("stop point")
			// }
			gaps := tc.intvls.Gaps()
			unexpectedLength := false
			if len(gaps) != len(tc.expectedGaps) {
				t.Errorf("unexpected gaps length: expected %d; got %d", len(tc.expectedGaps), len(gaps))
				unexpectedLength = true
			}
			if !unexpectedLength {
				for i, gap := range gaps {
					if gap.Low != tc.expectedGaps[i].Low || gap.High != tc.expectedGaps[i].High {
						t.Errorf("unexpected gap[%d]: expected %v; got %v", i, tc.expectedGaps[i], gap)
					}
				}
			}
		})
	}
}
