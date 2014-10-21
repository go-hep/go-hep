package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func gohepMakeCmdDeps() *commander.Command {
	cmd := &commander.Command{
		Run:       gohepRunCmdDeps,
		UsageLine: "deps",
		Short:     "print dependencies and exit",
		Long: fmt.Sprintf(`
print dependencies and exit.

ex:
 $ go-hep deps
 $ go-hep deps github.com/go-hep/go-hep
`),
		Flag: *flag.NewFlagSet("go-hep-deps", flag.ExitOnError),
	}
	return cmd
}

func gohepRunCmdDeps(cmd *commander.Command, args []string) error {
	pkgs := []string{"github.com/go-hep/go-hep"}
	if len(args) > 0 {
		pkgs = args
	}

	subargs := []string{
		"list",
		"-f",
		`{{ join .Deps "\n"}}`,
	}
	subargs = append(subargs, pkgs...)
	sub := exec.Command("go", subargs...)
	sub.Stdout = os.Stdout
	sub.Stderr = os.Stderr
	sub.Stdin = os.Stdin

	err := sub.Run()
	return err
}

// EOF
