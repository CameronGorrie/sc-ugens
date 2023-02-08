package mpe

import (
	"github.com/CameronGorrie/sc"
)

func BPF(p sc.Params) sc.Ugen {
	var (
		amp  = p.Add("amp", 0.3)
		atk  = p.Add("atk", 0.01)
		bend = p.Add("bend", 0)
		del  = p.Add("del", 1)
		freq = p.Add("freq", 440)
		gate = p.Add("gate", 0)
		out  = p.Add("out", 0)
		rel  = p.Add("rel", 0)
		sus  = p.Add("sus", 1)
	)

	env := sc.EnvGen{
		Gate: gate,
		Done: sc.FreeEnclosing,
		Env: sc.EnvADSR{
			A: atk,
			D: del,
			S: sus,
			R: rel,
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
		Bus:      out,
		Channels: sc.Multi(sig, sig),
	}.Rate(sc.AR)
}
