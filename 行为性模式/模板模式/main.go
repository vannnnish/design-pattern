package main

func main() {
	var httpDownloader Downloader = NewHttpDownload()
	httpDownloader.Download("http://tsinghua.edu.cn/ai.zip")
	var ftpDownload Downloader = NewFtpDownload()
	ftpDownload.Download("https://www.baidu.com")
}
