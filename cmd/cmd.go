package cmd

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"golife/cli"
	"golife/server"
)

type CLI struct{}

func (c *CLI) Name() string             { return "cli" }
func (c *CLI) Synopsis() string         { return "run CLI lifegame" }
func (c *CLI) Usage() string            { return "cli [args]" }
func (c *CLI) SetFlags(f *flag.FlagSet) {}
func (c *CLI) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	cli.Run()
	return subcommands.ExitSuccess
}

type Server struct{}

func (s *Server) Name() string             { return "server" }
func (s *Server) Synopsis() string         { return "run echo lifegame API server" }
func (s *Server) Usage() string            { return "server" }
func (s *Server) SetFlags(f *flag.FlagSet) {}
func (s *Server) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	server.Run()
	return subcommands.ExitSuccess
}
