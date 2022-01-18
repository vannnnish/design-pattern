package main

import (
	"pattern/creational/abs-factory/abstract"
	"pattern/creational/abs-factory/abstract/impl"
)

/*
	 工厂个抽象工厂的典型区别，就是，普通工厂，一个工厂只能生产一个商品，
	抽象工厂则是在之上进行改进，可以生产多个商品。

	利用golang的组合特性，可以很轻松的实现这一特性

	当需要拓展功能的时候， 在工厂里面添加一个你想要需要的商品， 然后，对应在abstract 里面实现一个相应的对象即可
*/

func main() {
	var factory abstract.DAOFactory
	factory = &impl.MySQLFactory{}
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}
