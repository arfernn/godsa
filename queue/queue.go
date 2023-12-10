package queue

import (
	"fmt"
	"godsa/slinkedlist"
)

type Queue[T any] struct {
	llist *slinkedlist.LinkedList[T]
}

func New[T any](inputs ...T) *Queue[T] {
	temp := &Queue[T]{}
	temp.llist = slinkedlist.New[T]()
	temp.Push(inputs...)

	return temp
}

func (q *Queue[T]) Push(inputs ...T) {
	q.llist.Append(inputs...)
	fmt.Println("hello")
}

func (q *Queue[T]) Pop() (T, error) {
	val, err := q.llist.Get(0)
	if err != nil {
		err = fmt.Errorf("Pop failed, queue is empty")
	} else {
		q.llist.Remove(0)
	}

	return val, err
}

func (q *Queue[T]) Peek() (T, error) {
	val, err := q.llist.Get(0)
	if err != nil {
		err = fmt.Errorf("Peek failed, queue is empty")
	}
	return val, err
}

func (q *Queue[T]) ToArray() []T {
	return q.llist.ToArray()
}
