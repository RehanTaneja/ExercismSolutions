package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	rate := float64(productionRate)
    success := successRate/float64(100)
    return rate*success
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    carsPerMinute :=int(carsPerHour/float64(60))
    return carsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	cnt := float64(carsCount)
    groups_of_ten := uint(cnt/float64(10))
    individuals := uint(int(cnt)%10)
    var cost uint = (groups_of_ten*95000)+(individuals*10000)
    return cost
}
