/*
** parser.go
** Author: Marin Alcaraz
** Mail   <marin.alcaraz@gmail.com>
** Started on  Mon Feb 09 18:47:17 2015 Marin Alcaraz
** Last update Tue Feb 10 19:19:47 2015 Marin Alcaraz
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
	"fmt"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("\\s+")

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
		//Now the same for CURATEDINPUT for the ) symbol
		curatedInput = strings.Replace(curatedInput, ")", " ) ", -1)
		//Now the same for CURATEDINPUT for the ) symbol
		fmt.Printf("|%s|\n", curatedInput)
		//Append to the slice that contains our code, we will preserve it
		//on every iteration
		curatedInput = re.ReplaceAllString(curatedInput, " ")
		curatedInput = strings.TrimSpace(curatedInput)
		program = append(program, strings.Split(curatedInput, " ")...)
	}
	return program
}

//Can I use pointers?
func pop(slice []string) (string, []string) {
	fmt.Println(slice)
	element := slice[0]
	slice = slice[1:]
	return element, slice
}

func buildAST(tokens []string) ([]string, error) {
	//Read an expression from a sequence of tokens."
	if len(tokens) == 0 {
		return nil, fmt.Errorf("Unexpected EOF")
	}
	_, tokens = pop(tokens)
	for tokens != nil {
		//fmt.Println("CurrentToken:", currentToken)
		buildAST(tokens)
	}
	return tokens, nil
}

// Parse function triggers the sequence of building an AST, it returns the AST
func Parse() []string {
	tokens := tokenize()
	ast, err := buildAST(tokens)
	if err != nil {
		fmt.Println("[!] Parser Error: %s", err)
	}
	return ast
}
