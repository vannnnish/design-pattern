package impl

import "fmt"

type MySQLEditDAO struct {
}

func (*MySQLEditDAO) SaveOrderEdit() {
	fmt.Println("MySQL Edit save")
}
