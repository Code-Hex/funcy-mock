package funcy

import (
	"fmt"
	"go/format"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func Run() int {
	if e := run(); e != nil {
		exitCode, err := UnwrapErrors(e)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return exitCode
		}
	}
	return 0
}

func run() error {
	file, dest, err := prepare()
	if err != nil {
		return errors.Wrap(err, "Failed to prepare")
	}
	f, err := parse(file)
	if err != nil {
		return errors.Wrap(err, "Failed to parse go file")
	}
	f.getInterfaces()
	bytes := f.generate()
	formatted, err := format.Source(bytes)
	if err != nil {
		return errors.Wrap(err, "Failed ast format")
	}
	fi, err := os.Create(dest)
	if err != nil {
		return errors.Wrapf(err, "Failed to create %s", dest)
	}
	if _, err := fi.Write(formatted); err != nil {
		return errors.Wrap(err, "Failed to write to %s")
	}
	if err := fi.Close(); err != nil {
		return errors.Wrapf(err, "Failed to close a %s", dest)
	}
	return nil
}

func prepare() (string, string, error) {
	var opts Options
	args, err := parseOptions(&opts, os.Args[1:])
	if err != nil {
		return "", "", err
	}
	if len(args) == 0 || opts.Help {
		return "", "", makeUsageError(errors.New(opts.usage()))
	}
	return args[0], getDest(opts, args[0]), nil
}

func getDest(opts Options, file string) string {
	if opts.Dest != "" {
		return opts.Dest
	}
	sep := strings.Split(file, ".")
	return strings.Join([]string{sep[0], "_mock", sep[1]}, ".")
}
