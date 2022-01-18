package abstract

import "pattern/creational/abs-factory/factory"

type DAOFactory interface {
	CreateOrderMainDAO() factory.OrderMainDAO
	CreateOrderDetailDAO() factory.OrderDetailDAO
	CreateOrderEditDao() factory.OrderEditDao
}
