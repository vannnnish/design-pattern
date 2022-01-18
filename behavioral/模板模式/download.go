package main

import "fmt"

type Downloader interface {
	Download(url string)
}

type template struct {
	implement
	url string
}

type implement interface {
	download()
	save()
}

func NewTemplate(impl implement) *template {
	return &template{implement: impl}
}

func (t *template) Download(url string) {
	fmt.Println("模板下载 start")
	t.url = url
	t.implement.download()
	t.implement.save()
	fmt.Println("模板下载 end")
}
func (t *template) save() {
	fmt.Println("模板下载 保存")
}
