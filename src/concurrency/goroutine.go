package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers1() {
	for i := 0; i < 10; i++ {
		// fmt.Printf("%/d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++ {
		// fmt.Printf("%c ", i)
	}
}

func printNumbers2() {
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Millisecond)

		// fmt.Printf("%d ", i)
	}
}

func printLetters2() {
	for i := 'A'; i < 'A'+100; i++ {
		time.Sleep(1 * time.Millisecond)

		// fmt.Printf("%c ", i)
	}
}

func printNumbers3(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Millisecond)

		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func printLetters3(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+100; i++ {
		time.Sleep(1 * time.Millisecond)

		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func printNumbers4(w chan bool) {
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Millisecond)

		fmt.Printf("%d ", i)
	}
	w <- true
}

func printLetters4(w chan bool) {
	for i := 'A'; i < 'A'+100; i++ {
		time.Sleep(1 * time.Millisecond)

		fmt.Printf("%c ", i)
	}
	w <- true
}

func print1() {
	printNumbers1()
	printLetters1()
}

func goPrint1() {
	go printNumbers1()
	go printLetters1()
}

func print2() {
	printNumbers2()
	printLetters2()
}

func goPrint2() {
	go printNumbers2()
	go printLetters2()
}

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("Threw >>", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("Caught <<", num)
	}
}

func callectA(c chan string) {
	c <- "Hello, world"
	close(c)
}

func callectB(c chan string) {
	c <- "Hola Mundo"
	close(c)
}
func main1() {
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go printNumbers3(&wg)
	// go printLetters3(&wg)
	// wg.Wait()

	// w1, w2 := make(chan bool), make(chan bool)
	// go printLetters4(w1)
	// go printNumbers4(w2)
	// <-w1
	// <-w2
	// c := make(chan int, 3)
	// go thrower(c)
	// go catcher(c)
	// time.Sleep(100 * time.Millisecond)
	a, b := make(chan string), make(chan string)
	go callectA(a)
	go callectB(b)
	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2 {
		select {
		case msg, ok1 = <-a:
			if ok1 {
				fmt.Printf("%s from A\n", msg)
			}
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("%s from B\n", msg)
			}
			// default:
			// 	fmt.Println("default")
		}

	}
}
