package ugens

import (
	"github.com/CameronGorrie/sc"
	"github.com/CameronGorrie/ugens/ambient"
	"github.com/CameronGorrie/ugens/pads"
	"github.com/CameronGorrie/ugens/samples"
	"github.com/CameronGorrie/ugens/simple"
	"github.com/CameronGorrie/ugens/soundin"
)

var Lib = map[string]sc.UgenFunc{
	"buf":               samples.Buf,
	"delay":             simple.Delay,
	"detune_distortion": ambient.DetuneDistortion,
	"mic":               soundin.Mic,
	"pr_drone":          pads.PolyRhythmicDrone,
	"saw":               simple.SimpleSaw,
	"sine":              simple.SimpleSine,
	"space_engine":      ambient.SpaceEngine,
	"wobble":            ambient.Wobble,
}
