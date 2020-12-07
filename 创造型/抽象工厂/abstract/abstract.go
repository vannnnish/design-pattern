package abstract

import "pattern/创造型/抽象工厂/factory"

type DAOFactory interface {
	CreateOrderMainDAO() factory.OrderMainDAO
	CreateOrderDetailDAO() factory.OrderDetailDAO
	CreateOrderEditDao() factory.OrderEditDao
}
