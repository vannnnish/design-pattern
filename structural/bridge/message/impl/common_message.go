package impl

import (
	"fmt"
	"pattern/structural/bridge/method"
)

type CommonMessage struct {
	method method.MessageImlementer
}


func NewCommonMessage(imlementer method.MessageImlementer)*CommonMessage {
	return &CommonMessage{method: imlementer}
}

func(m *CommonMessage)SendMessage(text,to string){
	m.method.Send(fmt.Sprintf("发送到[%s]",text),to)
}
