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
// v1.1 - 2020/06/20 - Bug fixes, enhancements, added insurance and double-down.
// v1.WIP -- Work in Progress. Watch for falling code.

package main

import (
	"fmt"
	"kutil"
	"math"
	"randomness"
	"strings"
)

// GLOBAL VARIABLES

// Credits - Amount of Credits that a player has. Initialized to 100 credits at start.
var Credits float64 = 100

// Bet - Amount of Credits that the player has bet. Initialized to  credits.
var Bet float64 = 0

// OnCard - Which is the next card that we are drawing from the deck?
// -- Initialized to 51 (arrays start at 0, 52 cards in standard deck)
var OnCard int = 51

// Deck -- Our shuffled deck of cards.
var Deck [52][2]int

// DealerTotal -- Total number of points that the dealer has.
var DealerTotal int

// PlayerTotal -- Total number of points that the player has.
var PlayerTotal int

// Soft17 -- Does the dealer hit on a Soft 17?
var Soft17 bool = false

func main() {

	// Run Game - Making this modular in case I decide to break this out into
	// another project like GoCasino. Super easy to call from a main menu that way.
	gobj()

}

func gobj() { // OK! Let's start up the BlackJack Routine!

	// Start Game Loop

	var playing bool = true   // Create boolean to let us know if still playing.
	var firstloop bool = true // Is this the first loop? If so, ask for Soft17.
	var softyn string = "N"   // Input for Soft17 Check

	for playing { // Play hands until either out of credits or quit.

		// Single deck Blackjack. Shuffle before each hand.
		shuffledeck()
		OnCard = 51

		// Check to see if the player wants the dealer to hit on a Soft 17.
		if firstloop {

			// Print the title.
			titleprint()

			// Ask. Defaults to "No".
			fmt.Println("Dealer Hit on Soft 17?")
			fmt.Print("(y/N) >> ")
			fmt.Scan(&softyn)

			// Set case of input string to upper case.
			softyn = strings.ToUpper(softyn)

			// If player enters "Y", then Soft 17 is a go, otherwise, defaulted
			// to false.
			if softyn == "Y" {
				Soft17 = true
			}

			// Set the "firstloop" variable to "false" so we aren't asking this
			// question twice.
			firstloop = false

		}

		// Return flag to let loop know that we're still playing.
		playing = playhand()

	}

}

func titleprint() { // Print the title of the game.

	// Clear the screen.
	kutil.ClearScreen()

	// Print out the game title. We'll probably use this a lot.
	fmt.Println("GoJack v1.WIP -- Blackjack for the Terminal")
	fmt.Println()
}

func shuffledeck() { // Shuffle a deck of cards.
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

		Deck[shufdeck][0] = unshuf[pickcard][0] // Add card to shuffled deck.
		Deck[shufdeck][1] = unshuf[pickcard][1] // Add suit to shuffled card.

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

}

// playhand - Plays a single hand of GoBJ then returns game status, credits,
//            and current deck of playing cards.
func playhand() bool {
	var stillplaying bool     // Is the player still playing?
	var validchoice bool      // Is the choice returned valid?
	var fraccent bool = false // Is the bet a fractional cent?

	validchoice = false
	for !validchoice { // Loop until a valid choice is made.
		titleprint()       // Clear the screen and print the title.
		fraccent = wager() // Wager prompt. Returns "0" for bet if quitting.

		if Bet >= 0 && Bet <= Credits && !fraccent {
			// OK, we're going to play BJ now!
			// Unless bet = 0, then we're quitting.
			validchoice = true
			if Bet > 0 {
				fmt.Println("Amount wagered:", Bet)
				playblackjack()
			}
			kutil.Pause(2)
		}

		/* Only a bet that is "0" to quit or an amount that's within the
		   credit range of the player is valid. Anything else, send them
		   back to try again! */

	}

	switch { // OK, let's see what we do after the game is played.
	case Bet == 0: // We bet "0" on the wager screen.
		// Chuck us back to the calling routine and
		// say we aren't playing any more. Since it hits
		// this check first, it'll quit regadless of how
		// many credits we have.
		stillplaying = false
	case Credits > 0: // We still have credits left and didn't bet zero.
		// Keep playing.
		stillplaying = true
	case Credits <= 0: // We're out of credits! No more game for us!
		stillplaying = false
		fmt.Println("You've run out of credits! Thanks for playing!")
		kutil.Pause(5)
	}

	return stillplaying
	// Return the flag that lets us know if we are still playing.
}

