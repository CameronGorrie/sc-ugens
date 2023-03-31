package ambient

import "github.com/CameronGorrie/sc"

func SpaceEngine(p sc.Params) sc.Ugen {
	var (
		out = p.Add("out", 0)
	)

	w2 := sc.SinOsc{
		Freq: sc.C(101),
		Phase: sc.Saw{
			Freq: sc.C(0.12345),
		}.Rate(sc.KR).Mul(sc.C(678)).Add(sc.C(9)),
	}.Rate(sc.KR).Mul(sc.C(0.2)).Add(sc.C(0.8))

	w3 := sc.Pulse{
		Freq:  sc.C(25),
		Width: sc.C(0.25),
	}.Rate(sc.KR).Mul(sc.C(0.125)).Min(sc.C(0.25))

	w1 := sc.SinOsc{
		Freq: sc.C(51),
	}.Rate(sc.KR).Mul(w2).Add(w3)

	sig := sc.SinOsc{
		Freq: w1.Mul(sc.C(50)).Add(sc.C(10)),
	}.Rate(sc.AR).Mul(sc.C(0.5))

	return sc.Out{
		Bus:      out,
		Channels: sc.Multi(sig, sig),
	}.Rate(sc.AR)
}
