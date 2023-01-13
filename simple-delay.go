package ugens

import "github.com/CameronGorrie/sc"

func SimpleDelay(p sc.Params) sc.Ugen {
	var (
		amp    = p.Add("amp", 1)
		dTime  = p.Add("dTime", 0.9)
		in     = p.Add("in", 0)
		mdTime = p.Add("mdTime", 1)
		out    = p.Add("out", 0)
	)

	inSig := sc.In{
		Bus: in,
	}.Rate(sc.AR).Mul(amp)

	sig := sc.Delay{
		In:            inSig,
		Interpolation: sc.InterpolationLinear,
		MaxDelayTime:  mdTime,
		DelayTime:     dTime,
	}.Rate(sc.AR)

	return sc.Out{
		Bus:      out,
		Channels: sc.Multi(sig, sig),
	}.Rate(sc.AR)
}
