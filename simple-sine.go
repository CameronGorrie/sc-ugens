package ugens

import "github.com/CameronGorrie/sc"

func SimpleSine(p sc.Params) sc.Ugen {
	var (
		bus  = p.Add("bus", 0)
		freq = p.Add("freq", 440)
		gate = p.Add("gate", 1)
		bend = p.Add("bend", 1)
		amp  = p.Add("amp", 0.3)
	)

	env := sc.EnvGen{
		Gate: gate,
		Done: sc.FreeEnclosing,
		Env: sc.EnvADSR{
			A: sc.C(0.01),
			D: sc.C(1),
			S: sc.C(1),
			R: sc.C(1),
		},
	}.Rate(sc.KR)

	sig := sc.SinOsc{
		Freq: freq.Mul(bend),
	}.Rate(sc.AR).Mul(env).Mul(amp)

	return sc.Out{
		Bus:      bus,
		Channels: sc.Multi(sig, sig),
	}.Rate(sc.AR)
}
