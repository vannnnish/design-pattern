package main

import "fmt"

type MySQLMainDAO struct {
}

func (*MySQLMainDAO) SaveOrderMain() {
	fmt.Println("MySQL main save")
}

type MySQLDetailDAO struct {
}

func (*MySQLDetailDAO) SaveOrderDetail() {
	fmt.Println("Mysql order detail save")
}
