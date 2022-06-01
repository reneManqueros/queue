package queue

import (
	"sync"
)

type Queue struct {
	Items  []any
	Mutex  sync.Mutex
	Hashes map[string]bool
}

func (q *Queue) Push(thisItem any) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	q.Items = append(q.Items, thisItem)
}

func (q *Queue) Pop() (item any) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	item, q.Items = q.Items[0], q.Items[1:]
	return
}

func (q *Queue) Length() int {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	return len(q.Items)
}
