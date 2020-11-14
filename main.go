package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const kimBusy = true

// Student is a data type passed to the kim goroutine from a student.
type Student struct {
	name      string
	busyCount int
	response  chan string
}

func newStudent(name string) Student {
	s := Student{name: name}
	s.busyCount = 0
	s.response = make(chan string)
	return s
}

func kim(hand chan Student) {
	time.Sleep(time.Duration(rand.Intn(5)+5) * time.Second)
	fmt.Println("Kim's all ready to help!")

	for {
		somebody := <-hand
		fmt.Printf("Kim sees %v's raised hand!\n", somebody.name)
		if kimBusy {
			if somebody.busyCount < rand.Intn(5)+2 {
				// Students get a 'Busy' a certain number of times
				fmt.Println("Kim says 'Busy!'")
				somebody.response <- "Busy"
			} else {
				// After a while, Kim tells them they're 'On Deck'
				fmt.Println("Kim says 'You're on Deck!'")
				somebody.response <- "On Deck"
			}
		} else {
			fmt.Println("Something went wrong! Kim is always busy!")
		}
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	}
}

func student(wg *sync.WaitGroup, name string, hand chan Student) {
	defer wg.Done()
	me := newStudent(name)
	fmt.Printf("%v has started his work!\n", me.name)
	time.Sleep(time.Duration(rand.Intn(15)+5) * time.Second)

	// Time for the futile attempt at getting some help from Kim
	fmt.Printf("%v raises his hand for some help!\n", me.name)
	for {
		select {
		case hand <- me: // Kim actually bothers to notice you
			kimSez := <-me.response
			if kimSez == "Busy" {
				me.busyCount++
			} else if kimSez == "On Deck" {
				fmt.Printf("%v is on deck and accepts his fate.\n", me.name)
				return
			}
			fmt.Printf("%v waits a bit before trying again.\n", me.name)
			time.Sleep(time.Duration(rand.Intn(5)+5) * time.Second)
			fmt.Printf("%v raises his hand again!\n", me.name)
		default: // You can't even get Kim's attention
			time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
		}
	}
}

func classroom(class []string) {
	var wg sync.WaitGroup      // Wait group for checking how many students are still working
	hand := make(chan Student) // Channel for Kim seeing raised hands
	fmt.Println("The bell has rung! Time for class!")

	go kim(hand) // Goroutine for Kim
	for _, name := range class {
		go student(&wg, name, hand) // Goroutine for each student
	}

	wg.Add(len(class)) // Set the wait group to the number of students
	wg.Wait()          // Don't end class until everyone's spirit is broken
	fmt.Println("The bell has rung! Class dismissed!")
}

func main() {
	names := []string{
		"Jon", "Peter", "Matt", "Mike",
		"Andrew", "Nima", "Dan", "Ben",
		"Bill", "Ted", "Marty", "Biff",
	}
	rand.Seed(time.Now().UnixNano())
	classroom(names)
}
