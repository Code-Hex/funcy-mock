package funcy

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func (f *file) generate() []byte {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package %s\n\n", f.pkg)
	f.genimports(&buf)

	for k, data := range f.data {
		fmt.Fprintf(&buf, "type %sMock struct {\n", k)
		for _, d := range data {
			fmt.Fprintf(&buf, "%s func(%s) %s\n",
				d.PrivateName(),
				d.Param.TypeOnly,
				d.Return.Type,
			)
		}
		fmt.Fprint(&buf, "}\n\n")

		fmt.Fprintf(&buf, "func New%sMock() *%sMock {\n", k, k)
		fmt.Fprintf(&buf, "return &%sMock{\n", k)
		for _, d := range data {
			fmt.Fprintf(&buf, "%s: func(%s) %s { return %s },\n",
				d.PrivateName(),
				d.Param.TypeOnly,
				d.Return.Type,
				d.Return.Value,
			)
		}
		fmt.Fprintf(&buf, "}\n}\n\n")

		for _, d := range data {
			lower := strings.ToLower(k)
			fmt.Fprintf(&buf, "func (%c *%sMock) %s(%s) %s {\n",
				lower[0],
				k,
				d.Name,
				d.Param.Field,
				d.Return.Type,
			)
			fmt.Fprintf(&buf, "return %c.%s(%s)\n}\n\n",
				lower[0],
				d.PrivateName(),
				d.Param.NameOnly,
			)

			fmt.Fprintf(&buf, "func (%c *%sMock) Set%s(f func(%s) %s) {\n",
				lower[0],
				k,
				d.Name,
				d.Param.TypeOnly,
				d.Return.Type,
			)
			fmt.Fprintf(&buf, "if f == nil {\npanic(\"You should specify the mock function\")\n}\n")
			fmt.Fprintf(&buf, "%c.%s = f\n}\n\n", lower[0], d.PrivateName())
		}
	}
	return buf.Bytes()
}

func (f *file) genimports(buf io.Writer) {
	imports := f.File.Imports
	if ln := len(imports); ln > 1 {
		fmt.Fprintf(buf, "import (\n")
		for _, i := range imports {
			fmt.Fprintf(buf, "%s\n", i.Path.Value)
		}
		fmt.Fprintf(buf, ")\n\n")
	} else if ln == 1 {
		fmt.Fprintf(buf, "import %s\n\n", imports[0].Path.Value)
	}
}
