package queue

import (
	"crypto/rand"
	"encoding/base32"
)

func (q *Queue) GetHash() string {
	randomBytes := make([]byte, 32)
	rand.Read(randomBytes)
	return base32.StdEncoding.EncodeToString(randomBytes)[:32]
}

func (q *Queue) HashExists(hash string) bool {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	_, ok := q.Hashes[hash]
	return ok
}

func (q *Queue) SetHash(hash string) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	q.Hashes[hash] = true
}

func (q *Queue) ClearHash(hash string) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	delete(q.Hashes, hash)
}
