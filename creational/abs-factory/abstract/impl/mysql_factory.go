package impl

import (
	"pattern/creational/abs-factory/factory"
	"pattern/creational/abs-factory/factory/impl"
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
