package sensors

import (
	"github.com/CameronGorrie/sc"
)

func Potentiometer(p sc.Params) sc.Ugen {
	var (
		out  = p.Add("out", 0)
		aPin = p.Add("aPin", 1)
	)

	gain := sc.AnalogIn{AnalogPin: aPin}.Rate(sc.AR)
	sig := sc.SinOsc{Freq: sc.C(440)}.Rate(sc.AR).Mul(gain)

	return sc.Out{Bus: out, Channels: sc.Multi(sig, sig)}.Rate(sc.AR)
}
