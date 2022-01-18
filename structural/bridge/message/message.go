package message

type AbstractMessage interface {
	SendMessage (text,to string) // 发送快，发送不同
}