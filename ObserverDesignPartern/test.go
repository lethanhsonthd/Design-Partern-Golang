package main

import (
	"fmt"
	"strconv"
)

type ObserverInterface interface {
	Update()
}
type Observer struct {
	ObserverInterface
	subject Subject
}

func (o *Observer) Update() {
}

type Subject struct {
	Observers []ObserverInterface
	State     int
}

func (s *Subject) getState() int {
	return s.State
}
func (s *Subject) setState(state int) {
	s.State = state
}
func (s *Subject) attach(observer Observer) {
	s.Observers = append(s.Observers, observer)
}
func (s *Subject) notifyAllObservers() {
	for _, element := range s.Observers {
		element.Update()
	}
}

type BinaryObserver struct {
	Observer
}

func (b *BinaryObserver) Update() {
	fmt.Println("Binary String: " + strconv.FormatInt(int64(b.subject.getState()), 2))
}
func (b *BinaryObserver) BinaryObserver(subject Subject) {
	b.subject = subject
	b.subject.attach(Observer(b)) // khong ép kiểu được
}
func main() {
	fmt.Println(strconv.FormatInt(255, 8))
}
