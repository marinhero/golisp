/*
** golisp.go
** Author: Marin Alcaraz
** Mail   <marin.alcaraz@gmail.com>
** Started on  Tue Feb 10 16:14:02 2015 Marin Alcaraz
** Last update Tue Feb 10 16:35:02 2015 Marin Alcaraz
 */

package main

import "fmt"
import "golisp/parser"

func main() {
	ast := parser.Parse()
	fmt.Println(ast)
}
