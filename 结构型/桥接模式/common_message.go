package main

import "fmt"

type CommonMessage struct {
	method MessageImlementer
}


func NewCommonMessage(imlementer MessageImlementer)*CommonMessage{
	return &CommonMessage{method:imlementer}
}

func(m *CommonMessage)SendMessage(text,to string){
	m.method.Send(fmt.Sprintf("发送到[%s]",text),to)
}
