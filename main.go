package main

import (
	"fmt"
	"math/rand"
	"time"
)

const TESTS = 10000
const LIMIT = 10000000

func main() {
	inc := 0
	rnd := 0
	equal := 0
	for i := 0; i < TESTS; i++ {
		random := rand.Intn(LIMIT)
		inc_time := guessInc(random)
		rnd_time := guessRandom(random)
		if inc_time < rnd_time {
			inc++
		} else if inc_time > rnd_time {
			rnd++
		} else {
			equal++
		}
	}
	fmt.Printf("Out of a %d tests : \n Inc was faster in %d \n Random was faster in %d \n They were equal in %d", TESTS, inc, rnd, equal)
}

func guessRandom(correct int) int64 {
	var guessed [LIMIT]int
	start := time.Now()
	guess := rand.Intn(LIMIT)
	for guess != correct && guessed[guess] == 1 {
		guess = rand.Intn(LIMIT)
	}
	elapsed := time.Since(start)
	return elapsed.Nanoseconds()
}

func guessInc(correct int) int64 {
	start := time.Now()
	for i := 0; i < LIMIT; i++ {
		if i == correct {
			break
		}
	}
	elapsed := time.Since(start)
	return elapsed.Nanoseconds()
}
