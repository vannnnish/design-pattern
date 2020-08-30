package main

func main() {
	c1 := NewProjectManagerChain()
	c2 := NewDepManagerChain()
	c3 := NewGeneraManagerChain()
	c1.SetSuccessor(c2)
	c2.SetSuccessor(c3)
	var c Manager = c1
	c.HandleFeeRequest("hupeng", 500)
	c.HandleFeeRequest("hupeng", 1500)
	c.HandleFeeRequest("lixiang", 4500)
	c.HandleFeeRequest("lixiang", 11500)
}
