package main


type OpenTvCommand struct {
	tv *TV
}

func (o *OpenTvCommand) Execute() {
	o.tv.Open()
}

type CloseTvCommand struct {
	tv *TV
}

func (c *CloseTvCommand) Execute() {
	c.tv.Close()
}

type ChangeTvCommand struct {
	tv *TV
}

func (c *ChangeTvCommand) Execute() {
	c.tv.Change()
}

