/*
** parser.go
** Author: Marin Alcaraz
** Mail   <marin.alcaraz@gmail.com>
** Started on  Mon Feb 09 18:47:17 2015 Marin Alcaraz
** Last update Thu Mar 05 15:31:17 2015 Marin Alcaraz
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

//Node are items conformed by a type and a value
type Node struct {
	number int
	symbol string
	child  []Node
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

func showAst(tree []Node) {
	//Todo
}

func atomize(token string) Node {
	//Safe string to int conversion
	if n, err := strconv.Atoi(token); err == nil {
		return Node{number: n, symbol: "NAS"}
	}
	return Node{number: -1, symbol: token}
}

func myAppend(expression []Node, l Node) []Node {
	for key, value := range expression {
		if value.number == 0 &&
			value.symbol == "" {
			expression[key] = l
			return expression
		}
	}
	return append(expression, l)
}

func parseTokens(tokens []string) (Node, []string) {
	currentToken, tokens := pop(tokens)
	switch currentToken {
	case "(":
		var parent Node
		for tokens[0] != ")" {
			var child Node
			child, tokens = parseTokens(tokens)
			parent.child = append(parent.child, child)
		}
		_, tokens = pop(tokens)
		return parent, tokens
	default:
		return atomize(currentToken), tokens
	}
}

//buildAst parses the program slice and defines the AST structure
func buildAST(tokens []string) Node {
	//Read an expression from a sequence of tokens."
	if len(tokens) == 0 {
		return Node{}
	}
	ast, rest := parseTokens(tokens)
	if len(rest) != 0 {
		fmt.Println("[!]Error: malformed expression->", rest)
	}
	return ast
}

// Parse function triggers the sequence of building an AST, it returns the AST
func Parse() Node {
	tokens := tokenize()
	ast := buildAST(tokens)
	fmt.Println(len(ast.child))
	fmt.Println(ast)
	return ast
}
