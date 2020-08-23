package main

import "fmt"

type MessageSMS struct {
}

func ViaSMS() MessageImlementer {
	return &MessageSMS{}
}

func (msms *MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS", text, to)
}
