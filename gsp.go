package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gsp-lang/gsp/generator"
	"github.com/gsp-lang/gsp/parser"
	_ "github.com/gsp-lang/stdlib/fmt"
	"golang.org/x/tools/go/ast/astutil"
)

func args(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	p := parser.ParseFromString(filename, string(b)+"\n")

	a := generator.GenerateAST(p)

	fset := token.NewFileSet()

	defaultImports := []string{"github.com/gsp-lang/stdlib/prelude", "github.com/gsp-lang/gsp/core"}
	for _, defaultImport := range defaultImports {
		split := strings.Split(defaultImport, "/")
		pkgName := split[len(split)-1]
		if !(a.Name.Name == "prelude" && pkgName == "prelude") {
			astutil.AddImport(fset, a, defaultImport)
		}
	}

	if a.Name.Name != "prelude" {
		a.Decls = append(a.Decls, &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{&ast.ValueSpec{
				Names:  []*ast.Ident{&ast.Ident{Name: "_"}},
				Values: []ast.Expr{&ast.Ident{Name: "prelude.Len"}},
			}},
		})
	}

	var buf bytes.Buffer
	printer.Fprint(&buf, fset, a)
	fmt.Printf("%s\n", buf.String())
}

func main() {
	if len(os.Args) > 1 {
		args(os.Args[1])
		return
	}

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		line, _, _ := r.ReadLine()
		p := parser.ParseFromString("<REPL>", string(line)+"\n")
		fmt.Println(p)

		// a := generator.GenerateAST(p)
		a := generator.EvalExprs(p)
		fset := token.NewFileSet()

		ast.Print(fset, a)

		var buf bytes.Buffer
		printer.Fprint(&buf, fset, a)
		fmt.Printf("%s\n", buf.String())
	}
}
