package funcy

import (
	"bytes"
	"fmt"
	"reflect"

	flags "github.com/jessevdk/go-flags"
)

const indent = "    "

type Options struct {
	Help    bool   `short:"h" long:"help" description:"show this message"`
	Package string `short:"p" long:"pkg" default:"main" description:"specify the package name"`
	Dest    string `short:"d" long:"dest" description:"specify the output destination"`
}

func parseOptions(opts *Options, argv []string) ([]string, error) {
	p := flags.NewParser(opts, flags.None)
	args, err := p.ParseArgs(argv)
	if err != nil {
		return nil, err
	}
	return args, nil
}

func (opts Options) usage() string {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, `Usage: funcygen [TARGET FILE] [OPTIONS]
Options:
`)

	t := reflect.TypeOf(opts)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag
		desc := tag.Get("description")
		var o string
		if s := tag.Get("short"); s != "" {
			o = fmt.Sprintf("-%s, --%s", tag.Get("short"), tag.Get("long"))
		} else {
			o = fmt.Sprintf("--%s", tag.Get("long"))
		}
		fmt.Fprintf(&buf, "  %-21s %s\n", o, desc)

		if deflt := tag.Get("default"); deflt != "" {
			fmt.Fprintf(&buf, "  %-21s    default: --%s='%s'\n", indent, tag.Get("long"), deflt)
		}
	}

	return buf.String()
}
