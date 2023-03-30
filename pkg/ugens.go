package ugens

import (
	"github.com/CameronGorrie/sc"
	"github.com/CameronGorrie/ugens/pkg/ambient"
	"github.com/CameronGorrie/ugens/pkg/pads"
	"github.com/CameronGorrie/ugens/pkg/samples"
	"github.com/CameronGorrie/ugens/pkg/simple"
	"github.com/CameronGorrie/ugens/pkg/soundin"
)

var Lib = map[string]sc.UgenFunc{
	"buf":               samples.Buf,
	"delay":             simple.Delay,
	"detuneDistortion":  ambient.DetuneDistortion,
	"mic":               soundin.Mic,
	"polyRhythmicDrone": pads.PolyRhythmicDrone,
	"saw":               simple.SimpleSaw,
	"sine":              simple.SimpleSine,
	"wobble":            ambient.Wobble,
}
