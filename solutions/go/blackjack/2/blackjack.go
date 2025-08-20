package blackjack
import "fmt"
// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace": return 11
        case "ten", "jack", "queen", "king": return 10
        case "nine" : return 9
        case "eight" : return 8
        case "seven" : return 7
        case "six" : return 6
        case "five" : return 5
        case "four" : return 4 
        case "three" : return 3
        case "two" : return 2
        default: return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    myCardSum := ParseCard(card1) + ParseCard(card2)
    fmt.Sprintf("value of myCard = %d, dealerSum = %d", myCardSum, ParseCard(dealerCard))
	switch {
        case myCardSum == 22: return "P"
        case myCardSum == 21:
            if ParseCard(dealerCard) < 10 {
                return "W"
            } else if ParseCard(dealerCard) == 11 || ParseCard(dealerCard) == 10 {
                return "S"
            }
    	case myCardSum >= 17 && myCardSum <= 20 || (myCardSum >= 12 && myCardSum <= 16 && ParseCard(dealerCard) < 7):
        	return "S"
        case myCardSum <= 11 || (myCardSum >= 12 && myCardSum <= 16 && ParseCard(dealerCard) >= 7):
        	return "H"
    }
    return ""
}
