package main

import (
	"fmt"
	"os"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

var gCmd *commander.Command

func init() {
	gCmd = &commander.Command{
		UsageLine: "go-hep",
		Short:     "go-hep manages the go-hep distribution",
		Subcommands: []*commander.Command{
			gohepMakeCmdVersion(),
		},
		Flag: *flag.NewFlagSet("go-hep", flag.ExitOnError),
	}
}

func main() {
	err := gCmd.Flag.Parse(os.Args[1:])
	handleErr(err)

	args := gCmd.Flag.Args()
	err = gCmd.Dispatch(args)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(1)
	}
}
