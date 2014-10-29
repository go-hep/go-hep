package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func gohepMakeCmdUpdate() *commander.Command {
	cmd := &commander.Command{
		Run:       gohepRunCmdUpdate,
		UsageLine: "update [package1 [package2 [...]]]",
		Short:     "update a (list of) package(s) and exit",
		Long: fmt.Sprintf(`
update a (list of) package(s) and exit.

ex:
 $ go-hep update
 $ go-hep update github.com/go-hep/fwk
 $ go-hep update github.com/go-hep/fwk/... github.com/go-hep/fads
`),
		Flag: *flag.NewFlagSet("go-hep-update", flag.ExitOnError),
	}
	return cmd
}

func gohepRunCmdUpdate(cmd *commander.Command, args []string) error {
	pkgs := Deps
	if len(args) > 0 {
		pkgs = args
	}

	subargs := []string{
		"get",
		"-u",
		"-v",
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
