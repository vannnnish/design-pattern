package main

func main() {
	instance := GetMediatorInstance()
	instance.changed(&CPU{})
}
