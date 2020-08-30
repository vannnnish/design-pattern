package main

// 每个星期行为
type Week interface {
	Today()
	Next(*DayContext)
}

// 星期的数据结构
type DayContext struct {
	today Week
}

func NewDayContext() *DayContext {
	return &DayContext{today: &Sunday{}}
}

func (dc *DayContext) Today() {
	dc.today.Today()
}

func (dc *DayContext) Next() {
	dc.today.Next(dc)
}
