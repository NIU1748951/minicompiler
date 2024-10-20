package main

import (
	"fmt"
	"os"
)

func GenerateAsm(node ASTNode, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("section .text\n")
	file.WriteString("global  _start\n")
	file.WriteString("\n")
	file.WriteString("_start:\n")

	switch n := node.(type) {
	case *NumberNode:
		fmt.Fprintf(file, "    mov eax, %d\n", n.Value)
	case *BinaryOperatorNode:
		generateBinaryOpNode(file, n)
	}

	file.WriteString("\n")
	file.WriteString("    mov ebx, eax            ; move result to ebx\n")
	file.WriteString("    mov eax, 1              ; sys_exit\n")
	file.WriteString("    int 0x80                ; calling sys_exit\n")

	return nil
}

func generateBinaryOpNode(file *os.File, node *BinaryOperatorNode) {
	switch left := node.Left.(type) {
	case *NumberNode:
		fmt.Fprintf(file, "    mov eax, %d\n", left.Value)
	case *BinaryOperatorNode:
		// Unimplemented
	}

	switch right := node.Right.(type) {
	case *NumberNode:
		fmt.Fprintf(file, "    mov ebx, %d\n", right.Value)
	case *BinaryOperatorNode:
		// Unimplemented
	}

	switch node.Op {
	case "+":
		file.WriteString("    add eax, ebx\n")
	case "-":
		file.WriteString("    sub eax, ebx\n")
	case "*":
		file.WriteString("    imul eax, ebx\n")
	case "/":
		file.WriteString("    cdq\n")
		file.WriteString("    idiv ebx\n")
	}

}
