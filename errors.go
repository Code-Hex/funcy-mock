package funcy

import "os"

type causer interface{ Cause() error }
type exiter interface{ ExitCode() int }
type usage struct{ err error }

func (u usage) Error() string { return u.err.Error() }

func makeUsageError(err error) error { return usage{err: err} }

// UnwrapErrors get important message from wrapped error message
func UnwrapErrors(err error) (int, error) {
	for e := err; e != nil; {
		switch e.(type) {
		case usage:
			os.Stderr.WriteString(e.Error())
			return 0, nil
		case exiter:
			return e.(exiter).ExitCode(), e
		case causer:
			e = e.(causer).Cause()
		default:
			return 1, e // default error
		}
	}
	return 0, nil
}
