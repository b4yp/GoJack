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
	"randomness"
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

	// Initialize Card Arrays
	/* As of right now, only one deck. May add another dimension to store
	   multiple decks in a shoe. */
	var deck [52][2]int // Multi-Dimension Array to Store Cards/Suits

	// Shuffle Cards
	deck = shuffledeck()

	// DECKSHOW -- Deal 5 Cards FOR TESTING %%
	deckshow := 5

	fmt.Println("First 5 Cards in the Deck:")

	for deckshow > -1 {
		fmt.Println(deck[deckshow][0], deck[deckshow][1])
		deckshow = deckshow - 1
	}

	// DECKSHOW -- Deal 5 cards FOR TESTING %%

	// Start Game Loop

	// Print out final credit total and goodbye.
}

/* The function below shuffles a deck of cards. Right now, it will only
   shuffle a full deck of cards. It does not take into account any cards
   that have already been dealt. This may cause inaccuracies on a re-shuffle
   because there are cards already on the table and it will, in essence,
   re-add those cards to the shuffled deck as well as have those same cards
   on the table. For now, to get around this, we will re-shuffle between every
   hand.  */

func shuffledeck() [52][2]int { // Shuffle a deck of cards.
	var dck [52][2]int    // Shuffled Deck
	var unshuf [52][2]int // Unshuffled Deck (Cards in numerical/suit order)
	var intdeck int = 0   // Counter for initializing the deck.
	var shufdeck int = 51 // Counter for shuffling the deck.
	var reindex int       // Counter to re-index unshuf so no duplicate cards.
	var pickcard int      // Random Number of picked card.

	for intdeck < 52 { // Loop to create all cards in the array.

		switch {
		case intdeck < 13: // Suit 1 (Clubs)
			unshuf[intdeck][0] = intdeck + 1
			unshuf[intdeck][1] = 1
		case intdeck > 12 && intdeck < 26: // Suit 2 (Diamonds)
			unshuf[intdeck][0] = intdeck - 12
			unshuf[intdeck][1] = 2
		case intdeck > 25 && intdeck < 39: // Suit 3 (Hearts)
			unshuf[intdeck][0] = intdeck - 25
			unshuf[intdeck][1] = 3
		case intdeck > 38: // Suit 4 (Spades)
			unshuf[intdeck][0] = intdeck - 38
			unshuf[intdeck][1] = 4
		}

		intdeck = intdeck + 1
	}

	for shufdeck > -1 { // Here is where the cards are shuffled.
		pickcard = randomness.GetRandomZ(shufdeck + 1) // Grab random card.

		dck[shufdeck][0] = unshuf[pickcard][0] // Add card to shuffled deck.
		dck[shufdeck][1] = unshuf[pickcard][1] // Add suit to shuffled card.

		reindex = pickcard // Starting card for re-index.

		for reindex < (shufdeck) { // Re-index unshuffled array. Collapse space.
			unshuf[reindex][0] = unshuf[reindex+1][0]
			unshuf[reindex][1] = unshuf[reindex+1][1]
			reindex = reindex + 1
		}

		unshuf[shufdeck][0] = 0 // Zero out last card. (Removed card to end.)
		unshuf[shufdeck][1] = 0 // Zero out last suit. (No card, no suit.)

		shufdeck = shufdeck - 1
	}

	return dck // Return the shuffled deck to the calling routine.
}
