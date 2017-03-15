package main

import (
	"flag"
	"runtime"

	"github.com/sickyoon/govideo/cmd"
)

var config = flag.String("config", "config.toml", "configuration file")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
