package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
        case "ace": return 11
        case "two": return 2
        case "three": return 3
        case "four": return 4
        case "five": return 5
        case "six": return 6
        case "seven": return 7
        case "eight": return 8
        case "nine": return 9
        case "ten": return 10
        case "jack": return 10
        case "queen": return 10
        case "king": return 10
        case "other": return 0
        default: return 0
    }
    return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    card1Value := ParseCard(card1)
    card2Value := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
    totalPoints := card1Value + card2Value

    if totalPoints == 21 {
        if dealerCardValue == 10 || dealerCardValue == 11 {
            return "S"
        }
        return "W"
    } else if totalPoints >= 17 && totalPoints <= 20 {
        return "S"
    } else if totalPoints >= 12 && totalPoints <= 16 {
        if dealerCardValue >= 7 {
            return "H"
        }
        return "S"
    } else if card1Value == 11 && card2Value == 11 {
        return "P"
    }
	return "H"
}
