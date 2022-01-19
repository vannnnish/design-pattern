package impl

import "fmt"

type MySQLMainDAO struct {
}

func (*MySQLMainDAO) SaveOrderMain() {
	fmt.Println("MySQL main save")
}
