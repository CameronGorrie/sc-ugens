package ugens

import "github.com/CameronGorrie/sc"

func SimpleSaw(p sc.Params) sc.Ugen {
	var (
		bus   = p.Add("bus", 0)
		freq  = p.Add("freq", 440)
		gate  = p.Add("gate", 1)
		phase = p.Add("phase", 0)
		width = p.Add("width", 0.05)
		bend  = p.Add("bend", 1)
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

	sig := sc.VarSaw{
		Freq:   freq.Mul(bend),
		IPhase: phase,
		Width:  width,
	}.Rate(sc.AR).Mul(env)

	return sc.Out{
		Bus:      bus,
		Channels: sc.Multi(sig, sig),
	}.Rate(sc.AR)
}
