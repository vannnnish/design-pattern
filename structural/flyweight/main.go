package main

/*
	享元模式

	节约内存的一种做法。
*/
func main() {
	v1 := NewImageViewer("1.jpg")
	v1.Display()
}
