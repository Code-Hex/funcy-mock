package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"path"
	"path/filepath"
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
	info *types.Info
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
	af, err := parser.ParseFile(fset, dir, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	buildPkg, err := build.ImportDir(filepath.Dir(dir), 0)
	if err != nil {
		panic(err)
	}
	pp.Println(filepath.Dir(dir))
	astFiles := make([]*ast.File, 0, 1+len(buildPkg.GoFiles)+len(buildPkg.CgoFiles))
	astFiles = append(astFiles, af)
	d := filepath.Dir(dir)
	base := filepath.Base(dir)
	for _, files := range [...][]string{buildPkg.GoFiles, buildPkg.CgoFiles} {
		for _, filename := range files {
			if filename == base {
				// already parsed this file above
				continue
			}
			file, err := parser.ParseFile(fset, path.Join(d, filename), nil, 0)
			if err != nil {
				panic(err)
			}
			astFiles = append(astFiles, file)
		}
	}
	info := &types.Info{
		Uses: map[*ast.Ident]types.Object{},
	}
	var conf types.Config
	conf.Importer = importer.Default()
	if _, err := conf.Check(dir, fset, astFiles, info); err != nil {
		panic(err)
	}
	return &file{
		File: af,
		info: info,
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
						Param:       f.getFields(v.Params.List),
						ReturnType:  f.getFields(v.Results.List),
						ReturnValue: f.getValues(v.Results.List),
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

	imports := f.File.Imports
	if ln := len(imports); ln > 1 {
		fmt.Fprintf(&buf, "import (\n")
		for _, i := range imports {
			fmt.Fprintf(&buf, "%s\n", i.Path.Value)
		}
		fmt.Fprintf(&buf, ")\n\n")
	} else if ln == 1 {
		pp.Println(imports)
		fmt.Fprintf(&buf, "import %s\n\n", imports[0].Path.Value)
	}
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

func (f *file) getFields(list []*ast.Field) string {
	params := make([]string, 0, len(list))
	for _, p := range list {
		if len(p.Names) > 0 {
			params = append(params, p.Names[0].Name+" "+f.getType(p.Type))
		} else {
			params = append(params, f.getType(p.Type))
		}
	}
	return strings.Join(params, ", ")
}

func (f *file) getValues(list []*ast.Field) string {
	params := make([]string, 0, len(list))
	for _, p := range list {
		params = append(params, f.getZeroValue(p.Type))
	}
	return strings.Join(params, ", ")
}

func (f *file) getZeroValue(expr ast.Expr) string {
	switch v := expr.(type) {
	case *ast.StarExpr, *ast.SliceExpr, *ast.ArrayType, *ast.MapType, *ast.FuncType,
		*ast.ChanType, *ast.StructType, *ast.InterfaceType:
		return "nil"
	case *ast.SelectorExpr:
		return f.getZeroValue(v.Sel)
	case *ast.Ident:
		return f.getBuiltinZeroValue(v)
	}
	return "nil"
}

func (f *file) getBuiltinZeroValue(ident *ast.Ident) string {
	switch f.info.TypeOf(ident).Underlying().String() {
	case "uint8", "uint16", "uint32", "uint64", "uint", "uintptr",
		"int8", "int16", "int32", "int64", "int", "byte", "rune",
		"float32", "float64",
		"complex64", "complex128":
		return "0"
	case "bool":
		return "false"
	case "string":
		return `""`
	default:
		//pp.Println(t)
		return "nil"
	}
}

func (f *file) getType(expr ast.Expr) string {
	switch v := expr.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.SelectorExpr:
		return f.getType(v.X) + "." + v.Sel.Name
	case *ast.StarExpr:
		return "*" + f.getType(v.X)
	case *ast.SliceExpr:
		return "[]" + f.getType(v.X)
	case *ast.MapType:
		return "map[" + f.getType(v.Key) + "]" + f.getType(v.Value)
	case *ast.FuncType:
		return "func(" + f.getFields(v.Params.List) + ") " + f.getFields(v.Results.List)
	}
	//pp.Println(expr)
	return "nil"
}

// walker adapts a function to satisfy the ast.Visitor interface.
// The function return whether the walk should proceed into the node's children.
type walker func(ast.Node) bool

func (w walker) Visit(node ast.Node) ast.Visitor {
	//pp.Println(node)
	if w(node) {
		return w
	}
	return nil
}
