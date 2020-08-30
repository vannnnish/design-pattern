package main

import "fmt"

type ImageViewer struct {
	*ImageFlyWeight
}

func NewImageViewer(filename string)*ImageViewer{
	image:=GetImageFlyWeightFactory().Get(filename)
	return &ImageViewer{
		image,
	}
}

func(i *ImageViewer)Display(){
	fmt.Printf("Display %s\n",i.Data())
}