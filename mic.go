package ugens

import "github.com/CameronGorrie/sc"

func Mic(p sc.Params) sc.Ugen {
	var (
		in  = p.Add("in", 0)
		out = p.Add("out", 0)
	)

	sig := sc.SoundIn{
		Bus: in,
	}.Rate(sc.AR)

	return sc.Out{
		Bus:      out,
		Channels: sc.Multi(sig, sig),
	}.Rate(sc.AR)
}
