package main

import "sync"

var (
	imgflyfactory *ImageFlyWeightFactory
	doOnce        sync.Once
)

type ImageFlyWeightFactory struct {
	maps map[string]*ImageFlyWeight
}

//
func GetImageFlyWeightFactory() *ImageFlyWeightFactory {
	if imgflyfactory == nil {
		doOnce.Do(func() {
			imgflyfactory = &ImageFlyWeightFactory{maps: make(map[string]*ImageFlyWeight)}
		})
	}
	return imgflyfactory
}

func (imf *ImageFlyWeightFactory) Get(filename string) *ImageFlyWeight {
	image := imf.maps[filename]
	if image == nil {
		image = NewImageFlyweight(filename)
		imf.maps[filename] = image
	}
	return image
}
