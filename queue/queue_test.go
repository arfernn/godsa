package queue

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
		outputArr: []int{1, 3, 4, 5, 6, 2},
	}, {
		input:     2,
		inputArr:  []int{},
		outputArr: []int{2},
	}}

	for _, test := range tests {
		queue := New(test.inputArr...)

		queue.Push(test.input)

		if !slices.Equal(test.outputArr, queue.ToArray()) {
			t.Errorf("Error, got %v expected %v", queue.ToArray(), test.outputArr)
		}
	}

}

func TestPeek(t *testing.T) {
	type test struct {
		inputArr []int
		output   int
	}
	tests := []test{{
		inputArr: []int{1, 3, 4, 5, 6},
		output:   1,
	}, {
		inputArr: []int{},
		output:   0,
	}}

	for _, test := range tests {
		queue := New(test.inputArr...)

		val, _ := queue.Peek()

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
		queue := New(test.input...)
		temp := []int{}
		for {
			val, err := queue.Pop()
			if err != nil {
				break
			}
			temp = append(temp, val)
		}

		if !slices.Equal(temp, test.input) {
			t.Errorf("Error, got %v expected %v", temp, test.input)
		}
	}
}
