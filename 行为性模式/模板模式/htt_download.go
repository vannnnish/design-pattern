package main

import "fmt"

type HttpDownloader struct {
	 *template
}

func NewHttpDownload() Downloader {
	downloader := &HttpDownloader{}
	template:=NewTemplate(downloader)
	downloader.template=template
	return downloader
}

func (hd *HttpDownloader) download() {
	fmt.Printf("http down %s", hd.url)
}

func (hd *HttpDownloader) save() {
	fmt.Printf("http save %s\n", hd.url)
}
