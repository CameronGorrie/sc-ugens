package ugens

import "github.com/CameronGorrie/sc"

var CompleteDictionary = map[string]sc.UgenFunc{
	"bpf":          BPF,
	"mic":          Mic,
	"simple_delay": SimpleDelay,
	"simple_saw":   SimpleSaw,
	"simple_sine":  SimpleSine,
}
