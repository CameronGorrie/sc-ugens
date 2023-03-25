package ugens

import (
	"github.com/CameronGorrie/sc"
	"github.com/CameronGorrie/ugens/pkg/ambient"
	"github.com/CameronGorrie/ugens/pkg/samples"
	"github.com/CameronGorrie/ugens/pkg/simple"
	"github.com/CameronGorrie/ugens/pkg/soundin"
)

var Lib = map[string]sc.UgenFunc{
	"buf":              samples.Buf,
	"detuneDistortion": ambient.DetuneDistortion,
	"mic":              soundin.Mic,
	"delay":            simple.Delay,
	"saw":              simple.SimpleSaw,
	"sine":             simple.SimpleSine,
	"wobble":           ambient.Wobble,
}
