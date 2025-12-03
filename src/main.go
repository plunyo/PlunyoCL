package main

import (
	"bufio"
	"fmt"
	"os"
	"pcl/src/frontend"
	"pcl/src/runtime"
)

func processSource(source string) {
	lexer := frontend.NewLexer(source)
	tokens := lexer.Tokenize()

    /*
	for _, token := range tokens {
		fmt.Println(token.String())
	}

	fmt.Println()
    */

	parser := frontend.NewParser(tokens)
	ast := parser.GenerateAST()

	// fmt.Println(ast)

    interpreter := runtime.NewInterpreter()
    result := interpreter.Evaluate(ast)

    fmt.Println(result)
}

func main() {
	if len(os.Args) > 2 {
		fmt.Printf("usage: %s [sourceFile]\n", os.Args[0])
		return
	}

	if len(os.Args) == 2 {
		sourceFile := os.Args[1]
		sourceCode, err := os.ReadFile(sourceFile)
		if err != nil {
			fmt.Printf("error reading file: %v\n", err)
			return
		}
		processSource(string(sourceCode))
	} else {
		fmt.Println("entering repl mode. type 'exit' to quit.")
		scanner := bufio.NewScanner(os.Stdin)

		for {
			fmt.Print(">> ")
			if !scanner.Scan() {
				break // eof or error
			}

			input := scanner.Text()
			if input == "exit" {
				break
			}

			processSource(input)
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("input error: %v\n", err)
		}
	}
}
