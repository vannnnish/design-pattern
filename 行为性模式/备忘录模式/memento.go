package main

/*
	被网络模式： 包含两个角色,
	原始状态角色- 角色状态管理人
	原始状态角色里面维护状态对象， 且提供一个输出当前状态对象的方法。
	原始状态角色 将输出的状态对象保存到 角色状态管理人里面。
*/
type Memento struct {
	state string
}

func (m *Memento) SetState(state string) {
	m.state = state
}

func (m *Memento) GetState() string {
	return m.state
}

// 发起人
type Originator struct {
	state string
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{o.state}
}

// 负责人
type Caretaker struct {
	memento *Memento
}

func (c *Caretaker) GetMemento() *Memento {
	return c.memento
}

func (c *Caretaker) SetMemento(m *Memento) {
	c.memento = m
}


