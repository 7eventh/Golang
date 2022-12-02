// Stefan Nilsson 2013-03-13
// This is a testbed to help you understand channels better.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

/*
// OG CODE
	func main() {
		// Use different random numbers each time this program is executed.
		rand.Seed(time.Now().Unix())

		const strings = 32
		const producers = 4
		const consumers = 2

		before := time.Now()
		ch := make(chan string)
		wgp := new(sync.WaitGroup)
		wgp.Add(producers)
		for i := 0; i < producers; i++ {
			go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
		}
		for i := 0; i < consumers; i++ {
			go Consume("c"+strconv.Itoa(i), ch)
		}
		wgp.Wait() // Wait for all producers to finish.
		close(ch)
		fmt.Println("time:", time.Now().Sub(before))
	}

// Produce sends n different strings on the channel and notifies wg when done.

	func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
		for i := 0; i < n; i++ {
			RandomSleep(100) // Simulate time to produce data.
			ch <- id + ":" + strconv.Itoa(i)
		}
		wg.Done()
	}

// Consume prints strings received from the channel until the channel is closed.

	func Consume(id string, ch <-chan string) {
		for s := range ch {
			fmt.Println(id, "received", s)
			RandomSleep(100) // Simulate time to consume data.
		}
	}

// RandomSleep waits for x ms, where x is a random number, 0 ≤ x < n,
// and then returns.

	func RandomSleep(n int) {
		time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
	}
*/

// ------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------

// Q: What happens if you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function?
// A: Program has a panic attack.
// ##### ERROR #####
// panic: send on closed channel
// goroutine 6 [running]:
// main.Produce({0xc000016088, 0x2}, 0x8, 0x0?, 0x0?)
// c:/Users/leo01/OneDrive/Desktop/Code/Golang/School/uppgift_5.go:109 +0x9f
// created by main.main
// c:/Users/leo01/OneDrive/Desktop/Code/Golang/School/uppgift_5.go:95 +0xce
// exit status 2

/*
func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 2

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgp.Add(producers)
	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch)
	}
	close(ch)
	wgp.Wait() // Wait for all producers to finish.
	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}
	wg.Done()
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string) {
	for s := range ch {
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
}

// RandomSleep waits for x ms, where x is a random number, 0 ≤ x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}*/

// ------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------

// Q: What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?
// A: The last producer trys to sent data but the chanel is closed by the previous one. Program panics and shuts down.
// ##### ERROR 1 #####
// panic: send on closed channel
// goroutine 7 [running]:
// main.Produce({0xc00001608a, 0x2}, 0x8, 0x0?, 0x0?)
// c:/Users/leo01/OneDrive/Desktop/Code/Golang/School/uppgift_5.go:167 +0x9e
// created by main.main
// c:/Users/leo01/OneDrive/Desktop/Code/Golang/School/uppgift_5.go:154 +0xce
// panic: send on closed channel
// ##### ERROR 2 #####
// goroutine 8 [running]:
// main.Produce({0xc00001608c, 0x2}, 0x8, 0x0?, 0x0?)
// c:/Users/leo01/OneDrive/Desktop/Code/Golang/School/uppgift_5.go:167 +0x9e
// created by main.main
// c:/Users/leo01/OneDrive/Desktop/Code/Golang/School/uppgift_5.go:154 +0xce
// exit status 2

/*
func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 2

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgp.Add(producers)
	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch)
	}
	wgp.Wait() // Wait for all producers to finish.
	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}
	close(ch)
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string) {
	for s := range ch {
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
}

// RandomSleep waits for x ms, where x is a random number, 0 ≤ x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}*/

// ------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------

// Q: What happens if you remove the statement close(ch) completely?
// A: LGTM. Code runs smoothly. Since the channel is empty in the end and there are no more data left to take the program shuts down properly(?).

/*
func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 2

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgp.Add(producers)
	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch)
	}
	wgp.Wait() // Wait for all producers to finish.
	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}
	wg.Done()
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string) {
	for s := range ch {
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
}

// RandomSleep waits for x ms, where x is a random number, 0 ≤ x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}*/

// ------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------

// Q: What happens if you increase the number of consumers from 2 to 4?
// A: We just increase the consumers so now the data will be shared amongst the 4 consumers instead of only 2

/*
func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 4

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgp.Add(producers)
	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch)
	}
	wgp.Wait() // Wait for all producers to finish.
	close(ch)
	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}
	wg.Done()
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string) {
	for s := range ch {
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
}

// RandomSleep waits for x ms, where x is a random number, 0 ≤ x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}*/

// Q: Can you be sure that all strings are printed before the program stops?
// A: Idk programming is like a box of chocolates, you never know what you're going to get by the time you hit run. 

func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 2

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgp.Add(producers)
	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch)
	}
	wgp.Wait() // Wait for all producers to finish.
	close(ch)
	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}
	wg.Done()
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string) {
	for s := range ch {
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
}

// RandomSleep waits for x ms, where x is a random number, 0 ≤ x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}
