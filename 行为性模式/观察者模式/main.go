package main

func main() {
	subject := NewSubject()
	r1 := NewReader("lixiang")
	r2 := NewReader("yangjie")
	r3 := NewReader("hupeng")
	subject.Attach(r1)
	subject.Attach(r2)
	subject.UpdateContext("妹子来了")
	subject.Attach(r3)
	subject.UpdateContext("妹子来了")
}
