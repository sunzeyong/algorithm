package queue

import "testing"

func TestCycleQueue(t *testing.T) {
	cycleQ := NewCycleQueue(11)

	for i := 0; i < 10; i++ {
		cycleQ.EnQueue(i)
	}

	if cycleQ.EnQueue(10) {
		t.Errorf("cycleQ should not EnQueue again")
	}

	for i := 0; i < 10; i++ {
		output, ok := cycleQ.DeQueue()
		if !ok {
			t.Errorf("cycleQ should get value")
		}
		if output != i {
			t.Errorf("cycleQ should get wrong value, expect: %v, actual: %v", i, output)
		}
	}

	_, ok := cycleQ.DeQueue()
	if ok {
		t.Errorf("cycleQ should not get value")
	}

}
