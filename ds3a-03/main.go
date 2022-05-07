package main

import (
	"errors"
	"fmt"
)

type queue struct {
	//buf存放队列元素
	buf []interface{}
	//队列大小
	Size int
}

func newQueue(size int) *queue {
	return &queue{
		buf:  make([]interface{},0),
		Size: size,
	}
}

func (q *queue) enqueue(elem int) (err error) {
	if len(q.buf) < q.Size {
		q.buf = append(q.buf, elem)
		return nil
	}
	return errors.New("q 满")
}

func (q *queue) dequeue() interface{} {
    if len(q.buf) == 0 {
        return nil
	}
	res := q.buf[0]
	q.buf = q.buf[1:]
	return res
}

func (q *queue) Print() {
	for _,elem := range q.buf {
		fmt.Println(elem)

	}
}

func main() {

	q := newQueue(5)

	q.enqueue(5)

	// s := q.dequeue()
    // fmt.Println(s)

	q.Print()
}
