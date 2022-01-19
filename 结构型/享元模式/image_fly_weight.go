package main

import "fmt"

type ImageFlyWeight struct {
	data string
}

func NewImageFlyweight(filename string) *ImageFlyWeight {
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyWeight{data: data}
}

func (i *ImageFlyWeight) Data() string {
	return i.data
}
