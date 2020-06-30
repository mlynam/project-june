package shared

import "sync"

// Task provides signaling channels for a goroutine
type Task struct {
	mutx   sync.Mutex
	result interface{}

	Done   chan bool
	Error  error
	Status chan string
}

// NewTask creates an empty task
func NewTask() Task {
	return Task{
		Done:   make(chan bool),
		Error:  nil,
		Status: make(chan string),
	}
}

// SetResult safely sets the result no matter the calling routine
func (t *Task) SetResult(any interface{}) {
	t.mutx.Lock()
	t.result = any
	t.mutx.Unlock()
}

// Result returns the current result for the task
func (t *Task) Result() interface{} {
	return t.result
}
