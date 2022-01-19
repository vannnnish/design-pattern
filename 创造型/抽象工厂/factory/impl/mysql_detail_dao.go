package impl

import "fmt"

type MySQLDetailDAO struct {
}

func (*MySQLDetailDAO) SaveOrderDetail() {
	fmt.Println("Mysql order detail save")
}
