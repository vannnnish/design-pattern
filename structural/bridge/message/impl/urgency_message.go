package impl

import (
	"fmt"
	"pattern/structural/bridge/method"
)

type UrgencyMessage struct {
	method method.MessageImlementer
}

func NewUrgencyMessage(method method.MessageImlementer) *UrgencyMessage {
	return &UrgencyMessage{method: method}
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("发送到[%s]", text), to) // 很快速度发送
}
