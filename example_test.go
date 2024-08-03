package queue

import "fmt"

func ExampleQueue_Add() {
	q := Queue[int]{}
	q.Add(1)
	q.Add(2)
	q.Add(3)
}

func ExampleQueue_Remove() {
	q := Queue[int]{}
	q.Add(1)
	q.Add(2)

	fmt.Println(q.Remove())
	// Output: 1
}

func ExampleQueue_Peek() {
	q := Queue[int]{}
	q.Add(1)
	q.Add(2)

	fmt.Println(q.Peek())
	// Output: 1
}

func ExampleQueue_IsEmpty() {
	q := Queue[int]{}
	q.Add(1)
	q.Add(2)

	fmt.Println(q.IsEmpty())
	// Output: false

	q.Remove()
	q.Remove()

	fmt.Println(q.IsEmpty())
	// Output: true
}
