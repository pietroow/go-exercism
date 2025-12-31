package interest

const (
    lessThanZero = float32(3.213)
    lessThanThousand = float32(0.5)
    thousandPlusAndLessThanFiveThousand = float32(1.621)
    greaterThanFiveThousand = float32(2.475)
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    switch {
        case balance < 0: return lessThanZero
        case balance >= 0 && balance < 1000: return lessThanThousand
        case balance >= 1000 && balance < 5000: return thousandPlusAndLessThanFiveThousand
        case balance >= 5000: return greaterThanFiveThousand
    }
    return float32(0)
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    return float64((balance * float64(InterestRate(balance))) / 100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
    return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	if balance >= targetBalance {
        return 0
    }

    var count int
    for {
        count++
        balance += Interest(balance)
        if balance >= targetBalance {
            return count
        }
    }
    return count
}
