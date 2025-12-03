package main

import (
	"fmt"
	"os"
	"pcl/src/frontend"
)

func main() {
    if len(os.Args) > 2 {
        fmt.Printf("Usage: %s [sourceFile]\n", os.Args[0])
        return
    }

    if len(os.Args) == 2 {
        sourceFile := os.Args[1]
        sourceCode, err := os.ReadFile(sourceFile)

        if err != nil {
            fmt.Printf("Error reading file: %v\n", err)
            return
        }
        
        lexer := frontend.NewLexer(string(sourceCode))
        tokens := lexer.Tokenize()

        for _, token := range tokens {
            fmt.Println(token.String())
        }

        fmt.Println()

        parser := frontend.NewParser(tokens)
        ast := parser.GenerateAST()
        
        fmt.Println(ast)
    } else {
        // repl mode
        fmt.Println("Entering REPL mode. Type 'exit' to quit.")

        for {
            fmt.Print(">> ")
            var input string
            fmt.Scanln(&input)

            if input == "exit" {
                break
            }

            lexer := frontend.NewLexer(input)
            tokens := lexer.Tokenize()

            for _, token := range tokens {
                fmt.Println(token.String())
            }

            fmt.Println()

            parser := frontend.NewParser(tokens)
            ast := parser.GenerateAST()

            fmt.Println(ast)
        }
    }
}