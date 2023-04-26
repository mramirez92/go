# Exercism: Conditionals Switch with Blackjack

## Instructions

In this exercise we will simulate the first turn of a Blackjack game.

You will receive two cards and will be able to see the face up card of the dealer. All cards are represented using a string such as "ace", "king", "three", "two", etc

Depending on your two cards and the card of the dealer, there is a strategy for the first turn of the game, in which you have the following options:

Stand (S)
Hit (H)
Split (P)
Automatically win (W)
Although not optimal yet, you will follow the strategy your friend Alex has been developing, which is as follows:

- If you have a pair of aces you must always split them.

- If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.

- If your cards sum up to a value within the range [17, 20] you should always stand.
- If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.

- If your cards sum up to 11 or lower you should always hit.

## Functions

1. Calculate the value of any given card.
Implement a function to calculate the numerical value of a card

2. Implement the decision logic for the first turn.
Write a function that implements the decision logic as described above:
