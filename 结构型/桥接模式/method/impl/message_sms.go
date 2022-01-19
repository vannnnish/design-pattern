package impl

import (
	"fmt"
	"pattern/结构型/桥接模式/method"
)

type MessageSMS struct {
}

func ViaSMS() method.MessageImlementer {
	return &MessageSMS{}
}

func (msms *MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS\n", text, to)
}
