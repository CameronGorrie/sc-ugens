package ugens

import (
	"github.com/CameronGorrie/sc"
)

func SimpleSine(p sc.Params) sc.Ugen {
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

	sig := sc.SinOsc{
		Freq: freq.Mul(bend.Midiratio()),
	}.Rate(sc.AR).Mul(env).Mul(amp)

	return sc.Out{
		Bus:      out,
		Channels: sc.Multi(sig, sig),
	}.Rate(sc.AR)
}
