package sort

import (
	"reflect"
	"testing"
)

var input = []int{10, 4, 6, 1, 7, 9}
var output = []int{1, 4, 6, 7, 9, 10}

func TestBubbleSort(t *testing.T) {
	bubbleSort(input)
	if !reflect.DeepEqual(input, output) {
		t.Errorf("bubble sort failed, expected: %v, infact: %v", output, input)
	}
}

func TestInsertionSort(t *testing.T) {
	insertionSort(input)
	if !reflect.DeepEqual(input, output) {
		t.Errorf("insertion sort failed, expected: %v, infact: %v", output, input)
	}
}

func TestSelectionSort(t *testing.T) {
	selectionSort(input)
	if !reflect.DeepEqual(input, output) {
		t.Errorf("selection sort failed, expected: %v, infact: %v", output, input)
	}
}
