package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func removePath(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("empty path")
	}

	name := filepath.Base(path)
	ext := filepath.Ext(name)

	return name[:len(name)-len(ext)], nil
}

func main() {
	if len(os.Args) != 2 {
		fileName, _ := removePath(os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage: %s <path_to_file>\n", fileName)
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := scanner.Text()

		lexer := NewLexer(input)
		parser := NewParser(lexer)

		ast, err := parser.Parse()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing file: %v", err)
			os.Exit(1)
		}

		asmFile := "out/output.asm"

		err = GenerateAsm(ast, asmFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating ASM file: %v", err)
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v", err)
		os.Exit(1)
	}
}
