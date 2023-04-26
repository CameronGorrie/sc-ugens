package ambient

import "github.com/CameronGorrie/sc"

func SpaceEngine(p sc.Params) sc.Ugen {
	var (
		amp  = p.Add("amp", 0.1)
		out  = p.Add("out", 0)
		gate = p.Add("gate", 0)
	)

	env := sc.EnvGen{
		Gate: gate,
		Done: sc.FreeEnclosing,
		Env: sc.EnvADSR{
			A: sc.C(0.1),
			D: sc.C(0.1),
			S: sc.C(0.5),
			R: sc.C(0.1),
		},
	}.Rate(sc.KR)

	w2 := sc.SinOsc{
		Freq: sc.C(800),
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
	}.Rate(sc.AR).Mul(amp).Mul(env)

	return sc.Out{
		Bus:      out,
		Channels: sc.Multi(sig, sig),
	}.Rate(sc.AR)
}
