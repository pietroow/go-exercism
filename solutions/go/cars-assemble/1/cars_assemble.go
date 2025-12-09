package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(productionRate) * (successRate / float64(100))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    perHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    return int(perHour) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// each 10.000
    // 10 cars = 95.000
    // 37 cars = 10 x 3 = 95 * 3 = 285.000 + 70.000 = 355.000

    unitsOfTen := carsCount / 10
    priceOfGroupedCars := unitsOfTen * 95000

    unitsLeft := carsCount % 10
    priceOfUnits := unitsLeft * 10000

    return uint(priceOfGroupedCars + priceOfUnits)
}
