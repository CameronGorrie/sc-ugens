package play

import (
	"fmt"
	"time"

	"github.com/CameronGorrie/sc"
)

// NewClient creates a new sc client.
func NewClient(local, scsynth string) (*sc.Client, error) {
	c, err := sc.NewClient("udp", local, scsynth, 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("creating sc client")
	}

	if _, err := c.AddDefaultGroup(); err != nil {
		return nil, fmt.Errorf("adding default group")
	}

	return c, nil
}
