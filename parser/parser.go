/*
** parser.go
** Author: Marin Alcaraz
** Mail   <marin.alcaraz@gmail.com>
** Started on  Mon Feb 09 18:47:17 2015 Marin Alcaraz
** Last update Tue Feb 10 16:51:05 2015 Marin Alcaraz
 */

// Based on:
// Rob Pike's Lexical analysis in Go
// Mary Rose's Little Lisp
// Peter Norvig's (How to Write a (Lisp) Interpreter (in Python))

// What does a language interpreter does?

// Parsing: build an AST based on the input.
// 			the AST mirrors the nested structure of the lang.

// Execution: Evaluation of the expressions in the AST.

package parser

import (
	"bufio"
	"os"
	"strings"
)

// Lexems are items conformed by a type and a value

//Lexem type

type lLexem struct {
	lType string
	//Not shure of this...
	value interface{}
}

func tokenize() []string {

	var currentStr string
	var curatedInput string
	var program []string

	//Create reader from Stdin
	reader := bufio.NewReader(os.Stdin)

	//Send the reader to the scanner
	scanner := bufio.NewScanner(reader)

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		currentStr = scanner.Text()

		//Curate the CURRENTSTRING
		curatedInput = strings.Replace(currentStr, "(", " ( ", -1)
		//Curate the CURATEDSTRING to preserve the old changes
		curatedInput = strings.Replace(curatedInput, ")", " ) ", -1)

		//Append to the slice that contains our code, we will preserve it
		//on every iteration
		program = append(program, strings.Split(curatedInput, " ")...)
	}
	return program
}

func buildAST(tokens []string) []string {
	return tokens
}

// Parse function triggers the sequence of building an AST, it returns the AST
func Parse() []string {
	tokens := tokenize()
	ast := buildAST(tokens)
	return ast
}
