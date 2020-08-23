package main

import "fmt"

type UrgencyMessage struct {
	method MessageImlementer
}

func NewUrgencyMessage(method MessageImlementer) *UrgencyMessage {
	return &UrgencyMessage{method: method}
}

func (m  *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("发送到[%s]",text),to)// 很快速度发送
}
