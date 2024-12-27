package main

import (
	"blom/cli"
	"os"

	"github.com/gookit/goutil/dump"
)

func main() {
	dump.Config(func(o *dump.Options) {
		o.MaxDepth = 100
	})

	cli.Run(os.Args)
}
