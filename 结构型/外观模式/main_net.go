package main

type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

func (imp *aModuleImpl) TestA() string {
	return "主网已经运行"
}
