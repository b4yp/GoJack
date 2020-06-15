# GoJack
Simple terminal version of Blackjack written in Go. Mainly an exercise to practice Go arrays.

## Release Notes:
### v1.0 - 2020/06/15
So this took me on-and-off about 2 weeks to complete and test. I'm pretty happy with it! Likely going to add more features and tweak stuff, but that'll be something for the future. 

Right now, it's a simple Blackjack game. You start with 100 credits, you wager them, you win or lose. As of now, you can only "Hit" or "Stay". Also AI isn't very robust (pretty much just hit if less than 16, stay if higher).

Enjoy!

## Known issues:
   * Chance that the game /may/ crash if there are 5ish cards left in the deck when the game starts. This is because you need 4 cards to start the game, the game only re-shuffles between hands if there are less than 4 cards left in the deck. If you hit once, then there will be no cards left in the deck and the next hit will cause an array out of bounds error. Still thinking of the best way to circumvent this. For now, though, I've played for an hour or so and haven't hit this problem so.. YMMV.
   * Game doesn't clear the screen on OS's other than Linux or Windows. Low priority.
   * Ran into an issue where the game would print out an invalid card on bust. Only encountered once. Haven't been able to reproduce.

## ToDo:
   * Come up with a solution for where/when to shuffle.
   * Revise the code to use global variables and simplify it a bit.
   * Perhaps figure out how to do formatted prints to add to the screen rather than clear it between each action.
   * Add the ability to "Split" on hands.
   * Add the option to purchase "Insurance" for possible dealer Blackjack.
   * Refine the dealer AI to be more nuanced instead of being as simple as it is. Make the game more challenging.
   * Fix bugs as I find them.