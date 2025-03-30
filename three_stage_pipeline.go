package main

import "fmt"

//TODO:Three stage Pipeline

/*
The first stage, gen, is a function that converts a list of integers to a channel that emits the integers in the list.
The gen function starts a goroutine that sends the integers on the channel and closes the channel when all the values have been sent:
*/
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

/*
The main function sets up the pipeline and runs the final stage: it receives values from
the second stage and prints each one, until the channel is closed:
*/
func ThreeStagePipe() {

	//TODO: 1 way of consuming
	c := gen(2, 3)
	out := sq(c)
	//
	//fmt.Println(<-out)
	//fmt.Println(<-out)
	//TODO: second way fo consuming
	// Set up the pipeline and consume the output.
	for n := range out {
		fmt.Println(n) // 16 then 81
	}
}
