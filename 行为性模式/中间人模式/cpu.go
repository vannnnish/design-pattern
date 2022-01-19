package main

import "fmt"

type CPU struct {
	data string
}

func (c *CPU) Process(data string) {
	fmt.Println("CPU is running ", data)
}
