package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	items map[string]string
	mu    sync.Mutex
}

func (s *Singleton) Set(key, data string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = data
}
func (s *Singleton) Get(key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, ok := s.items[key]
	if !ok {
		return "", fmt.Errorf("The %s is not presented", key)
	}
	return item, nil
}

var (
	s *Singleton
)

func NewSingleton() *Singleton {
	if s == nil {
		s = &Singleton{items: make(map[string]string)}
	}
	return s
}
func main() {
	t := NewSingleton()
	t.items["123"] = "456"
	fmt.Println(t.Get("123"))
}
