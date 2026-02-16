package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func speaker(name string, debatte chan int) {
	for {
		microphone := <-debatte // Auf Mikro warten

		fmt.Printf("Turn %v: %v says '%v'\n", microphone, name, randomAnswer())
		time.Sleep(400 * time.Millisecond)

		microphone++
		debatte <- microphone
	}
}

func randomAnswer() string {
	answers := []string{"I'm right.", "Take this argument.", "But I've got rhat.", "You b***"}
	return answers[rand.IntN((len(answers) - 1))]
}

func main() {
	debatte := make(chan int)

	go speaker("Jackie", debatte)
	go speaker("Frank", debatte)

	microphone := 1

	debatte <- microphone

	time.Sleep(2 * time.Second)

	<-debatte
}
