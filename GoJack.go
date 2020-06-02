// GoJack - Simple console Blackjack game in Go.
//          Mainly written to practice arrays and (maybe) slices.
//
//			Thinking about it, this may be rolled into a
//          GoCasino project eventually. TBD.
//
// KEC - 06/2020

package main

import (
	"fmt"
	"kutil"
)

func main() {
	// Clear Screen - From K-Util Package
	kutil.ClearScreen()

	// Print Game Title
	titleprint()

	// Run Game - Making this modular in case I decide to brek this out into
	// another project like GoCasino.
	gobj()
}

func titleprint() {
	// Function that exists exclusively to print the title.

	fmt.Println("GoJack -- Blackjack for the Terminal")
	fmt.Println("KEC - 2020")
	fmt.Println()
}

func gobj() { // Yay! It's time to play Blackjack!

	// Print out intro and initialize money to "100" credits.

	var credits int = 100

	fmt.Println("Welcome to GoJack -- Blackjack for the terminal!")
	fmt.Println("You start with", credits, "credits.")
	fmt.Println("Let's see what you do with them. :)")
	fmt.Println()
	kutil.Pause(5)

	// Initialize Card Arrays

	// Shuffle Cards

	// Start Game Loop

	// Print out final credit total and goodbye.
}
