package list

import "fmt"

type LinkedList[T any] struct {
	firstNode *listNode[T]
	length    int
	tail      int
}

type listNode[T any] struct {
	value T
	next  *listNode[T]
}

func ToList[T any](input []T) *LinkedList[T] {
	newList := &LinkedList[T]{}
	for _, num := range input {
		newList.Append(num)
	}
	return newList
}

func (l *LinkedList[T]) Get(position int) (T, error) {
	_, target, error := l.traverse(position)
	return target.value, error
}

func (l *LinkedList[T]) Remove(position int) error {
	previous, target, error := l.traverse(position)
	l.length--
	if previous != nil {
		previous.next = target.next
	} else {
		l.firstNode = target.next
	}

	return error
}

func (l *LinkedList[T]) Insert(position int, value T) error {
	l.length++
	prev, curr, error := l.traverse(position)
	var new = &listNode[T]{value: value, next: curr}
	if prev == nil {
		l.firstNode = new
	} else {
		prev.next = new
	}
	return error
}

func (l *LinkedList[T]) Prepend(value T) {
	l.length++
	newNode := &listNode[T]{value: value, next: l.firstNode}
	l.firstNode = newNode
}

func (l *LinkedList[T]) Append(value T) {
	newNode := &listNode[T]{value: value, next: nil}
	if l.length == 0 {
		l.firstNode = newNode
	} else {

		l.firstNode.last().next = newNode
	}
	l.length++
}

func (l *LinkedList[T]) ToArray() []T {
	array := make([]T, 0, l.length)

	if l.length > 0 {
		currentNode := l.firstNode
		array = append(array, currentNode.value)

		for currentNode.next != nil {
			currentNode = currentNode.next
			array = append(array, currentNode.value)
		}
	}
	return array
}

func (l *LinkedList[T]) String() string {
	return fmt.Sprint(l.ToArray())
}

func (l *listNode[T]) last() *listNode[T] {
	var result *listNode[T] = l
	for result.next != nil {
		result = result.next
	}

	return result
}

func (l *LinkedList[T]) traverse(position int) (*listNode[T], *listNode[T], error) {
	if position >= l.length || position < 0 {
		return nil, nil, fmt.Errorf("%d out of range", position)
	}

	var current *listNode[T] = l.firstNode
	var prev *listNode[T] = nil
	for i := 0; i < position; i++ {
		prev = current
		current = current.next
	}

	return prev, current, nil
}
