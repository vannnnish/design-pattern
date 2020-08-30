package main


type Observer interface {
	Update(subject *Subject)
}
