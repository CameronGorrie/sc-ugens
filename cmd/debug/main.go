package main

import play "github.com/CameronGorrie/ugens/internal"

func main() {
	c, err := play.NewClient("0.0.0.0:0", "0.0.0.0:57110")
	if err != nil {
		play.ErrorAndExit("[client]", err)
	}

	play.Cmd(c)
}
