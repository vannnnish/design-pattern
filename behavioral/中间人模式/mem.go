package main

import "fmt"

type Mem struct {
	data string
}

func (m *Mem) Dump(data string) {
	fmt.Println("Mem is running ", data)
}
