package util

import "testing"

func TestQueue(t *testing.T) {

	q := Initialize()
	Enqueue(q, "a")
	Enqueue(q, "b")

	a := Dequeue(q)
	b := Dequeue(q)

	if a != "a" {
		t.Errorf("a != %s", a)
	}
	if b != "b" {
		t.Errorf("b != %s", b)
	}
}

func producer(q *Queue, n int) {
	for i := 0; i < n; i++ {
		Enqueue(q, "a")
	}
}

func consumer(q *Queue, c chan int) {
	Dequeue(q)
}

func TestQueueCosume(t *testing.T) {
	q := Initialize()
	c := make(chan int)
	go producer(q, 100)
	go consumer(q)
}
