package main

// 命令的执行者：向TV发起命令 , 也就是调用者
type TVRemote struct {
	open   *OpenTvCommand
	change *ChangeTvCommand
	close  *CloseTvCommand
}

func (tv *TVRemote) Open() {
	tv.open.Execute()
}
func (tv *TVRemote) Change() {
	tv.change.Execute()
}
func (tv *TVRemote) Close() {
	tv.close.Execute()
}

func NewTVRemote() *TVRemote {
	// 创建接收者
	rece := &TV{}
	// 创建命令对象
	openComm := &OpenTvCommand{rece}
	changeComm := &ChangeTvCommand{rece}
	closeComm := &CloseTvCommand{rece}

	// 创建请求者，把命令对象设置进去
	return &TVRemote{
		open:   openComm,
		change: changeComm,
		close:  closeComm,
	}
}
