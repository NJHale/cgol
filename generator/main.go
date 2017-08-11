package main

import (
	"fmt"
	"time"
	"math/rand"

	"tfs.ups.com/tfs/UpsProd/P08SGIT_EA_CDP/_git/cgol/generator/life"
)

func main() {
	fmt.Printf("Hello World!")
	uni := life.NewUniverse(10, 10)
	fmt.Println(uni)

	uni := life.NewUniverse(11, 18)
	var period time.Duration = 60 * time.Millisecond
	unis := make(chan *life.Universe)
	defer close(unis)
	end := make(chan int)
	defer close(end)
	//r := rand.New(rand.NewSource(99))
	//uni.RandomSeed(200, r)

	uni.Space = [][]int {
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}


	// Run Life in a go-routine
	go life.Play(uni, period, unis, end)
	// Wait for stuff to end and close
	timer := time.NewTimer(time.Second * 120)
	<-timer.C
	// End the thing!
	end <- 1
}