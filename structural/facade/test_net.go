package main

type BModuleAPI interface {
	TestB() string
}


type bModuleImpl struct{}


func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

func (imp *bModuleImpl) TestB() string {
	return "测试网已经运行"
}
