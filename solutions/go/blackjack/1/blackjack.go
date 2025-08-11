package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var val int
	switch card {
	case "ace":
		val = 11
	case "two":
		val = 2
	case "three":
		val = 3
	case "four":
		val = 4
	case "five":
		val = 5
	case "six":
		val = 6
	case "seven":
		val = 7
	case "eight":
		val = 8
	case "nine":
		val = 9
	case "ten":
		fallthrough
	case "jack":
		fallthrough
	case "king":
		fallthrough
	case "queen":
		val = 10
	default:
		val = 0
	}

	return val
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	/*
			- Stand (S)
		- Hit (H)
		- Split (P)
		- Automatically win (W)

		Although not optimal yet, you will follow the strategy your friend Alex has been developing, which is as follows:

		- If you have a pair of aces you must always split them.
		- If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace,
			a figure or a ten then you automatically win. If the dealer does have any of those cards then
			you'll have to stand and wait for the reveal of the other card.
		- If your cards sum up to a value within the range [17, 20] you should always stand.
		- If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7
			or higher, in which case you should always hit.
		- If your cards sum up to 11 or lower you should always hit.
	*/
	parsedCardOne, parsedCardTwo, parsedDealerCard := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
	switch {
	case parsedCardOne == parsedCardTwo && parsedCardOne == 11:
		return "P"
	case ((parsedCardOne + parsedCardTwo) == 21) && (parsedDealerCard < 10):
		return "W"
	case ((parsedCardOne + parsedCardTwo) == 21) && (parsedDealerCard >= 10):
		return "S"
	case ((parsedCardOne + parsedCardTwo) >= 12) && ((parsedCardOne + parsedCardTwo) <= 16) && parsedDealerCard >= 7:
		return "H"
	case ((parsedCardOne + parsedCardTwo) >= 12) && ((parsedCardOne + parsedCardTwo) <= 16):
		return "S"
	case ((parsedCardOne + parsedCardTwo) <= 11):
		return "H"
	}
	return "S"
}
