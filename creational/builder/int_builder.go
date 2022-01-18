package main

type IntBuilder struct {
	result int
}

func (sb *IntBuilder) Part1() {
	sb.result += 1
}

func (sb *IntBuilder) Part2() {
	sb.result += 2
}

func (sb *IntBuilder) Part3() {
	sb.result += 3
}

func (sb *IntBuilder) GetResult() int {
	return sb.result
}
