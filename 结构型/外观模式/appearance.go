package main

import "fmt"

// 子系统中的一组接口 提供同一个入口，此模式定义一个高层接口，
type TAPI interface {
	Test() string
}
type APICall struct {
	a AModuleAPI
	b BModuleAPI
}

func NewTAPI() TAPI {
	return &APICall{NewAModuleAPI(),NewBModuleAPI()}

}
func (api *APICall) Test() string {
	return fmt.Sprintf("%s\n%s", api.a.TestA(), api.b.TestB())
}


func main(){
	api:=NewTAPI()
	fmt.Println(api.Test())
}