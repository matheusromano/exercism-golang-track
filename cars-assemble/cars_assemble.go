package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successWorkingCars := float64(productionRate) * successRate / 100
	return successWorkingCars
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerMinute := CalculateWorkingCarsPerHour(productionRate, successRate) / 60
	return int(workingCarsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	costIndividualCar := 10000
	costGroupTenCars := 95000
	groupByten := carsCount / 10
	remainder := carsCount % 10
	amoutOfCost := groupByten*costGroupTenCars + costIndividualCar*remainder
	return uint(amoutOfCost)
}
