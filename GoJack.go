// GoJack - Simple console Blackjack game in Go.
//          Mainly written to practice arrays and (maybe) slices.
//
//			Thinking about it, this may be rolled into a
//          GoCasino project eventually. TBD.
//
// KEC (@b4yp) - 2020
// ----------
// Release Log:
//
// v1.0 - 2020/06/15 - Initial release!

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

func gobj() { // OK! Let's start up the BlackJack Routine!

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

}

func titleprint() { // Print the title of the game.

	// Clear the screen.
	kutil.ClearScreen()

	// Print out the game title. We'll probably use this a lot.
	fmt.Println("GoJack v1.0 -- Blackjack for the Terminal")
	fmt.Println("Branch - WIP")
	fmt.Println()
}

/* The function below shuffles a deck of cards. Right now, it will only
   shuffle a full deck of cards. It does not take into account any cards
   that have already been dealt. This may cause inaccuracies on a re-shuffle
   because there are cards already on the table and it will, in essence,
   re-add those cards to the shuffled deck as well as have those same cards
   on the table. For now, to get around this, we will re-shuffle when there
   are 4 or less cards left between hands. This is only a temporary solution,
   since there will likely be times where there are 5 cards remaining in the
   deck and we need more than that. */

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
			kutil.Pause(2)
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
		fmt.Println()
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

