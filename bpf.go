package ugens

import (
	"github.com/CameronGorrie/sc"
)

func BPF(p sc.Params) sc.Ugen {
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

	carrier := sc.SinOsc{
		Freq: freq.Mul(bend.Midiratio()),
	}.Rate(sc.AR)

	sigMod := sc.SinOsc{
		Freq: sc.XLine{
			Start: freq,
			End:   freq.Add(sc.C(100)),
			Dur:   sc.C(20),
		}.Rate(sc.KR),
	}.Rate(sc.KR).MulAdd(sc.C(3600), sc.C(4000))

	sig := sc.BPF{
		In:   carrier,
		Freq: sigMod,
		RQ:   sc.C(3),
	}.Rate(sc.AR).Mul(amp).Mul(env)

	return sc.Out{
		Bus:      bus,
		Channels: sc.Multi(sig, sig),
	}.Rate(sc.AR)
}
