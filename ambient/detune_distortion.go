package ambient

import (
	"math"
	"math/rand"
	"time"

	"github.com/CameronGorrie/sc"
)

func DetuneDistortion(p sc.Params) sc.Ugen {
	var (
		amp   = p.Add("amp", 0.3)
		atk   = p.Add("atk", 0.01)
		del   = p.Add("del", 1)
		freq  = p.Add("freq", 40)
		gate  = p.Add("gate", 0)
		out   = p.Add("out", 0)
		rel   = p.Add("rel", 0)
		sus   = p.Add("sus", 1)
		width = p.Add("width", 0.5)
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

	rand.Seed(time.Now().UnixNano())
	freqArr := randFloats(1, 20, 40, mulEachByFreq(freq))
	pulses := make([]sc.Input, len(freqArr))

	for i, f := range freqArr {
		pulses[i] = sc.Pulse{
			Freq:  f,
			Width: width,
		}.Rate(sc.AR).Mul(sc.C(0.05))
	}

	sig := sc.Mix(sc.AR, pulses).Mul(amp).Mul(env)

	return sc.Out{
		Bus:      out,
		Channels: sc.Multi(sig, sig),
	}.Rate(sc.AR)
}

func mulEachByFreq(f sc.Input) func(x float64) sc.Input {
	return func(x float64) sc.Input {
		return sc.C(x).Mul(f)
	}
}

func randFloats(min, max float64, n int, f func(x float64) sc.Input) []sc.Input {
	res := make([]sc.Input, n)
	for i := range res {
		x := min + rand.Float64()*(max-min)/1000
		x = roundFloat(x, 1000)
		res[i] = f(x)
	}
	return res
}

func roundFloat(n, p float64) float64 {
	return math.Round(n*p) / p
}
