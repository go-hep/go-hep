go-hep
======

`go-hep` is a simple command to gather all the `go-hep` dependencies
and thus create a release.

## Installation

```sh
$ go get github.com/go-hep/go-hep
```

## Usage

### Version

```sh
$ go-hep version
go-hep-0.0.1 (0000000)
```

### Dependencies

Displays all the dependencies of a package (`github.com/go-hep/go-hep`
by default):

```sh
$ go-hep deps fmt
errors
io
math
os
reflect
runtime
strconv
sync
sync/atomic
syscall
time
unicode/utf8
unsafe
```

### Help

```sh
$ go-hep help
go-hep - go-hep manages the go-hep distribution

Commands:

    deps        print dependencies and exit
    version     print version and exit

Use "go-hep help <command>" for more information about a command.
```
