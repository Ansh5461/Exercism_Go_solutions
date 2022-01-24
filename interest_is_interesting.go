package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
        return 3.213
    } else if balance >= 0 && balance < 1000 {
    return 0.5
    } else if balance >= 1000 && balance < 5000 {
    return 1.621
    } else {
    return 2.475
    }
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	var in float64
    in = balance * float64(InterestRate(balance)/100)
    return in
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	var f float64
    f = balance + Interest(balance)
    return f
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var t int
    for {
        balance = AnnualBalanceUpdate(balance)
        if balance <= targetBalance {
            t = t+1
        } else {
        break
        }
    }
return t+1
}
