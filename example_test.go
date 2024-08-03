package queue

import "fmt"

func ExampleQueue_Add() {
	s := Queue[int]{}
	s.Add(1)
	s.Add(2)
	s.Add(3)
}

func ExampleQueue_Remove() {
	s := Queue[int]{}
	s.Add(1)
	s.Add(2)

	fmt.Println(s.Remove())
	// Output: 1
}

func ExampleQueue_Peek() {
	s := Queue[int]{}
	s.Add(1)
	s.Add(2)

	fmt.Println(s.Peek())
	// Output: 1
}

func ExampleQueue_IsEmpty() {
	s := Queue[int]{}
	s.Add(1)
	s.Add(2)

	fmt.Println(s.IsEmpty())
	// Output: false

	s.Remove()
	s.Remove()

	fmt.Println(s.IsEmpty())
	// Output: true
}
