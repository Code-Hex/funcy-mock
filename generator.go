package funcy

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func (f *file) generate() []byte {
	var buf bytes.Buffer
	buf.WriteString("package main\n\n")
	f.genimports(&buf)

	for k, data := range f.data {
		fmt.Fprintf(&buf, "type %sMock struct {\n", k)
		for _, d := range data {
			fmt.Fprintf(&buf, "%sMock func() %s\n", d.Name, d.ReturnField())
		}
		fmt.Fprint(&buf, "}\n\n")

		fmt.Fprintf(&buf, "func New%sMock() *%sMock {\n", k, k)
		fmt.Fprintf(&buf, "return &%sMock{\n", k)
		for _, d := range data {
			fmt.Fprintf(&buf, "%sMock: func() %s { return %s },\n", d.Name, d.ReturnField(), d.ReturnValue)
		}
		fmt.Fprintf(&buf, "}\n}\n\n")

		for _, d := range data {
			lower := strings.ToLower(k)
			fmt.Fprintf(&buf, "func (%c *%sMock) %s(%s) %s {\n",
				lower[0],
				k,
				d.Name,
				d.Param,
				d.ReturnField(),
			)
			fmt.Fprintf(&buf, "return %c.%sMock()\n}\n\n", lower[0], d.Name)
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
