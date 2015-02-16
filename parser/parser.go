/*
** parser.go
** Author: Marin Alcaraz
** Mail   <marin.alcaraz@gmail.com>
** Started on  Mon Feb 09 18:47:17 2015 Marin Alcaraz
** Last update Mon Feb 16 18:36:53 2015 Marin Alcaraz
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
	"strconv"
	"strings"
)

var re = regexp.MustCompile("\\s+")

// nodes are items conformed by a type and a value

type node struct {
	number int
	symbol string
	child  []node
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
		//Append to the slice that contains our code, we will preserve it
		//on every iteration
		curatedInput = re.ReplaceAllString(curatedInput, " ")
		curatedInput = strings.TrimSpace(curatedInput)

		fmt.Printf("INPUT: |%s|\n", curatedInput)
		//Finally consider the curatedInput as part of the program
		program = append(program, strings.Split(curatedInput, " ")...)
	}
	return program
}

//A classic pop function for a slice
func pop(slice []string) (string, []string) {
	element := slice[0]
	slice = slice[1:]
	return element, slice
}

func showAst(tree []node) {
	//Todo
}

func atomize(token string) node {
	//Safe string to int conversion
	if n, err := strconv.Atoi(token); err == nil {
		return node{number: n, symbol: "NAS"}
	}
	return node{number: -1, symbol: token}
}

func myAppend(expression []node, l node) []node {
	for key, value := range expression {
		if value.number == 0 &&
			value.symbol == "" {
			expression[key] = l
			return expression
		}
	}
	return append(expression, l)
}

//buildAst parses the program slice and defines the AST structure
func buildAST(tokens []string) []string {
	//Read an expression from a sequence of tokens."
	if len(tokens) == 0 {
		fmt.Printf("[Parser] Error: Unexpected EOF")
		os.Exit(1)
	}
	currentToken, tokens := pop(tokens)
	fmt.Println(currentToken)
	if currentToken == "(" {
		if currentToken == ")" {
			fmt.Errorf("[Parser] Error Unexpected )")
			os.Exit(1)
		}
		level := make([]node, 1) //Minimun size of a valid LISP expression
		for _, t := range tokens {
			fmt.Println(t)
			if t == ")" {
				break
			}
			level = myAppend(level, atomize(t))
		}
		fmt.Println(level)
	}
	return tokens
}

// Parse function triggers the sequence of building an AST, it returns the AST
func Parse() []string {
	tokens := tokenize()
	ast := buildAST(tokens)
	return ast
}
