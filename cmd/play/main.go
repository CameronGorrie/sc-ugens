package main

import (
	"context"

	client "github.com/CameronGorrie/scc"
)

// Intended as a simple way to test local ugens.
func main() {
	c, err := client.NewClient("0.0.0.0:0", "127.0.0.1:57110")
	if err != nil {
		panic(err)
	}

	err = c.Play(
		context.Background(),
		"saw",
		[]string{
			"out", "0",
		},
	)

	if err != nil {
		panic(err)
	}
}
