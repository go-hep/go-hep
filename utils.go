package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

// godeps returns the list of packages package pkg depends on.
//
// FIXME(sbinet) use go/build.Context: build.Default.Import(pkg, "", build.AllowBinary)
//  instead of shell-ing out. But: Context.Import doesn't handle "foo/bar/..."
func godeps(pkg string) ([]string, error) {
	args := []string{
		"list",
		"-f",
		`{{ join .Deps " "}}`,
		pkg,
	}

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	cmd := exec.Command("go", args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("go-hep: error running 'go list deps %s':\n%s\nerr=%v\n",
			pkg,
			string(stderr.Bytes()),
			err,
		)
	}

	set := make(map[string]struct{})
	for _, dep := range strings.Split(string(stdout.Bytes()), " ") {
		dep = strings.TrimSpace(dep)
		if dep == "" {
			continue
		}
		set[dep] = struct{}{}
	}

	deps := make([]string, 0, len(set))
	for dep := range set {
		deps = append(deps, dep)
	}
	sort.Strings(deps)

	return deps, err
}
