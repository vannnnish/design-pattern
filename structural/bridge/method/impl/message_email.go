package impl

import (
	"fmt"
	"pattern/structural/bridge/method"
)

type MessageEmail struct {
}

func ViaEmail() method.MessageImlementer {
	return &MessageEmail{}
}

func (msms *MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email\n", text, to)
}
