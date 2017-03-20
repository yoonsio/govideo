package main

import (
	"runtime"

	"github.com/sickyoon/govideo/cmd"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	cmd.Execute()
}