// Yay!!! We get to play BlackJack! At last!
func playblackjack(curcredits int, curbet int, curcard int, curdeck [52][2]int) (int, int, [52][2]int) {

	var c int                // Array Counter
	var cardsout bool = true // Are there cards that still need to be played?
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
	// Total amount of the player's and dealer's hands.
	var dealertotal int
	var playertotal int
	// Did the player or dealer bust?
	var playerbust bool = false
	var dealerbust bool = false
	// Define whether or not it's the dealer's turn to play.
	// If "false" then the first card that the dealer plays is covered.
	var dealerturn bool = false
	// Define if the dealer is done with their turn.
	var dealerdone bool = false
	// Did the player get blackjack on the initial deal?
	var playerbj bool = false

	// Set card numbers and add current cards to arrays.
	c = 0
	for c < 4 {
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

	// Start by doing a check of the initial totals to see if the player
	// got a blackjack and populate the totals for the screen.
	playertotal, playerbust = cardcount(playerhand, playercards)
	dealertotal, dealerbust = cardcount(dealerhand, dealercards)

	// If player gets a Blackjack, set a flag to let the cardsout loop know.
	if playertotal == 21 {
		playerbj = true
	}

	// Now that all of that setup is done, let's actually play the game!
	for cardsout {
		// Show the player and dealer's hand on the screen.
		showhand(dealerhand, playerhand, dealertotal, playertotal,
			dealercards, playercards, dealerturn)
		fmt.Println()

		// Determine who's turn it is and call the appropriate routine to process same.
		switch {
		case playerbj:
			// If player gets a BlackJack, let them know, pay them, and exit the loop.
			// No game to play.
			fmt.Println("BLACKJACK! W00T!")
			cardsout = false
			curcredits = curcredits + (curbet * 2)
			kutil.Pause(2)
		case dealerturn == false: // It's the player's turn.
			playercards = playerselect(playercards)
			// Check the array. If the card in the highest place in the array
			// is zero, that means that the player has drawn a card. Add the card
			// to the player's hand.
			if playerhand[playercards-1][0] == 0 {
				playerhand[playercards-1][0] = curdeck[curcard][0]
				playerhand[playercards-1][1] = curdeck[curcard][1]
				curcard = curcard - 1
			} else {
				// If the array returns non-zero, that means that the player did not
				// draw a card and their turn is over.
				dealerturn = true
			}
		case dealerturn == true: // Now, it's the dealer's turn.
			dealercards = dealerselect(dealercards, dealertotal)
			// Check the array. If the card in the highest place in the array
			// is zero, that means that the dealer has drawn a card. Add the card
			// to the dealer's hand.
			if dealerhand[dealercards-1][0] == 0 {
				dealerhand[dealercards-1][0] = curdeck[curcard][0]
				dealerhand[dealercards-1][1] = curdeck[curcard][1]
				curcard = curcard - 1
			} else {
				// If the array returns non-zero, that means that the dealer did not
				// draw a card and their turn is over.
				dealerdone = true
			}
		}

		// Total up the player and dealer hands. Determine if either have busted.
		playertotal, playerbust = cardcount(playerhand, playercards)
		dealertotal, dealerbust = cardcount(dealerhand, dealercards)

		// If either the player or dealer bust, clear the screen, show the updated
		// hands and totals, then process outcome.
		if playerbust || dealerbust {
			// Show the player and dealer's hand on the screen.
			showhand(dealerhand, playerhand, dealertotal, playertotal,
				dealercards, playercards, dealerturn)
			fmt.Println()

			// If the player busts, let them know, subtract the credits that they bet
			// and set the flag to exit the game loop.
			if playerbust {
				fmt.Println("PLAYER HAS BUSTED!")
				curcredits = curcredits - curbet
				cardsout = false
				kutil.Pause(2)
			}

			// If the dealer busts, let the player know, add the credits that they bet
			// and set the flag to exit the game loop.
			if dealerbust {
				fmt.Println("DEALER HAS BUSTED!")
				curcredits = curcredits + curbet
				cardsout = false
				kutil.Pause(2)
			}

		}

		// If the dealer completes their turn successfully without busting,
		// then it's time to determine who won the game based on score.
		if dealerdone {

			// Show the player and dealer's hand on the screen.
			showhand(dealerhand, playerhand, dealertotal, playertotal,
				dealercards, playercards, dealerturn)
			fmt.Println()

			switch {
			case playertotal > dealertotal:
				// Player has more points than the dealer. They win!
				fmt.Println("Congratulations, you won!")
				curcredits = curcredits + curbet
				kutil.Pause(2)
			case dealertotal > playertotal:
				// Dealer has more points than the player. Dealer wins!
				fmt.Println("Sorry, you lost!")
				curcredits = curcredits - curbet
				kutil.Pause(2)
			case dealertotal == playertotal:
				// Dealer and player have the same. No win, no loss.
				fmt.Println("Push. No win, no loss.")
				kutil.Pause(2)
			}

			cardsout = false // Set flag to exit loop since both players are done.
		}

	}

	return curcredits, curcard, curdeck

}

// This function exists to print out the hands of the dealer and the player.
// Modified: Now also prints out the totals as well.
func showhand(dlrhand [12][2]int, plrhand [12][2]int, dtotal int, ptotal int,
	dlrcards int, plrcards int, dlrturn bool) {
	var d int // Counter for Card Deal

	realcard := [4]string{"J", "Q", "K", "A"} // Array of Face Cards
	suit := [4]string{"♣", "♦", "♥", "♠"}     // Suits Expressed as Pictures

	// Clear Screen and Print Title
	kutil.ClearScreen()
	titleprint()

	// Print Dealer Hand
	d = 0
	fmt.Println("Dealer's Hand:")
	for d < dlrcards {
		if dlrturn == false && d == 0 {
			// If it isn't the dealer's turn yet, hide the first card.
			fmt.Print("[X X]")
		} else {
			// If it is the dealer's turn, show all of the cards.
			switch {
			case dlrhand[d][0] > 10: // If the card is > 10, it's a face card. Account for this.
				fmt.Print("[", realcard[(dlrhand[d][0]-11)], " ", suit[dlrhand[d][1]], "]")
			case dlrhand[d][0] == 1: // If the card is a "1", it's an Ace. Account for this.
				fmt.Print("[", realcard[3], " ", suit[dlrhand[d][1]], "]")
			default: // Otherwise it's a regular old number card. Just print it out as-is.
				fmt.Print("[", dlrhand[d][0], " ", suit[dlrhand[d][1]], "]")
			}
		}
		d = d + 1
	}

	fmt.Println()

	if dlrturn == false {
		// If it's the player's turn, we don't know the dealer's total because the
		// first card is hidden.. so display question marks.
		fmt.Println("Dealer Total: ??")
	} else {
		// If it it's the dealer's turn, we want to show the total for the dealer.
		fmt.Println("Dealer Total:", dtotal)
	}
	fmt.Println()
	// Print Player Hand
	d = 0
	fmt.Println("Player's Hand:")
	for d < plrcards {
		switch {
		case plrhand[d][0] > 10: // If the card is > 10, it's a face card. Account for this.
			fmt.Print("[", realcard[(plrhand[d][0]-11)], " ", suit[plrhand[d][1]], "]")
		case plrhand[d][0] == 1: // If the card is a "1", it's an Ace. Account for this.
			fmt.Print("[", realcard[3], " ", suit[plrhand[d][1]], "]")
		default: // Otherwise, it's a regular old number card. Just print it out as-is.
			fmt.Print("[", plrhand[d][0], " ", suit[plrhand[d][1]], "]")
		}
		d = d + 1
	}
	// Print out the player's total. We always do this.
	fmt.Println()
	fmt.Println("Player Total:", ptotal)
}

// This function totals the score for all of the cards and determines if someone busts.
func cardcount(counthand [12][2]int, countcards int) (int, bool) {
	var d int = 0          // Array counter/variable.
	var acecount int = 0   // Number of aces that we encounter that have been drawn.
	var totalcount int = 0 // Total value of the cards in the hand.
	var bust bool = false  // Did the we bust?

	for d < countcards {
		switch {
		case counthand[d][0] == 1: // This card is an Ace. Register it as so and add 11.
			totalcount = totalcount + 11
			acecount = acecount + 1
		case counthand[d][0] > 9: // This card is a face card (or a 10). Add 10 to the total.
			totalcount = totalcount + 10
		case counthand[d][0] < 10: // This card is less than 10. Add the value as-is.
			totalcount = totalcount + counthand[d][0]
		}

		// If we reach a total that's greater than 21, check to see if we have any aces.
		// If so, reduce the count by 10 so the ace counts as "1" instead of 11.
		// Also, decrement the number of aces that we have to use for this.
		if totalcount > 21 && acecount > 0 {
			totalcount = totalcount - 10
			acecount = acecount - 1
		}

		if totalcount > 21 { // Total > 21? Sorry. Ya busted.
			bust = true
		}

		d = d + 1
	}
	return totalcount, bust
}

// This function is to allow the player to select their
func playerselect(incard int) int {
	var newcard int   // Number of cards in the hand.
	var selection int // Menu selection.

	// Print out the menu and get the selection from it.
	fmt.Println("1. Hit")
	fmt.Println("2. Stay")
	fmt.Print(">> ")
	fmt.Scan(&selection)

	// Process our selection.
	switch {
	case selection == 1: // Hit. Add a card to the hand.
		newcard = incard + 1
		fmt.Println("Requesting a card.")
		kutil.Pause(2)
	case selection == 2: // Stay. No added card.
		newcard = incard
		fmt.Println("No more cards for you!")
		kutil.Pause(2)
	default:
		newcard = incard // Invalid. No added card.
		fmt.Println("Invalid selection, assuming Stay.")
		kutil.Pause(5)
	}

	return newcard
}

// This function is the "AI" for the dealer. It's a pretty dumb AI right now.
// Basically, hit on anything less than 16.
func dealerselect(indcard int, indtotal int) int {
	var newdcard int // Number of cards in the hand.

	switch {
	case indtotal < 16: // Less than 16? Hit. Add a new card.
		newdcard = indcard + 1
		fmt.Println("Dealer takes a card.")
		kutil.Pause(2)
	case indtotal >= 16: // Greater or equal to 16? Stay. No new card.
		newdcard = indcard
		fmt.Println("Dealer stays.")
		kutil.Pause(2)
	}

	return newdcard
}
