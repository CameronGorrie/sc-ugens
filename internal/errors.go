package play

import (
	"fmt"
	"os"
)

func ErrorAndExit(slug string, err error) {
	fmt.Fprintf(os.Stderr, "%s %s\n", slug, err.Error())
	os.Exit(1)
}
