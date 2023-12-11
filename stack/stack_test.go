package stack

import (
	"slices"
	"testing"
)

func TestPush(t *testing.T) {

	type test struct {
		input     int
		inputArr  []int
		outputArr []int
	}

	tests := []test{{
		input:     2,
		inputArr:  []int{1, 3, 4, 5, 6},
		outputArr: []int{2, 6, 5, 4, 3, 1},
	}, {
		input:     2,
		inputArr:  []int{},
		outputArr: []int{2},
	}}

	for _, test := range tests {
		stack := New(test.inputArr...)

		stack.Push(test.input)

		if !slices.Equal(test.outputArr, stack.ToArray()) {
			t.Errorf("Error, got %v expected %v", stack.ToArray(), test.outputArr)
		}
	}

}

func TestSinglePush(t *testing.T) {
	testStack := New([]int{}...)
	testStack.Push(1)
	testStack.Push(2)
	testStack.Push(3)

	val, _ := testStack.Pop()
	if val != 3 {
		t.Errorf("Unexpected result %d", val)
	}

	val, _ = testStack.Pop()
	if val != 2 {
		t.Errorf("Unexpected result %d", val)
	}

	val, _ = testStack.Pop()
	if val != 1 {
		t.Errorf("Unexpected result %d", val)
	}
}

func TestPeek(t *testing.T) {
	type test struct {
		inputArr []int
		output   int
	}
	tests := []test{{
		inputArr: []int{1, 3, 4, 5, 6},
		output:   6,
	}, {
		inputArr: []int{},
		output:   0,
	}}

	for _, test := range tests {
		stack := New(test.inputArr...)

		val, _ := stack.Peek()

		if val != test.output {
			t.Errorf("Error, got %d expected %d", test.output, val)
		}
	}

}

func TestPop(t *testing.T) {
	type test struct {
		input  []int
		output []int
	}
	tests := []test{{
		input: []int{1, 3, 4, 5, 6},
	}, {
		input: []int{},
	}}

	for _, test := range tests {
		stack := New(test.input...)
		temp := []int{}
		for {
			val, err := stack.Pop()
			if err != nil {
				break
			}
			temp = append(temp, val)
		}

		slices.Reverse(test.input)
		if !slices.Equal(temp, test.input) {
			t.Errorf("Error, got %v expected %v", temp, test.input)
		}
	}
}
