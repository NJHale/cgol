package life

import (
	//"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestLife(t *testing.T) {
	uni := NewUniverse(11, 18)
	var period time.Duration = 60 * time.Millisecond
	unis := make(chan *Universe)
	defer close(unis)
	end := make(chan int)
	defer close(end)
	var r *rand.Rand
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
	go Play(uni, period, r, unis, end)
	// Wait for stuff to end and close
	timer := time.NewTimer(time.Second * 120)
	<-timer.C
	// End the thing!
	end <- 1
}

// func TestLiveNeighbors(t *testing.T) {
// 	uni := NewUniverse(10)
// 	r := rand.New(rand.NewSource(99))
// 	uni.Seed(50, r)
//
// 	fmt.Println(uni)
//
// 	for i := 0; i < uni.Size; i++ {
// 		for j := 0; j < uni.Size; j++ {
// 			fmt.Printf("LiveNeigbors(%d, %d): %d\n", i, j, uni.LiveNeighbors(i, j))
// 		}
// 	}
// }