func wager() bool { // This routine grabs our bet! Or quits.
	var partcent float64     // Integer value from bet.
	var badcent bool = false // Return "true" if fractions of a cent.

	fmt.Println("How much would you like to wager?")
	fmt.Println("Credits:", fmt.Sprintf("%.2f", Credits))
	fmt.Printf("Amount (0 to exit) > ")
	fmt.Scan(&Bet) // Get our bet amount from the terminal.

	// This bit of code determines if we have any values less than
	// one cent entered as a bet.
	// Logic -- Multiply by 100 to get an integer value and truncate
	//          the rest. Then compare if the integer value is less than
	//          the true value. If it's a whole cent, they should be equal.
	//          If it's a fraction of a cent, then "Bet * 100" will be greater
	//          because there are values past the whole cent.
	partcent = math.Trunc(Bet * 100)
	if partcent < (Bet * 100) {
		badcent = true
	}

	switch { // OK.. let's see what we bet.
	case Bet == 0: // Bet is 0. Quit the game.
		fmt.Println()
		fmt.Println("Thanks for playing!")
	case Bet < 0: // Bet is negative. Invalid.
		fmt.Println("Sorry, you can't wager negative amounts of credits.")
	case Bet > Credits: // Bet is more than we have. Also invalid.
		fmt.Println("Sorry, you can't wager more credits than you have!")
	case badcent: // Bet contains a fraction of a cent. Sorry, fam.
		fmt.Println("Bets must not contain fractional cent values.")
	case Bet > 0 && Bet <= Credits: // Bet is valid amount. Continue.
		fmt.Println("Ok, you're betting", fmt.Sprintf("%.2f", Bet), "credits!")
	}

	kutil.Pause(2) // Pause for a couple of seconds to allow user to read the outcome.

	return badcent

}

