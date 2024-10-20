package main

import (
	"fmt"
	"strconv"
)

type Parser struct {
	lexer    *Lexer
	curToken Token
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer: lexer, curToken: lexer.nextToken()}
}

// Se vienen los porros
func (p *Parser) Parse() (ASTNode, error) {
	return p.parseExpression()
}

// Hay que arreglar esto, no es nada escalable y muy probablemente hayan errores
func (p *Parser) parseExpression() (ASTNode, error) {
	// De momento únicamente BinaryOperators siendo cada operando un numero

	left, err := p.parseNumber()
	if err != nil {
		return nil, err
	}

	operator := p.curToken
	if operator.Type != TokenOperator {
		return nil, fmt.Errorf("operator expected, instead got: %s", operator.Value)
	}
	p.nextToken()

	right, err := p.parseNumber()
	if err != nil {
		return nil, err
	}

	return &BinaryOperatorNode{
		Left:  left,
		Op:    operator.Value,
		Right: right,
	}, nil
}

func (p *Parser) parseNumber() (ASTNode, error) {
	if p.curToken.Type != TokenNumber {
		return nil, fmt.Errorf("number expected, got %s", p.curToken.Value)
	}

	value, err := strconv.Atoi(p.curToken.Value)
	if err != nil {
		return nil, fmt.Errorf("error converting to int: %v", err)
	}

	p.nextToken()
	return &NumberNode{Value: value}, nil
}

func (p *Parser) nextToken() {
	p.curToken = p.lexer.nextToken()
}
