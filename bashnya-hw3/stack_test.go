package main

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()

	if stack == nil {
		t.Error("NewStack() вернул nil")
	}

	if !stack.IsEmpty() || stack.Size() != 0 {
		t.Error("Новый стек должен быть пустым")
	}
}

func TestStack_Push(t *testing.T) {
	tests := []struct {
		name         string
		inputs       []int
		expectedSize int
	}{
		{"One element", []int{10}, 1},
		{"Multiple elements", []int{10, 20, 30}, 3},
		{"Negative numbers", []int{-1, -2, -3}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack()

			for _, input := range tt.inputs {
				stack.Push(input)
			}

			if stack.Size() != tt.expectedSize {
				t.Errorf("Size() = %d, expected %d", stack.Size(), tt.expectedSize)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name         string
		setup        func(*Stack)
		expectedVal  int
		expectedErr  error
		expectedSize int
	}{
		{
			name:         "Empty stack",
			setup:        func(s *Stack) {},
			expectedVal:  0,
			expectedErr:  ErrStackEmpty,
			expectedSize: 0,
		},
		{
			name: "Single element",
			setup: func(s *Stack) {
				s.Push(42)
			},
			expectedVal:  42,
			expectedErr:  nil,
			expectedSize: 0,
		},
		{
			name: "Multiple elements",
			setup: func(s *Stack) {
				s.Push(10)
				s.Push(20)
				s.Push(30)
			},
			expectedVal:  30,
			expectedErr:  nil,
			expectedSize: 2,
		},
		{
			name: "After multiple pops",
			setup: func(s *Stack) {
				s.Push(1)
				s.Push(2)
				s.Pop()
			},
			expectedVal:  1,
			expectedErr:  nil,
			expectedSize: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack()
			tt.setup(stack)

			val, err := stack.Pop()

			if err != tt.expectedErr {
				t.Errorf("Pop() error = %v, expectedErr %v", err, tt.expectedErr)
			}

			if val != tt.expectedVal {
				t.Errorf("Pop() = %d, expected %d", val, tt.expectedVal)
			}

			if stack.Size() != tt.expectedSize {
				t.Errorf("Size() after Pop = %d, expected %d", stack.Size(), tt.expectedSize)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name         string
		setup        func(*Stack)
		expectedVal  int
		expectedErr  error
		expectedSize int
	}{
		{
			name:         "Empty stack",
			setup:        func(s *Stack) {},
			expectedVal:  0,
			expectedErr:  ErrStackEmpty,
			expectedSize: 0,
		},
		{
			name: "Single element",
			setup: func(s *Stack) {
				s.Push(100)
			},
			expectedVal:  100,
			expectedErr:  nil,
			expectedSize: 1,
		},
		{
			name: "Multiple elements",
			setup: func(s *Stack) {
				s.Push(1)
				s.Push(2)
				s.Push(3)
			},
			expectedVal:  3,
			expectedErr:  nil,
			expectedSize: 3,
		},
		{
			name: "After Pop",
			setup: func(s *Stack) {
				s.Push(10)
				s.Push(20)
				s.Pop()
			},
			expectedVal:  10,
			expectedErr:  nil,
			expectedSize: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack()
			tt.setup(stack)

			initialSize := stack.Size()
			val, err := stack.Peek()

			if err != tt.expectedErr {
				t.Errorf("Peek() error = %v, expectedErr %v", err, tt.expectedErr)
			}

			if val != tt.expectedVal {
				t.Errorf("Peek() = %d, expected %d", val, tt.expectedVal)
			}

			if stack.Size() != initialSize {
				t.Errorf("Size() changed after Peek: %d -> %d", initialSize, stack.Size())
			}

			if stack.Size() != tt.expectedSize {
				t.Errorf("Size() = %d, expected %d", stack.Size(), tt.expectedSize)
			}
		})
	}
}