// Yay!!! We get to play BlackJack! At last!
func playblackjack() {

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
	// Did the player double down on their turn?
	var doubledown bool = false
	// Did the player or dealer bust?
	var playerbust bool = false
	var dealerbust bool = false
	// Define whether or not it's the dealer's turn to play.
	// If "false" then the first card that the dealer plays is covered.
	var dealerturn bool = false
	// Define if the dealer is done with their turn.
	var dealerdone bool = false
	// Is this a soft hand for the player? (Contains > 0 Aces)
	var softhandp bool = false
	// Is this a soft hand for the dealer? (Contains > 0 Aces)
	var softhandd bool = false
	// Did the player get blackjack on the initial deal?
	var playerbj bool = false
	// Did the dealer get blackjack on the initial deal?
	var dealerbj bool = false
	// Did we process insurance already (if necessary)
	var insflag bool = false

	// Set card numbers and add current cards to arrays.
	c = 0
	for c < 4 {
		switch {
		case c == 0 || c == 1:
			dealerhand[c][0] = Deck[OnCard-c][0]
			dealerhand[c][1] = Deck[OnCard-c][1]
		case c == 2 || c == 3:
			playerhand[c-2][0] = Deck[OnCard-c][0]
			playerhand[c-2][1] = Deck[OnCard-c][1]
		}
		c = c + 1
	}

	// 4 cards were dealt out, so account for those 4 cards in our counter.
	OnCard = OnCard - 4

	// Start by doing a check of the initial totals to see if the player
	// got a blackjack and populate the totals for the screen.
	PlayerTotal, playerbust, softhandp = cardcount(playerhand, playercards)
	DealerTotal, dealerbust, softhandd = cardcount(dealerhand, dealercards)

	// If player gets a Blackjack, set a flag to let the cardsout loop know.
	if PlayerTotal == 21 {
		playerbj = true
	}

	// If dealer gets a Blackjack, set a flag. Only really relevant if
	// player and dealer both get a Blackjack on the initial deal.
	if DealerTotal == 21 {
		dealerbj = true
	}

	// Now that all of that setup is done, let's actually play the game!
	for cardsout {

		// Show the player and dealer's hand on the screen.
		showhand(dealerhand, playerhand, dealercards, playercards, dealerturn)
		fmt.Println()

		// Determine who's turn it is and call the appropriate routine to process same.
		// Also determine initial blackjack and and if we're going to ask if the
		// player wants insurance.
		switch {
		case playerbj && !dealerbj:
			// If player gets a BlackJack, let them know, pay them, and exit the loop.
			// No game to play.
			fmt.Println("BLACKJACK! W00T!")
			cardsout = false
			Credits = Credits + (Bet * 2)
			kutil.Pause(2)
		case playerbj && dealerbj:
			// If player gets a Blackjack AND the dealer gets a blackjack, push,
			// no payout, exit the loop. No game to play.
			dealerturn = true
			showhand(dealerhand, playerhand, dealercards, playercards, dealerturn)
			fmt.Println()
			fmt.Println("Player and dealer have Blackjack. Push.")
			kutil.Pause(2)
			cardsout = false
		case dealerhand[1][0] == 1 && !insflag:
			// If the the second card in the dealer card array is "1" that means
			// that the card showing is an Ace. Ask if the player wants insurance.
			insflag = true
			insurance()
		case !dealerturn: // It's the player's turn.
			playercards, doubledown = playerselect(playercards, softhandp)
			// Check the array. If the card in the highest place in the array
			// is zero, that means that the player has drawn a card. Add the card
			// to the player's hand.
			if playerhand[playercards-1][0] == 0 {
				playerhand[playercards-1][0] = Deck[OnCard][0]
				playerhand[playercards-1][1] = Deck[OnCard][1]
				OnCard = OnCard - 1
				if doubledown { // Did the player double down? If so, end their turn.
					dealerturn = true
				}
			} else {
				// If the array returns non-zero, that means that the player did not
				// draw a card and their turn is over.
				dealerturn = true
			}
		case dealerturn: // Now, it's the dealer's turn.
			dealercards = dealerselect(dealercards, softhandd)
			// Check the array. If the card in the highest place in the array
			// is zero, that means that the dealer has drawn a card. Add the card
			// to the dealer's hand.
			if dealerhand[dealercards-1][0] == 0 {
				dealerhand[dealercards-1][0] = Deck[OnCard][0]
				dealerhand[dealercards-1][1] = Deck[OnCard][1]
				OnCard = OnCard - 1
			} else {
				// If the array returns non-zero, that means that the dealer did not
				// draw a card and their turn is over.
				dealerdone = true
			}
		}

		// Total up the player and dealer hands. Determine if either have busted.
		PlayerTotal, playerbust, softhandp = cardcount(playerhand, playercards)
		DealerTotal, dealerbust, softhandd = cardcount(dealerhand, dealercards)

		// If either the player or dealer bust, clear the screen, show the updated
		// hands and totals, then process outcome.
		if playerbust || dealerbust {
			// Show the player and dealer's hand on the screen.
			showhand(dealerhand, playerhand, dealercards, playercards, dealerturn)
			fmt.Println()

			// If the player busts, let them know, subtract the credits that they bet
			// and set the flag to exit the game loop.
			if playerbust {
				fmt.Println("PLAYER HAS BUSTED!")
				Credits = Credits - Bet
				cardsout = false
				kutil.Pause(2)
			}

			// If the dealer busts, let the player know, add the credits that they bet
			// and set the flag to exit the game loop.
			if dealerbust {
				fmt.Println("DEALER HAS BUSTED!")
				Credits = Credits + Bet
				cardsout = false
				kutil.Pause(2)
			}

		}

		// If the dealer completes their turn successfully without busting,
		// then it's time to determine who won the game based on score.
		if dealerdone {

			// Show the player and dealer's hand on the screen.
			showhand(dealerhand, playerhand, dealercards, playercards, dealerturn)
			fmt.Println()

			switch {
			case PlayerTotal > DealerTotal:
				// Player has more points than the dealer. They win!
				fmt.Println("Congratulations, you won!")
				Credits = Credits + Bet
				kutil.Pause(2)
			case DealerTotal > PlayerTotal:
				// Dealer has more points than the player. Dealer wins!
				fmt.Println("Sorry, you lost!")
				Credits = Credits - Bet
				kutil.Pause(2)
			case DealerTotal == PlayerTotal:
				// Dealer and player have the same. No win, no loss.
				fmt.Println("Push. No win, no loss.")
				kutil.Pause(2)
			}

			cardsout = false // Set flag to exit loop since both players are done.
		}

	}

}

