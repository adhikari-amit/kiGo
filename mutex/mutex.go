package mutex

import (
	"fmt"
	"sync"
)

var waitgroup sync.WaitGroup
var message string

func updatemessage(s string, m *sync.Mutex) {
	waitgroup.Done()
	m.Lock()
	message = s
	m.Unlock()
}

func Mutex() {
	message = "Hello World"
	var mutex sync.Mutex
	waitgroup.Add(2)
	go updatemessage("Hello Sun", &mutex)
	go updatemessage("Hello Moon", &mutex)
	waitgroup.Wait()

	fmt.Println(message)
}
