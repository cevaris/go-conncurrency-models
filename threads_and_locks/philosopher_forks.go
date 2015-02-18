package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var NUM_OF_PHILOSOPHERS = 5

type Fork struct {
	Id int
	mutex *sync.Mutex
}

type Philosopher struct {
	Id int
	First *Fork
	Second *Fork
	ThinkCount int64
}
// Philosoper constructor
// Accepts two Forks and a unique Integer
func NewPhilosopher(left *Fork, right *Fork, id int) *Philosopher {
	p := &Philosopher{ Id: id, ThinkCount: 0 }

	// Bread and butter, solution to Philosopeher Dining problem
	if left.Id < right.Id {
		p.First = left
		p.Second = right
	} else {
		p.First = right
		p.Second = left
	}
	return p
}

// When 2 forks are free, acquire forks and start eating
// After eating, think for a random period of time, and repeat
func (p *Philosopher) run() {
	for { // Forever

		// Think some
		p.ThinkCount += 1
		if p.ThinkCount % 5 == 0 {
			fmt.Println("Philosopher", p.Id, "has thought", p.ThinkCount, "times")
		}
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))

		// Attempt and acquire locks two froks
		p.First.mutex.Lock()
		p.Second.mutex.Lock()

		// Eat some
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))

		// Release forks, and repeat
		p.First.mutex.Unlock()
		p.Second.mutex.Unlock()
	}
}


func main() {
	var wg sync.WaitGroup
	wg.Add(NUM_OF_PHILOSOPHERS)
	
	// Create Forks
	f1 := &Fork{mutex: &sync.Mutex{}, Id: 1}
	f2 := &Fork{mutex: &sync.Mutex{}, Id: 2}

	// Add N Philosopers
	for i := 0; i < NUM_OF_PHILOSOPHERS; i++ {
		// Spawn new thread for each Philosopher and start think/eating
		go NewPhilosopher(f1, f2, i).run()
	}
	 
	// Run indefinitely
	wg.Wait()
}
