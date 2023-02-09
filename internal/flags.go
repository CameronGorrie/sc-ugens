package play

import (
	"flag"
	"fmt"
)

var HelpMsg = `NAME:
ugens - A command line app for building and playing SynthDefs from the github.com/CameronGorrie/ugens library.

USAGE:
ugens [command options] [arguments...]

OPTIONS:
-h							Print this help message
-l              List the available sounds
-s SYNTH        Play SYNTH
`

func CreateFlagSet() *flag.FlagSet {
	fs := flag.NewFlagSet("ugens", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println(HelpMsg)
	}

	return fs
}
