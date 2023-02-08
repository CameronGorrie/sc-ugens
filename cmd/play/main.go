package main

import (
	"fmt"
	"time"

	"github.com/CameronGorrie/sc"
	ugens "github.com/CameronGorrie/ugens/pkg"
)

func main() {
	c, err := sc.NewClient("udp", "0.0.0.0:0", "192.168.7.2:57110", 5*time.Second)
	if err != nil {
		fmt.Println("Could not connect to client", err)
	}

	def := sc.NewSynthdef("potentiometer", ugens.Lib["potentiometer"])

	if err := c.SendDef(def); err != nil {
		fmt.Println("Could not send def to server", err)
	}

	c.Synth(
		"potentiometer",
		c.NextSynthID(),
		sc.AddToTail,
		sc.DefaultGroupID,
		map[string]float32{},
	)
}
