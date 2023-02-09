package play

import (
	"fmt"
	"os"

	"github.com/CameronGorrie/sc"
	ugens "github.com/CameronGorrie/ugens/pkg"
)

func Cmd(c *sc.Client) {
	fs := CreateFlagSet()

	if len(os.Args) <= 1 {
		fmt.Println(HelpMsg)
	}

	app := NewApp(c, fs)

	for name, def := range ugens.Lib {
		if err := app.Add(name, def); err != nil {
			ErrorAndExit("[def]", err)
		}
	}

	if err := fs.Parse(os.Args[1:]); err != nil {
		ErrorAndExit("[parse]", err)
	}

	os.Exit(app.Run(os.Args[2:]))
}
