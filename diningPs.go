package main

import (
	"fmt"
	"sync"
	"math/rand"
)

type chopS struct{ sync.Mutex }

type philO struct {
	id string
	// required to make random which hand the philosopher is dominant with
	first, second *chopS
	meals int
}

func NewSticks() []*chopS {
	sticks := make([]*chopS, 5)
	for i := 0; i < 5; i++ {
		sticks[i] = new(chopS)
	}
	return sticks
}

func NewP(number int, Cs1, Cs2 *chopS) philO {
	var p philO
	randy := rand.Intn(2)
	// 
	if randy < 1 {
		p.first = Cs1
		p.second = Cs2
	} else {
		p.second = Cs1
		p.first = Cs2
	}
	p.meals = 0			
	switch number {
	case 0:
		p.id = "1 Rorty"
	case 1:
		p.id = "2 Confusius"
	case 2:
		p.id = "3 Kant"
	case 3:
		p.id = "4 Nietzsche"
	case 4: 
		p.id = "5 Wittgenstein"
	}
	return p
}

// Goes from thinking to hungry to eating and done eating then starts over.
// Adapt the pause values to increased or decrease contentions
// around the forks.
func (philo philO) eat(maxeat chan string) {
	for x := 0; x < 3; x++ {
		maxeat <- philo.id
		philo.meals++
		philo.first.Lock()
		philo.second.Lock()
		fmt.Printf("Philosopher #%s is eating a meal of rice\n", philo.id)
		philo.first.Unlock()
		philo.second.Unlock()
		fmt.Printf("Philosopher #%s is done eating a meal of rice\n", philo.id)
		<-maxeat
		if philo.meals == 3 {
			fmt.Printf("Philosopher #%s is finished eating!!!!!!\n", philo.id)
		}
	}
}

var wg sync.WaitGroup

func main() {
	const count = 5
	var philos []philO

	CSticks := NewSticks()
	for i := 0; i < 5; i++ {
		philos = append(philos, NewP(i, CSticks[i], CSticks[(i+1)%5]))
	}

	var wg sync.WaitGroup
	maxeat := make(chan string, 2) 
	done := 0
	for e := 0; e < count; e++ {
		wg.Add(1)
		go func(e int) {
			philos[e].eat(maxeat)
			done++
			fmt.Printf("-----%v Philosopher(s) have now finished eating-----\n", done)
			wg.Done()
		}(e)
	}

	wg.Wait()
}
