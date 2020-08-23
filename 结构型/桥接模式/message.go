package main


type AbstractMessage interface {
	SendMessage (text,to string) // 发送快，发送不同
}

type MessageImlementer interface {
	Send(text,to string)// 短信 邮件
}

func main(){
	m:=NewCommonMessage(ViaSMS())
	m.SendMessage("nihao","西西")
	
}