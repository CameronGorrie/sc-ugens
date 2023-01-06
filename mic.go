package ugens

import "github.com/CameronGorrie/sc"

func Mic(p sc.Params) sc.Ugen {
	out := p.Add("out", 0)
	sig := sc.SoundIn{Bus: sc.C(0)}.Rate(sc.AR)

	return sc.Out{
		Bus:      out,
		Channels: sc.Multi(sig, sig),
	}.Rate(sc.AR)
}
