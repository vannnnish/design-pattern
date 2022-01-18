package main

/*
	使用场景: 当系统中有一系列的对象需要相互调用的时候，
	这个时候为了弱化对象之间的依赖关系，使得对象之间松耦合，就可使用中介者模式

	比如说， A 对象需要调用B 的方法， 然后再调用C的方法，
	这样写的话， ABC 之间的耦合就非常的严重， 此时，我们引入，中介者，
	他拥有ABC三个对象， 这时由中介者去 调节ABC三者的调用的关系，则可以降低，ABC之间的耦合。
*/
func main() {
	instance := GetMediatorInstance()
	instance.changed(&CPU{})
}
