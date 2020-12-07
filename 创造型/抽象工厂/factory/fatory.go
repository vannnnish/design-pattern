package factory

type OrderMainDAO interface { // 订单记录
	SaveOrderMain()
}

type OrderDetailDAO interface { // 订单详情
	SaveOrderDetail()
}

type OrderEditDao interface { // 订单修改
	SaveOrderEdit()
}
