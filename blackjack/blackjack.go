package blackjack

// cardsTable["ace"] = 11
// cardsTable["two"] = 2
// cardsTable["three"] = 3
// cardsTable["four"] = 4
// cardsTable["five"] = 5
// cardsTable["six"] = 6
// cardsTable["seven"] = 7
// cardsTable["eight"] = 8
// cardsTable["nine"] = 9
// cardsTable["ten"] = 10
// cardsTable["jack"] = 10
// cardsTable["queen"] = 10
// cardsTable["king"] = 10

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardsTable := make(map[string]int)
	cardsTable["ace"] = 11
	cardsTable["two"] = 2
	cardsTable["three"] = 3
	cardsTable["four"] = 4
	cardsTable["five"] = 5
	cardsTable["six"] = 6
	cardsTable["seven"] = 7
	cardsTable["eight"] = 8
	cardsTable["nine"] = 9
	cardsTable["ten"] = 10
	cardsTable["jack"] = 10
	cardsTable["queen"] = 10
	cardsTable["king"] = 10

	return cardsTable[card]

}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerCount := ParseCard(card1) + ParseCard(card2)
	dealerCount := ParseCard(dealerCard)

	var returnString string

	switch {
	case playerCount == 22:
		returnString = "P"
		return returnString
	case playerCount == 21 && dealerCount <= 9:
		returnString = "W"
		return returnString
	case playerCount == 21 && dealerCount > 9:
		returnString = "S"
		return returnString
	case playerCount >= 17 && playerCount <= 20:
		returnString = "S"
		return returnString
	case playerCount >= 12 && playerCount <= 16 && dealerCount < 7:
		returnString = "S"
		return returnString
	case playerCount >= 12 && playerCount <= 16 && dealerCount >= 7:
		returnString = "H"
		return returnString
	case playerCount <= 11:
		returnString = "H"
		return returnString
	}
	return returnString
}
