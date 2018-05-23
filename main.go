package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"

	"github.com/k0kubun/pp"
)

type Interface struct {
	Name        string
	Param       string
	ReturnType  string
	ReturnValue string
}

func (i *Interface) ReturnField() string {
	if strings.Contains(i.ReturnType, ",") {
		return "(" + i.ReturnType + ")"
	}
	return i.ReturnType
}

type file struct {
	*ast.File
	data map[string][]*Interface
}

func (f *file) walk(fn func(ast.Node) bool) {
	ast.Walk(walker(fn), f.File)
}

func main() {
	f := MustParse(`tmp/interface.go`)
	f.GetInterfaces()
	f.MockGen()
}

func MustParse(dir string) *file {
	fset := token.NewFileSet()
	af, err := parser.ParseFile(fset, dir, nil, parser.Mode(0))
	if err != nil {
		panic(err)
	}
	return &file{
		File: af,
		data: make(map[string][]*Interface),
	}
}

func (f *file) GetInterfaces() {
	key := ""
	result := make([]*Interface, 0)
	f.walk(func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.TypeSpec:
			if _, ok := v.Type.(*ast.InterfaceType); ok {
				key = v.Name.Name
			}
		case *ast.InterfaceType:
			// Do not check interface method names.
			// They are often constrainted by the method names of concrete types.
			for _, x := range v.Methods.List {
				switch v := x.Type.(type) {
				case *ast.FuncType:
					i := &Interface{
						Name:        x.Names[0].Name,
						Param:       getFields(v.Params.List),
						ReturnType:  getFields(v.Results.List),
						ReturnValue: getValues(v.Results.List),
					}
					result = append(result, i)
				}
			}
			f.data[key] = result
		}
		return true
	})
}

func (f *file) MockGen() {
	var buf bytes.Buffer
	buf.WriteString("package main\n\n")
	for k, data := range f.data {
		fmt.Fprintf(&buf, "type %sMock struct {\n", k)
		for _, d := range data {
			fmt.Fprintf(&buf, "\t%sMock func() %s\n", d.Name, d.ReturnField())
		}
		fmt.Fprint(&buf, "}\n\n")

		fmt.Fprintf(&buf, "func New%sMock() *%sMock {\n", k, k)
		fmt.Fprintf(&buf, "\treturn &%sMock{\n", k)
		for _, d := range data {
			fmt.Fprintf(&buf, "\t\t%sMock: func() %s { return %s },\n", d.Name, d.ReturnField(), d.ReturnValue)
		}
		fmt.Fprintf(&buf, "\t}\n}\n\n")

		for _, d := range data {
			lower := strings.ToLower(k)
			fmt.Fprintf(&buf, "func (%c *%sMock) %s(%s) %s {\n",
				lower[0],
				k,
				d.Name,
				d.Param,
				d.ReturnField(),
			)
			fmt.Fprintf(&buf, "\treturn %c.%sMock()\n}\n\n", lower[0], d.Name)
		}
	}
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	fi, err := os.Create("mock.go")
	if err != nil {
		log.Fatal(err)
	}
	fi.Write(formatted)
	fi.Close()
}

func getFields(list []*ast.Field) string {
	params := make([]string, 0, len(list))
	for _, p := range list {
		if len(p.Names) > 0 {
			params = append(params, p.Names[0].Name+" "+getType(p.Type))
		} else {
			params = append(params, getType(p.Type))
		}
	}
	return strings.Join(params, ", ")
}

func getValues(list []*ast.Field) string {
	params := make([]string, 0, len(list))
	for _, p := range list {
		params = append(params, getZeroValue(p.Type))
	}
	return strings.Join(params, ", ")
}

func getZeroValue(expr ast.Expr) string {
	switch v := expr.(type) {
	case *ast.StarExpr, *ast.ArrayType, *ast.MapType, *ast.FuncType:
		return "nil"
	case *ast.SelectorExpr:
		return getZeroValue(v.X)
	case *ast.Ident:
		switch v.Name {
		case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "byte", "rune", "uint", "int", "uintptr",
			"float32", "float64",
			"complex64", "complex128":
			return "0"
		case "bool":
			return "false"
		case "string":
			return `""`
		case "error":
			return "nil"
		default:
			return "nil"
		}
	}
	pp.Println(expr)
	return "???"
}

func getType(expr ast.Expr) string {
	switch v := expr.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.SelectorExpr:
		return getType(v.X) + "." + v.Sel.Name
	case *ast.StarExpr:
		return "*" + getType(v.X)
	case *ast.ArrayType:
		return "[]" + getType(v.Elt)
	case *ast.MapType:
		return "map[" + getType(v.Key) + "]" + getType(v.Value)
	case *ast.FuncType:
		return "func(" + getFields(v.Params.List) + ") " + getFields(v.Results.List)
	}
	pp.Println(expr)
	return "???"
}

// walker adapts a function to satisfy the ast.Visitor interface.
// The function return whether the walk should proceed into the node's children.
type walker func(ast.Node) bool

func (w walker) Visit(node ast.Node) ast.Visitor {
	if w(node) {
		return w
	}
	return nil
}
