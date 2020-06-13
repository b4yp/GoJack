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

	// Run Game - Making this modular in case I decide to break this out into
	// another project like GoCasino. Super easy to call from a main menu that way.
	gobj()

}

func titleprint() {

	// Clear the screen.

	kutil.ClearScreen()

	// Print out the game title. We'll probably use this a lot.

	fmt.Println("GoJack -- Blackjack for the Terminal")
	fmt.Println("KEC - 2020")
	fmt.Println()
}

func gobj() { // Yay! It's time to play Blackjack!

	// Print out intro and initialize money to "100" credits.

	var credits int = 100
	var oncard int = 51

	// Initialize Card Arrays
	/* As of right now, only one deck. May add another dimension to store
	   multiple decks in a shoe. */
	var deck [52][2]int // Multi-Dimension Array to Store Cards/Suits

	// Shuffle Cards
	deck = shuffledeck()

	// Start Game Loop

	var playing bool = true // Create boolean to let us know if still playing.

	for playing == true { // Play hands until either out of credits or quit.
		// Pass credits and deck deck of remaining cards.
		// Return playing flag, new credit amount, and current remainig deck.
		playing, credits, oncard, deck = playhand(credits, oncard, deck)

		// If there are less than 4 cards (Standard to start game), reshuffle.
		if oncard < 4 {
			deck = shuffledeck()
			oncard = 51
		}

	}

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
			unshuf[intdeck][1] = 0
		case intdeck > 12 && intdeck < 26: // Suit 2 (Diamonds)
			unshuf[intdeck][0] = intdeck - 12
			unshuf[intdeck][1] = 1
		case intdeck > 25 && intdeck < 39: // Suit 3 (Hearts)
			unshuf[intdeck][0] = intdeck - 25
			unshuf[intdeck][1] = 2
		case intdeck > 38: // Suit 4 (Spades)
			unshuf[intdeck][0] = intdeck - 38
			unshuf[intdeck][1] = 3
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

// playhand - Plays a single hand of GoBJ then returns game status, credits,
//            and current deck of playing cards.
func playhand(cred int, card int, dck [52][2]int) (bool, int, int, [52][2]int) {
	var stillplaying bool // Is the player still playing?
	var validchoice bool  // Is the choice returned valid?
	var bet int           // This is the amount of the bet.

	validchoice = false
	for validchoice == false { // Loop until a valid choice is made.
		titleprint()      // Clear the screen and print the title.
		bet = wager(cred) // Wager prompt. Returns "0" for bet if quitting.

		if bet >= 0 && bet <= cred {
			// OK, we're going to play BJ now!
			// Unless bet = 0, then we're quitting.
			validchoice = true
			if bet > 0 {
				fmt.Println("Amount wagered:", bet)
				cred, card, dck = playblackjack(cred, bet, card, dck)
			}
			kutil.Pause(5)
		}

		/* Only a bet that is "0" to quit or an amount that's within the
		   credit range of the player is valid. Anything else, send them
		   back to try again! */

	}

	switch { // OK, let's see what we do after the game is played.
	case bet == 0: // We bet "0" on the wager screen.
		// Chuck us back to the calling routine and
		// say we aren't playing any more. Since it hits
		// this check first, it'll quit regadless of how
		// many credits we have.
		stillplaying = false
	case cred > 0: // We still have credits left and didn't bet zero.
		// Keep playing.
		stillplaying = true
	case cred <= 0: // We're out of credits! No more game for us!
		stillplaying = false
		fmt.Println("You've run out of credits! Thanks for playing!")
		kutil.Pause(5)
	}

	return stillplaying, cred, card, dck
	// Return the flag that lets us know if we are still playing,
	// our leftover creds, and what's left of the deck of cards.
}

func wager(creds int) int { // This routine grabs our bet! Or quits.
	var betamt int // Bet amount.

	fmt.Println("How much would you like to wager?")
	fmt.Println("Credits:", creds)
	fmt.Printf("Amount (0 to exit) > ")
	fmt.Scan(&betamt) // Get our bet amount from the terminal.

	switch { // OK.. let's see what we bet.
	case betamt == 0: // Bet is 0. Quit the game.
		fmt.Println("Thanks for playing!")
	case betamt < 0: // Bet is negative. Invalid.
		fmt.Println("Sorry, you can't wager negative amounts of credits.")
	case betamt > creds: // Bet is more than we have. Also invalid.
		fmt.Println("Sorry, you can't wager more credits than you have!")
	case betamt > 0 && betamt <= creds: // Bet is valid amount. Continue.
		fmt.Println("Ok, you're betting", betamt, "credits!")
	}

	kutil.Pause(2) // Pause for a couple of seconds to allow user to read the outcome.

	return betamt // Return the amount that was bet.
}

func playblackjack(curcredits int, curbet int, curcard int, curdeck [52][2]int) (int, int, [52][2]int) {

	var c int // Array Counter
	var cardsout bool = true
	// Define Dealer and Player Hands -- Using 12 cards as maximum. Reason?
	// 1x4 + 2x4 + 3x3 = 21 , Add 1 additional card and player goes over 21.
	// Player/Dealer would need to be incredibly lucky to get that hand, but
	// I want to code for it because you never know.
	var dealerhand [12][2]int // Dealer's Hand
	var playerhand [12][2]int // Player's Hand
	// Define cards in hand for loop so we only loop as many times as there are cards.
	// Initialized to "2" each because initial deal is only 2 cards each,
	var dealercards int = 2 // Number of cards that the dealer has.
	var playercards int = 2 // Number of cards that the player has.
	// Define whether or not it's the dealer's turn to play.
	// If "false" then the first card that the dealer plays is covered.
	var dealerturn bool = false

	// Set card numbers and add current cards to arrays.

	c = 0
	for c < 4 {
		fmt.Println(curcard)
		switch {
		case c == 0 || c == 1:
			dealerhand[c][0] = curdeck[curcard-c][0]
			dealerhand[c][1] = curdeck[curcard-c][1]
		case c == 2 || c == 3:
			playerhand[c-2][0] = curdeck[curcard-c][0]
			playerhand[c-2][1] = curdeck[curcard-c][1]
		}
		c = c + 1
	}

	curcard = curcard - 4

	for cardsout {
		showhand(dealerhand, playerhand, dealercards, playercards, dealerturn)
		cardsout = false
	}

	curcredits = curcredits - curbet

	return curcredits, curcard, curdeck

}

func showhand(dlrhand [12][2]int, plrhand [12][2]int, dlrcards int, plrcards int, dlrturn bool) {
	var d int // Counter for Card Deal

	suit := [4]string{"C", "D", "H", "S"}

	// Clear Screen and Print Title
	kutil.ClearScreen()
	titleprint()

	// Print Dealer Hand
	d = 0
	fmt.Println("Dealer's Hand:")
	for d < dlrcards {
		if dlrturn == false && d == 0 {
			fmt.Print("[X X]")
		} else {
			fmt.Print("[", dlrhand[d][0], " ", suit[dlrhand[d][1]], "]")
		}
		d = d + 1
	}

	fmt.Println()
	fmt.Println()
	// Print Player Hand
	d = 0
	fmt.Println("Player's Hand:")
	for d < dlrcards {
		fmt.Print("[", plrhand[d][0], " ", suit[plrhand[d][1]], "]")
		d = d + 1
	}
}
