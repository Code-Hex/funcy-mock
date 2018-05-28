package funcy

import (
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"
	"strings"
)

type file struct {
	*ast.File
	pkg  string
	info *types.Info
	data map[string][]*Interface
}

func parse(fi string) (*file, error) {
	dir := filepath.Dir(fi)
	fset := token.NewFileSet()
	af, err := parser.ParseFile(fset, fi, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	astFiles, err := loadStdlib(fset, af, fi, dir)
	if err != nil {
		return nil, err
	}
	info, err := check(fset, astFiles, dir)
	if err != nil {
		return nil, err
	}
	return &file{
		File: af,
		pkg:  af.Name.Name,
		info: info,
		data: make(map[string][]*Interface),
	}, nil
}

func check(fset *token.FileSet, astFiles []*ast.File, dir string) (*types.Info, error) {
	info := &types.Info{
		Uses: map[*ast.Ident]types.Object{},
	}
	var conf types.Config
	conf.Importer = importer.Default()
	if _, err := conf.Check(dir, fset, astFiles, info); err != nil {
		return nil, err
	}
	return info, nil
}

func loadStdlib(fset *token.FileSet, af *ast.File, path, dir string) ([]*ast.File, error) {
	buildPkg, err := build.ImportDir(dir, 0)
	if err != nil {
		return nil, err
	}
	base := filepath.Base(path)
	astFiles := make([]*ast.File, 0, 1+len(buildPkg.GoFiles)+len(buildPkg.CgoFiles))
	astFiles = append(astFiles, af)
	for _, files := range [...][]string{buildPkg.GoFiles, buildPkg.CgoFiles} {
		for _, file := range files {
			if file == base {
				// already parsed this file above
				continue
			}
			file, err := parser.ParseFile(fset, filepath.Join(dir, file), nil, 0)
			if err != nil {
				return nil, err
			}
			astFiles = append(astFiles, file)
		}
	}
	return astFiles, nil
}

func (f *file) getInterfaces() {
	name := "" // interface name
	result := make([]*Interface, 0)
	f.walk(func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.TypeSpec:
			if _, ok := v.Type.(*ast.InterfaceType); ok {
				name = v.Name.Name
			}
		case *ast.InterfaceType:
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
			f.data[name] = result
		}
		return true
	})
}

func (f *file) walk(fn func(ast.Node) bool) {
	ast.Walk(walker(fn), f.File)
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
