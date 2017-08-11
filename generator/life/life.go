package life

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * [sem description]
 * @type {[type]}
 */
func Play(uni *Universe, period time.Duration, r *rand.Rand,
	unis chan *Universe, end chan int) {
	// Use a ticker to run Tick every period milliseconds
	ticker := time.NewTicker(period)

	// Print the initial Universe
	fmt.Println(uni)

	// Fire off a go-routine to call Tick at the desired intervals
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			Tick(uni)
			// If we were given a random variable, seed it
			if r != nil {
				uni.RandomSeed(5, r)
			}
			fmt.Println(uni)
		}
	}()
	// Wait for the end
	val, ok := <-end
	if !ok {
		// End has been closed
		fmt.Println("Ending channel has been closed before ending message sent...")
	}

	fmt.Println("Ending msg:", val)

}

/**
 * [sem description]
 * @type {[type]}
 */
func Tick(uni *Universe) {
	// Create semaphore(s) to wait for all cells to be calculated
	analyzed := make(chan int, uni.Height*uni.Width)
	mutating := make(chan int, uni.Height*uni.Width)
	mutated := make(chan int, uni.Height*uni.Width)

	// Send each cell to a go routine and generate the new universe
	for i := 0; i < uni.Height; i++ {
		for j := 0; j < uni.Width; j++ {
			go func(i, j int) {
				//fmt.Print("Checking (", i, ",", j, ")\t")
				// Figure out whether the cell should live or die
				shouldLive := ShouldLive(i, j, uni)
				// Signal that the cell has been analyzed
				analyzed <- 1
				// Wait until we can mutate
				<-mutating
				// Update cell
				if shouldLive {
					uni.Live(i, j)
				} else {
					uni.Die(i, j)
				}
				// Signal that we are done mutating
				mutated <- 1
			}(i, j)
		}
	}

	// Wait for all the go routines to finish analyzing
	for s := 0; s < uni.Height*uni.Width; s++ {
		<-analyzed
	}

	// Single go routines to begin mutating
	for s := 0; s < uni.Height*uni.Width; s++ {
		mutating <- 1
	}

	// Wait for go routines to
	for s := 0; s < uni.Height*uni.Width; s++ {
		<-mutated
	}

	close(analyzed)
	close(mutating)
	close(mutated)

	// Iterate the generation
	uni.Generation++
}

/**
 * [cell description]
 * @type {[type]}
 */
func ShouldLive(i, j int, uni *Universe) bool {
	// Declare live flag
	var live bool
	// Get a pointer to the cell
	cell, err := uni.Get(i, j)
	// Get the number of live neighbors
	living := uni.LiveNeighbors(i, j)

	//fmt.Printf("(%d, %d) has %d live neighbors\n", i, j, living)

	// Apply cgol rules on cell
	if err == nil && cell > 0 {
		// Live cell rules
		if living < 2 || living > 3 {
			// Die by under or overpopulation
			live = false
		} else {
			// Live by comfort
			live = true
		}
	} else {
		// Dead cell rule
		if living == 3 {
			// Live by reproduction
			live = true
		} else {
			// Remain dead
			live = false
		}
	}

	return live
}
