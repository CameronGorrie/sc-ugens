package ugens

import "github.com/CameronGorrie/sc"

func Buf(p sc.Params) sc.Ugen {
	var (
		bufNum  = p.Add("bufNum", 0)
		loop    = p.Add("loop", 0)
		out     = p.Add("out", 0)
		speed   = p.Add("speed", 1)
		start   = p.Add("start", 0)
		trigger = p.Add("trigger", 0)
	)

	sig := sc.PlayBuf{
		NumChannels: 2,
		BufNum:      bufNum,
		Speed:       speed.Midiratio(),
		Trigger:     trigger,
		Start:       start,
		Loop:        loop,
		Done:        sc.FreeEnclosing,
	}.Rate(sc.AR)

	return sc.Out{Bus: out, Channels: sig}.Rate(sc.AR)
}
