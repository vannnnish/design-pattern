package main

import "fmt"

type Sunday struct {
}

func (sd *Sunday) Today() {
	fmt.Println("Sunday")
}

func (sd *Sunday) Next(ctx *DayContext) {
	ctx.today = &Monday{}
}

type Monday struct {
}

func (sd *Monday) Today() {
	fmt.Println("Monday")
}

func (sd *Monday) Next(ctx *DayContext) {
	ctx.today = &Friday{}
}

type Friday struct {
}

func (sd *Friday) Today() {
	fmt.Println("Friday")
}

func (sd *Friday) Next(ctx *DayContext) {
	ctx.today = &Sunday{}
}
