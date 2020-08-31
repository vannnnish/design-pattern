package main

import "fmt"

type FtpDownloader struct {
	tmp *template
}

func NewFtpDownload() Downloader {
	downloader := &HttpDownloader{}
	return &FtpDownloader{
		tmp: NewTemplate(downloader),
	}
}

func (hd *FtpDownloader) download() {
	fmt.Printf("ftp down %s", hd.tmp.url)
}

func (hd *FtpDownloader) save() {
	fmt.Printf("ftp save %s\n", hd.tmp.url)
}

func (hd *FtpDownloader) Download(url string) {
	hd.tmp.url = url
	hd.download()
	hd.save()
}

