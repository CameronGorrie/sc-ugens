package main

import (
	"github.com/CameronGorrie/sc"
	play "github.com/CameronGorrie/ugens/internal"
)

func main() {
	c, err := play.NewClient(sc.DefaultLocalAddr, sc.DefaultScsynthAddr)
	if err != nil {
		play.ErrorAndExit("[client]", err)
	}

	play.Cmd(c)
}
