package main

import (
	"strconv"
	"strings"
)

type Parser struct {
	exp   []string
	index int
	prev  Node
}

func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")
	for {
		if p.index >= len(p.exp) {
			return
		}
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAddNode()
		case "-":
			p.prev = p.newSubNode()
		default:
			p.prev = p.newValNode()

		}
	}
}

func (p *Parser) newAddNode() Node {
	p.index++
	return &AddNode{left: p.prev, right: p.newValNode()}
}

func (p *Parser) newSubNode() Node {
	p.index++
	return &SubNode{left: p.prev, right: p.newValNode()}

}

func (p *Parser) newValNode() Node {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &ValNode{v}
}

func (p *Parser) Result() Node {
	return p.prev
}
