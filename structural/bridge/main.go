package main

import (
	"pattern/structural/bridge/message/impl"
	impl2 "pattern/structural/bridge/method/impl"
)

/*
	桥接: 定义抽象调用， 具体实现则根据具体对象。
*/
func main() {
	commonMsgbySms := impl.NewCommonMessage(impl2.ViaSMS())
	commonMsgbySms.SendMessage("nihao", "西西")

	commonMsgByEmail := impl.NewCommonMessage(impl2.ViaEmail())
	commonMsgByEmail.SendMessage("nihao", "西西")
}
