package impl

import (
	"pattern/创造型/抽象工厂/factory"
	"pattern/创造型/抽象工厂/factory/impl"
)

type MySQLFactory struct {
}

func (*MySQLFactory) CreateOrderMainDAO() factory.OrderMainDAO {
	return &impl.MySQLMainDAO{}
}

func (*MySQLFactory) CreateOrderDetailDAO() factory.OrderDetailDAO {
	return &impl.MySQLDetailDAO{}
}

func (*MySQLFactory) CreateOrderEditDao() factory.OrderEditDao {
	return &impl.MySQLEditDAO{}
}
