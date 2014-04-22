package tracks

import (
	"testing"
)

var stationTrack = &TrackData{
	Elements: []TrackElement{
		TrackElement{
			Segment: TS_MAP[ELEM_END_STATION],
		},
	},
}

// The simplest possible circuit
var completeTrack = &TrackData{
	Elements: []TrackElement{
		TrackElement{
			Segment: TS_MAP[ELEM_END_STATION],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_FLAT],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
	},
}

var hillyTrack = &TrackData{
	Elements: []TrackElement{
		TrackElement{
			Segment: TS_MAP[ELEM_END_STATION],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_FLAT_TO_25_DEG_UP],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_25_DEG_UP_TO_FLAT],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_FLAT_TO_25_DEG_DOWN],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_25_DEG_DOWN_TO_FLAT],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_FLAT],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_FLAT],
		},
		TrackElement{
			Segment: TS_MAP[ELEM_FLAT],
		},
	},
}

func TestCircuit(t *testing.T) {
	td := &TrackData{}
	if IsCircuit(td) {
		t.Errorf("empty ride should not be a circuit.")
	}

	if IsCircuit(stationTrack) {
		t.Errorf("track with single station shouldn't be a circuit")
	}

	if !IsCircuit(completeTrack) {
		t.Errorf("Complete track should be marked as a circuit but wasn't")
	}

	if !IsCircuit(hillyTrack) {
		t.Errorf("Hilly track should be marked as a circuit but wasn't")
	}
}