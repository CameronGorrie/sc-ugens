package soundin

import "github.com/CameronGorrie/sc"

func Mic(p sc.Params) sc.Ugen {
	var (
		out = p.Add("out", 0)
	)

	sig := sc.SoundIn{
		// SoundIn busses must be constant or an array of constants.
		Bus: sc.C(0),
	}.Rate(sc.AR)

	return sc.Out{
		Bus:      out,
		Channels: sc.Multi(sig, sig),
	}.Rate(sc.AR)
}
