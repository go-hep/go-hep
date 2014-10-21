package main

import (
	"fmt"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func gohepMakeCmdVersion() *commander.Command {
	vers := Version
	rev := Revision
	cmd := &commander.Command{
		Run:       gohepRunCmdVersion,
		UsageLine: "version",
		Short:     "print version and exit",
		Long: fmt.Sprintf(`
print version and exit.

ex:
 $ go-hep version
 go-hep-%s (%s)
`, vers, rev),
		Flag: *flag.NewFlagSet("go-hep-version", flag.ExitOnError),
	}
	return cmd
}

func gohepRunCmdVersion(cmd *commander.Command, args []string) error {
	vers := Version
	rev := Revision
	fmt.Printf("go-hep-%s (%s)\n", vers, rev)
	return nil
}

// EOF
