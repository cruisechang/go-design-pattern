package main

import "sync"

type List interface {
	next() interface{}
	previous() interface{}
	last() interface{}
	first() interface{}
}

type list struct {
	mutex     *sync.RWMutex
	container []interface{}
	index     int
}

func (s *list) first() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.container) > 0 {
		s.index = 0
		return s.container[0]
	}
	return nil
}
func (s *list) next() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.container) >= s.index+1 {
		s.index++
		return s.container[s.index]
	}
	return nil
}
func (s *list) previous() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if (s.index - 1) >= 0 {
		s.index--
		return s.container[s.index]
	}
	return nil
}
func (s *list) last() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	l := len(s.container)
	if l > 0 {
		s.index = l - 1 //last index
		return s.container[s.index]
	}
	return nil
}
func newList() List {
	return &list{
		index:     0,
		mutex:     &sync.RWMutex{},
		container: []interface{}{},
	}
}

func main() {

	l := newList()
	f := l.first()

}
