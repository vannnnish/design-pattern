package main

import "sync"

type Single struct {
}

var singleton *Single

var once sync.Once

func GetInterface() *Single {
	once.Do(func() {
		singleton = &Single{}
	})
	return singleton
}
