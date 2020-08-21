package main

type MySQLFactory struct {

}

func (*MySQLFactory)CreateOrderMainDAO()OrderMainDAO{
	return &MySQLMainDAO{}
}

func(*MySQLFactory)CreateOrderDetailDAO()OrderDetailDAO{
	return &MySQLDetailDAO{}
}