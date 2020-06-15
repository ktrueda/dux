package util

import (
	"fmt"
)

type Queue []string

func Initialize() *Queue {
	var q Queue = make(Queue, 0)
	return &q
}

func Enqueue(q *Queue, e string) {
	*q = append(*q, e)
}

func Dequeue(q *Queue) string {
	fmt.Println(*q)

	v := (*q)[0]
	*q = (*q)[1:]
	return v
}
