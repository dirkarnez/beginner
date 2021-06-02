package main

import (
	"fmt"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/parser"
	"log"
)

func main() {
	src := `
		var obj = {
			a: 123,
			b: "this is a string"
		};
	`

	// Parse some JavaScript, yielding a *ast.Program and/or an ErrorList
	program, err := parser.ParseFile(nil, "", src, 0)
	if err != nil {
		log.Fatal(err)
	}

	for _, declaration := range program.DeclarationList {

		s, ok := declaration.(*ast.VariableDeclaration)
		if !ok {

		} else {
			element := s.List[0]
			fmt.Println(element.Name)

			objectLiteral, ok := element.Initializer.(*ast.ObjectLiteral)
			if !ok {

			} else {
				for _, valueItem := range objectLiteral.Value {
					fmt.Print(valueItem.Key)
					fmt.Print(": ")
					number, ok := valueItem.Value.(*ast.NumberLiteral)
					if ok {
						fmt.Println(number.Literal)
					}

					stringValue, ok := valueItem.Value.(*ast.StringLiteral)
					if ok {
						fmt.Println(stringValue.Literal)
					}
				}
			}
		}
	}
}