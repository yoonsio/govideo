package main

import (
	"flag"
	"runtime"

	"github.com/sickyoon/govideo/govideo"
)

var config = flag.String("config", "config.toml", "configuration file")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	app := govideo.NewApp(*config)
	app.Run()
}
