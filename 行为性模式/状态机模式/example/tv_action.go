package main

import "fmt"

// 测试

var (
	StateOpen  = 1
	StateClose = 2
	StatePlay  = 3
)

func JudgeState1(state int) {
}

func JudgeState(state int) int {
	// 如果开启则，进行播放
	if state == StateOpen {
		state = StatePlay
		fmt.Println("开始播放")
	} else if state == StateClose {
		// 如果关闭打开
		state = StateClose
		fmt.Println("关闭")
	} else {
		// 如果播放，则关闭
		state = StateClose
		fmt.Println("播放中")
	}
	return state
}

type StateContext struct {
	st IStater
}

func (s *StateContext) ToOpen(c *StateContext) {
	s.st.ToOpen(c)
}

func (s *StateContext) ToClose(c *StateContext) {
	s.st.ToClose(c)
}

func (s *StateContext) ToPlay(c *StateContext) {
	s.st.ToPlay(c)
}
func (s *StateContext) CurrentState() string {
	return s.st.CurrentState()
}

// 改写该没有if else 形式
func (s *StateContext) ChangeSate(st IStater) {
	s.st = st
}

type IStater interface {
	ToOpen(c *StateContext)
	ToClose(c *StateContext)
	ToPlay(c *StateContext)
	CurrentState() string
}

type OpenState struct {
	c *StateContext
}

func (s *OpenState) ToOpen(c *StateContext) {}

func (s *OpenState) ToClose(c *StateContext) {
	s.c = c
	fmt.Println("关闭")
	s.c.ChangeSate(&CloseState{})
}
func (s *OpenState) ToPlay(c *StateContext) {
	s.c = c
	fmt.Println("播放中")
	s.c.ChangeSate(&PlayState{})
}
func (s *OpenState) CurrentState() string {
	return "open"
}

type CloseState struct {
	c *StateContext
}

func (s *CloseState) ToOpen(c *StateContext) {
	s.c = c
	fmt.Println("打开中")
	s.c.ChangeSate(&OpenState{})
}

func (s *CloseState) ToClose(c *StateContext) {
}
func (s *CloseState) ToPlay(c *StateContext) {}
func (s *CloseState) CurrentState() string {
	return "close"
}

type PlayState struct {
	c *StateContext
}

func (s *PlayState) ToOpen(c *StateContext) {}

func (s *PlayState) ToClose(c *StateContext) {
	s.c = c
	fmt.Println("关闭中")
	s.c.ChangeSate(&CloseState{})
}
func (s *PlayState) ToPlay(c *StateContext) {
	panic("state err")
}

func (s *PlayState) CurrentState() string {
	return "play"
}

func NewStateFactory(state string) *StateContext {
	var ctx *StateContext
	switch state {
	case "open":
		ctx = &StateContext{&OpenState{}}
	case "close":
		ctx = &StateContext{&CloseState{}}
	case "play":
		ctx = &StateContext{&CloseState{}}
	}
	return ctx
}
