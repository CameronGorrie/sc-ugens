package ugens

import (
	"github.com/CameronGorrie/sc"
)

func SimpleSine(p sc.Params) sc.Ugen {
	var (
		a    = p.Add("a", 0.01)
		amp  = p.Add("amp", 0.3)
		bend = p.Add("bend", 0)
		bus  = p.Add("bus", 0)
		d    = p.Add("d", 1)
		freq = p.Add("freq", 440)
		gate = p.Add("gate", 0)
		r    = p.Add("r", 0)
		s    = p.Add("s", 1)
	)

	env := sc.EnvGen{
		Gate: gate,
		Done: sc.FreeEnclosing,
		Env: sc.EnvADSR{
			A: a,
			D: d,
			S: s,
			R: r,
		},
	}.Rate(sc.KR)

	sig := sc.SinOsc{
		Freq: freq.Mul(bend.Midiratio()),
	}.Rate(sc.AR).Mul(env).Mul(amp)

	return sc.Out{
		Bus:      bus,
		Channels: sc.Multi(sig, sig),
	}.Rate(sc.AR)
}
