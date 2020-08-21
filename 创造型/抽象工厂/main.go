package main

func main() {
	var factory DAOFactory
	factory = &MySQLFactory{}
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}
