package main

import play "github.com/CameronGorrie/ugens/internal"

func main() {
	c, err := play.NewClient("0.0.0.0:0", "192.168.7.2:57110")
	if err != nil {
		play.ErrorAndExit("[client]", err)
	}

	play.Cmd(c)
}
