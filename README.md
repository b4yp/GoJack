# GoJack
Simple terminal version of Blackjack written in Go. Mainly an exercise to practice Go arrays.

## Release Notes:
### v1.0 - 2020/06/15
So this took me on-and-off about 2 weeks to complete and test. I'm pretty happy with it! Likely going to add more features and tweak stuff, but that'll be something for the future. 

Right now, it's a simple Blackjack game. You start with 100 credits, you wager them, you win or lose. As of now, you can only "Hit" or "Stay". Also AI isn't very robust (pretty much just hit if less than 16, stay if higher).

Enjoy!

## Known issues:
   * Game doesn't clear the screen on OS's other than Linux or Windows. Low priority.

## ToDo:
   * Perhaps figure out how to do formatted prints to add to the screen rather than clear it between each action.
   * Add the ability to "Split" on hands.
   * Fix bugs as I find them.