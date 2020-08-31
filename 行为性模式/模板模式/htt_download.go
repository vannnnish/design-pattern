package main

import "fmt"

type HttpDownloader struct {
	tmp *template
}

func NewHttpDownload() Downloader {
	downloader := &HttpDownloader{}
	return &HttpDownloader{
		tmp: NewTemplate(downloader),
	}
}

func (hd *HttpDownloader) download() {
	fmt.Printf("http down %s", hd.tmp.url)
}

func (hd *HttpDownloader) save() {
	fmt.Printf("http save %s\n", hd.tmp.url)
}

func (hd *HttpDownloader) Download(url string) {
	hd.tmp.url = url
	hd.download()
	hd.save()
}
