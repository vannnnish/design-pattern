package main

import "fmt"

type FtpDownloader struct {
	*template
}

func NewFtpDownload() Downloader {
	downloader := &FtpDownloader{}
	template := NewTemplate(downloader)
	downloader.template = template
	return downloader
}

func (hd *FtpDownloader) download() {
	fmt.Printf("ftp down %s", hd.url)
}

func (hd *FtpDownloader) save() {
	fmt.Printf("ftp save %s\n", hd.url)
}