// This function exists to print out the hands of the dealer and the player.
// Modified: Now also prints out the totals as well.
func showhand(dlrhand [12][2]int, plrhand [12][2]int,
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
		if !dlrturn && d == 0 {
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

	if !dlrturn {
		// If it's the player's turn, we don't know the dealer's total because the
		// first card is hidden.. so display question marks.
		fmt.Println("Dealer Total: ??")
	} else {
		// If it it's the dealer's turn, we want to show the total for the dealer.
		fmt.Println("Dealer Total:", DealerTotal)
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
	fmt.Println("Player Total:", PlayerTotal)
}

// This function totals the score for all of the cards and determines if someone busts.
func cardcount(counthand [12][2]int, countcards int) (int, bool, bool) {
	var d int = 0          // Array counter/variable.
	var acecount int = 0   // Number of aces that we encounter that have been drawn.
	var totalcount int = 0 // Total value of the cards in the hand.
	var bust bool = false  // Did the we bust?
	var soft bool = false  // Does this hand contain an ace?

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

		for acecount > 0 && totalcount > 21 {

			totalcount = totalcount - 10
			acecount = acecount - 1

		}

		// If there are more than one aces available, this is a soft hand, report
		// as true if so.
		if acecount > 0 {
			soft = true
		}

		if totalcount > 21 { // Total > 21? Sorry. Ya busted.
			bust = true
		}

		d = d + 1
	}
	return totalcount, bust, soft
}

// This function is to ask the player if they want insurance and process if they do.
func insurance() {
	var selection string = "N" // Menu Selection
	var insbet float64         // Insurance Bet (Up to half of Bet)
	var partcent float64       // Make sure we don't have any partial cents in bet.
	var validsel bool = false  // Did the user make a valid selection?

	// Print out the menu and get the selection from it.
	fmt.Println("Insurance?")
	fmt.Print("(y/N) >> ")
	fmt.Scan(&selection)

	selection = strings.ToUpper(selection)

	for !validsel { // Loop until we have a valid selection.
		switch {
		case selection == "Y": // Yes, we want insurance!
			// Print initial bet and ask for insurance bet.
			fmt.Println()
			fmt.Println("Initial Bet:", Bet)
			fmt.Println("Insurance amount (up to half of bet) > ")
			fmt.Scan(&insbet)

			// Set up "partcent" for check of partial cents in switch cases.
			partcent = math.Trunc(insbet * 100)

			switch {
			case insbet > (Bet / 2): // Is insurance > (bet / 2)?
				fmt.Println("Insurance can't be more than half of your bet!")

			case partcent < (insbet * 100): // Is bet in partial cents?
				fmt.Println("Bet can't use fractions of a cent. Try again.")

			case DealerTotal == 21: // If the dealer DOES have Blackjack, pay out.
				fmt.Println("Dealer has blackjack, you win",
					(insbet * 2), "credits!")
				Credits = Credits + (insbet * 2)
				kutil.Pause(2)
				validsel = true

			default:
				// If we get this far, the bet is valid, and the dealer does NOT
				// have Blackjack. Player lost the bet. Subtract money from total credits.
				fmt.Println("No Blackjack for the dealer, sorry.")
				Credits = Credits - insbet
				kutil.Pause(2)
				validsel = true
			}
		case selection == "N": // No insurance. Set valid selection flag, notify, and exit.
			validsel = true
			fmt.Println()
			fmt.Println("OK, no insurance!")
			kutil.Pause(2)
		}
	}

}

// This function is to allow the player to select their card(s).
// FYI: Whether the hand is soft or not is passed. Not used. It is
//      only included to stop the compiler from complaining.
//      Likely not usable for this function.
func playerselect(incard int, soft bool) (int, bool) {
	var newcard int           // Number of cards in the hand.
	var selection int = 0     // Menu selection.
	var validsel bool = false // Valid selection?
	var double bool = false   // Did the player double down?

	// Print out the menu and get the selection from it. Cycle if invalid.
	for !validsel {
		fmt.Println("1. Hit")
		fmt.Println("2. Stay")
		if incard == 2 {
			// Only offer option to Double Down on first round.
			// Remove on subsequent rounds.
			fmt.Println("3. Double Down")
		}

		fmt.Print(">> ")
		fmt.Scan(&selection)

		if incard > 2 && selection == 3 {
			// If this is any other round than the first round, selecting
			// "3" is invalid. Set selection to "0" to be picked up later.
			selection = 0
		}

		if selection < 1 || selection > 3 {
			// If selection is NOT 1, 2, or 3 (in certain cases),
			// then it's invalid, report and make the user select again.
			fmt.Println("Invalid selection, please try again!")
			kutil.Pause(2)
			fmt.Println()
		} else { // Otherwise, our selection is valid.
			validsel = true
		}

	}

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
	case selection == 3 && incard == 2:
		// Double down. Only an option first time. Take one card, double bet.
		if (Bet * 2) > Credits { // Player doesn't have enough credits to double down..
			newcard = incard + 1
			fmt.Println("Sorry, not enough credits to double down.")
			fmt.Println("Hitting instead.")
			kutil.Pause(2)
		} else { // Enough credits.. OK, let's do this!
			newcard = incard + 1
			Bet = Bet * 2
			double = true
			fmt.Print("Doubling down. Drawing card, bet now ",
				fmt.Sprintf("%.2f", Bet), "!")
		}
		kutil.Pause(2)
	default:
		// We should never ever get to this code.
		// If we do, something has gone horribly wrong earlier.
		fmt.Println("Something has gone horribly wrong.")
	}

	return newcard, double
}

// This function is the "AI" for the dealer. It's a pretty dumb AI right now.
// Changed to casino rules. Dealer hits if less than 17. May update to hit on
// soft 17 as well in the future.. or add option.
func dealerselect(indcard int, soft bool) int {
	var newdcard int // Number of cards in the hand.

	switch {
	case DealerTotal == 17 && soft && Soft17:
		// Dealer hits on Soft 17
		newdcard = indcard + 1
		fmt.Println("Dealer takes a card.")
		kutil.Pause(2)
	case DealerTotal < 17:
		// Less than 17? Hit. Unless tied or winning.
		newdcard = indcard + 1
		fmt.Println("Dealer takes a card.")
		kutil.Pause(2)
	case DealerTotal >= 17:
		// If dealer total greater or equal to 17, stay.
		newdcard = indcard
		fmt.Println("Dealer stays.")
		kutil.Pause(2)
	}

	return newdcard
}
