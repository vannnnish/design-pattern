package main

func main() {
	ctx := NewLeaderContext("Mayun", 53, &Manager{})
	ctx.Way()
}
