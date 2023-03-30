package pads

import (
	"fmt"
	"strconv"

	"github.com/CameronGorrie/sc"
)

func PolyRhythmicDrone(p sc.Params) sc.Ugen {
	var (
		amp    = p.Add("amp", 0.15)
		freq   = p.Add("freq", 440)
		gate   = p.Add("gate", 0)
		modAmp = p.Add("modAmp", 0.8)
		out    = p.Add("out", 0)
	)

	env := sc.EnvGen{
		Done: sc.FreeEnclosing,
		Gate: gate,
		Env: sc.Env{
			Levels: []sc.Input{sc.C(0), sc.C(0.7), sc.C(0.7), sc.C(0)},
			Times:  []sc.Input{sc.C(18), sc.C(18), sc.C(18)},
		},
	}.Rate(sc.KR)

	modIdx := sc.LFNoise{
		Interpolation: sc.NoiseLinear,
		Freq:          sc.C(6),
	}.Rate(sc.KR)

	car := make([]sc.Input, 6)
	for i := 0; i < len(car); i++ {
		modFreq, _ := strconv.ParseFloat(fmt.Sprintf("%d.%d", i, i), 32)

		mod := sc.SinOsc{
			Freq: freq.Mul(sc.C(i)),
		}.Rate(sc.AR).Mul(freq).Mul(modIdx).Mul(sc.C(i)).Mul(modAmp)

		car[i] = sc.SinOsc{
			Freq: sc.C(modFreq).Add(mod),
		}.Rate(sc.AR)
	}

	sig := sc.Mix(sc.AR, car).Mul(amp).Mul(env)

	return sc.Out{Bus: out, Channels: sc.Multi(sig, sig)}.Rate(sc.AR)
}
