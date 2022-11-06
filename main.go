package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"golife/cmd"
)

func init() {
	subcommands.Register(&cmd.Server{}, "")
	subcommands.Register(&cmd.CLI{}, "")

	flag.Parse()
}

func main() {
	ctx := context.Background()
	subcommands.Execute(ctx)
}
