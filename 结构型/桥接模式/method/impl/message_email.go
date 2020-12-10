package impl

import (
	"fmt"
	"pattern/结构型/桥接模式/method"
)

type MessageEmail struct {
}

func ViaEmail() method.MessageImlementer {
	return &MessageEmail{}
}

func (msms *MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email\n", text, to)
}
