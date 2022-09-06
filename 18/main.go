package main

import (
	"fmt"
	"sync"
)

func main() {
	c := 0
	n := 100
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			m.Lock()
			c++
			m.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(c)
}

/* sync with channel
c := 0
n := 100

ch := make(chan struct{}, n)
chanWg := sync.WaitGroup{}
chanWg.Add(1)
go func() {
	for range ch {
		c++
	}
	chanWg.Done()
}()

wg := sync.WaitGroup{}
wg.Add(n)
for i := 0; i < n; i++ {
	go func(i int) {
		ch <- struct{}{}
		wg.Done()
	}(i)
}
wg.Wait()
close(ch)
chanWg.Wait()

fmt.Println(c)
*/

/* sync with atomic
c := int32(0)
n := 200

wg := sync.WaitGroup{}
wg.Add(n)
for i := 0; i < n; i++ {
	go func(i int) {
		atomic.AddInt32(&c, 1)
		wg.Done()
	}(i)
}
wg.Wait()

fmt.Println(c)
*/
