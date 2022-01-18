package main

import "fmt"

type GPU struct {
	data string
}

func (g *GPU) Display(data string) {
	fmt.Println("GPU is running ", data)
}
