package main

import "errors"

type Stack struct {
	items []int
}

var ErrStackEmpty = errors.New("СТЕК ПУСТ")

func NewStack() *Stack {
	return &Stack{
		items: []int{},
	}
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrStackEmpty
	}

	index := len(s.items) - 1
	deleted := s.items[index]
	s.items = s.items[:index]
	return deleted, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Clear() {
	s.items = []int{}
}

func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrStackEmpty
	}

	return s.items[len(s.items)-1], nil
}
