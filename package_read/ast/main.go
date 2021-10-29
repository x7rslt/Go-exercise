package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)
//AST，它的全名是abstract syntax tree(抽象语法树),其实就是使用树状结构表示源代码的语法结构，树的每一个节点就代表源代码中的一个结构。
func main() {
	// src is the input for which we want to print the AST.
	src := `
package main
func main() {
	println("Hello, World!")
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)

}
