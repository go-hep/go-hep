package main

import (
	"fmt"
	"sort"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func gohepMakeCmdDeps() *commander.Command {
	cmd := &commander.Command{
		Run:       gohepRunCmdDeps,
		UsageLine: "deps [package1 [package2 [...]]]",
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
	pkgs := Deps
	if len(args) > 0 {
		pkgs = args
	}

	depch := make(chan []string)
	errch := make(chan error)

	set := make(map[string]struct{})
	for _, pkg := range pkgs {
		go func(pkg string) {
			deps, err := godeps(pkg)
			if err != nil {
				errch <- err
				return
			}
			depch <- deps
		}(pkg)
	}

	for _ = range pkgs {
		select {
		case err := <-errch:
			return err
		case deps := <-depch:
			for _, dep := range deps {
				set[dep] = struct{}{}
			}
		}
	}

	deps := make([]string, 0, len(set))
	for dep := range set {
		deps = append(deps, dep)
	}

	sort.Strings(deps)
	for _, dep := range deps {
		fmt.Printf("%s\n", dep)
	}

	return nil
}

// EOF
