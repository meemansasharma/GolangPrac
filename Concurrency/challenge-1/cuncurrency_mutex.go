package main

import "sync"

var count int

func mutexExample() {

	count = 0
	var mutex sync.Mutex
	wg.Add(1)
	go updateCount(2, &mutex)
	wg.Wait()

	wg.Add(1)
	go updateCount(7, &mutex)
	wg.Wait()

	wg.Add(1)
	go updateCount(9, &mutex)
	wg.Wait()

}
func updateCount(x int, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	count += x
	mutex.Unlock()
}
