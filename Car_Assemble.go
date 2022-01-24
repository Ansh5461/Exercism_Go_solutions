package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var res float64 
    var t float64 = float64(productionRate)
    res = t*successRate/100
    return res
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var res int
    var r float64
    var t float64 = float64(productionRate)
    r = t*successRate/100
    res = int(r/60)
    return res
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
    var rem, q, res int
    var r uint
    rem = carsCount%10
    q = carsCount/10
    res = q*95000 + rem*10000
    r = uint(res)
    return r
}
