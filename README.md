# GoJack
Simple terminal version of Blackjack written in Go. Mainly an exercise to practice Go arrays.

## NOTICE!!! 
Getting a new server in the next week. This project /may/ be moving to a private (yet still accessible to the public) GitLab repository if I can get it set up properly. Stay tuned!

## Release Notes:

### v1.1 - 2020/06/20
I was really in the mood to get some changes done to make this game more "complete" so dug into it this week. Most of the changes were "under the hood" changes, but I did add a couple new features. Now, players can double down if they want as well as make insurance bets if they'd like! 

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
