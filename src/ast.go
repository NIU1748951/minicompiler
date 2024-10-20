package main

// El objetivo de crear un Árbol de Sintaxis Abstracta (AST) es representar la estructura
// de las expresiones de forma jerárquica. Esto permite analizar las expresiones y la
// relación entre ellas de forma mucho más fácil a la hora de generar el código.
//
// Un AST es un modelo que refleja la estructura gramatical de un lenguaje de programación.
// Se compone de nodos, donde cada uno representa un elemento del lenguaje.
//
// Por ejemplo, una operación binaria (BinaryOperation) involucra el uso de dos
// operandos (números, otra operación binaria, etc.) "unidos" por un operador:
// - **Operando izquierdo**: representa el primer operando de la operación.
// - **Operador**: representa la operación que se va a realizar (como suma, resta, etc.).
// - **Operando derecho**: representa el segundo operando de la operación.
//
// Aquí, tanto el operando izquierdo como el operando derecho están relacionados por el
// operador. Por ejemplo, en la expresión `3 + 4`, el nodo correspondiente en el AST
// tendría un operando izquierdo que es el número `3`, el operador que es el símbolo `+`,
// y un operando derecho que es el número `4`:
//
//             BinaryOperation
//           /        |       \
//       3 (Izq)   + (Op)   4 (Der)
//

import "fmt"

// Se podria hacer directamente una struct de cada tipo de nodo pero seguramente así
// sea más fácil de mantener en un futuro y manejar teniendolo todo en una misma interfaz
type ASTNode interface {
	String() string
}

type NumberNode struct {
	Value int
}

func (n *NumberNode) String() string {
	return fmt.Sprintf("%d", n.Value)
}

type BinaryOperatorNode struct {
	Left  ASTNode // No number node porque también puede ser otra expresion!
	Op    string
	Right ASTNode
}

func (n *BinaryOperatorNode) String() string {
	return fmt.Sprintf("%s %s %s", n.Left, n.Op, n.Right)
}
