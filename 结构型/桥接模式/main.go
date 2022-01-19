package main

import (
	"pattern/结构型/桥接模式/message/impl"
	impl2 "pattern/结构型/桥接模式/method/impl"
)

func main() {
	commonMsgbySms := impl.NewCommonMessage(impl2.ViaSMS())
	commonMsgbySms.SendMessage("nihao", "西西")

	commonMsgByEmail := impl.NewCommonMessage(impl2.ViaEmail())
	commonMsgByEmail.SendMessage("nihao", "西西")
}
