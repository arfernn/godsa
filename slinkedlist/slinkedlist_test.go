package slinkedlist

import (
	"fmt"
	"slices"
	"testing"
)

func TestGetStr(t *testing.T) {
	type test[T any] struct {
		input  []T
		index  int
		output T
	}
	inputStrArray := []string{"hello", "how", "you", "doing"}
	tests := []test[string]{
		{
			input:  inputStrArray,
			index:  1,
			output: "how",
		},
	}
	for _, test := range tests {
		list := New(test.input...)
		result, err := list.Get(test.index)
		if err != nil {
			t.Error(err)
		}
		if result != test.output {
			t.Errorf("Error, got %s expected %s", test.output, result)
		}
	}

}

func TestGet(t *testing.T) {
	type test struct {
		input  []int
		index  int
		output int
	}
	inputArray := []int{0, 11, 2, 55, 2, 3, 123}
	tests := []test{
		{
			input:  inputArray,
			index:  1,
			output: 11,
		},
		{
			input:  inputArray,
			index:  3,
			output: 55,
		},
		{
			input:  inputArray,
			index:  0,
			output: 0,
		},
		{
			input:  inputArray,
			index:  6,
			output: 123,
		},
	}

	for _, test := range tests {
		list := New(test.input...)
		result, err := list.Get(test.index)
		if err != nil {
			t.Error(err)
		}
		if result != test.output {
			t.Errorf("Error, got %d expected %d", test.output, result)
		}
	}

}

func TestRemove(t *testing.T) {

	type test struct {
		array    []int
		position int
		output   []int
	}

	tests := []test{
		{
			array:    []int{0, 11, 2, 55, 2, 3, 123},
			position: 1,
			output:   []int{0, 2, 55, 2, 3, 123},
		},
		{
			array:    []int{0, 11, 2, 55, 2, 3, 123},
			position: 0,
			output:   []int{11, 2, 55, 2, 3, 123},
		},
		{
			array:    []int{0, 11, 2, 55, 2, 3, 123},
			position: 6,
			output:   []int{0, 11, 2, 55, 2, 3},
		},
		{
			array:    []int{0},
			position: 0,
			output:   []int{},
		},
	}

	for _, test := range tests {

		linkedList := New(test.array...)

		linkedList.Remove(test.position)
		if !slices.Equal(linkedList.ToArray(), test.output) {
			t.Errorf("Error, %v not equal to %v", test.array, test.output)
		}
	}

}
func TestPrepend(t *testing.T) {

	type test struct {
		array  []int
		toadd  int
		output []int
	}

	tests := []test{
		{
			array:  []int{0, 11, 2, 55, 2, 3, 123},
			toadd:  1,
			output: []int{1, 0, 11, 2, 55, 2, 3, 123},
		},
		{
			array:  []int{0},
			toadd:  1,
			output: []int{1, 0},
		},
		{
			array:  []int{},
			toadd:  1,
			output: []int{1},
		},
	}

	for _, test := range tests {
		linkedList := New(test.array...)
		linkedList.Prepend(test.toadd)

		if !slices.Equal(linkedList.ToArray(), test.output) {
			t.Errorf("Error, %v not equal to %v", linkedList.ToArray(), test.output)
		}
	}

}
func TestInsert(t *testing.T) {

	type test struct {
		array    []int
		toadd    int
		position int
		output   []int
	}

	tests := []test{
		{
			array:    []int{0, 11, 2, 55, 2, 3, 123},
			toadd:    1,
			position: 0,
			output:   []int{1, 0, 11, 2, 55, 2, 3, 123},
		}, {
			array:    []int{0, 11, 2, 55, 2, 3, 123},
			toadd:    1,
			position: 1,
			output:   []int{0, 1, 11, 2, 55, 2, 3, 123},
		},
		{
			array:    []int{0, 11, 2, 55, 2, 3, 123},
			toadd:    1,
			position: 6,
			output:   []int{0, 11, 2, 55, 2, 3, 1, 123},
		},
		{
			array:    []int{0},
			toadd:    1,
			position: 0,
			output:   []int{1, 0},
		},
	}

	for _, test := range tests {
		linkedList := New(test.array...)

		linkedList.Insert(test.position, test.toadd)

		if !slices.Equal(linkedList.ToArray(), test.output) {
			t.Errorf("Error, %v not equal to %v", linkedList.ToArray(), test.output)
		}
	}

}
func TestToArray(t *testing.T) {
	type test struct {
		array []int
	}

	tests := []test{
		{
			array: []int{0, 11, 2, 55, 2, 3, 123},
		},
		{
			array: []int{0},
		}, {
			array: []int{},
		},
	}

	for _, test := range tests {
		linkedList := New(test.array...)

		if !slices.Equal(linkedList.ToArray(), test.array) {
			t.Errorf("Error, %v not equal to %v", test.array, test.array)
		}
	}
}
func TestTraverse(t *testing.T) {
	input := New[int]()
	input.Append(2)
	input.Append(2)

	if input.length != 2 {
		t.Errorf("Unexpected length %d", input.length)
	}

	prev, cur, err := input.traverse(-1)

	if prev != nil || cur != nil || err == nil {
		t.Errorf("Traverse %v to pos %d should return an error", input, -1)
	}

	prev, cur, err = input.traverse(2)
	if prev != nil || cur != nil || err == nil {
		t.Errorf("Traverse %v to pos %d should return an error", input, 2)
	}

}
func TestToString(t *testing.T) {
	type test struct {
		input  []int
		output string
	}

	tests := []test{
		{
			input:  []int{0, 11, 2, 55, 2, 3, 123},
			output: "[0 11 2 55 2 3 123]",
		},
		{
			input:  []int{},
			output: "[]",
		},
	}

	for _, test := range tests {
		list := New(test.input...)
		result := fmt.Sprint(list)
		if result != test.output {
			t.Errorf("Error, expected %s got %s", test.output, result)
		}
	}
}
