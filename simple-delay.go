package ugens

import "github.com/CameronGorrie/sc"

func SimpleDelay(p sc.Params) sc.Ugen {
	var (
		in        = p.Add("in", 0)
		out       = p.Add("out", 0)
		amp       = p.Add("amp", 1)
		delayTime = p.Add("delayTime", 1)
	)

	sig := sc.In{Bus: in}.Rate(sc.AR).Mul(amp)
	dSig := sc.Delay{
		In:            sig,
		Interpolation: sc.InterpolationLinear,
		MaxDelayTime:  delayTime,
		DelayTime:     sc.C(0.9),
	}.Rate(sc.AR)

	return sc.Out{
		Bus:      out,
		Channels: sc.Multi(dSig, dSig),
	}.Rate(sc.AR)
}
