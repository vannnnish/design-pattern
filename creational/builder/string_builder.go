package main

type StringBuilder struct {
	result string
}

func (sb *StringBuilder) Part1() {
	sb.result += "1"
}

func (sb *StringBuilder) Part2() {
	sb.result += "2"
}

func (sb *StringBuilder) Part3() {
	sb.result += "3"
}

func (sb *StringBuilder) GetResult() string {
	return sb.result
}
