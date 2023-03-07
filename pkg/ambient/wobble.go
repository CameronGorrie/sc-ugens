package ambient

import "github.com/CameronGorrie/sc"

func Wobble(p sc.Params) sc.Ugen {
	var (
		atk    = p.Add("atk", 0.01)
		cf     = p.Add("cf", 100)
		freq   = p.Add("freq", 440)
		gate   = p.Add("gate", 0)
		out    = p.Add("out", 0)
		rel    = p.Add("rel", 4)
		t_bd   = p.Add("t_bd", 0)
		t_sd   = p.Add("t_sd", 0)
		width  = p.Add("width", 0.4)
		wobble = p.Add("wobble", 3)
	)

	freqModArr := []float64{0.99, 0.5, 1.01}
	pulses := make([]sc.Input, len(freqModArr))

	for i, f := range freqModArr {
		pulse := sc.Pulse{
			Freq:  freq.Mul(sc.C(f)),
			Width: width,
		}.Rate(sc.AR)

		lpf := sc.RLPF{
			In:   pulse,
			Freq: cf,
		}.Rate(sc.AR)

		mod := sc.SinOsc{
			Freq: wobble,
		}.Rate(sc.KR)

		pulses[i] = lpf.Mul(mod).Sin()
	}

	base := sc.Mix(sc.AR, pulses)

	env := sc.EnvGen{
		Gate: gate,
		Done: sc.FreeEnclosing,
		Env: sc.EnvLinen{
			Attack:  atk,
			Release: rel,
		},
	}.Rate(sc.KR)

	bd := sc.Ringz{
		In: sc.LPF{
			In: sc.Trig{
				In:  t_bd,
				Dur: sc.SampleDur{}.Rate(sc.IR),
			}.Rate(sc.AR),
			Freq: sc.C(1000),
		}.Rate(sc.AR),
		Freq:      sc.C(30),
		DecayTime: sc.C(0.5),
	}.Rate(sc.AR).Mul(sc.C(7)).Sin().Tanh()

	sd := sc.Ringz{
		In: sc.LPF{
			In: sc.Trig{
				In:  t_sd,
				Dur: sc.SampleDur{}.Rate(sc.IR),
			}.Rate(sc.AR),
			Freq: sc.C(1000),
		}.Rate(sc.AR),
		Freq:      sc.C(120),
		DecayTime: sc.C(0.75),
	}.Rate(sc.AR).Mul(sc.PinkNoise{}.Rate(sc.AR).Mul(sc.C(2))).Sin().Mul(sc.C(2))

	hpf := sc.HPF{
		In:   sd,
		Freq: sc.C(60),
	}.Rate(sc.AR)

	sig := sc.GVerb{
		In: sc.HPF{
			In:   base,
			Freq: sc.C(30),
		}.Rate(sc.AR),
		RoomSize: sc.C(70),
		RevTime:  sc.C(11),
		Damping:  sc.C(0.15),
	}.Rate(sc.AR).Mul(sc.C(0.5)).Add(base).Add(bd).Add(hpf).Mul(env).Tanh()

	return sc.Out{Bus: out, Channels: sc.Multi(sig, sig)}.Rate(sc.AR)
}
