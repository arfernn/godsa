package list

import "fmt"

type LinkedList struct {
	firstNode *listNode
	length    int
}

type listNode struct {
	value int
	next  *listNode
}

func ToList(input []int) *LinkedList {
	newList := &LinkedList{}
	for _, num := range input {
		newList.Append(num)
	}
	return newList
}

func (l *LinkedList) Get(position int) (int, error) {
	_, target, error := l.traverse(position)
	return target.value, error
}

func (l *LinkedList) Remove(position int) error {
	previous, target, error := l.traverse(position)
	l.length--
	if previous != nil {
		previous.next = target.next
	} else {
		l.firstNode = target.next
	}

	return error
}

func (l *LinkedList) Insert(position int, value int) error {
	l.length++
	prev, curr, error := l.traverse(position)
	var new = &listNode{value: value, next: curr}
	if prev == nil {
		l.firstNode = new
	} else {
		prev.next = new
	}
	return error
}

func (l *LinkedList) Prepend(value int) {
	l.length++
	newNode := &listNode{value: value, next: l.firstNode}
	l.firstNode = newNode
}

func (l *LinkedList) Append(value int) {
	newNode := &listNode{value: value, next: nil}
	if l.length == 0 {
		l.firstNode = newNode
	} else {

		l.firstNode.last().next = newNode
	}
	l.length++
}

func (l *LinkedList) ToArray() []int {
	array := make([]int, 0, l.length)

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

func (l *LinkedList) String() string {
	return fmt.Sprint(l.ToArray())
}

func (l *listNode) last() *listNode {
	var result *listNode = l
	for result.next != nil {
		result = result.next
	}

	return result
}

func (l *LinkedList) traverse(position int) (*listNode, *listNode, error) {
	if position >= l.length || position < 0 {
		return nil, nil, fmt.Errorf("%d out of range", position)
	}

	var current *listNode = l.firstNode
	var prev *listNode = nil
	for i := 0; i < position; i++ {
		prev = current
		current = current.next
	}

	return prev, current, nil
}
